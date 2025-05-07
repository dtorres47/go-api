package controller

import (
	"encoding/json"
	"go-api/service"
	"log"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) HelloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if _, err := responseWriter.Write([]byte("Hello from the service layer!")); err != nil {
		log.Printf("Error writing hello response: %v", err)
	}
}

func (uc *UserController) UserHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	users, err := uc.userService.ListUsers()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(responseWriter).Encode(users); err != nil {
		log.Printf("Error encoding users: %v", err)
	}

	//switch request.Method {
	//case http.MethodGet:
	//	user, err := uc.userService.GetUser()
	//	if err != nil {
	//		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	if err := json.NewEncoder(responseWriter).Encode(user); err != nil {
	//		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	//	}
	//
	//case http.MethodPost:
	//	var user service.User
	//	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
	//		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//	created, err := uc.userService.CreateUser(user)
	//	if err != nil {
	//		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	if err := json.NewEncoder(responseWriter).Encode(created); err != nil {
	//		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	//	}
	//
	//default:
	//	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (uc *UserController) UsersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")

	users, err := uc.userService.ListUsers()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(responseWriter).Encode(users); err != nil {
		log.Printf("Error encoding users list: %v", err)
	}
}
