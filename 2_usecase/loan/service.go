package loan

import (
	"errors"
	"fmt"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user"
)

type Loan struct {
	user user.UseCase
	book book.UseCase
}

func NewLoan(u user.UseCase, b book.UseCase) *Loan {
	return &Loan{user: u, book: b}
}

func (l *Loan) Borrow(userID, bookID int) error {
	u, err := l.user.GetByIDUser(userID)
	if err != nil {
		if err == entity.ErrNotFound {
			return fmt.Errorf("user %w", entity.ErrNotFound)
		}
		return err
	}

	b, err := l.book.GetByIDBook(bookID)
	if err != nil {
		if err == entity.ErrNotFound {
			return fmt.Errorf("book %w", entity.ErrNotFound)
		}
		return err
	}

	if b.Quantity <= 0 {
		return errors.New("not enough books")
	}

	err = u.AddBook(bookID)
	if err != nil {
		return err
	}

	err = l.user.UpdateUser(u)
	if err != nil {
		return err
	}

	b.Quantity--
	err = l.book.UpdateBook(b)
	if err != nil {
		return err
	}

	return nil
}

func (l *Loan) Return(userID, bookID int) error {
	u, err := l.user.GetByIDUser(userID)
	if err != nil {
		if err == entity.ErrNotFound {
			return fmt.Errorf("user %w", entity.ErrNotFound)
		}
		return err
	}

	b, err := l.book.GetByIDBook(bookID)
	if err != nil {
		if err == entity.ErrNotFound {
			return fmt.Errorf("book %w", entity.ErrNotFound)
		}
		return err
	}

	err = u.RemoveBook(bookID)
	if err != nil {
		return err
	}

	err = l.user.UpdateUser(u)
	if err != nil {
		return err
	}

	b.Quantity++
	err = l.book.UpdateBook(b)
	if err != nil {
		return err
	}

	return nil
}
