package main

import "golang-products-api/server"

func main() {
	server := server.CreateServer()
	server.Run()
}