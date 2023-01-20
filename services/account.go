package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/storyofhis/go-payment/controllers/views"
	"github.com/storyofhis/go-payment/repositories"
	"github.com/storyofhis/go-payment/repositories/models"
)

type AccountSvc struct {
	Repo repositories.Repositories
}

func NewAccountSvc(repo repositories.Repositories) *AccountSvc {
	return &AccountSvc{
		Repo: repo,
	}
}

func (asvc *AccountSvc) CreateAccountResponse(account_id, status, description string) *views.Account {
	return &views.Account{
		AccountId:   account_id,
		Status:      status,
		Description: description,
	}
}

func (asvc *AccountSvc) GenerateAccount() *models.Account {
	identifier := rand.Intn(4666778181156223-4666000000000000) + 4666000000000000
	return &models.Account{
		Id:               fmt.Sprintf("%v", identifier),
		Available:        0,
		Blocked:          0,
		Deposited:        0,
		Currency:         "GBP",
		CardName:         "Mr Payment",
		CardType:         "VISA",
		CardNumber:       identifier,
		CardExpiryMonth:  rand.Intn(12-1) + 1,
		CardExpiryYear:   rand.Intn(24-18) + 18,
		CardSecurityCode: rand.Intn(999-100) + 100,
		CreationTime:     time.Now(),
	}
}
