package entity

import "time"

type Book struct {
	ID        int       `json:"ID"`
	Tittle    string    `json:"Tittle"`
	Author    string    `json:"Author"`
	Pages     int       `json:"Pages"`
	Quantity  int       `json:"Quantity"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
