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
	db.Init()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	printTemp := func() {
		temp := healthcheck.GetGpuTemp()
		db.SaveTemperatureReadout(temp)
		fmt.Println(temp)
	}

	fmt.Println("Starting server on port 4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println("Error starting server", err)
		return
	}

	go scheduler.RunPeriodically(ctx, 2*time.Second, printTemp)

	fmt.Println("getting readouts")

	readouts, err := db.GetLastWeekTemperatureReadings()

	if err != nil {
		return
	}

	fmt.Println(len(readouts))
}
