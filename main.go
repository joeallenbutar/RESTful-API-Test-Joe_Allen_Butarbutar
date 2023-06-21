package main

import (
	database "RESTful-API-Test-Joe_Allen_Butarbutar/configs"
	"RESTful-API-Test-Joe_Allen_Butarbutar/routers"
	"log"
)

func main() {
	database.Setup()
	r := routers.Setup()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}