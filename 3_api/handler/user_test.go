package handler

import (
	"encoding/json"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/custom_mocks"
	mock_user "github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

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

	m := mock_user.NewMockUseCase(controller)
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
		t.Errorf("user ID wanted = %d, got = %d", u1.ID, u2.ID)
	}
	if writer.Code != http.StatusOK {
		t.Errorf("http status wanted = %v, got = %v", http.StatusOK, writer.Code)
	}
}

func TestGetByIDHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	m := mock_user.NewMockUseCase(controller)
	h := NewUserHandler(m)
	r := mux.NewRouter()
	h.MakeUserHandler(r)

	expUser := &entity.User{ID: 1}
	m.EXPECT().GetByIDUser(expUser.ID).Return(expUser, nil)

	testServ := httptest.NewServer(r)
	defer testServ.Close()
	resp, _ := http.Get(testServ.URL + "/user/" + strconv.Itoa(expUser.ID))

	var resUser entity.User
	respBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(respBody, &resUser)

	assert.NotNil(t, resUser)
	assert.Equal(t, expUser.ID, resUser.ID)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

}
