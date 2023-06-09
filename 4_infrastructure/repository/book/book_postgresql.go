package repositoryBook

import (
	"database/sql"
	"fmt"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"time"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *PostgreSQL {
	return &PostgreSQL{db: db}
}

func (r *PostgreSQL) Create(b *entity.Book) error {
	_, err := r.db.Exec("INSERT INTO books (id, tittle, author, pages, quantity, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)",
		b.ID, b.Tittle, b.Author, b.Pages, b.Quantity, b.CreatedAt, time.Time{})
	return err
}

func (r *PostgreSQL) GetByID(id int) (*entity.Book, error) {
	var book entity.Book
	row := r.db.QueryRow("SELECT id, tittle, author, pages, quantity, created_at, updated_at FROM books WHERE id = $1", id)
	err := row.Scan(&book.ID, &book.Tittle, &book.Author, &book.Pages, &book.Quantity, &book.CreatedAt, &book.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, entity.ErrNotFound
	}
	return &book, err
}

func (r *PostgreSQL) GetAll() ([]*entity.Book, error) {
	rows, err := r.db.Query("SELECT id, tittle, author, pages, quantity, created_at, updated_at FROM books")
	if err != nil {
		return nil, err
	}

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Tittle, &book.Author, &book.Pages, &book.Quantity, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (r *PostgreSQL) Update(e *entity.Book) error {
	res, err := r.db.Exec("UPDATE books SET tittle = $1, author = $2, pages = $3, quantity = $4, updated_at = $5 WHERE id = $6",
		e.Tittle, e.Author, e.Pages, e.Quantity, e.UpdatedAt, e.ID)
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

func (r *PostgreSQL) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
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
