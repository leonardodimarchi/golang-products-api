package main

import (
	"golang-products-api/database"
	"golang-products-api/server"
)

func main() {
	database.StartDatabase()

	server := server.CreateServer()
	server.Run()
}