package controllers

import (
	"encoding/json"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/midtrans"
	"github.com/shaderboi/koffie-backend/api/settings"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {

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

	var data midtrans.Payment

	if err := json.Unmarshal(reqBody, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := midtrans.GetTransactionDetails("KOFFIE-ORDER-"+string(rand.Int()), data.Amount, data.Phone)

	snapResp, _ := settings.S.CreateTransaction(req)

	err = json.NewEncoder(w).Encode(snapResp)

	if err != nil {
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
