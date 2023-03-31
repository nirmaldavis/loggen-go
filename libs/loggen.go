package libs

import (
	"log"
	"time"
)

func CreateLogs(count int, duration string) {
	if len(duration) > 0 {
		createLogsTillDuration(duration)
	} else {
		createLogsTillCount(count)
	}
}

func createLogsTillCount(count int) {
	number := 1
	for {
		if number > count {
			break
		}
		createLog(number)
		number++
	}
}

func createLogsTillDuration(duration string) {

	if len(duration) > 0 {
		timeDuration, _ := time.ParseDuration(duration)
		log.Println("Time Duration : ", timeDuration)
		endTime := time.Now().Add(timeDuration)
		log.Println("Need to run till : ", endTime)

		number := 1
		for {
			if time.Now().After(endTime) {
				break
			}
			createLog(number)
			number++
		}
	}
}

func createLog(number int) {
	log.Printf("Hello Log at %d\n", number)
}
