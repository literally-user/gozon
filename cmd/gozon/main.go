package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	"github.com/literally_user/gozon/internal/config"
	"github.com/literally_user/gozon/internal/infrastructure/mock/publisher"
	"github.com/literally_user/gozon/internal/infrastructure/mock/repositories"
	userPresentation "github.com/literally_user/gozon/internal/presentation/controllers/manageUser"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	run(ctx)
}

func run(ctx context.Context) {
	// Config reader
	var (
		configReader = config.NewReader("your config path")
		configData   = configReader.Read()
	)

	// Http server
	var (
		mux = http.NewServeMux()
		srv = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", configData.API.Host, configData.API.Port),
			Handler: mux,
		}
	)

	// Repositories
	var (
		userRepository = repositories.NewInMemoryUserRepository(repositories.NewUserStorage())
		mockPublisher  = publisher.NewMockPublisher()
	)

	// Interactors
	var (
		createUserInteractor = userApplication.CreateUserInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
		deleteUserInteractor = userApplication.DeleteUserInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
	)

	// Controllers
	var (
		createUserController = userPresentation.CreateUserController{
			CreateUserInteractor: createUserInteractor,
		}
		deleteUserController = userPresentation.DeleteUserController{
			DeleteUserInteractor: deleteUserInteractor,
		}
	)

	// Register handlers
	mux.HandleFunc("POST /users/", createUserController.Execute)
	mux.HandleFunc("DELETE /users/", deleteUserController.Execute)

	// Serve
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Server started at http://%s:%d", configData.API.Host, configData.API.Port)

	// Wait for sigterm/sigint signal from OS
	<-ctx.Done()

	log.Println("Shutting down gracefully")
	shutdownContext, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := srv.Shutdown(shutdownContext); err != nil {
		log.Println(err)
	}
}
