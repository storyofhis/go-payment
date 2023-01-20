package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/storyofhis/go-payment/controllers/params"
	"github.com/storyofhis/go-payment/controllers/views"
	"github.com/storyofhis/go-payment/repositories"
	"github.com/storyofhis/go-payment/repositories/models"
	"github.com/storyofhis/go-payment/utils"
)

type PaymentSvc struct {
	repo repositories.Repositories
}

func NewPaymentSvc(repo repositories.Repositories) *PaymentSvc {
	return &PaymentSvc{repo: repo}
}

func (psvc *PaymentSvc) CreatePaymentResponse(reference_id, status, description string) *views.Payment {
	return &views.Payment{
		ReferenceId: reference_id,
		Status:      status,
		Description: description,
	}
}

func (psvc *PaymentSvc) CreateAuthorizationPayment(authRequest params.Authorization,
	personalAccount, businessAccount models.Account,
	status, description string,
) *models.Payment {
	return &models.Payment{
		Id:              uuid.New().String(),
		BusinessId:      businessAccount.Id,
		OrderId:         authRequest.OrderID,
		Operation:       utils.AUTHORIZATION,
		OriginalAmount:  authRequest.Amount,
		CurrentAmount:   authRequest.Amount,
		Status:          status,
		Description:     description,
		Currency:        authRequest.Currency,
		CardName:        authRequest.CardName,
		CardType:        personalAccount.CardType,
		CardNumber:      int64(personalAccount.CardNumber),
		CardExpiryMonth: personalAccount.CardExpiryMonth,
		CardExpiryYear:  personalAccount.CardExpiryYear,
		CreationTime:    time.Now(),
	}
}

func (psvc *PaymentSvc) CreateSuccessivePayment(successiveReq params.SuccessiveReq,
	referencedPayment models.Payment,
	status string,
	description string,
) *models.Payment {
	return &models.Payment{
		Id:              uuid.New().String(),
		BusinessId:      referencedPayment.BusinessId,
		OrderId:         successiveReq.OrderID,
		Operation:       successiveReq.Type,
		OriginalAmount:  successiveReq.Amount,
		CurrentAmount:   successiveReq.Amount,
		Status:          status,
		Description:     description,
		Currency:        referencedPayment.Currency,
		CardName:        referencedPayment.CardName,
		CardType:        referencedPayment.CardType,
		CardNumber:      referencedPayment.CardNumber,
		CardExpiryMonth: referencedPayment.CardExpiryMonth,
		CardExpiryYear:  referencedPayment.CardExpiryYear,
		CreationTime:    time.Now(),
	}
}
