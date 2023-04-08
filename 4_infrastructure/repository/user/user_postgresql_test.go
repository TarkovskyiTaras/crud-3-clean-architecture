package repositoryUser

import (
	"database/sql"
	entity "github.com/TarasTarkovskyi/crud-3-clean-architecture/1_entity"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/5_pkg/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var db *sql.DB

type userTest struct {
	args userArgs
	want userWant
}
type userArgs struct {
	user *entity.User
}
type userWant struct {
	user  *entity.User
	users []*entity.User
	err   error
}

func setUp() {
	var err error
	db, err = database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	if err != nil {
		log.Fatal(err)
	}
	var initialUser = &entity.User{ID: 1, FirstName: "Taras", LastName: "Tarkovskyi", DOB: time.Date(1992, 01, 23, 0, 0, 0, 0, time.UTC), Location: "Ukraine", CellPhoneNumber: "0933115485", Email: "taras6317492@gmail.com", Password: "12345qwerty", Books: []int{1, 2, 3}}

	userRepo := NewUsers(db)
	_, err = userRepo.db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
	_, err = userRepo.db.Exec("INSERT INTO users (id, first_name, last_name, dob, location, cellphone_number, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		initialUser.ID, initialUser.FirstName, initialUser.LastName, initialUser.DOB, initialUser.Location, initialUser.CellPhoneNumber, initialUser.Email, initialUser.Password, time.Time{}, time.Time{})
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	defer db.Close()
	userRepo := NewUsers(db)

	_, err := userRepo.db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
	tearDown()
}

func TestCreateUser(t *testing.T) {
	userRepo := NewUsers(db)
	userArg1 := &entity.User{ID: 2, FirstName: "Sergey", LastName: "Onishenko", DOB: time.Date(1990, 12, 28, 0, 0, 0, 0, time.UTC), Location: "Ukraine", CellPhoneNumber: "0935554422", Email: "sergeypoc@gmail.com", Password: "12345qwerty"}
	tests := []userTest{
		{args: userArgs{user: userArg1}, want: userWant{user: userArg1, err: nil}},
	}

	for _, ut := range tests {
		errGot := userRepo.Create(ut.args.user)
		userGot, err := userRepo.GetByID(ut.args.user.ID)
		if err != nil {
			log.Fatal(err)
		}

		userGot.DOB = userGot.DOB.UTC()
		userGot.CreatedAt = userGot.CreatedAt.UTC()
		userGot.UpdatedAt = userGot.UpdatedAt.UTC()

		assert.Equal(t, ut.want.user, userGot)
		assert.Equal(t, ut.want.err, errGot)
	}
}

func TestGetByIDUser(t *testing.T) {
	userRepo := NewUsers(db)
	userArg1 := &entity.User{ID: 1}
	tests := []userTest{
		{args: userArgs{user: userArg1}, want: userWant{user: userArg1, err: nil}},
	}

	for _, ut := range tests {
		userGot, errGot := userRepo.GetByID(ut.args.user.ID)

		assert.Equal(t, ut.want.user.ID, userGot.ID)
		assert.Equal(t, ut.want.err, errGot)
	}
}

func TestGetAllUsers(t *testing.T) {
	userRepo := NewUsers(db)

	tests := []userTest{
		{want: userWant{err: nil}},
	}

	for _, ut := range tests {
		usersGot, errGot := userRepo.GetAll()

		assert.NotNil(t, usersGot)
		assert.Equal(t, ut.want.err, errGot)
	}
}

func TestUpdateUser(t *testing.T) {
	userRepo := NewUsers(db)
	userArg1 := &entity.User{ID: 1, FirstName: "UPD_Taras", LastName: "UPD_Tarkovskyi", DOB: time.Date(1992, 01, 23, 0, 0, 0, 0, time.UTC), Location: "Ukraine", CellPhoneNumber: "0933115485", Email: "taras6317492@gmail.com", Password: "12345qwerty", Books: []int{4, 5, 6}}
	tests := []userTest{
		{args: userArgs{user: userArg1}, want: userWant{user: userArg1, err: nil}},
	}

	for _, ut := range tests {
		errGot := userRepo.Update(ut.args.user)
		userGot, err := userRepo.GetByID(ut.args.user.ID)
		if err != nil {
			log.Fatal(err)
		}

		userGot.DOB = userGot.DOB.UTC()
		userGot.CreatedAt = userGot.CreatedAt.UTC()
		userGot.UpdatedAt = userGot.UpdatedAt.UTC()

		assert.Equal(t, ut.want.user, userGot)
		assert.Equal(t, ut.want.err, errGot)
	}
}

func TestDeleteUser(t *testing.T) {
	userRepo := NewUsers(db)
	userArg1 := &entity.User{ID: 1}
	tests := []userTest{
		{args: userArgs{user: userArg1}, want: userWant{err: nil}},
	}

	for _, ut := range tests {
		errGot := userRepo.Delete(ut.args.user.ID)
		userGot, err := userRepo.GetByID(ut.args.user.ID)

		assert.NotNil(t, err)
		assert.Nil(t, userGot)
		assert.Equal(t, ut.want.err, errGot)
	}
}
