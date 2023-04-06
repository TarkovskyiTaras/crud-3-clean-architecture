// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package bmock is a generated GoMock package.
package bmock

import (
	reflect "reflect"

	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(b *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), b)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll() ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(id int) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockRepository) Update(b *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), b)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockUseCase) CreateBook(b *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockUseCaseMockRecorder) CreateBook(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockUseCase)(nil).CreateBook), b)
}

// DeleteBook mocks base method.
func (m *MockUseCase) DeleteBook(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockUseCaseMockRecorder) DeleteBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockUseCase)(nil).DeleteBook), id)
}

// GetAllBooks mocks base method.
func (m *MockUseCase) GetAllBooks() ([]*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks")
	ret0, _ := ret[0].([]*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockUseCaseMockRecorder) GetAllBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockUseCase)(nil).GetAllBooks))
}

// GetByIDBook mocks base method.
func (m *MockUseCase) GetByIDBook(id int) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDBook", id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDBook indicates an expected call of GetByIDBook.
func (mr *MockUseCaseMockRecorder) GetByIDBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDBook", reflect.TypeOf((*MockUseCase)(nil).GetByIDBook), id)
}

// UpdateBook mocks base method.
func (m *MockUseCase) UpdateBook(b *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockUseCaseMockRecorder) UpdateBook(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockUseCase)(nil).UpdateBook), b)
}