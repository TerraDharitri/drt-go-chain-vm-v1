package main

import (
	"os"

	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/vmserver"
)

var log = logger.GetOrCreate("vmserver")

const (
	// ErrCodeSuccess signals success
	ErrCodeSuccess = iota
	// ErrCodeCriticalError signals a critical error
	ErrCodeCriticalError
)

func main() {
	logger.ToggleLoggerName(true)
	_ = logger.SetLogLevel("*:TRACE")

	facade := vmserver.NewDebugFacade()
	app := initializeCLI(facade)

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(ErrCodeCriticalError)
	}

	os.Exit(ErrCodeSuccess)
}
