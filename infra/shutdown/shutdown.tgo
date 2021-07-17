package shutdown

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"time"

    "github.com/{{.Author}}/{{.RepositoryName}}/infra/config"
    "github.com/{{.Author}}/{{.RepositoryName}}/infra/logger"
)

type ShutdownFunction func(context.Context) error

func GracefulShutdownThatExpectsSignal(conf config.Config, functions ...ShutdownFunction) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
    <-quit

	GracefulShutdown(conf, functions...)
}

func GracefulShutdown(conf config.Config, functions ...ShutdownFunction) {
	var (
		log      = logger.Get(conf)
		pkgNames = pkgList(functions...)
		timeout  = 10
	)

	if conf.Env == config.Test {
		timeout = 0
	}

	log.Infof("Shutting down modules [%s]", pkgNames)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	for _, function := range functions {
		if err := function(ctx); err != nil {
			log.Errorf("Error shutting down the module %s server - %+v",
				runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), err)
		}
	}
    <-ctx.Done()

	log.Infof("Modules [%s] are down...", pkgNames)
}

func pkgList(functions ...ShutdownFunction) string {
	pkgs := ""
	for i, fn := range functions {
		nameFull := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		nameSlice := strings.Split(nameFull, "/")
		pkgSlice := strings.Split(nameSlice[len(nameSlice)-1], ".")
		name := fmt.Sprintf("%s%s", pkgSlice[0], pkgSlice[1])

		if i == len(functions)-1 {
			pkgs += name
			continue
		}
		pkgs += name + ", "
	}
	return pkgs
}
