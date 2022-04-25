package controllers

import (
	"encoding/json"
	"fmt"
	mid "github.com/midtrans/midtrans-go"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/midtrans"
	"github.com/shaderboi/koffie-backend/api/settings"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type PaymentRequest struct {
	Payment  midtrans.Payment `json:"payment"`
	Quantity int32            `json:"quantity"`
}

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

	var d PaymentRequest

	if err := json.Unmarshal(reqBody, &d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := d.Payment

	details := []mid.ItemDetails{}

	for _, s := range data.Products {
		itemDetails := mid.ItemDetails{
			Name:  s.Name,
			Price: int64(s.Price),
			Qty:   d.Quantity,
		}
		details = append(details, itemDetails)
	}

	req := midtrans.GetTransactionDetails(fmt.Sprintf("KOFFIE-ORDER-%d", rand.Int()), data.Amount, data.Phone)

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
