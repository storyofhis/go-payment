package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/storyofhis/go-payment/config"
	"github.com/storyofhis/go-payment/controllers"
	"github.com/storyofhis/go-payment/repositories"
	"github.com/storyofhis/go-payment/services"
)

func main() {
	// setup to connect database
	db, err := config.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		repo = repositories.NewRepositories(db)

		accSvc = services.NewAccountSvc(repo)
		paySvc = services.NewPaymentSvc(repo)

		ctrll = controllers.NewControllers(repo, *paySvc, *accSvc)
	)

	// router
	router := mux.NewRouter()
	route := config.NewRouter(db, router, ctrll)

	// server configuration
	srv := &http.Server{
		Handler:      route.Start(),
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
