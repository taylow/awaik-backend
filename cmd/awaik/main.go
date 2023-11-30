package main

import (
	"flag"
	"fmt"

	"github.com/taylow/awaik-backend/services"

	_ "github.com/taylow/awaik-backend/services/health/monitoring"
	_ "github.com/taylow/awaik-backend/services/health/recovery"
	_ "github.com/taylow/awaik-backend/services/notification"
	_ "github.com/taylow/awaik-backend/services/task/editing"
	_ "github.com/taylow/awaik-backend/services/task/execution"
	_ "github.com/taylow/awaik-backend/services/task/scheduling"
)

// ServicesToRun holds a list of service to run.
var (
	// ServicesToRun is a slice of service names, passed in as flags, that will be run.
	ServicesToRun services.SliceFlag

	// RunAll determines whether to run all initialised services.
	RunAll bool
)

// init initialises all flags.
func init() {
	flag.Var(&ServicesToRun, "service", "determines which services are run")
	flag.BoolVar(&RunAll, "all", false, "runs all services")
	flag.Parse()
}

// main is the entrypoint to the program.
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// run runs the application with error handling.
func run() error {
	services := services.Services
	fmt.Printf("Found %d service(s)\n", len(services))

	filteredServices := services.Filter(services, ServicesToRun, RunAll)
	fmt.Printf("Running %d service(s)\n", len(filteredServices))

	for _, service := range filteredServices {
		fmt.Println("running", service.Name())
		go func() {
			if err := service.Start(); err != nil {
				panic(err)
			}
		}()
		defer service.Stop()
	}

	return nil
}
