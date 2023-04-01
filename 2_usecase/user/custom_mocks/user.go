package custom_mocks

import (
	"errors"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
)

type FakeUser struct {
	something string
}

func NewFakeUser() *FakeUser {
	return &FakeUser{}
}

func (f *FakeUser) CreateUser(e *entity.User) error {
	return nil
}

func (f *FakeUser) GetByIDUser(id int) (*entity.User, error) {
	if id == 1 {
		return &entity.User{ID: 1}, nil
	}
	if id == 3 {
		return nil, entity.ErrNotFound
	}
	if id == 999 {
		return nil, errors.New("internal server error")
	}
	return nil, nil
}

func (f *FakeUser) GetAllUsers() ([]*entity.User, error) {
	return nil, nil
}

func (f *FakeUser) UpdateUser(e *entity.User) error {
	return nil
}

func (f *FakeUser) DeleteUser(id int) error {
	return nil
}
