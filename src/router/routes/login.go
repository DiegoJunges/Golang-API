package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Routes{
	URI:                  "/login",
	Method:               http.MethodPost,
	Function:             controllers.Login,
	AuthenticationNeeded: false,
}
