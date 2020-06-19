package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dillonkkoval/MeetingsApp/calendar"
	"github.com/peterbourgon/ff"
)

func main() {
	var (
		jackCredentials = flag.String("jack.credentials", "./credentials/jack.json", "Credentials file for Jack's calendar")
		// dillonCredentials = flag.String("dillon.credentials", "./credentials/dillon.json", "Credentials file for Dillon's calendar")
		// nickCredentials   = flag.String("nick.credentials", "./credentials/nick.json", "Credentials file for Nick's calendar")
	)

	if err := ff.Parse(flag.CommandLine, os.Args[1:], ff.WithEnvVarNoPrefix()); err != nil {
		log.Fatalf("Error parsing flags: %s", err)
	}

	jackCalendar, err := calendar.GetCalendarService(*jackCredentials)
	if err != nil {
		log.Fatalf("error reading from jack's calendar: %v", err)
	}

	jackEvents, err := calendar.ReadCalendar(jackCalendar)
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

	// dillonCalendar, err := calendar.GetCalendarService(*dillonCredentials)
	// if err != nil {
	// 	log.Fatalf("error reading from dillon's calendar")
	// }

	// nickCalendar, err := calendar.GetCalendarService(*nickCredentials)
	// if err != nil {
	// 	log.Fatalf("error reading from nick's calendar")
	// }

	//lights.ChangeLight(lights.STAIRS, lights.GREEN)

}
