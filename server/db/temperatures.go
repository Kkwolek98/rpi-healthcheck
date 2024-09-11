package db

import "time"

type TemperatureReadout struct {
	Temp      float64   `json:"temp"`
	Timestamp time.Time `json:"ts"`
}

func SaveTemperatureReadout(temp float64) {
	readout := &TemperatureReadout{Temp: temp, Timestamp: time.Now()}
	DB.Create(&readout)
}

func GetLastWeekTemperatureReadings() ([]TemperatureReadout, error) {
	var readouts []TemperatureReadout

	now := time.Now()
	weekAgo := now.Add(-7 * 24 * time.Hour)

	r := DB.Where("timestamp > ?", weekAgo).Find(&readouts)

	if r.Error != nil {
		return nil, r.Error
	}

	return readouts, nil
}
