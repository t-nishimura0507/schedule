package schedule

//var GOOGLE_CALENDAR_ID string = os.Getenv("GOOGLE_CALENDAR_ID")

type Schedule struct {
	task string
	time string
}

func Get(date string) []Schedule {

	var schedules = []Schedule{
		{
			"Study",
			"20:00:00",
		},
		{
			"Cycling",
			"50:00:00",
		},
	}
	return schedules
}
