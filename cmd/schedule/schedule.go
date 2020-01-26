package schedule

import (
	"fmt"
	"github.com/t-nishimura0507/schedule/pkg/googleAPI"
	"google.golang.org/api/calendar/v3"
	"os"
	"regexp"
)

// Schedule format.
type Schedule struct {
	Unit      string
	Summary   string
	Detail    string
	StartTime string
	EndTime   string
}

// Get a schedule for a specified date.
func Get(date string) ([]Schedule, error) {

	// Validation.
	var schedules = []Schedule{}
	if isDateFormat(date) == false {
		return schedules, fmt.Errorf("param `date` is validation error")
	}

	// Get Client.
	client, err := googleAPI.GetClient()
	if err != nil {
		return nil, err
	}

	// Get Calender Service.
	service, err := calendar.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}

	// Public Holiday Schedule.
	//	var holidayCalendarID = os.Getenv("PUBLIC_HOLIDAY_CALENDAR_ID")
	var holidayCalendarID = "ja.japanese#holiday@group.v.calendar.google.com"
	if holidayCalendarID == "" {
		return nil, fmt.Errorf("Unset env `PUBLIC_HOLIDAY_CALENDAR_ID`.")
	}

	// Get Calendar Events.
	holidyEvents, err := service.Events.List(holidayCalendarID).
		SingleEvents(true).
		TimeMax(date + `T23:59:59+09:00`).
		TimeMin(date + `T00:00:00+09:00`).
		OrderBy("startTime").
		Do()
	if err != nil {
		return schedules, err
	}
	if len(holidyEvents.Items) == 0 {
		fmt.Println("No upcoming events found. ")
	} else {
		for _, item := range holidyEvents.Items {
			schedules = append(schedules, Schedule{
				"day",
				item.Summary,
				item.Description,
				"",
				"",
			})
		}
		return schedules, nil
	}

	// My Schedule.
	var calendarID = os.Getenv("GOOGLE_CALENDAR_ID")
	if calendarID == "" {
		return nil, fmt.Errorf("Unset env `GOOGLE_CALENDAR_ID`.")
	}

	// Get Calendar Events.
	events, err := service.Events.List(calendarID).
		SingleEvents(true).
		TimeMax(date + `T23:59:59+09:00`).
		TimeMin(date + `T00:00:00+09:00`).
		OrderBy("startTime").
		Do()
	if err != nil {
		return schedules, err
	}
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found. ")
	} else {
		for _, item := range events.Items {
			// スケジュールの単位（時間単位 or 日単位）を指定
			var unit = "hour"
			if item.Start.DateTime == "" && item.End.DateTime == "" {
				unit = "day"
			}

			schedules = append(schedules, Schedule{
				unit,
				item.Summary,
				item.Description,
				item.Start.DateTime,
				item.End.DateTime,
			})
		}
	}

	return schedules, nil
}

// isDateFormat is validation method of `date`. ex. `2020-01-21`
func isDateFormat(date string) bool {
	if m, _ := regexp.MatchString(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`, date); m {
		return true
	}
	return false
}
