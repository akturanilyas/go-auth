package main

import (
	"auth-go/enums"
	"auth-go/pkg/auth"
	"auth-go/pkg/middleware"
	"auth-go/pkg/route"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var Routes = []route.Route{
	{
		Method:      http.MethodPost,
		Path:        enums.Login,
		Handler:     auth.Login,
		Middlewares: []fiber.Handler{middleware.ValidateRequest(&auth.LoginValidation{})},
	},
	{
		Method:      http.MethodPost,
		Path:        enums.Register,
		Handler:     auth.Register,
		Middlewares: []fiber.Handler{middleware.ValidateRequest(&auth.CreateUserValidation{})},
	},
}
