package main

import (
	"davisbento/golang-api/api/handlers"
	db "davisbento/golang-api/db"
	"davisbento/golang-api/db/repository"
	"fmt"
	"net/http"
)

func main() {
	newDB := db.NewDB()

	conn, err := newDB.Connect("../sqlite/test.db")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("server up")

	// instance the net/http server and attach the handlers
	userRepo := repository.NewUserRepository(conn)
	userHandler := handlers.NewUserHandler(userRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.FindAll)
	mux.HandleFunc("/user", userHandler.FindById)
	mux.HandleFunc("/user/create", userHandler.Create)

	// start the server
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}

}
