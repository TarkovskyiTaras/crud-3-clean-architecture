package repository

import (
	"database/sql"
	"fmt"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"time"
)

type UserPostgreSQL struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *UserPostgreSQL {
	return &UserPostgreSQL{db: db}
}

func (u *UserPostgreSQL) Create(user *entity.User) error {
	_, err := u.db.Exec("INSERT INTO users (id, first_name, last_name, dob, location, cellphone_number, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		(*user).ID, user.FirstName, user.LastName, user.DOB, user.Location, user.CellPhoneNumber, user.Email, user.Password, user.CreatedAt.Format("2006-01-02 15:04:05"), time.Time{})
	return err
}

func (u *UserPostgreSQL) GetByID(id int) (*entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT id, first_name, last_name, dob, location, cellphone_number, email, password, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.DOB, &user.Location, &user.CellPhoneNumber, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, entity.ErrNotFound
	}

	//new code
	rows, _ := u.db.Query("SELECT id_book FROM users_books WHERE id_user = $1", id)
	for rows.Next() {
		var i int
		rows.Scan(&i)
		user.Books = append(user.Books, i)
	}
	//end of new code

	return &user, err
}

func (u *UserPostgreSQL) GetAll() ([]*entity.User, error) {
	rows, err := u.db.Query("SELECT id, first_name, last_name, dob, location, cellphone_number, email, password, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.DOB, &user.Location, &user.CellPhoneNumber, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	//new code
	for _, user := range users {
		rows, _ = u.db.Query("SELECT id_book FROM users_books WHERE id_user = $1", user.ID)
		for rows.Next() {
			var b int
			rows.Scan(&b)
			user.Books = append(user.Books, b)
		}
	}
	//end of new code

	return users, nil
}

func (u *UserPostgreSQL) Update(user *entity.User) error {
	res, err := u.db.Exec("UPDATE users SET first_name = $1, last_name = $2, dob = $3, location = $4, cellphone_number = $5, email = $6, password = $7, updated_at = $8 WHERE id = $9",
		user.FirstName, user.LastName, user.DOB, user.Location, user.CellPhoneNumber, user.Email, user.Password, user.UpdatedAt.Format("2006-01-02 15:04:05"), user.ID)
	if err != nil {
		return err
	}

	//new code
	u.db.Exec("DELETE FROM users_books WHERE id_user = $1", user.ID)
	for _, bookId := range user.Books {
		u.db.Exec("INSERT INTO users_books (id_user, id_book) VALUES($1, $2)", user.ID, bookId)
	}
	//end of new code

	rowsAff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAff != 1 {
		return fmt.Errorf("weird behavior, total rows affected = %d", rowsAff)
	}

	return nil
}

func (u *UserPostgreSQL) Delete(id int) error {
	res, err := u.db.Exec("DELETE FROM users WHERE id = $1 ", id)
	if err != nil {
		return err
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAff != 1 {
		return fmt.Errorf("weird behavior, total rows affected = %d", rowsAff)
	}

	return nil
}
