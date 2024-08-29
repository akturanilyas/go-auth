package routes

import (
	authcontroller "go-test/controllers"
	"go-test/enums"
	"go-test/types"
	"net/http"
)

var AuthRoutes = []types.Route{
	{
		Method:  http.MethodGet,
		Path:    enums.Login,
		Handler: authcontroller.Login,
	},
	{
		Method:  http.MethodPost,
		Path:    enums.Register,
		Handler: authcontroller.Register,
	},
}
