package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/expose443/forum/backend/api"
	"github.com/expose443/forum/backend/pkg/configs"
	"github.com/expose443/forum/backend/pkg/logger"
)

func main() {
	logger := logger.New()
	cfg := configs.NewConfig(logger)

	server := api.NewServer(*cfg, *logger)

	go func() {
		logger.Info(fmt.Sprintf("server started at %s", server.Addr))
		if err := server.ListenAndServeTLS("", ""); err != nil {
			logger.ErrorLog.Printf("listen: %s\n", err)
		}
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info("Shutting down servers..")
	if err := server.Shutdown(ctx); err != nil {
		logger.ErrorLog.Printf("server shutdown: %s\n", err)
	} else {
		logger.Info("Server gracefully stoped")
	}

}
