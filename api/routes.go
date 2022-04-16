package api

import (
	"github.com/gorilla/mux"
	"github.com/shaderboi/koffie-backend/api/controllers"
	"github.com/shaderboi/koffie-backend/api/middleware"
	"log"
	"net/http"
	"os"
)

func Routes() {
	r := mux.NewRouter()
	requireAuth := r.Methods(http.MethodPost).Subrouter()
	requireAuth.HandleFunc("/api/v1/payment", controllers.ProcessPayment).Methods("POST")
	requireAuth.HandleFunc("/api/v1/products", controllers.CreateProduct).Methods("POST")
	requireAuth.Use(middleware.AuthMiddleware)
	r.HandleFunc("/api/v1/products", controllers.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/v1/products/{code}", controllers.GetProduct).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
