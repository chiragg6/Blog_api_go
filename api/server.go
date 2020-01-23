package api

import (
	"github.com/crud_api/api/controllers"
	"github.com/crud_api/api/seed"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "aicumendeveloper"
	Dbname   = "postgres"
	Password = "dev"
	Sslmode  = "disable"
)

var server = controllers.Server{}

func Run() {

	var err error

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env,not comming through %v", err)

	// } else {
	// 	fmt.Println("We are getting the env values")

	// }
	server.Initalize(Host, User, Dbname, Password, Sslmode, Port)

	seed.Load(server.DB)
	server.Run(":8080")

}
