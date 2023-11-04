package enity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Device   string    `json:"device"`
	Ip       string    `json:"ip"`
}
