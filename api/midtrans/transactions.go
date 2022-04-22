package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func GetTransactionDetails(order_id string, amount int64, phone string, details *[]midtrans.ItemDetails) *snap.Request {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order_id,
			GrossAmt: amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Phone: phone,
		},
		Items: details,
	}

	return req
}
