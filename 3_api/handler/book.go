package handler

import (
	"encoding/json"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookUseCase book.UseCase
}

func NewBookHandler(b book.UseCase) *BookHandler {
	return &BookHandler{bookUseCase: b}
}

func (h *BookHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var book entity.Book
	err = json.Unmarshal(reqBody, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.bookUseCase.CreateBook(&book)
	if err != nil {
		if err == entity.ErrConflict {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
			return
		}

		if err == entity.ErrInvalidEntity {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BookHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := h.bookUseCase.GetByIDBook(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bookJson, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bookJson)
}

func (h *BookHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookUseCase.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	booksJson, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(booksJson)
}

func (h *BookHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var book entity.Book
	err = json.Unmarshal(reqBody, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.bookUseCase.UpdateBook(&book)
	if err != nil {
		if err == entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err == entity.ErrInvalidEntity {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.bookUseCase.DeleteBook(id)
	if err != nil {
		if err == entity.ErrNotFound {
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

func (h *BookHandler) MakeBookHandler(r *mux.Router) {
	r.HandleFunc("/book", h.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/book/{id:[0-9]+}", h.GetByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/book", h.GetAllHandler).Methods(http.MethodGet)
	r.HandleFunc("/book", h.UpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/book/{id:[0-9]+}", h.DeleteHandler).Methods(http.MethodDelete)
}
