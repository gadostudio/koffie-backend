package midtrans

import (
	"github.com/shaderboi/koffie-backend/api/products"
	"time"
)

type Payment struct {
	Amount    int64              `json:"amount"`
	Phone     string             `json:"phone"`
	CreatedAt *time.Time         `json:"created_at"`
	Products  []products.Product `json:"products"`
}
