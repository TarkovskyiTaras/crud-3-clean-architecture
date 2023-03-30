package book

import (
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"time"
)

type Books struct {
	repo Repository
}

func NewService(repo Repository) *Books {
	return &Books{repo: repo}
}

func (u *Books) CreateBook(book *entity.Book) error {
	_, err := u.repo.GetByID(book.ID)
	if err != entity.ErrNotFound {
		return entity.ErrConflict
	}

	err = ValidateInput(book)
	if err != nil {
		return err
	}

	book.CreatedAt = time.Now()
	return u.repo.Create(book)
}

func (u *Books) GetByIDBook(id int) (*entity.Book, error) {
	return u.repo.GetByID(id)
}

func (u *Books) GetAllBooks() ([]*entity.Book, error) {
	return u.repo.GetAll()
}

func (u *Books) UpdateBook(book *entity.Book) error {
	_, err := u.repo.GetByID(book.ID)
	if err != nil {
		return err
	}

	err = ValidateInput(book)
	if err != nil {
		return err
	}

	book.UpdatedAt = time.Now()
	return u.repo.Update(book)
}

func (u *Books) DeleteBook(id int) error {
	_, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	return u.repo.Delete(id)
}

func ValidateInput(b *entity.Book) error {
	if b.ID <= 0 || b.Tittle == "" || b.Author == "" || b.Pages <= 0 || b.Quantity <= 0 {
		return entity.ErrInvalidEntity
	}
	return nil
}
