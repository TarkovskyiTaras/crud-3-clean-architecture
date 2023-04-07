package loan

import (
	"errors"
	"fmt"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	bmock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book/mocks"
	umock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type loanTest struct {
	user          *entity.User
	book          *entity.Book
	errGetUser    error
	errGetBook    error
	errUpdateUser error
	errUpdateBook error
	times         timesToCall
	want          testWant
}

type testWant struct {
	user     *entity.User
	book     *entity.Book
	errFinal error
}

type timesToCall struct {
	ttcGetUser    int
	ttcGetBook    int
	ttcUpdateUser int
	ttcUpdateBook int
}

var errUseCase = errors.New("some usecase error")

func TestBorrow_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m1 := umock.NewMockUseCase(controller)
	m2 := bmock.NewMockUseCase(controller)
	l := NewLoan(m1, m2)

	u1 := &entity.User{ID: 1}
	b1 := &entity.Book{ID: 3, Quantity: 5}
	u1Want := &entity.User{ID: 1, Books: []int{3}}
	b1Want := &entity.Book{ID: 3, Quantity: 4}

	tests := []loanTest{
		{user: u1, book: b1, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, want: testWant{user: u1Want, book: b1Want, errFinal: nil}},
	}

	for _, lt := range tests {
		m1.EXPECT().GetByIDUser(lt.user.ID).Return(lt.user, lt.errGetUser)
		m2.EXPECT().GetByIDBook(lt.book.ID).Return(lt.book, lt.errGetBook)
		m1.EXPECT().UpdateUser(lt.user).Return(lt.errUpdateUser)
		m2.EXPECT().UpdateBook(lt.book).Return(lt.errUpdateBook)

		errGot := l.Borrow(lt.user.ID, lt.book.ID)

		assert.Equal(t, lt.want.errFinal, errGot)
		assert.Equal(t, lt.want.user, lt.user)
		assert.Equal(t, lt.want.book, lt.book)
	}

}

func TestBorrow_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m1 := umock.NewMockUseCase(controller)
	m2 := bmock.NewMockUseCase(controller)
	l := NewLoan(m1, m2)

	tests := []loanTest{
		{user: &entity.User{ID: 1}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: errUseCase, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 0, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: errUseCase, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: nil, errUpdateUser: errUseCase, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 1, ttcUpdateBook: 0}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: errUseCase, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 1, ttcUpdateBook: 1}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: entity.ErrNotFound, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 0, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: fmt.Errorf("user %w", entity.ErrNotFound)}},
		{user: &entity.User{ID: 1}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: entity.ErrNotFound, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: fmt.Errorf("book %w", entity.ErrNotFound)}},
		{user: &entity.User{ID: 2}, book: &entity.Book{ID: 4, Quantity: 0}, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errors.New("not enough books")}},
		{user: &entity.User{ID: 3, Books: []int{5}}, book: &entity.Book{ID: 5, Quantity: 14}, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errors.New("book already borrowed")}},
	}

	for i, lt := range tests {
		fmt.Println("****", i, "****")
		m1.EXPECT().GetByIDUser(lt.user.ID).Return(lt.user, lt.errGetUser)
		m2.EXPECT().GetByIDBook(lt.book.ID).Return(lt.book, lt.errGetBook).Times(lt.times.ttcGetBook)
		m1.EXPECT().UpdateUser(lt.user).Return(lt.errUpdateUser).Times(lt.times.ttcUpdateUser)
		m2.EXPECT().UpdateBook(lt.book).Return(lt.errUpdateBook).Times(lt.times.ttcUpdateBook)

		errGot := l.Borrow(lt.user.ID, lt.book.ID)
		assert.Equal(t, lt.want.errFinal, errGot)
	}
}

func TestReturn_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m1 := umock.NewMockUseCase(controller)
	m2 := bmock.NewMockUseCase(controller)
	l := NewLoan(m1, m2)

	tests := []loanTest{
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, want: testWant{user: &entity.User{ID: 1, Books: []int{}}, book: &entity.Book{ID: 3, Quantity: 6}, errFinal: nil}},
	}

	for _, lt := range tests {
		m1.EXPECT().GetByIDUser(lt.user.ID).Return(lt.user, lt.errGetUser)
		m2.EXPECT().GetByIDBook(lt.book.ID).Return(lt.book, lt.errGetBook)
		m1.EXPECT().UpdateUser(lt.user).Return(lt.errUpdateUser)
		m2.EXPECT().UpdateBook(lt.book).Return(lt.errUpdateBook)

		errGot := l.Return(lt.user.ID, lt.book.ID)
		assert.Equal(t, lt.want.errFinal, errGot)
		assert.Equal(t, lt.want.user, lt.user)
		assert.Equal(t, lt.want.book, lt.book)
	}
}

func TestReturn_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m1 := umock.NewMockUseCase(controller)
	m2 := bmock.NewMockUseCase(controller)
	l := NewLoan(m1, m2)

	tests := []loanTest{
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: errUseCase, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 0, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: errUseCase, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: nil, errUpdateUser: errUseCase, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 1, ttcUpdateBook: 0}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: errUseCase, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 1, ttcUpdateBook: 1}, want: testWant{errFinal: errUseCase}},
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: entity.ErrNotFound, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 0, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: fmt.Errorf("user %w", entity.ErrNotFound)}},
		{user: &entity.User{ID: 1, Books: []int{3}}, book: &entity.Book{ID: 3, Quantity: 5}, errGetUser: nil, errGetBook: entity.ErrNotFound, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: fmt.Errorf("book %w", entity.ErrNotFound)}},
		{user: &entity.User{ID: 3, Books: []int{}}, book: &entity.Book{ID: 5, Quantity: 14}, errGetUser: nil, errGetBook: nil, errUpdateUser: nil, errUpdateBook: nil, times: timesToCall{ttcGetBook: 1, ttcUpdateUser: 0, ttcUpdateBook: 0}, want: testWant{errFinal: errors.New("book was never borrowed")}},
	}

	for i, lt := range tests {
		fmt.Println("****", i, "****")
		m1.EXPECT().GetByIDUser(lt.user.ID).Return(lt.user, lt.errGetUser)
		m2.EXPECT().GetByIDBook(lt.book.ID).Return(lt.book, lt.errGetBook).Times(lt.times.ttcGetBook)
		m1.EXPECT().UpdateUser(lt.user).Return(lt.errUpdateUser).Times(lt.times.ttcUpdateUser)
		m2.EXPECT().UpdateBook(lt.book).Return(lt.errUpdateBook).Times(lt.times.ttcUpdateBook)

		errGot := l.Return(lt.user.ID, lt.book.ID)
		assert.Equal(t, lt.want.errFinal, errGot)
	}
}
