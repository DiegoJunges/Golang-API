package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Routes{
	{
		URI:                  "/posts",
		Method:               http.MethodPost,
		Function:             controllers.CreatePost,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/posts",
		Method:               http.MethodGet,
		Function:             controllers.FindAllPosts,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/posts/{postId}",
		Method:               http.MethodGet,
		Function:             controllers.FindPost,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/posts/{postId}",
		Method:               http.MethodPut,
		Function:             controllers.UpdatePost,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/posts/{postId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeletePost,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}/posts",
		Method:               http.MethodGet,
		Function:             controllers.FindPostsByUser,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/posts/{postId}/like",
		Method:               http.MethodPost,
		Function:             controllers.LikePost,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/posts/{postId}/dislike",
		Method:               http.MethodPost,
		Function:             controllers.DislikePost,
		AuthenticationNeeded: true,
	},
}
