package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/taylow/awaik-backend/internal/tui"
	"github.com/taylow/awaik-backend/services"

	_ "github.com/taylow/awaik-backend/services/monitor/command"
	_ "github.com/taylow/awaik-backend/services/monitor/query"
	_ "github.com/taylow/awaik-backend/services/test/ping"
	// _ "github.com/taylow/awaik-backend/services/health/monitoring"
	// _ "github.com/taylow/awaik-backend/services/health/recovery"
	// _ "github.com/taylow/awaik-backend/services/notification"
	// _ "github.com/taylow/awaik-backend/services/task/command"
	// _ "github.com/taylow/awaik-backend/services/task/execution"
	// _ "github.com/taylow/awaik-backend/services/task/query"
	// _ "github.com/taylow/awaik-backend/services/task/scheduling"
)

var (
	// ServicesToRun is a slice of service names, passed in as flags, that will be run
	ServicesToRun services.SliceFlag

	// RunAll determines whether to run all registered services
	RunAll bool

	// tuiMode determines whether to run in TUI mode
	tuiMode bool
)

// init initialises all flags
func init() {
	flag.Var(&ServicesToRun, "service", "determines which services are run - use this flag more than once for multiple services")
	flag.BoolVar(&RunAll, "all", false, "runs all services")
	flag.BoolVar(&tuiMode, "tui", false, "runs in TUI mode")
	flag.Parse()
}

// main is the entrypoint to the program
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// run runs the application with error handling
func run() error {
	var err error

	registeredServices := services.Services
	fmt.Printf("🔎 found %d service(s)\n", len(registeredServices))

	selectedServices := registeredServices
	if len(ServicesToRun) != 0 {
		selectedServices = registeredServices.Filter(ServicesToRun, RunAll)
	}

	if tuiMode {
		registeredServices, err = tui.SelectServices(registeredServices, selectedServices)
		if err != nil {
			return err
		}
	} else {
		registeredServices = selectedServices
	}

	fmt.Printf("🏃‍♀️ starting %d service(s)\n", len(registeredServices))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("🛑 received %s signal\n", sig)
		for _, service := range registeredServices {
			fmt.Println("⚠️  stopping", service.Name())
			if err := service.Stop(); err != nil {
				fmt.Printf("‼️ failed to stop service %q: %v", service.Name(), err)
			}
		}
		done <- true
	}()

	if err := runServices(registeredServices, done); err != nil {
		return err
	}

	return nil
}

func runServices(registry services.ServiceRegistry, done chan bool) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0

	for _, service := range registry {
		fmt.Printf("🚀 starting %s on %s\n", service.Name(), service.Address())
		wg.Add(1)

		go func(service services.Service) {
			defer wg.Done()

			if err := service.Start(); err != nil {
				fmt.Printf("🚨 failed to start %s: %v\n", service.Name(), err)
				return
			}

			fmt.Println("✅ started", service.Name())

			mu.Lock()
			count++
			mu.Unlock()
		}(service)
	}

	wg.Wait()

	if count == 0 {
		fmt.Println("🤷‍♀️ no services running")
	} else {
		fmt.Printf("🚀 running %d service(s)\n", count)
		<-done
	}
	fmt.Println("👋 bye bye")
	return nil
}
