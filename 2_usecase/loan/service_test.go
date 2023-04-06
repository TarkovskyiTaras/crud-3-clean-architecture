package loan

import (
	"fmt"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	bmock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book/mocks"
	umock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestBorrow_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m1 := umock.NewMockUseCase(controller)
	m2 := bmock.NewMockUseCase(controller)

	l := NewLoan(m1, m2)
	u1 := &entity.User{ID: 1}
	b1 := &entity.Book{ID: 3, Quantity: 5}
	u2 := &entity.User{ID: 3333}

	m1.EXPECT().GetByIDUser(u1.ID).Return(u1, nil)
	m2.EXPECT().GetByIDBook(b1.ID).Return(b1, nil)

	m1.EXPECT().UpdateUser(u1).Return(nil)
	m2.EXPECT().UpdateBook(b1).Return(nil)

	errGot := l.Borrow(1, 3)

	fmt.Println(errGot)

	fmt.Println("User1:", u1.ID, u1.Books)
	fmt.Println("Book:", b1.ID, b1.Quantity)
	fmt.Println("User2:", u2.ID, u2.Books)

}
