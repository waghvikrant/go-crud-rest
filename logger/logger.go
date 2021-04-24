package logger

import (
	colorable "github.com/mattn/go-colorable"
	logger "github.com/sirupsen/logrus"
)

func Init() {
	logger.SetReportCaller(true)
	logger.SetFormatter(&logger.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		ForceColors:            true,
		DisableLevelTruncation: true,
	})
	logger.SetOutput(colorable.NewColorableStdout())
}
