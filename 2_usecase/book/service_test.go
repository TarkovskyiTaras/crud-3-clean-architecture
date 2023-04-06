package book

import (
	"errors"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	bmock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type bookTest struct {
	book *entity.Book
	want wantBook
	t    timesToCall
}
type wantBook struct {
	book          *entity.Book
	books         []*entity.Book
	errFromCreate error
	errFromGet    error
	errFromGetAll error
	errFromUpdate error
	errFromDelete error
	errFinal      error
}

type timesToCall struct {
	ttcCreate int
	ttcUpdate int
	ttcDelete int
}

func TestCreateBook_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1, Tittle: "God's Little Acre", Author: "Erskine Caldwell", Pages: 224, Quantity: 5}

	tests := []bookTest{
		{book: b1, want: wantBook{book: nil, errFromGet: entity.ErrNotFound, errFromCreate: nil, errFinal: nil}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		m.EXPECT().Create(bt.book).Return(bt.want.errFromCreate)

		errGot := b.CreateBook(bt.book)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestCreateBook_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1, Tittle: "God's Little Acre", Author: "Erskine Caldwell", Pages: 224, Quantity: 5}
	b2 := &entity.Book{ID: 1, Tittle: "", Author: "", Pages: 300, Quantity: 3}

	tests := []bookTest{
		{book: b1, want: wantBook{book: b1, errFromGet: nil, errFromCreate: nil, errFinal: entity.ErrConflict}, t: timesToCall{ttcCreate: 0}},
		{book: b2, want: wantBook{book: nil, errFromGet: entity.ErrNotFound, errFromCreate: nil, errFinal: entity.ErrInvalidEntity}, t: timesToCall{ttcCreate: 0}},
		{book: b1, want: wantBook{book: nil, errFromGet: entity.ErrNotFound, errFromCreate: errors.New("some database error"), errFinal: errors.New("some database error")}, t: timesToCall{ttcCreate: 1}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		m.EXPECT().Create(bt.book).Return(bt.want.errFromCreate).Times(bt.t.ttcCreate)

		errGot := b.CreateBook(bt.book)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestGetByIDBook_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1}

	tests := []bookTest{
		{book: b1, want: wantBook{book: b1, errFromGet: nil, errFinal: nil}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		bookGot, errGot := b.GetByIDBook(bt.book.ID)

		assert.Equal(t, bt.want.book, bookGot)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestGetByIDBook_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1}

	tests := []bookTest{
		{book: b1, want: wantBook{book: nil, errFromGet: errors.New("some database error"), errFinal: errors.New("some database error")}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		bookGot, errGot := b.GetByIDBook(bt.book.ID)

		assert.Equal(t, bt.want.book, bookGot)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestGetAllBooks_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	books := []*entity.Book{{ID: 1}, {ID: 2}, {ID: 3}}

	tests := []bookTest{
		{want: wantBook{books: books, errFromGetAll: nil, errFinal: nil}},
	}

	for _, bt := range tests {
		m.EXPECT().GetAll().Return(bt.want.books, bt.want.errFromGetAll)
		booksGot, errGot := b.GetAllBooks()

		assert.Equal(t, bt.want.books, booksGot)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestGetAllBooks_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	tests := []bookTest{
		{want: wantBook{books: nil, errFromGetAll: errors.New("some database error"), errFinal: errors.New("some database error")}},
	}

	for _, bt := range tests {
		m.EXPECT().GetAll().Return(bt.want.books, bt.want.errFromGetAll)
		booksGot, errGot := b.GetAllBooks()

		assert.Equal(t, bt.want.books, booksGot)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestUpdateBook_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1, Tittle: "God's Little Acre", Author: "Erskine Caldwell", Pages: 224, Quantity: 5}

	tests := []bookTest{
		{book: b1, want: wantBook{book: b1, errFromGet: nil, errFromUpdate: nil, errFinal: nil}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		m.EXPECT().Update(bt.book).Return(bt.want.errFromUpdate)

		errGot := b.UpdateBook(bt.book)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestUpdateBook_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1, Tittle: "God's Little Acre", Author: "Erskine Caldwell", Pages: 224, Quantity: 5}
	b2 := &entity.Book{ID: 2, Tittle: "", Author: "", Pages: 300, Quantity: 3}

	tests := []bookTest{
		{book: b1, want: wantBook{book: nil, errFromGet: entity.ErrNotFound, errFromUpdate: nil, errFinal: entity.ErrNotFound}, t: timesToCall{ttcUpdate: 0}},
		{book: b2, want: wantBook{book: b2, errFromGet: nil, errFromUpdate: nil, errFinal: entity.ErrInvalidEntity}, t: timesToCall{ttcUpdate: 0}},
		{book: b1, want: wantBook{book: b1, errFromGet: nil, errFromUpdate: errors.New("some database error"), errFinal: errors.New("some database error")}, t: timesToCall{ttcUpdate: 1}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		m.EXPECT().Update(bt.book).Return(bt.want.errFromUpdate).Times(bt.t.ttcUpdate)

		errGot := b.UpdateBook(bt.book)
		assert.Equal(t, bt.want.errFinal, errGot)

	}
}

func TestDeleteBook_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1}

	tests := []bookTest{
		{book: b1, want: wantBook{book: b1, errFromGet: nil, errFromDelete: nil, errFinal: nil}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		m.EXPECT().Delete(bt.book.ID).Return(bt.want.errFromDelete)

		errGot := b.DeleteBook(bt.book.ID)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}

func TestDeleteBook_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockRepository(controller)
	b := NewService(m)

	b1 := &entity.Book{ID: 1}

	tests := []bookTest{
		{book: b1, want: wantBook{book: nil, errFromGet: entity.ErrNotFound, errFromDelete: nil, errFinal: entity.ErrNotFound}, t: timesToCall{ttcDelete: 0}},
		{book: b1, want: wantBook{book: b1, errFromGet: nil, errFromDelete: errors.New("some database error"), errFinal: errors.New("some database error")}, t: timesToCall{ttcDelete: 1}},
	}

	for _, bt := range tests {
		m.EXPECT().GetByID(bt.book.ID).Return(bt.want.book, bt.want.errFromGet)
		m.EXPECT().Delete(bt.book.ID).Return(bt.want.errFromDelete).Times(bt.t.ttcDelete)

		errGot := b.DeleteBook(bt.book.ID)
		assert.Equal(t, bt.want.errFinal, errGot)
	}
}
