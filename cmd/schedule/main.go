package schedule

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

// Schedule format.
type Schedule struct {
	task string
	time string
}

// Get a schedule for a specified date.
func Get(date string) ([]Schedule, error) {

	// Validation.
	var schedules = []Schedule{}
	if isDateFormat(date) == false {
		return schedules, fmt.Errorf("param `date` is validation error")
	}

	// Get Client.
	client := getClient()

	// Get Calender Service.
	service, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// Get Calendar Events.
	events, err := service.Events.List(os.Getenv("GOOGLE_CALENDAR_ID")).
		SingleEvents(true).
		TimeMax(date + `T23:59:59+09:00`).
		TimeMin(date + `T00:00:00+09:00`).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Printf(err.Error())
		return schedules, fmt.Errorf("errro...")
		//log.Fatalf("Unable to get events..", err)
	}
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found. ")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
			fmt.Printf("%v (%v)\n", item.Start.DateTime, date)
			fmt.Printf("%v (%v)\n", item.End.DateTime, date)
		}
	}


	// Create Response.
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
	if m, _ := regexp.MatchString(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`, date); m {
		return true
	}
	return false
}

// getClient
func getClient() *http.Client {
	// credentials.jsonを読み込み
	b, err := ioutil.ReadFile("../../config/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// クライアント作成
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to cofing: %v", err)
	}

	tokenFile := "../../config/token.json"
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}

func tokenFromFile(fileName string) (*oauth2.Token, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
/*
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
*/
    authCode := "4/vwHhC5qPP_v1k_qP1sHvV7HOEV9Whjhx0Pg5UTcEA_yb16ZUTOAcdH0"
	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return token
}

func saveToken(path string, token *oauth2.Token) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(token)
}

func getDuration(startTime string, endTime string) time.Duration {
	// Duration..
	st, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	et, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return et.Sub(st)
}