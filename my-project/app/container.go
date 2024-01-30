package app

import (
	"fmt"
	"my-project/config"
	"my-project/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	userHandler user.Handler
)

// Init setups resources for container(configuration, logger, storage, etc.)
func Init() error {
	fmt.Println("initial database")
	return nil
}

// Start starts services and server
func Start() {

	orgRepository := user.NewRepository()
	orgService := user.NewService(orgRepository)
	userHandler = user.NewHandler(orgService)

	serverErr := Serve()

	// wait for server error or terminate signal
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

loop:
	for {
		select {
		case errServer := <-serverErr:
			fmt.Printf("error on server, %s\n", errServer.Error())
			break loop
		case sig := <-gracefulStop:
			fmt.Printf("caught signal[%v] then gracefully stop\n", sig)
			break loop
		}
	}
	Destroy()
}

func Serve() chan error {
	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(fmt.Sprintf(":%d", config.App.Server.Port), routes())
	}()

	return errChan
}

// Destroy clears all resources on container termination
func Destroy() {
	disconnectDB()
}

// disconnectServer gracefully stops server
func disconnectDB() {
	fmt.Printf("disconnect database")
}
