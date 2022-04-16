package users

import "time"

type User struct {
	PhoneNumber string
	Gender      int
	BirthDate   time.Time
}
