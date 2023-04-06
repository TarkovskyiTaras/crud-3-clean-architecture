package handler

import (
	"fmt"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	lmock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/loan/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type loanTest struct {
	uID  string
	bID  string
	want wantLoan
}
type wantLoan struct {
	err        error
	statusCode int
}

func TestBorrowHandler_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := lmock.NewMockUseCase(controller)
	h := NewLoanHandler(m)
	r := mux.NewRouter()
	h.MakeLoanHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []loanTest{
		{uID: "1", bID: "1", want: wantLoan{err: nil, statusCode: http.StatusOK}},
	}

	for _, lt := range tests {
		uIDInt, err := strconv.Atoi(lt.uID)
		assert.NoError(t, err)
		bIDInt, err := strconv.Atoi(lt.bID)
		assert.NoError(t, err)

		m.EXPECT().Borrow(uIDInt, bIDInt).Return(lt.want.err)
		resp, err := http.Get(fmt.Sprintf("%s/loan/borrow/%s/%s", testServ.URL, lt.uID, lt.bID))
		assert.NoError(t, err)

		assert.Equal(t, lt.want.statusCode, resp.StatusCode)
	}
}

func TestBorrowHandler_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := lmock.NewMockUseCase(controller)
	h := NewLoanHandler(m)
	r := mux.NewRouter()
	h.MakeLoanHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []loanTest{
		{uID: "1", bID: "1", want: wantLoan{err: fmt.Errorf("user %w", entity.ErrNotFound), statusCode: http.StatusNotFound}},
		{uID: "2", bID: "2", want: wantLoan{err: fmt.Errorf("book %w", entity.ErrNotFound), statusCode: http.StatusNotFound}},
		{uID: "3", bID: "3", want: wantLoan{err: fmt.Errorf("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	for _, lt := range tests {
		uIDInt, err := strconv.Atoi(lt.uID)
		assert.NoError(t, err)
		bIDInt, err := strconv.Atoi(lt.bID)
		assert.NoError(t, err)

		m.EXPECT().Borrow(uIDInt, bIDInt).Return(lt.want.err)
		resp, err := http.Get(fmt.Sprintf("%s/loan/borrow/%s/%s", testServ.URL, lt.uID, lt.bID))
		assert.NoError(t, err)

		assert.Equal(t, lt.want.statusCode, resp.StatusCode)
	}
}

func TestReturnHandler_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := lmock.NewMockUseCase(controller)
	h := NewLoanHandler(m)
	r := mux.NewRouter()
	h.MakeLoanHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []loanTest{
		{uID: "1", bID: "1", want: wantLoan{err: nil, statusCode: http.StatusOK}},
	}

	for _, lt := range tests {
		uIDInt, err := strconv.Atoi(lt.uID)
		assert.NoError(t, err)
		bIDInt, err := strconv.Atoi(lt.bID)
		assert.NoError(t, err)

		m.EXPECT().Return(uIDInt, bIDInt).Return(lt.want.err)
		resp, err := http.Get(fmt.Sprintf("%s/loan/return/%s/%s", testServ.URL, lt.uID, lt.bID))
		assert.NoError(t, err)

		assert.Equal(t, lt.want.statusCode, resp.StatusCode)
	}

}

func TestReturnHandler_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := lmock.NewMockUseCase(controller)
	h := NewLoanHandler(m)
	r := mux.NewRouter()
	h.MakeLoanHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	tests := []loanTest{
		{uID: "1", bID: "1", want: wantLoan{err: fmt.Errorf("user %w", entity.ErrNotFound), statusCode: http.StatusNotFound}},
		{uID: "2", bID: "2", want: wantLoan{err: fmt.Errorf("book %w", entity.ErrNotFound), statusCode: http.StatusNotFound}},
		{uID: "3", bID: "3", want: wantLoan{err: fmt.Errorf("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	for _, lt := range tests {
		uIDInt, err := strconv.Atoi(lt.uID)
		assert.NoError(t, err)
		bIDInt, err := strconv.Atoi(lt.bID)
		assert.NoError(t, err)

		m.EXPECT().Return(uIDInt, bIDInt).Return(lt.want.err)
		resp, err := http.Get(fmt.Sprintf("%s/loan/return/%s/%s", testServ.URL, lt.uID, lt.bID))
		assert.NoError(t, err)

		assert.Equal(t, lt.want.statusCode, resp.StatusCode)
	}
}
