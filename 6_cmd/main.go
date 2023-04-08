package main

import (
	"fmt"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/book"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/loan"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/3_api/handler"
	repositoryBook "github.com/TarasTarkovskyi/crud-3-clean-architecture/4_infrastructure/repository/book"
	repositoryUser "github.com/TarasTarkovskyi/crud-3-clean-architecture/4_infrastructure/repository/user"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/5_pkg/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	userRepo := repositoryUser.NewUsers(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	bookRepo := repositoryBook.NewBooks(db)
	bookService := book.NewService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	loanService := loan.NewLoan(userService, bookService)
	loanHandler := handler.NewLoanHandler(loanService)

	r := mux.NewRouter()
	userHandler.MakeUserHandler(r)
	bookHandler.MakeBookHandler(r)
	loanHandler.MakeLoanHandler(r)

	serv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
