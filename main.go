package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func main() {
	router := mux.NewRouter()

	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	posts = append(posts, Post{ID: "2", Title: "My second post", Body: "This is the content of my second post"})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPosts).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

func createPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

// Update Post

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)

			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)

			return
		}
	}

	json.NewEncoder(w).Encode(posts)
}

// delete post

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(posts)
}
