package helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	log.Info("Logger initialized")
	return log
}
