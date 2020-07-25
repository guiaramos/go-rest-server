package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity"
	"./repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(res http.ResponseWriter, req *http.Request) {
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error getting the posts"}`))
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error decoding the request`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}
