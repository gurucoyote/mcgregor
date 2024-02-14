package main

import (
	"bufio"
	"bytes"
	"strings"
	"fmt"
	"sort"
	ics "github.com/arran4/golang-ical"
	"log"
	"os"
	"os/exec"
	"time"
)

type Event struct {
	Summary       string
	Description   string
	StartAt       time.Time
	EndAt         time.Time
	AllDayStartAt time.Time
	AllDayEndAt   time.Time
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myevents <path-to-directory>")

		os.Exit(1)
	}

	path := os.Args[1]

	files, err := findICSFiles(path)
	if err != nil {
		log.Fatalf("Error finding .ics files: %v", err)
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		calendar, err := ics.ParseCalendar(file)
		if err != nil {
			// fmt.Printf("Error parsing calendar: %v\n", err)
			continue
		}
		var events []Event
		// Use the calendar name if needed
		// Example: fmt.Printf("Calendar Name: %s\n", getCalendarName(calendar))

		for _, component := range calendar.Components {
			switch event := component.(type) {
			case *ics.VEvent:
				var e Event

				summaryProp := event.GetProperty(ics.ComponentPropertySummary)
				if summaryProp != nil {
					e.Summary = ics.FromText(summaryProp.Value)
				}

				descriptionProp := event.GetProperty(ics.ComponentPropertyDescription)
				if descriptionProp != nil {
					e.Description = formatDescription(descriptionProp.Value)
				}

				e.StartAt, _ = event.GetStartAt()

				e.EndAt, _ = event.GetEndAt()

				e.AllDayStartAt, _ = event.GetAllDayStartAt()

				e.AllDayEndAt, _ = event.GetAllDayEndAt()

				events = append(events, e)
			}
		}
		 sort.Slice(events, func(i, j int) bool {
			startI := events[i].StartAt
			if events[i].AllDayStartAt.IsZero() == false {
				startI = events[i].AllDayStartAt
			}
			startJ := events[j].StartAt
			if events[j].AllDayStartAt.IsZero() == false {
				startJ = events[j].AllDayStartAt
			}
			return startI.Before(startJ)
		})
		if true {
		today := time.Now().Truncate(24 * time.Hour)
		events = filterEvents(events, today)
	}

		// Output the events in the following format:
		// startdatetime,weekday, summary, description, enddatetime on one line

		for _, e := range events {
			weekday := e.StartAt.Weekday().String()
			startDate := e.StartAt.Format("2006-01-02")
			endDate := e.EndAt.Format("2006-01-02")
			fmt.Printf("%s,%s,%s,%s,%s\n", startDate, weekday, e.Summary, e.Description, endDate)
		}
	}
}
func filterEvents(events []Event, after time.Time) []Event {
	var filtered []Event
	for _, e := range events {
		if e.EndAt.After(after) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
func findICSFiles(path string) ([]string, error) {
	cmd := exec.Command("find", path, "-type", "f", "-name", "*.ics")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	files := []string{}
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}
	return files, scanner.Err()
}
func formatDescription(description string) string {
	description = strings.ReplaceAll(description, "\n", " ")
	description = strings.ReplaceAll(description, "\r", " ")
	if len(description) > 100 {
		description = description[:100]
	}
	return description
}
