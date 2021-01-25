package main

import (
	routes "CRUD-Operation/routes"
)

func main() {
	router := routes.StartGin()
	router.Run("qbtl.fr:8081")
}
