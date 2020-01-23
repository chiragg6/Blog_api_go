package controllers

import (
	"github.com/crud_api/api/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/crud_api/models"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
		return
	}

	user.Preapare() 
	err = user.Validate("")
	if err != nil {
		panic(err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		panic(err)
		return
	}

	// w.Header().Set()	
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.FindAllUser(server.DB)
	if err != nil {
		panic(err)
	}

	responses.JSON(w, http.StatusOK, users)


}
