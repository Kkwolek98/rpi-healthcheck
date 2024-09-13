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

	routes := enableCORS(initializeRoutes())

	fmt.Println("Starting server on port 3000")
	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runPeriodicalTask(ctx)

	server.ListenAndServe()
}

func initializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/temperature-readouts/last-week", temperaturecontroller.GetWeeklyHandler)
	mux.HandleFunc("GET /api/temperature-readouts/live", temperaturecontroller.GetLive)

	return mux
}

func runPeriodicalTask(ctx context.Context) {
	printTemp := func() {
		temp := healthcheck.GetGpuTemp()
		db.SaveTemperatureReadout(temp)
		str := fmt.Sprintf("%f", temp)
		bytes := []byte(str)
		temperaturecontroller.ConnectionManager.Broadcast(bytes)
	}
	go scheduler.RunPeriodically(ctx, 2*time.Second, printTemp)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
