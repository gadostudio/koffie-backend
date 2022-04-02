package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/products"
	"io/ioutil"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	conn, err := db.Connection()

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data products.Product

	if err := json.Unmarshal(reqBody, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := conn.Create(&data); err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	conn, _ := db.Connection()

	var all_products []products.Product

	conn.Find(&all_products)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(all_products)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	connect, _ := db.Connection()

	vars := mux.Vars(r)

	key := vars["code"]

	var _products []products.Product

	connect.Find(&_products, products.Product{Code: key})

	json.NewEncoder(w).Encode(&_products)
	w.WriteHeader(http.StatusOK)

}

//func GetUsers(w http.ResponseWriter, r *http.Request) {
//
//	ctx := context.Background()
//
//	client := settings.SetupFirebase(ctx)
//
//	u, err := client.GetUser(ctx, uid)
//
//	if err != nil {
//		log.Fatalf("error getting user %s: %v\n", uid, err)
//	}
//
//	log.Println("TES: %s", u.PhoneNumber)
//}
