package cmd

import (
	"context"

	"github.com/iwanjunaid/hexarch/configs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.sicepat.tech/platform/golib/log"
)

func InitConfigFromJson(cmd *cobra.Command) {
	configFile, err := cmd.Flags().GetString("config")

	if err != nil {
		panic(err)
	}

	err = configs.ConfigureFromJsonFile(configFile)

	if err != nil {
		panic(err)
	}
}

func InitLogger(ctx context.Context, serviceName, serviceVersion, level, format string) *logrus.Entry {
	log.SetLevel(format)
	logger := log.WithContext(ctx).WithFields(logrus.Fields{
		"service": serviceName,
		"version": serviceVersion,
	})

	if format == "json" {
		logger.Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	return logger
}
