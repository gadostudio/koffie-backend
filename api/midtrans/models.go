package midtrans

import "time"

type Payment struct {
	Amount    int64      `json:"amount"`
	Phone     string     `json:"phone"`
	CreatedAt *time.Time `json:"created_at"`
}
