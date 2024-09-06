package routes

import (
	authController "go-test/controllers"
	"go-test/enums"
	"go-test/types"
	"net/http"
)

var AuthRoutes = []types.Route{
	{
		Method:  http.MethodGet,
		Path:    enums.Login,
		Handler: authController.Login,
	},
	{
		Method:  http.MethodPost,
		Path:    enums.Register,
		Handler: authController.Register,
	},
}
