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
	"github.com/literally_user/gozon/internal/infrastructure/auth"
	"github.com/literally_user/gozon/internal/infrastructure/mock/publisher"
	"github.com/literally_user/gozon/internal/infrastructure/mock/repositories"
	userPresentation "github.com/literally_user/gozon/internal/presentation/controllers/manageUser"
	"github.com/literally_user/gozon/internal/presentation/middlewares"
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

	// Token Manager
	var (
		tokenManager = auth.TokenManager{
			SecretKey: []byte(configData.API.SecretKey),
		}
	)

	// Interactors
	var (
		createUserInteractor = userApplication.CreateUserInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
		getUserInteractor = userApplication.GetUserInteractor{
			Repository: &userRepository,
		}
		deleteUserInteractor = userApplication.DeleteUserInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
		changeUsernameInteractor = userApplication.ChangeUsernameInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
		changePasswordInteractor = userApplication.ChangePasswordInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
		changeEmailInteractor = userApplication.ChangeEmailInteractor{
			Repository: &userRepository,
			Publisher:  mockPublisher,
		}
	)

	// Controllers
	var (
		createUserController = userPresentation.CreateUserController{
			CreateUserInteractor: createUserInteractor,
			TokenManager:         tokenManager,
		}
		loginController = userPresentation.LoginController{
			GetUserInteractor: getUserInteractor,
			TokenManager:      tokenManager,
		}
		logoutController     = userPresentation.LogoutController{}
		deleteUserController = userPresentation.DeleteUserController{
			DeleteUserInteractor: deleteUserInteractor,
		}
		changeUsernameController = userPresentation.ChangeUsernameController{
			ChangeUsernameInteractor: changeUsernameInteractor,
		}
		changePasswordController = userPresentation.ChangePasswordController{
			ChangePasswordInteractor: changePasswordInteractor,
		}
		changeEmailController = userPresentation.ChangeEmailController{
			ChangeEmailInteractor: changeEmailInteractor,
		}
	)

	var (
		authMiddleware = middlewares.AuthMiddleware{
			TokenManager: tokenManager,
		}
	)

	// Register handlers
	mux.HandleFunc("POST /users/register", createUserController.Execute)
	mux.HandleFunc("POST /users/login", loginController.Execute)
	mux.Handle(
		"POST /users/me/logout",
		authMiddleware.Execute(http.HandlerFunc(logoutController.Execute)),
	)
	mux.Handle(
		"DELETE /users/me",
		authMiddleware.Execute(http.HandlerFunc(deleteUserController.Execute)),
	)
	mux.Handle(
		"POST /users/me/username",
		authMiddleware.Execute(http.HandlerFunc(changeUsernameController.Execute)),
	)
	mux.Handle(
		"POST /users/me/password",
		authMiddleware.Execute(http.HandlerFunc(changePasswordController.Execute)),
	)
	mux.Handle(
		"POST /users/me/email",
		authMiddleware.Execute(http.HandlerFunc(changeEmailController.Execute)),
	)
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
