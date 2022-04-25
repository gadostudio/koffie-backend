package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/products"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

	now := time.Now()

	data.CreatedAt = &now

	if err := conn.Create(&data); err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	conn, _ := db.Connection()

	var all_products []products.Product

	var cat []products.Category

	conn.Find(&cat)

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

	key, err := strconv.ParseUint(vars["code"], 10, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var _products []products.Product

	connect.Find(&_products, products.Product{Id: uint(key)})

	json.NewEncoder(w).Encode(&_products)
	w.WriteHeader(http.StatusOK)

}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	conn, _ := db.Connection()

	var all_categories []products.Category

	conn.Find(&all_categories)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(all_categories)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {

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

	var data products.Category

	if err := json.Unmarshal(reqBody, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now()

	data.CreatedAt = &now

	if err := conn.Create(&data); err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}
