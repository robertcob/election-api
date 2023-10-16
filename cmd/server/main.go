package main

import (
	"context"
	"election-api/internal/auth"
	"election-api/internal/config"
	customerr "election-api/internal/errors"
	"errors"
	"flag"
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/qiangxue/go-rest-api/pkg/accesslog"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Version = "1.0.0"
var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	logger := log.New().With(nil, "version", Version)
	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(logger, cfg),
	}
	// start the HTTP server with graceful shutdown
	go GracefulShutdown(hs, 10*time.Second, logger.Infof)
	logger.Infof("server %v is running at %v", Version, address)
	if err = hs.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err)
		os.Exit(-1)
	}

}

func buildHandler(logger log.Logger, cfg *config.Config) http.Handler {
	router := routing.New()
	router.Use(
		accesslog.Handler(logger),
		customerr.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)
	rg := router.Group("/v1")
	//authHandler := auth.Handler(cfg.JWTSigningKey)
	// register API handlers
	auth.RegisterHandlers(rg.Group(""),
		auth.NewService(cfg.JWTSigningKey, cfg.JWTExpiration, logger),
		logger,
	)
	return router

}

// GracefulShutdown shuts down the given HTTP server gracefully when receiving an os.Interrupt or syscall.SIGTERM signal.
// It will wait for the specified timeout to stop hanging HTTP handlers.
func GracefulShutdown(hs *http.Server, timeout time.Duration, logFunc func(format string, args ...interface{})) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	logFunc("shutting down server with %s timeout", timeout)
	if err := hs.Shutdown(ctx); err != nil {
		logFunc("errors while shutting down server: %v", err)
	} else {
		logFunc("server was shut down gracefully")
	}
}
