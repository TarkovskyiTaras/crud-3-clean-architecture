package user

import entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"

type Repository interface {
	Create(user *entity.User) error
	GetByID(id int) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Update(e *entity.User) error
	Delete(id int) error
}

type UseCase interface {
	CreateUser(user *entity.User) error
	GetByIDUser(id int) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id int) error
}
