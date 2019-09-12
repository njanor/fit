package main

import (
	"fmt"
	"time"

	"github.com/tormoder/fit"
)

// Lap contains beginning and end second for the given lap
type Lap struct {
	Beginning time.Time
	End       time.Time
}

// InternalActivity is a file-independent representation of a cyling activity
type InternalActivity struct {
	Laps []Lap
}

// ConvertToInternalActivity converts from a FIT activity to a InternalActivity
func convertToInternalActivity(fitActivty *fit.ActivityFile) *InternalActivity {
	internalActivity := InternalActivity{
		Laps: make([]Lap, 0, len(fitActivty.Laps)),
	}

	for _, fitLap := range fitActivty.Laps {
		lap := Lap{
			Beginning: fitLap.StartTime,
			End:       fitLap.StartTime.Add(time.Second * time.Duration(fitLap.TotalElapsedTime)),
		}
		internalActivity.Laps = append(internalActivity.Laps, lap)
	}

	return &internalActivity
}

// TryToGetThisToWork prints struff
func TryToGetThisToWork() {
	fmt.Println("From separate file!")
}
