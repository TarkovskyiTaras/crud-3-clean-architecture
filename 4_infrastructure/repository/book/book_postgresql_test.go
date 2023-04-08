package repositoryBook

import (
	"database/sql"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/5_pkg/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var db *sql.DB

type bookTest struct {
	args bookArgs
	want bookWant
}
type bookArgs struct {
	book *entity.Book
}
type bookWant struct {
	book  *entity.Book
	books []*entity.Book
	err   error
}

func setUp() {
	var err error
	db, err = database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	if err != nil {
		log.Fatal(err)
	}
	initialBook := &entity.Book{ID: 1, Tittle: "Concrete Design Handbook", Author: "Tarkovskyi T", Pages: 290, Quantity: 5}

	_, err = db.Exec("INSERT INTO books (id, tittle, author, pages, quantity, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)",
		initialBook.ID, initialBook.Tittle, initialBook.Author, initialBook.Pages, initialBook.Quantity, time.Time{}, time.Time{})
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	defer db.Close()

	_, err := db.Exec("DELETE FROM books")
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
}

func TestCreate(t *testing.T) {
	bookRepo := NewBooks(db)
	bookArg1 := &entity.Book{ID: 2, Tittle: "Handbook of Steel Construction", Author: "CISC ICCA", Pages: 354, Quantity: 10, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}
	tests := []bookTest{
		{args: bookArgs{book: bookArg1}, want: bookWant{book: bookArg1, err: nil}},
	}

	for _, bt := range tests {
		errGot := bookRepo.Create(bt.args.book)
		bookGot, err := bookRepo.GetByID(bt.args.book.ID)
		if err != nil {
			log.Fatal(err)
		}

		bookGot.CreatedAt = bookGot.CreatedAt.UTC()
		bookGot.UpdatedAt = bookGot.UpdatedAt.UTC()

		assert.Equal(t, bt.want.book, bookGot)
		assert.Equal(t, bt.want.err, errGot)
	}
}

func TestGetByID(t *testing.T) {
	bookRepo := NewBooks(db)
	bookArg1 := &entity.Book{ID: 1}
	tests := []bookTest{
		{args: bookArgs{book: bookArg1}, want: bookWant{book: bookArg1, err: nil}},
	}
	for _, bt := range tests {
		bookGot, errGot := bookRepo.GetByID(bt.args.book.ID)

		assert.Equal(t, bt.want.book.ID, bookGot.ID)
		assert.Equal(t, bt.want.err, errGot)
	}
}

func TestGetAll(t *testing.T) {
	bookRepo := NewBooks(db)

	tests := []bookTest{
		{want: bookWant{err: nil}},
	}
	for _, bt := range tests {
		booksGot, errGot := bookRepo.GetAll()

		assert.NotNil(t, booksGot)
		assert.Equal(t, bt.want.err, errGot)
	}
}

func TestUpdate(t *testing.T) {
	bookRepo := NewBooks(db)
	bookArg1 := &entity.Book{ID: 1, Tittle: "UPD_Concrete Design Handbook", Author: "UPD_Tarkovskyi T", Pages: 290, Quantity: 5}
	tests := []bookTest{
		{args: bookArgs{book: bookArg1}, want: bookWant{book: bookArg1, err: nil}},
	}

	for _, bt := range tests {
		errGot := bookRepo.Update(bt.args.book)
		bookGot, err := bookRepo.GetByID(bt.args.book.ID)
		if err != nil {
			log.Fatal(err)
		}

		bookGot.CreatedAt = bookGot.CreatedAt.UTC()
		bookGot.UpdatedAt = bookGot.UpdatedAt.UTC()

		assert.Equal(t, bt.want.book, bookGot)
		assert.Equal(t, bt.want.err, errGot)
	}
}

func TestDelete(t *testing.T) {
	bookRepo := NewBooks(db)
	bookArg1 := &entity.Book{ID: 1}
	tests := []bookTest{
		{args: bookArgs{book: bookArg1}, want: bookWant{err: nil}},
	}
	for _, bt := range tests {
		errGot := bookRepo.Delete(bt.args.book.ID)
		bookGot, err := bookRepo.GetByID(bt.args.book.ID)

		assert.NotNil(t, err)
		assert.Nil(t, bookGot)
		assert.Equal(t, bt.want.err, errGot)
	}
}
