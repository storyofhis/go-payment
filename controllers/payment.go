package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/storyofhis/go-payment/controllers/params"
	"github.com/storyofhis/go-payment/repositories"
	"github.com/storyofhis/go-payment/repositories/models"
	"github.com/storyofhis/go-payment/services"
)

type PaymentControllers struct {
	repo repositories.Repositories
	psvc services.PaymentSvc
	asvc services.AccountSvc
}

func NewControllers(repo repositories.Repositories, psvc services.PaymentSvc, asvc services.AccountSvc) *PaymentControllers {
	return &PaymentControllers{
		repo: repo,
		psvc: psvc,
		asvc: asvc,
	}
}

func (control *PaymentControllers) PaymentsAuthorization(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Request
	var authRequest params.Authorization
	err = json.Unmarshal(b, &authRequest)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Prepare Payment Response
	w.Header().Set("content-type", "application/json")

	// Basic Validation - Business Account
	var businessAccountId = r.Header.Get("From")
	var businessAccount models.Account
	if len(businessAccountId) > 0 {
		businessAccount, _ = control.repo.GetAccount(businessAccountId)
	} else {
		json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(authRequest.OrderID, "3", "Invalid Merchant"))
		return
	}
	if businessAccount.Id != businessAccountId {
		json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(authRequest.OrderID, "15", "No Such Issuer"))
		return
	}

	// Basic Validation - Personal Account
	var personalAccountId = fmt.Sprintf("%v", authRequest.CardNumber)
	var personalAccount models.Account
	if len(personalAccountId) > 0 {
		personalAccount, _ = control.repo.GetAccount(personalAccountId)
	} else {
		json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(authRequest.OrderID, "12", "Invalid Card Number"))
		return
	}
	if personalAccount.Id != personalAccountId {
		json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(authRequest.OrderID, "56", "No Card Record"))
		return
	}

	if personalAccount.CardNumber != authRequest.CardNumber ||
		personalAccount.CardSecurityCode != authRequest.CardSecurityCode ||
		personalAccount.CardExpiryYear != authRequest.CardExpiryYear ||
		personalAccount.CardExpiryMonth != authRequest.CardExpiryMonth {
		payment := control.psvc.CreateAuthorizationPayment(authRequest,
			personalAccount,
			businessAccount,
			"5",
			"Do not Honour",
		)
		control.repo.SavePayment(*payment)
		businessAccount.Statement = append(businessAccount.Statement, payment.Id)
		control.repo.SaveAccount(businessAccount)
		json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(payment.Id, "5", "Do not Honour"))
		return
	}
	if personalAccount.Available < authRequest.Amount {
		payment := control.psvc.CreateAuthorizationPayment(authRequest,
			personalAccount,
			businessAccount,
			"51",
			"Insufficient Funds",
		)
		control.repo.SavePayment(*payment)
		businessAccount.Statement = append(businessAccount.Statement, payment.Id)
		control.repo.SaveAccount(businessAccount)
		json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(payment.Id, "51", "Insufficient Funds"))
		return
	}

	// Successful Payment
	personalAccount.Available = personalAccount.Available - authRequest.Amount
	personalAccount.Blocked = personalAccount.Blocked + authRequest.Amount
	control.repo.SaveAccount(personalAccount)
	businessAccount.Blocked = businessAccount.Blocked + authRequest.Amount
	control.repo.SaveAccount(businessAccount)
	payment := control.psvc.CreateAuthorizationPayment(authRequest,
		personalAccount,
		businessAccount,
		"0",
		"Approved",
	)
	control.repo.SavePayment(*payment)
	businessAccount.Statement = append(businessAccount.Statement, payment.Id)
	control.repo.SaveAccount(businessAccount)
	personalAccount.Statement = append(personalAccount.Statement, payment.Id)
	control.repo.SaveAccount(personalAccount)
	json.NewEncoder(w).Encode(control.psvc.CreatePaymentResponse(payment.Id, "0", "Approved"))

	return
}
