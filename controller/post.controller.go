package controller

import (
	"encoding/json"
	"net/http"

	"../entity"
	"../errors"
	"../service"
)

type controller struct{}

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

var (
	postService service.PostService
)

// NewPostController create a new Post controller
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(res http.ResponseWriter, req *http.Request) {
	posts, err := postService.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func (*controller) AddPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error decoding the request"})
		return
	}

	err1 := postService.Validate(&post)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error on saving the post"})
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
