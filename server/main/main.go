package main

import (
	"context"
	"fmt"
	"net/http"
	"rpi-healthcheck/db"
	"rpi-healthcheck/healthcheck"
	"rpi-healthcheck/scheduler"
	temperaturecontroller "rpi-healthcheck/temperature_controller"
	"time"
)

func main() {
	db.Init()

	routes := initializeRoutes()

	fmt.Println("Starting server on port 3000")
	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// runPeriodicalTask(ctx)

	server.ListenAndServe()
}

func initializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/temperature-readouts/last-week", temperaturecontroller.GetWeeklyHandler)

	return mux
}

func runPeriodicalTask(ctx context.Context) {
	printTemp := func() {
		temp := healthcheck.GetGpuTemp()
		db.SaveTemperatureReadout(temp)
		fmt.Println(temp)
	}
	go scheduler.RunPeriodically(ctx, 2*time.Second, printTemp)
}
