package db

import (
	"math"
	"time"
)

type TemperatureReadout struct {
	Temp      float64   `json:"temp"`
	Timestamp time.Time `json:"ts"`
}

func SaveTemperatureReadout(temp float64) {
	readout := &TemperatureReadout{Temp: temp, Timestamp: time.Now()}
	DB.Create(&readout)
}

func GetLastWeekTemperatureReadings() ([]TemperatureReadout, error) {
	// TODO: average
	var readouts []TemperatureReadout

	now := time.Now()
	weekAgo := now.Add(-7 * 24 * time.Hour)

	r := DB.Where("timestamp > ?", weekAgo).Find(&readouts)

	if r.Error != nil {
		return nil, r.Error
	}

	avgReadouts := averageOutReadouts(&readouts, 8)

	return avgReadouts, nil
}

func averageOutReadouts(readouts *[]TemperatureReadout, sections int) []TemperatureReadout {
	if len(*readouts) == 0 || sections <= 0 {
		return nil
	}

	readoutsLen := len(*readouts)
	sectionSize := readoutsLen / sections

	if sectionSize <= 1 {
		return *readouts
	}

	var averagedReadouts []TemperatureReadout

	for i := 0; i < sections; i++ {
		start := i * sectionSize
		end := start + sectionSize

		if i == sections-1 {
			end = readoutsLen
		}

		currSection := (*readouts)[start:end]
		currSectionLen := len(currSection)

		var totalTemp float64
		for _, r := range currSection {
			totalTemp += r.Temp
		}
		avgTemp := totalTemp / float64(currSectionLen)
		roundedAvgTemp := math.Round(avgTemp*100) / 100

		midIndex := (start + currSectionLen) / 2
		midTimestamp := (*readouts)[midIndex].Timestamp

		averagedReadouts = append(averagedReadouts, TemperatureReadout{
			Temp:      roundedAvgTemp,
			Timestamp: midTimestamp,
		})
	}

	return averagedReadouts
}
