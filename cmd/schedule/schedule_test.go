package schedule

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func before() {
    err := godotenv.Load("config/test.env", os.Getenv("GOOGLE_CALENDAR_ID"))
    if err != nil {
    	log.Fatal("Error loadin .nev file..")
	}
}

func TestGetSuccess(t *testing.T) {

	// param
	var param = "2020-01-21"

	// exec
	schedules, err := Get(param)

	// assertion
	if err != nil {
		t.Error("ERROR Message:" + err.Error())
	}
	log.Printf("%v", schedules)
}

func TestGetRestSuccess(t *testing.T) {

	// param
	var param = "2020-01-22"

	// exec
	schedules, err := Get(param)

	// assertion
	if err != nil {
		t.Error("ERROR Message:" + err.Error())
	}
	log.Printf("%v", schedules)
}

func TestGetHoliDaySuccess(t *testing.T) {

	// param
	var param = "2020-01-01"

	// exec
	schedules, err := Get(param)

	// assertion
	if err != nil {
		t.Error("ERROR Message:" + err.Error())
	}
	log.Printf("%v", schedules)
}

func TestGetValidationError(t *testing.T) {

	// param
	var param = "間違ったパラメータ"

	// exec
	_, err := Get(param)

	// assertion
	if err == nil {
		t.Error("Response is Success...")
	}
}