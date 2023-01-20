package config

import (
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"github.com/storyofhis/go-payment/controllers"
)

type Routes struct {
	db    *bolt.DB
	route *mux.Router

	paymentController *controllers.PaymentControllers
}

func NewRouter(dbConfig *bolt.DB, route *mux.Router, payController *controllers.PaymentControllers) *Routes {
	return &Routes{
		db:                dbConfig,
		route:             route,
		paymentController: payController,
	}
}

func (r *Routes) Start() *mux.Router {

	r.route.HandleFunc("/v1/payments/authorization", r.paymentController.PaymentsAuthorization).Methods("POST")
	return r.route
}
