package entity

import "time"

type User struct {
	ID              int       `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	DOB             time.Time `json:"dob"`
	Location        string    `json:"location"`
	CellPhoneNumber string    `json:"cellphone_number"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
