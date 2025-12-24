package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/literally_user/gozon/config"
	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/infrastructure/mock/publisher"
	"github.com/literally_user/gozon/internal/infrastructure/mock/repositories"
	userPresentation "github.com/literally_user/gozon/internal/presentation/controllers/manageUser"
)

func main() {
	configReader := config.NewReader("")
	configData := configReader.Read()

	mux := http.NewServeMux()

	userRepository := repositories.NewInMemoryUserRepository()
	mockPublisher := publisher.NewMockPublisher()
	createUserInteractor := userApplication.CreateUserInteractor{
		Repository: userRepository,
		Publisher:  mockPublisher,
	}
	deleteUserInteractor := userApplication.DeleteUserInteractor{
		Repository: userRepository,
		Publisher:  mockPublisher,
	}
	createUserController := userPresentation.CreateUserController{
		CreateUserInteractor: createUserInteractor,
	}
	deleteUserController := userPresentation.DeleteUserController{
		DeleteUserInteractor: deleteUserInteractor,
	}

	mux.HandleFunc("POST /users/", createUserController.Execute)
	mux.HandleFunc("DELETE /users/", deleteUserController.Execute)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", configData.API.Host, configData.API.Port), mux))
}
