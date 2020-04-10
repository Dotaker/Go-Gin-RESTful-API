package main

import (
	routes "CRUD-Operation/routes"
)

func main() {
	router := routes.StartGin()
	router.Run("127.0.0.1:8081")
}
