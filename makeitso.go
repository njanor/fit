package main

import (
	"fmt"
	"os"
)

func main() {
	fitActivity, loadStatus := LoadFitActivityFromFile(os.Args[1])
	if loadStatus != nil {
		fmt.Println(loadStatus)
	} else {
		fmt.Println("File loaded.")
	}

	activity := convertToInternalActivity(fitActivity)

	fmt.Printf("The file contains %v laps\n", len(activity.Laps))
	for lapIndex, lap := range activity.Laps {
		fmt.Printf("Lap %v:\nBeginning: %v\nEnd: %v\n\n", lapIndex+1, lap.Beginning, lap.End)
	}
}
