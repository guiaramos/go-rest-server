package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"google.golang.org/api/option"

	"../entity"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "go-rest-server"
	collectionName string = "posts"
)

var (
	ctx = context.Background()
)

func getFirebaseClient() (*firestore.Client, error) {
	sa := option.WithCredentialsFile("/Users/guilhermeramos/Desktop/study/go/go-rest-server/go-rest-server-firebase-adminsdk-3unyw-36ea1d61ff.json")

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatalf("Failed to create the Firebase App: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	return client, err
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	client, _ := getFirebaseClient()
	defer client.Close()

	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	client, _ := getFirebaseClient()
	defer client.Close()

	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
