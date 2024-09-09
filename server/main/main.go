package main

import (
	"context"
	"fmt"
	"net/http"
	"rpi-healthcheck/db"
	"rpi-healthcheck/healthcheck"
	"rpi-healthcheck/scheduler"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	printTemp := func() {
		temp := healthcheck.GetGpuTemp()
		fmt.Printf("GPU Temp: %.2f'C\n", temp)
	}

	go scheduler.RunPeriodically(ctx, time.Second, printTemp)

	db.Init()
	fmt.Println("Starting server on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server", err)
	}
}
