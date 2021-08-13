package main

import (
	"github.com/rafamarquesrmb/webapi_books_crud_golang/database"
	"github.com/rafamarquesrmb/webapi_books_crud_golang/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
