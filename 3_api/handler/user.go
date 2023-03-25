package handler

import (
	"encoding/json"
	"github.com/TarasTarkovskyi/CRUD-6-template/1_entity"
	"github.com/TarasTarkovskyi/CRUD-6-template/2_usecase/user"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUsecase user.UseCase
}

func NewHandler(u user.UseCase) *UserHandler {
	return &UserHandler{userUsecase: u}
}

func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var u entity.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.userUsecase.CreateUser(&u)
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

func (h *UserHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	u, err := h.userUsecase.GetByIDUser(id)
	if err != nil {
		if err == entity.ErrUserNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userJson, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

func (h *UserHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	usersJson, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)
}

func (h *UserHandler) UpdateByIDHandler(w http.ResponseWriter, r *http.Request) {
	jsonUser, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var user entity.User
	err = json.Unmarshal(jsonUser, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.userUsecase.UpdateUser(&user)
	if err != nil {
		if err == entity.ErrUserNotFound {
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

func (h *UserHandler) DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = h.userUsecase.DeleteUser(id)
	if err != nil {
		if err == entity.ErrUserNotFound {
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

func (h *UserHandler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/user").Subrouter()

	s.HandleFunc("", h.CreateHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", h.GetByIDHandler).Methods(http.MethodGet)
	s.HandleFunc("/all", h.GetAllHandler).Methods(http.MethodGet)
	s.HandleFunc("/update", h.UpdateByIDHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", h.DeleteByIDHandler).Methods(http.MethodDelete)

	return s
}
