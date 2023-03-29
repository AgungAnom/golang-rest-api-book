package main

import (
	"golang-rest-api-book/routers"
)

const PORT = ":4000"
func main(){
	routers.StartServer().Run(PORT)
}
