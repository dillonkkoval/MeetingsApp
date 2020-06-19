package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	googleCalendar "google.golang.org/api/calendar/v3"

	"github.com/dillonkkoval/MeetingsApp/calendar"
	"github.com/peterbourgon/ff"
)

func main() {
	var (
		pollInterval    = flag.Duration("poll.interval", 5*time.Minute, "Calendar polling interval")
		jackCredentials = flag.String("jack.credentials", "./credentials/jack.json", "Credentials file for Jack's calendar")
		// dillonCredentials = flag.String("dillon.credentials", "./credentials/dillon.json", "Credentials file for Dillon's calendar")
		// nickCredentials   = flag.String("nick.credentials", "./credentials/nick.json", "Credentials file for Nick's calendar")
	)

	if err := ff.Parse(flag.CommandLine, os.Args[1:], ff.WithEnvVarNoPrefix()); err != nil {
		log.Fatalf("Error parsing flags: %s", err)
	}

	ticker := time.NewTicker(*pollInterval)

	jackCalendar, err := calendar.GetCalendarService("jack", *jackCredentials)
	if err != nil {
		log.Fatalf("error reading from jack's calendar: %v", err)
	}

	// dillonCalendar, err := calendar.GetCalendarService("dillon", *dillonCredentials)
	// if err != nil {
	// 	log.Fatalf("error reading from dillon's calendar")
	// }

	// nickCalendar, err := calendar.GetCalendarService("nick", *nickCredentials)
	// if err != nil {
	// 	log.Fatalf("error reading from nick's calendar")
	// }

	go poll(ticker, jackCalendar)
	// go poll(ticker, dillonCalendar)
	// go poll(ticker, nickCalendar)

	//lights.ChangeLight(lights.STAIRS, lights.GREEN)
	select {}
}

func poll(t *time.Ticker, c *googleCalendar.Service) {
	for {
		select {
		case tick := <-t.C:
			fmt.Printf("tick: %v\n", tick)
			jackEvents, err := calendar.ReadCalendar(c)
			if err != nil {
				log.Fatalf("error retrieving jack's calendar events: %v", err)
			}

			fmt.Println("Upcoming events:")
			if len(jackEvents.Items) == 0 {
				fmt.Println("No upcoming events found.")
			} else {
				for _, item := range jackEvents.Items {
					date := item.Start.DateTime
					if date == "" {
						date = item.Start.Date
					}
					fmt.Printf("%v (%v)\n", item.Summary, date)
				}
			}
		}
	}
}
