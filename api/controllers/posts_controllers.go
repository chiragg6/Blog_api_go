package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/crud_api/api/responses"

	"github.com/crud_api/models"
)

func (server *Server) createPost(w http.ResponseWriter, r *http.Request) {

	post := models.Post{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &post)
	if err != nil {
		panic(err)
	}
	post.Prepare()
	err = post.Validate()

	if err != nil {
		panic(err)
	}

	postCreated, err := post.SavePost(server.DB)
	// err = db.Debug().Model(&Post{}).Create(&p).Error
	// if err != nil {
	// 	return &Post{}, err
	// }
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Post{}, err
	// 	}
	// }
	// return p, nil
	if err != nil {
		panic(err)
	}

	responses.JSON(w, http.StatusCreated, postCreated)

}

// This function will return all Posts
func (server *Server) getAllPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}

	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		panic(err)
		return

	}

	responses.JSON(w, http.StatusOK, posts)
}

// This function will return specific post based on post ID
func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	id := vars["id"]

	post := models.Post{}
	PostFound, err := post.FindPostByID(server.DB, id)
	if err != nil {
		panic(err)
	}

	responses.JSON(w, http.StatusOK, postRecieved)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	post := models.Post{}

	//check if the post exits
	err := server.DB.Debug().Model(models.Post{}).Where("id = ?", id).Take(&post).Error
	if err != nil {
		// panic(err)
		fmt.Println("Post Not Found", err)
		return
	}

	if id != post.AuthorID {
		fmt.Println("not authorised user")
		return

	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
		return
	}

	// Request data - Update Procees

	postUpdate := models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		panic(err)
		return
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()
	if err != nil {
		panic(err)
		return
	}

	postUpdate.ID = post.ID //This is to maintain the original post id

	UpdatedPost, err := postUpdate.UpdateAPost(server.DB)
	if err != nil {
		panic(err)
		return
	}
	responses.JSON(w, http.StatusOK, UpdatedPost)
}
