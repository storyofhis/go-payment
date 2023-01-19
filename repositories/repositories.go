package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/storyofhis/go-payment/config"
	"github.com/storyofhis/go-payment/repositories/models"
)

type Repositories struct {
	db *bolt.DB
}

func NewRepositories(db *bolt.DB) *Repositories {
	return &Repositories{
		db: db,
	}
}

func (repo *Repositories) SaveAccount(account models.Account) error {
	var key = fmt.Sprintf("%v", account.Id)
	var value, err = json.Marshal(account)
	if err != nil {
		return fmt.Errorf("Could not json marshal entry: %v", err)
	}

	err = repo.db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(config.RootBucket).Bucket(config.AccountBucket).Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("Could not insert entry: %v", err)
		}
		return nil
	})

	fmt.Println("Added Account Entry " + account.Id)
	return err
}

func (repo *Repositories) GetAccount(id string) (models.Account, error) {
	var account models.Account
	var key = []byte(id)

	err := repo.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(config.RootBucket).Bucket(config.AccountBucket).Get(key)
		fmt.Sprintf("Found Account Entry %s", value)
		return json.Unmarshal(value, &account)
	})

	if err != nil {
		fmt.Printf("Could not get Account ID %s", id)
		return account, err
	}
	return account, nil
}

func (repo *Repositories) SavePayment(payment models.Payment) error {
	var key = fmt.Sprintf("%v", payment.Id)
	var value, err = json.Marshal(payment)
	if err != nil {
		return fmt.Errorf("Could not json marshal entry: %v", err)
	}

	err = repo.db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(config.RootBucket).Bucket(config.AccountBucket).Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("Could not insert entry: %v", err)
		}
		return nil
	})
	fmt.Println("Added Payment Entry " + key)
	return err
}

func (repo *Repositories) GetPayment(id string) (models.Payment, error) {
	var key = []byte(id)
	var payment models.Payment

	err := repo.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(config.RootBucket).Bucket(config.AccountBucket).Get(key)
		fmt.Sprintf("Found Payment Entry %s", value)
		return json.Unmarshal(value, &payment)
	})

	if err != nil {
		fmt.Printf("Could not get Payment ID %s", id)
		return payment, err
	}
	return payment, nil
}
