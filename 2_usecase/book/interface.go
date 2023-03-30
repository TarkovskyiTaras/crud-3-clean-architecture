package book

import entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"

type Repository interface {
	Create(b *entity.Book) error
	GetByID(id int) (*entity.Book, error)
	GetAll() ([]*entity.Book, error)
	Update(b *entity.Book) error
	Delete(id int) error
}

type UseCase interface {
	CreateBook(b *entity.Book) error
	GetByIDBook(id int) (*entity.Book, error)
	GetAllBooks() ([]*entity.Book, error)
	UpdateBook(b *entity.Book) error
	DeleteBook(id int) error
}
