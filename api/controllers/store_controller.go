package controllers

import (
	"encoding/json"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/stores"
	"io/ioutil"
	"net/http"
)

func CreateStore(w http.ResponseWriter, r *http.Request) {

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

	var data stores.Store

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

func GetAllStores(w http.ResponseWriter, r *http.Request) {

	conn, _ := db.Connection()

	var all_stores []stores.Store

	conn.Find(&all_stores)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(all_stores)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetStoreByCoord(w http.ResponseWriter, r *http.Request) {

	_, _ = db.Connection()

	q := r.URL.Query()
	lat := q["lat"]
	_ = q["lon"]

	//
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//var _store stores.Store
	//
	//connect.Where("coordinate = ?")

	json.NewEncoder(w).Encode(&lat)
	w.WriteHeader(http.StatusOK)

}
