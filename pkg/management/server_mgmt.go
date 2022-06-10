package management

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/millbj92/synctl/pkg/utils"
)

func StartServer(a *fiber.App, host string, port int) {
	exit := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Server refusing to shut down... Info: %v", err)
		}

		close(exit)
	}()

	// Build Fiber connection URL.
	fiberURL := BuildFiberURL(host, port)

	// Run server.
	if err := a.Listen(fiberURL); err != nil {
		log.Printf("Server refused to run. Info: %v", err)
	}

	<-exit
}

func BuildFiberURL (host string, port int) string {
	return fmt.Sprintf("http://%s:%d",  host, port)
}


func DebugStartServer(a *fiber.App) {
	// Build Fiber connection URL.
	fiberURL, _ := utils.ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberURL); err != nil {
		log.Printf("Server refused to run. Info: %v", err)
	}
}
