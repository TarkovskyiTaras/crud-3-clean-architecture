package handler

import (
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
	userID, _ := strconv.Atoi(vars["u_id"])
	bookID, _ := strconv.Atoi(vars["b_id"])

	l.LoanUseCase.Borrow(userID, bookID)

	w.WriteHeader(http.StatusOK)
}

func (l *LoanHandler) ReturnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["u_id"])
	bookID, _ := strconv.Atoi(vars["b_id"])

	l.LoanUseCase.Return(userID, bookID)

	w.WriteHeader(http.StatusOK)
}

func (l *LoanHandler) MakeLoanHandler(r *mux.Router) {
	r.HandleFunc("/loan/borrow/{u_id:[0-9]+}/{b_id:[0-9]+}", l.BorrowHandler)
	r.HandleFunc("/loan/return/{u_id:[0-9]+}/{b_id:[0-9]+}", l.ReturnHandler)
}
