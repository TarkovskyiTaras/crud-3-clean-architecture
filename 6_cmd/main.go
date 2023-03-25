package main

import (
	"fmt"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/2_usecase/user"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/3_api/handler"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/4_infrastructure/repository"
	"github.com/TarasTarkovskyi/crud-3-clean-architecture/5_pkg/database"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	userRepo := repository.NewUsers(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewHandler(userService)

	server := http.Server{
		Addr:    ":8080",
		Handler: userHandler.InitRouter(),
	}

	server.ListenAndServe()
}
