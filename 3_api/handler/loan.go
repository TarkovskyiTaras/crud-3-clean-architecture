package handler

import (
	"errors"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/loan"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type LoanHandler struct {
	LoanUseCase loan.UseCase
}

func NewLoanHandler(l loan.UseCase) *LoanHandler {
	return &LoanHandler{LoanUseCase: l}
}

func (l *LoanHandler) BorrowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["u_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	bookID, err := strconv.Atoi(vars["b_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = l.LoanUseCase.Borrow(userID, bookID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (l *LoanHandler) ReturnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["u_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	bookID, err := strconv.Atoi(vars["b_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = l.LoanUseCase.Return(userID, bookID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (l *LoanHandler) MakeLoanHandler(r *mux.Router) {
	r.HandleFunc("/loan/borrow/{u_id:[0-9]+}/{b_id:[0-9]+}", l.BorrowHandler).Methods(http.MethodGet)
	r.HandleFunc("/loan/return/{u_id:[0-9]+}/{b_id:[0-9]+}", l.ReturnHandler).Methods(http.MethodGet)
}
