package restserver

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/config"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/logger"
)

type Server struct {
	server *echo.Echo
	conf   config.Config
}

// Start HTTP server
func Start(conf config.Config) *Server {
	log := logger.Get(conf)
	api := &Server{echo.New(), conf}
	api.server.HideBanner = true
	api.server.HTTPErrorHandler = errorHandler(logger.Get(conf))

	switch strings.ToUpper(conf.Env) {
	case config.Development:
		Wrap(api.server)
	default:
		log.Infof("%s, profiling desabled", conf.Env)
	}
	log.Info("Starting HTTP server...")
	api.server.GET("/", api.health)
	api.server.GET("/health", api.health)
	api.server.GET("/ready", api.health)

	go func() {
		api.server.Logger.Fatal(api.server.Start(":" + conf.HTTPPort))
	}()
	return api
}

// Shutdown HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	log := logger.Get(s.conf)
	if err := s.server.Shutdown(ctx); err != nil {
		log.Errorf("Error stopping server - ERROR [%+v]", err)
		return err
	}
	log.Info("Server shutting down")
	return nil
}
