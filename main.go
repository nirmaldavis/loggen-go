package main

import (
	"flag"
	"log"

	"github.com/nirmaldavis/loggen-go/libs"
	"rsc.io/quote"
)

var (
	count    = flag.Int("count", 5, "Run for counts")
	duration = flag.String("duration", "", "Run for time duration")
	sleep    = flag.String("sleep", "", "Sleep between logging")
	listen   = flag.Bool("listen", false, "Listen for instructions")
)

func main() {

	flag.Parse()

	log.Println("count : ", *count)
	log.Println("duration : ", *duration)
	log.Println("sleep : ", *sleep)
	log.Println("listen : ", *listen)

	log.Println()

	log.Println("Hello, World!")
	log.Println(quote.Go())

	log.Println()

	libs.CreateLogs(*count, *duration, *sleep)
}
