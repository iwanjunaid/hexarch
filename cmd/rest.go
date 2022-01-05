package cmd

import (
	"context"
	"time"

	"github.com/iwanjunaid/hexarch/configs"
	iregistry "github.com/iwanjunaid/hexarch/internal/adapter/inbound/registry"
	"github.com/iwanjunaid/hexarch/internal/adapter/inbound/rest"
	mregistry "github.com/iwanjunaid/hexarch/internal/adapter/outbound/repository/memdb/registry"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use: "rest",
	Run: func(cmd *cobra.Command, args []string) {
		InitConfigFromJson(cmd)

		logLevel := configs.GetString("logger.level")
		logFormat := configs.GetString("logger.format")
		logger := InitLogger(context.Background(), "hexarch", "1.0.0", logLevel, logFormat)

		restOptions := rest.RestOptions{
			Port:            configs.GetInt("rest.port"),
			GracefulTimeout: time.Duration(configs.GetInt("rest.graceful_timeout")) * time.Second,
			ReadTimeout:     time.Duration(configs.GetInt("rest.read_timeout")) * time.Second,
			WriteTimeout:    time.Duration(configs.GetInt("rest.write_timeout")) * time.Second,
			ApiTimeout:      configs.GetInt("rest.api_timeout"),
		}

		// Init repository registry
		repositoryRegistry := mregistry.NewRepositoryRegistry()

		// Init service registry
		serviceRegistry := iregistry.NewServiceRegistry(repositoryRegistry)

		// Init rest api server
		restApi := rest.New(logger, &restOptions, serviceRegistry)

		go restApi.Serve()

		restApi.SignalCheck()
	},
}
