package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Caio-Nogueira/epoch/epoch"
)



func main() {
	wordPtr := flag.String("date", "2006-01-02T15:04", "date should be in one of the following formats: YYYY-MM-DDTHH:MM -> RFC3339")
	epochPtr := flag.Int("epoch", 0, "epoch time in seconds")

	flag.Parse()
	msg := "> %v\n"

	if *wordPtr != "" {
		epoch, err := epoch.Epoch(*wordPtr)
		if err != nil {
			log.Fatal(err)
		}
		println(fmt.Sprintf(msg, epoch))
	} else if *epochPtr != 0 {
		date := epoch.Date(*epochPtr)
		println(fmt.Sprintf(msg, date))
	} else {
		println(fmt.Sprintf(msg, epoch.CurrentTime()))
	}

}