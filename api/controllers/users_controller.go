package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/crud_api/api/responses"
	"github.com/gorilla/mux"

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

	_, err = user.SaveUser(server.DB)

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

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	value, _ := strconv.ParseUint(id, 10, 64)

	user := models.User{}
	UserDetail, err := user.FindUserByID(server.DB, value)
	if err != nil {
		panic(err)
	}
	responses.JSON(w, http.StatusOK, UserDetail)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["id"]
	id, err := strconv.ParseUint(uid, 10, 64)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// responses.JSON(w, http.StatusUnprocessableEntity, err)
		panic(err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		// responses.ERROR
		panic(err)
	}

	// user.Preapare()
	user.Preapare()
	err = user.Validate("update")
	if err != nil {
		panic(err)
	}
	updatedUser, err := user.UpdateAUser(server.DB, id)
	if err != nil {
		panic(err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	pid, _ := strconv.ParseUint(id, 10, 64)

	user := models.User{}
	_, err := user.DeleteAUser(server.DB, pid)
	if err != nil {
		panic(err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
