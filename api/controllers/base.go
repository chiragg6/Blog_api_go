package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crud_api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initalize() {
	var err error
	// if Dbdriver == "postgres" {
	// 	DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", DbHost, Port, DbUser, DbName, DbPassword)
	// server.DB, err = gorm.Open(Dbdriver, DBURL)
	// 	if err != nil {
	// 		fmt.Printf("Cannot connect to %s database", Dbdriver)
	// 		log.Fatal("This is the error:", err)
	// 	} else {
	// 		fmt.Printf("We are connected to the %s database", Dbdriver)
	// 	}
	// }

	server.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=aicumendeveloper dbname=postgres password=dev sslmode=disable")
	if err != nil {
		fmt.Printf("Cannot connect to databaase")
	}

	err = server.DB.DB().Ping()
	if err != nil {
		panic(err)
		fmt.Println("Not connected with DB")
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	//Database Migration
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to Port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
