package calendar

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"

	"github.com/dillonkkoval/MeetingsApp/auth"
)

func GetCalendarService(credentials string) (*calendar.Service, error) {
	b, err := ioutil.ReadFile(credentials)
	if err != nil {
		return nil, fmt.Errorf("Unable to read client secret file: %w", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %w", err)
	}
	client := auth.GetClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Calendar client: %w", err)
	}
	return srv, nil
}

func ReadCalendar(service *calendar.Service) (*calendar.Events, error) {
	t := time.Now().Format(time.RFC3339)
	events, err := service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve next ten of the user's events: %w", err)
	}

	return events, nil
}
