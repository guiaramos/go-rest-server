package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"./repository"
	"./service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
	// httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const PORT string = ":8000"

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)

	httpRouter.SERVER(PORT)
}
