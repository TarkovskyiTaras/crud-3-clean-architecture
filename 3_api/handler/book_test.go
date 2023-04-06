package handler

import (
	"encoding/json"
	"errors"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	bmock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type bookTest struct {
	id   string
	book string
	want wantBook
}

type wantBook struct {
	err        error
	statusCode int
	book       entity.Book
	books      []*entity.Book
}

func TestCreateHandler_Book_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	payload := `{"ID":1,"Tittle":"God's Little Acre","Author":"Erskine Caldwell","Pages":224,"Quantity":5,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`

	tests := []bookTest{
		{book: payload, want: wantBook{err: nil, statusCode: http.StatusCreated}},
	}

	for _, bt := range tests {
		m.EXPECT().CreateBook(gomock.Any()).Return(bt.want.err)
		resp, err := http.Post(testServ.URL+"/book", "application/json", strings.NewReader(bt.book))
		assert.NoError(t, err)
		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}
}

func TestCreateHandler_Book_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	//testing a failure before the usecase.book.CreateBook call
	tests := []bookTest{
		{book: "making unmarshalling fail", want: wantBook{statusCode: http.StatusBadRequest}},
	}
	for _, bt := range tests {
		resp, _ := http.Post(testServ.URL+"/book", "application/json", strings.NewReader(bt.book))
		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}

	payload := `{"ID":1,"Tittle":"God's Little Acre","Author":"Erskine Caldwell","Pages":224,"Quantity":5,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`
	tests = []bookTest{
		{book: payload, want: wantBook{err: entity.ErrConflict, statusCode: http.StatusConflict}},
		{book: payload, want: wantBook{err: entity.ErrInvalidEntity, statusCode: http.StatusUnprocessableEntity}},
		{book: payload, want: wantBook{err: errors.New("some internal server error"), statusCode: http.StatusInternalServerError}},
	}
	for _, bt := range tests {
		m.EXPECT().CreateBook(gomock.Any()).Return(bt.want.err)
		resp, err := http.Post(testServ.URL+"/book", "application/json", strings.NewReader(bt.book))
		assert.NoError(t, err)
		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}
}

func TestGetByIDHandler_Book_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []bookTest{
		{id: "1", want: wantBook{err: nil, statusCode: http.StatusOK, book: entity.Book{ID: 1}}},
	}

	for _, bt := range tests {
		idInt, _ := strconv.Atoi(bt.id)
		m.EXPECT().GetByIDBook(idInt).Return(&bt.want.book, bt.want.err)
		resp, err := http.Get(testServ.URL + "/book/" + bt.id)
		assert.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()
		var bookGot entity.Book
		err = json.Unmarshal(respBody, &bookGot)
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
		assert.Equal(t, bt.want.book, bookGot)
	}
}

func TestGetByIDHandler_Book_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []bookTest{
		{id: "1", want: wantBook{err: errors.New("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	for _, bt := range tests {
		idInt, _ := strconv.Atoi(bt.id)
		m.EXPECT().GetByIDBook(idInt).Return(&bt.want.book, bt.want.err)
		resp, err := http.Get(testServ.URL + "/book/" + bt.id)
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}
}

func TestGetAllHandler_Book_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []bookTest{
		{want: wantBook{books: []*entity.Book{{ID: 1}, {ID: 2}}, err: nil, statusCode: http.StatusOK}},
	}

	for _, bt := range tests {
		m.EXPECT().GetAllBooks().Return(bt.want.books, bt.want.err)
		resp, err := http.Get(testServ.URL + "/book")
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}
}

func TestGetAllHandler_Book_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []bookTest{
		{want: wantBook{err: errors.New("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	for _, bt := range tests {
		m.EXPECT().GetAllBooks().Return(bt.want.books, bt.want.err)
		resp, err := http.Get(testServ.URL + "/book")
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}

}

func TestDeleteByIDHandler_Book_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []bookTest{
		{id: "1", want: wantBook{err: nil, statusCode: http.StatusOK}},
	}

	for _, bt := range tests {
		idInt, err := strconv.Atoi(bt.id)
		assert.NoError(t, err)

		m.EXPECT().DeleteBook(idInt).Return(bt.want.err)

		req, err := http.NewRequest(http.MethodDelete, testServ.URL+"/book/"+bt.id, nil)
		assert.NoError(t, err)
		client := http.DefaultClient
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}

}

func TestDeleteByIDHandler_Book_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []bookTest{
		{id: "1", want: wantBook{err: entity.ErrNotFound, statusCode: http.StatusNotFound}},
		{id: "2", want: wantBook{err: errors.New("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	for _, bt := range tests {
		idInt, err := strconv.Atoi(bt.id)
		assert.NoError(t, err)

		m.EXPECT().DeleteBook(idInt).Return(bt.want.err)

		req, err := http.NewRequest(http.MethodDelete, testServ.URL+"/book/"+bt.id, nil)
		assert.NoError(t, err)
		client := http.DefaultClient
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}

}

func TestUpdateHandler_Book_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	payload := `{"ID":1,"Tittle":"God's Little Acre","Author":"Erskine Caldwell","Pages":224,"Quantity":5,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`
	tests := []bookTest{
		{book: payload, want: wantBook{err: nil, statusCode: http.StatusOK}},
	}

	for _, bt := range tests {
		m.EXPECT().UpdateBook(gomock.Any()).Return(bt.want.err)

		req, err := http.NewRequest(http.MethodPut, testServ.URL+"/book", strings.NewReader(bt.book))
		assert.NoError(t, err)
		client := http.DefaultClient
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}

}

func TestUpdateHandler_Book_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := bmock.NewMockUseCase(controller)
	h := NewBookHandler(m)
	r := mux.NewRouter()
	h.MakeBookHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	payload := `{"ID":1,"Tittle":"God's Little Acre","Author":"Erskine Caldwell","Pages":224,"Quantity":5,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`
	tests := []bookTest{
		{book: payload, want: wantBook{err: entity.ErrNotFound, statusCode: http.StatusNotFound}},
		{book: payload, want: wantBook{err: entity.ErrInvalidEntity, statusCode: http.StatusUnprocessableEntity}},
		{book: payload, want: wantBook{err: errors.New("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	for _, bt := range tests {
		m.EXPECT().UpdateBook(gomock.Any()).Return(bt.want.err)

		req, err := http.NewRequest(http.MethodPut, testServ.URL+"/book", strings.NewReader(bt.book))
		assert.NoError(t, err)
		client := http.DefaultClient
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, bt.want.statusCode, resp.StatusCode)
	}
}
