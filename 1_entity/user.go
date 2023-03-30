package entity

import (
	"errors"
	"time"
)

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
	Books           []int
}

func (u *User) AddBook(idBook int) error {
	for _, b := range u.Books {
		if b == idBook {
			return errors.New("book already borrowed")
		}
	}
	u.Books = append(u.Books, idBook)
	return nil
}

func (u *User) RemoveBook(idBook int) error {
	for i, bID := range u.Books {
		if bID == idBook {
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("book was never borrowed")
}
