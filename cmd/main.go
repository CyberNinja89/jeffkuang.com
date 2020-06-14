package main

import (
	"os"

	"github.com/cyberninja89/website/internal/consts"
	"github.com/cyberninja89/website/internal/server"

	"github.com/go-nacelle/log"
	"github.com/go-nacelle/nacelle"
)

func main() {
	// Configure environment variable sourcing
	sources := []nacelle.ConfigSourcer{
		nacelle.NewEnvSourcer(consts.ProcessName),
	}

	state := nacelle.NewBootstrapper(
		consts.ProcessName,
		func(process nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
			// Register additional processes and initializers here.

			// For now, just a web framework
			process.RegisterProcess(
				server.NewProcess(),
				nacelle.WithProcessName("api"),
			)
			return nil
		},
		nacelle.WithConfigSourcer(nacelle.NewMultiSourcer(sources...)),
		nacelle.WithLoggingFields(log.LogFields{
			"appName": consts.ProcessName,
		}),
	).Boot()

	if state != 0 {
		os.Exit(state)
	}
}
