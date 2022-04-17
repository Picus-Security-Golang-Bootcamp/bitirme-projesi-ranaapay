package graceful

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ShutdownGin shuts down the given HTTP server gracefully when receiving an os.Interrupt or syscall.SIGTERM signal.
// It will wait for the specified timeout to stop hanging HTTP handlers.
func ShutdownGin(srv *http.Server, timeout time.Duration) {
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Warn("Shutting down server...")

	// The context is used to inform the server it has timeout to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Warn("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
