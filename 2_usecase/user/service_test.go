package user

import (
	"errors"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	umock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type userTest struct {
	user *entity.User
	want wantUser
	t    timesToCall
}
type wantUser struct {
	user          *entity.User
	users         []*entity.User
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

func TestCreateUser_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1, FirstName: "Taras", LastName: "Tarkovskyi", DOB: time.Date(1992, 01, 23, 0, 0, 0, 0, time.UTC), Location: "Kyiv", CellPhoneNumber: "0933115485", Email: "taras6317492@gmail.com", Password: "12345"}

	tests := []userTest{
		{user: u1, want: wantUser{user: nil, errFromGet: entity.ErrNotFound, errFromCreate: nil, errFinal: nil}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		m.EXPECT().Create(ut.user).Return(ut.want.errFromCreate)

		errGot := u.CreateUser(ut.user)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestCreateUser_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1, FirstName: "Taras", LastName: "Tarkovskyi", DOB: time.Date(1992, 01, 23, 0, 0, 0, 0, time.UTC), Location: "Kyiv", CellPhoneNumber: "0933115485", Email: "taras6317492@gmail.com", Password: "12345"}
	u2 := &entity.User{ID: 2, FirstName: "Sergey", LastName: "Onishenko", DOB: time.Date(1990, 12, 28, 0, 0, 0, 0, time.UTC), Location: "Kyiv", CellPhoneNumber: "", Email: "", Password: ""}

	tests := []userTest{
		{user: u1, want: wantUser{user: u1, errFromGet: nil, errFromCreate: nil, errFinal: entity.ErrConflict}},
		{user: u2, want: wantUser{user: nil, errFromGet: entity.ErrNotFound, errFromCreate: nil, errFinal: entity.ErrInvalidEntity}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		m.EXPECT().Create(ut.user).Return(ut.want.errFromCreate).AnyTimes()

		errGot := u.CreateUser(ut.user)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestGetByIDUser_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1}

	tests := []userTest{
		{user: u1, want: wantUser{user: u1, errFromGet: nil, errFinal: nil}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		userGot, errGot := u.GetByIDUser(ut.user.ID)

		assert.Equal(t, ut.want.user, userGot)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestGetByIDUser_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1}

	someDBError := errors.New("some database error")
	tests := []userTest{
		{user: u1, want: wantUser{user: nil, errFromGet: someDBError, errFinal: someDBError}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		userGot, errGot := u.GetByIDUser(ut.user.ID)

		assert.Equal(t, ut.want.user, userGot)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestGetAllUsers_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	users := []*entity.User{{ID: 1}, {ID: 2}, {ID: 3}}

	tests := []userTest{
		{want: wantUser{users: users, errFinal: nil}},
	}

	for _, ut := range tests {
		m.EXPECT().GetAll().Return(ut.want.users, ut.want.errFinal)
		usersGot, errGot := u.GetAllUsers()

		assert.Equal(t, ut.want.users, usersGot)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestGetAllUsers_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	tests := []userTest{
		{want: wantUser{users: nil, errFinal: errors.New("some database error")}},
	}

	for _, ut := range tests {
		m.EXPECT().GetAll().Return(ut.want.users, ut.want.errFinal)
		usersGot, errGot := u.GetAllUsers()

		assert.Equal(t, ut.want.users, usersGot)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1, FirstName: "Taras", LastName: "Tarkovskyi", DOB: time.Date(1992, 01, 23, 0, 0, 0, 0, time.UTC), Location: "Kyiv", CellPhoneNumber: "0933115485", Email: "taras6317492@gmail.com", Password: "12345"}

	tests := []userTest{
		{user: u1, want: wantUser{user: u1, errFromGet: nil, errFromUpdate: nil, errFinal: nil}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		m.EXPECT().Update(ut.user).Return(ut.want.errFromUpdate)

		errGot := u.UpdateUser(ut.user)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestUpdateUser_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1, FirstName: "Taras", LastName: "Tarkovskyi", DOB: time.Date(1992, 01, 23, 0, 0, 0, 0, time.UTC), Location: "Kyiv", CellPhoneNumber: "0933115485", Email: "taras6317492@gmail.com", Password: "12345"}
	u2 := &entity.User{ID: 2, FirstName: "Sergey", LastName: "Onishenko", DOB: time.Date(1990, 12, 28, 0, 0, 0, 0, time.UTC), Location: "Kyiv", CellPhoneNumber: "", Email: "", Password: ""}

	tests := []userTest{
		{user: u1, want: wantUser{user: nil, errFromGet: entity.ErrNotFound, errFinal: entity.ErrNotFound}, t: timesToCall{ttcUpdate: 0}},
		{user: u2, want: wantUser{user: u2, errFromGet: nil, errFinal: entity.ErrInvalidEntity}, t: timesToCall{ttcUpdate: 0}},
		{user: u1, want: wantUser{user: u1, errFromGet: nil, errFromUpdate: errors.New("some database error"), errFinal: errors.New("some database error")}, t: timesToCall{ttcUpdate: 1}},
	}

	for _, ut := range tests {

		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		m.EXPECT().Update(ut.user).Return(ut.want.errFromUpdate).Times(ut.t.ttcUpdate)

		errGot := u.UpdateUser(ut.user)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1}

	tests := []userTest{
		{user: u1, want: wantUser{user: u1, errFromGet: nil, errFromDelete: nil, errFinal: nil}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		m.EXPECT().Delete(ut.user.ID).Return(ut.want.errFromDelete)

		errGot := u.DeleteUser(ut.user.ID)
		assert.Equal(t, ut.want.errFinal, errGot)
	}

}

func TestDeleteUser_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockRepository(controller)
	u := NewService(m)

	u1 := &entity.User{ID: 1}

	tests := []userTest{
		{user: u1, want: wantUser{user: nil, errFromGet: entity.ErrNotFound, errFinal: entity.ErrNotFound}, t: timesToCall{ttcDelete: 0}},
		{user: u1, want: wantUser{user: u1, errFromGet: nil, errFromDelete: errors.New("some database error"), errFinal: errors.New("some database error")}, t: timesToCall{ttcDelete: 1}},
	}

	for _, ut := range tests {
		m.EXPECT().GetByID(ut.user.ID).Return(ut.want.user, ut.want.errFromGet)
		m.EXPECT().Delete(ut.user.ID).Return(ut.want.errFromDelete).Times(ut.t.ttcDelete)

		errGot := u.DeleteUser(ut.user.ID)
		assert.Equal(t, ut.want.errFinal, errGot)
	}
}
