package libs

import (
	"log"
	"time"
)

func CreateLogs(count int, duration string, sleep string) {

	if len(sleep) == 0 {
		sleep = "0s"
	}

	sleepDuration, _ := time.ParseDuration(sleep)
	if len(duration) > 0 {
		createLogsTillDuration(duration, sleepDuration)
	} else {
		createLogsTillCount(count, sleepDuration)
	}
}

func createLogsTillCount(count int, sleepDuration time.Duration) {
	number := 1
	for {
		if number > count {
			break
		}
		createLog(number)
		number++
		time.Sleep(sleepDuration)
	}
}

func createLogsTillDuration(duration string, sleepDuration time.Duration) {

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
			time.Sleep(sleepDuration)
		}
	}
}

func createLog(number int) {
	log.Printf("Hello Log at %d\n", number)
}
