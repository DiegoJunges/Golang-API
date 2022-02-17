package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		AuthenticationNeeded: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.FindAll,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodGet,
		Function:             controllers.FindUser,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodPut,
		Function:             controllers.UpdateUser,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}/follow",
		Method:               http.MethodPost,
		Function:             controllers.FollowUser,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}/unfollow",
		Method:               http.MethodPost,
		Function:             controllers.UnfollowUser,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}/followers",
		Method:               http.MethodGet,
		Function:             controllers.GetFollowers,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}/following",
		Method:               http.MethodGet,
		Function:             controllers.GetFollowing,
		AuthenticationNeeded: true,
	},
	{
		URI:                  "/users/{userId}/update-password",
		Method:               http.MethodPost,
		Function:             controllers.UpdatePassword,
		AuthenticationNeeded: true,
	},
}
