package main

import (
	"fmt"
	"net/http"
	"rpi-healthcheck/healthcheck"
)

func main() {
	healthcheck.GetGpuTemp()
	fmt.Println("Starting server on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server", err)
	}
}
