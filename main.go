package main

import (
	"golang-rest-api-book/database"
	"golang-rest-api-book/routers"
)

const PORT = ":4000"
func main(){
	database.RunDB()
	routers.StartServer().Run(PORT)
}
