package user

import (
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"time"
)

type Users struct {
	repo Repository
}

func NewService(p Repository) *Users {
	return &Users{p}
}

func (u *Users) CreateUser(e *entity.User) error {
	_, err := u.repo.GetByID(e.ID)
	if err != entity.ErrUserNotFound {
		return entity.ErrConflict
	}

	err = ValidateInput(e)
	if err != nil {
		return err
	}

	e.CreatedAt = time.Now()
	return u.repo.Create(e)
}

func (u *Users) GetByIDUser(id int) (*entity.User, error) {
	return u.repo.GetByID(id)
}

func (u *Users) GetAllUsers() ([]*entity.User, error) {
	return u.repo.GetAll()
}

func (u *Users) UpdateUser(e *entity.User) error {
	_, err := u.repo.GetByID(e.ID)
	if err != nil {
		return err
	}

	err = ValidateInput(e)
	if err != nil {
		return err
	}

	e.UpdatedAt = time.Now()
	return u.repo.Update(e)
}

func (u *Users) DeleteUser(id int) error {
	_, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	return u.repo.Delete(id)
}

func ValidateInput(user *entity.User) error {
	if user.ID <= 0 || user.FirstName == "" || user.LastName == "" || user.DOB.IsZero() || user.Location == "" || user.CellPhoneNumber == "" || user.Email == "" || user.Password == "" {
		return entity.ErrInvalidEntity
	}
	return nil
}
