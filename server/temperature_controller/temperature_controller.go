package temperaturecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	connectionmanager "rpi-healthcheck/connection_manager"
	"rpi-healthcheck/db"

	"github.com/gorilla/websocket"
)

var ConnectionManager = connectionmanager.NewConnectionManager()

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

func GetLive(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error during websocket upgrade", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ConnectionManager.AddConnection(conn)
	defer ConnectionManager.RemoveConnection(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				fmt.Println("Connection closed", conn.RemoteAddr())
			} else {
				fmt.Println("Error reading message:", err)
			}
			break
		}
	}
}
