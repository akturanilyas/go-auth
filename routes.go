package main

import (
	"go-test/enums"
	"go-test/pkg/auth"
	"go-test/pkg/route"
	"net/http"
)

var Routes = []route.Route{
	{
		Method:  http.MethodPost,
		Path:    enums.Login,
		Handler: auth.Login,
	},
	{
		Method:  http.MethodPost,
		Path:    enums.Register,
		Handler: auth.Register,
	},
}
