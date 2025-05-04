// router/router.go
package router

import (
	"go-api/controller"
	"net/http"
)

func RegisterRoutes(userController *controller.UserController) {
	http.HandleFunc("/", userController.HelloHandler)
	http.HandleFunc("/user", userController.UserHandler)
}
