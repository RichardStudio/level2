package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalf("Error fetching NTP time: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(exactTime.Format(time.RFC1123))
}
