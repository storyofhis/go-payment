package config

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var (
	RootBucket              = []byte("DB")
	AccountBucket           = []byte("ACCOUNT")
	PaymentBucket           = []byte("PAYMENT")
	businessStatementBucket = []byte("BUSINESS_STATEMENT")
	personalStatementBucket = []byte("PERSONAL_STATEMENT")
)

func SetupDB() (*bolt.DB, error) {
	db, err := bolt.Open("main.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists(RootBucket)
		if err != nil {
			return fmt.Errorf("Could not Create Root Bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists(AccountBucket)
		if err != nil {
			return fmt.Errorf("Could not Create Account Bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists(PaymentBucket)
		if err != nil {
			return fmt.Errorf("Could not Create Payment Bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Could not set up bucket, %v", err)
	}
	return db, nil
}
