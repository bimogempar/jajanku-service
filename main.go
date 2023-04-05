package main

import (
	"log"
	"os"

	"github.com/maulanarisqimustofa/jajanku-project/app/config"
	"github.com/maulanarisqimustofa/jajanku-project/modules/routes"
)

func main() {
	db := config.InitDB()
	r := routes.InitRoute(db)
	log.Fatal(r.Listen(os.Getenv("APP_PORT")))
}
