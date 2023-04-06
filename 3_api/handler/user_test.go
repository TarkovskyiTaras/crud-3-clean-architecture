package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/custom_mocks"
	umock "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

type wantUser struct {
	err        error
	statusCode int
}

type userTest struct {
	id   int
	user string
	want wantUser
}

func TestCreateUserHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	u1 := `{"id":1,"first_name":"Peter","last_name":"Anderson","dob":"1990-01-15T00:00:00Z","location":"Canada","cellphone_number":"+16479150167","email":"Peter@gmail.com","password":"qwerty12345"}`
	u2 := `{"id":3,"first_name":"","last_name":"","dob":"0001-01-01T00:00:00Z","location":"","cellphone_number":"","email":"","password":""}`
	u3 := `{"id":1,"first_name":"Peter","last_name":"Anderson","dob":"1990-01-15T00:00:00Z","location":"Canada","cellphone_number":"+16479150167","email":"Peter@gmail.com","password":"qwerty12345"}`

	tests := []userTest{
		{user: u1, want: wantUser{err: nil, statusCode: http.StatusCreated}},
		{user: u2, want: wantUser{err: entity.ErrInvalidEntity, statusCode: http.StatusUnprocessableEntity}},
		{user: u3, want: wantUser{err: entity.ErrConflict, statusCode: http.StatusConflict}},
	}

	for _, ut := range tests {
		var userUnmarshalled entity.User
		err := json.Unmarshal([]byte(ut.user), &userUnmarshalled)
		assert.NoError(t, err)

		m.EXPECT().CreateUser(&userUnmarshalled).Return(ut.want.err)

		resp, err := http.Post(testServ.URL+"/user", "application/json", strings.NewReader(ut.user))
		assert.NoError(t, err)

		assert.Equal(t, ut.want.statusCode, resp.StatusCode)
	}

}

func TestGetByIDUserHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	expUser := &entity.User{ID: 1}
	m.EXPECT().GetByIDUser(expUser.ID).Return(expUser, nil)

	testServ := httptest.NewServer(r)
	defer testServ.Close()
	resp, err := http.Get(testServ.URL + "/user/" + strconv.Itoa(expUser.ID))
	assert.NoError(t, err)

	var resUser entity.User
	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, &resUser)
	assert.NoError(t, err)

	assert.NotNil(t, resUser)
	assert.Equal(t, expUser.ID, resUser.ID)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestGetAllUsersHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	users := []*entity.User{{ID: 1}, {ID: 2}, {ID: 3}}

	m.EXPECT().GetAllUsers().Return(users, nil)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	resp, _ := http.Get(testServ.URL + "/user")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateByIDUserHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	u1 := `{"id":1,"first_name":"Jonathan","last_name":"Adams","dob":"1987-03-21T00:00:00Z","location":"USA","cellphone_number":"+16479250145","email":"Jonathan@gmail.com","password":"pw124567"}`
	u2 := `{"id":2,"first_name":"Peter","last_name":"Anderson","dob":"1990-01-15T00:00:00Z","location":"Canada","cellphone_number":"+16479150167","email":"Peter@gmail.com","password":"qwerty12345"}`
	u3 := `{"id":3,"first_name":"","last_name":"","dob":"0001-01-01T00:00:00Z","location":"","cellphone_number":"","email":"","password":""}`

	tests := []userTest{
		{user: u1, want: wantUser{err: nil, statusCode: http.StatusOK}},
		{user: u2, want: wantUser{err: entity.ErrNotFound, statusCode: http.StatusNotFound}},
		{user: u3, want: wantUser{err: entity.ErrInvalidEntity, statusCode: http.StatusUnprocessableEntity}},
	}

	for _, ut := range tests {
		var userUnmarshalled entity.User
		err := json.Unmarshal([]byte(ut.user), &userUnmarshalled)
		assert.NoError(t, err)

		m.EXPECT().UpdateUser(&userUnmarshalled).Return(ut.want.err)

		req, err := http.NewRequest(http.MethodPut, testServ.URL+"/user", strings.NewReader(ut.user))
		assert.NoError(t, err)
		client := http.DefaultClient
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, ut.want.statusCode, resp.StatusCode)
	}

}

func TestDeleteByIDUserHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	tests := []userTest{
		{id: 1, want: wantUser{err: nil, statusCode: http.StatusOK}},
		{id: 2, want: wantUser{err: entity.ErrNotFound, statusCode: http.StatusNotFound}},
		{id: 3, want: wantUser{err: errors.New("some internal server error"), statusCode: http.StatusInternalServerError}},
	}

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	for _, ut := range tests {
		m.EXPECT().DeleteUser(ut.id).Return(ut.want.err)

		req, err := http.NewRequest(http.MethodDelete, testServ.URL+"/user/"+strconv.Itoa(ut.id), nil)
		assert.NoError(t, err)
		client := http.DefaultClient
		resp, err := client.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, ut.want.statusCode, resp.StatusCode)
	}

}

func TestGetByIDHandler1(t *testing.T) {

	pathes := []string{"/user/1", "/user/3", "/user/999"}

	fu := custom_mocks.NewFakeUser()
	fh := NewUserHandler(fu)
	r := mux.NewRouter()
	fh.MakeUserHandler(r)

	for i, p := range pathes {
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", p, nil)
		r.ServeHTTP(writer, request)

		switch i {
		case 0:
			var u entity.User
			json.Unmarshal(writer.Body.Bytes(), &u)
			if u.ID != 1 || writer.Code != http.StatusOK || writer.Header().Get("Content-Type") != "application/json" {
				t.Error("error case 1")
			}
		case 1:
			if writer.Body.String() != "not found" || writer.Code != http.StatusNotFound {
				t.Error("error case 2")
			}
		case 3:
			if writer.Body.String() != "internal server error" || writer.Code != http.StatusInternalServerError {
				t.Error("error case 3")
			}
		}
	}
}

func TestGetByIDHandler2(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	u1 := &entity.User{ID: 1}
	m.EXPECT().GetByIDUser(u1.ID).Return(u1, nil)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user/1", nil)
	r.ServeHTTP(writer, request)

	var u2 entity.User
	json.Unmarshal(writer.Body.Bytes(), &u2)

	if u2.ID != u1.ID {
		t.Errorf("userS ID wanted = %d, got = %d", u1.ID, u2.ID)
	}
	if writer.Code != http.StatusOK {
		t.Errorf("http statusCode wanted = %v, got = %v", http.StatusOK, writer.Code)
	}
}

func TestCreateHandler1(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := umock.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	testServ := httptest.NewServer(r)
	defer testServ.Close()

	u1 := &entity.User{ID: 1, FirstName: "Peter", LastName: "Anderson", DOB: time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC), Location: "Canada", CellPhoneNumber: "+16479150167", Email: "Peter@gmail.com", Password: "qwerty12345"}
	u2 := &entity.User{}
	u3 := &entity.User{ID: 1, FirstName: "Peter", LastName: "Anderson", DOB: time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC), Location: "Canada", CellPhoneNumber: "+16479150167", Email: "Peter@gmail.com", Password: "qwerty12345"}

	m.EXPECT().CreateUser(u1).Return(nil)
	m.EXPECT().CreateUser(u2).Return(entity.ErrInvalidEntity)
	m.EXPECT().CreateUser(u3).Return(entity.ErrConflict)

	u1Json, err := json.Marshal(u1)
	assert.NoError(t, err)
	reqBody1 := bytes.NewReader(u1Json)
	resp1, err := http.Post(testServ.URL+"/user", "application/json", reqBody1)
	assert.NoError(t, err)

	u2Json, err := json.Marshal(u2)
	assert.NoError(t, err)
	reqBody2 := bytes.NewReader(u2Json)
	resp2, err := http.Post(testServ.URL+"/user", "application/json", reqBody2)
	assert.NoError(t, err)

	u3Json, err := json.Marshal(u3)
	assert.NoError(t, err)
	reqBody3 := bytes.NewReader(u3Json)
	resp3, err := http.Post(testServ.URL+"/user", "application/json", reqBody3)
	assert.NoError(t, err)

	assert.Equal(t, resp1.StatusCode, http.StatusCreated)
	assert.Equal(t, resp2.StatusCode, http.StatusUnprocessableEntity)
	assert.Equal(t, resp3.StatusCode, http.StatusConflict)
}
