package loan

import (
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user"
)

type Loan struct {
	user user.Users
	book book.Books
}

func NewLoan(u *user.Users, b *book.Books) *Loan {
	return &Loan{user: *u, book: *b}
}

func (l *Loan) Borrow(userID, bookID int) error {
	u, _ := l.user.GetByIDUser(userID)
	b, _ := l.book.GetByIDBook(bookID)

	u.AddBook(bookID)
	l.user.UpdateUser(u)

	b.Quantity--
	l.book.UpdateBook(b)

	return nil
}

func (l *Loan) Return(userID, bookID int) error {
	u, _ := l.user.GetByIDUser(userID)
	b, _ := l.book.GetByIDBook(bookID)

	u.RemoveBook(bookID)
	l.user.UpdateUser(u)

	b.Quantity++
	l.book.UpdateBook(b)

	return nil
}
