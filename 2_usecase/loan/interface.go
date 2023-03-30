package loan

type UseCase interface {
	Borrow(userID, bookID int) error
	Return(userID, bookID int) error
}
