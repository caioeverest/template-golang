package main

import (
    "github.com/{{.Author}}/{{.RepositoryName}}/adapter/restserver"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/config"
	"github.com/{{.Author}}/{{.RepositoryName}}/infra/logger"
    "github.com/{{.Author}}/{{.RepositoryName}}/infra/shutdown"
)

func main() {
    conf := config.Get()
    log := logger.Get(conf)
    server := restserver.Start(conf)

    log.Infof("App started, to close it press CTRL-C")
    shutdown.GracefulShutdownThatExpectsSignal(conf, server.Shutdown)
}
