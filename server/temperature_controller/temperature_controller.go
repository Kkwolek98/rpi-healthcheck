package temperaturecontroller

import (
	"encoding/json"
	"net/http"
	"rpi-healthcheck/db"
)

func GetWeeklyHandler(w http.ResponseWriter, r *http.Request) {
	readouts, err := db.GetLastWeekTemperatureReadings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(readouts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}

}
