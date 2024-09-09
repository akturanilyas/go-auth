package route

import (
	"go-test/enums"
	"go-test/pkg/auth"
	"net/http"
)

var Routes = []Route{
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
