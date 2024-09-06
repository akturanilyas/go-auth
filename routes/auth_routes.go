package routes

import (
	authController "go-test/controllers"
	endpoints "go-test/enums"
	"go-test/types"
	"net/http"
)

var AuthRoutes = []types.Route{
	{
		Method:  http.MethodGet,
		Path:    endpoints.Login,
		Handler: authController.Login,
	},
	{
		Method:  http.MethodPost,
		Path:    endpoints.Register,
		Handler: authController.Register,
	},
}
