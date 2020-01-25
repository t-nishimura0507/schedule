package schedule

import (
	"fmt"
	"regexp"
)

// Schedule format.
type Schedule struct {
	task string
	time string
}

// Get a schedule for a specified date.
func Get(date string) ([]Schedule, error) {

	var schedules = []Schedule{}
	if isDateFormat(date) == false {
		return schedules, fmt.Errorf("param `date` is validation error")
	}


	schedules = []Schedule{
		{
			"Study",
			"20:00:00",
		},
		{
			"Cycling",
			"50:00:00",
		},
	}
	return schedules, nil
}

// isDateFormat is validation method of `date`.
func isDateFormat(date string) bool {
	// `2020-01-21`
	if m, err := regexp.MatchString(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`, date); m {
		if err != nil {
			fmt.Printf(err.Error())
		}
		return true
	}
	return false
}