package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

var events = make(map[int]Event)
var idCounter = 1

func main() {
	http.HandleFunc("/create_event", logRequest(createEventHandler))
	http.HandleFunc("/update_event", logRequest(updateEventHandler))
	http.HandleFunc("/delete_event", logRequest(deleteEventHandler))
	http.HandleFunc("/events_for_day", logRequest(eventsForDayHandler))
	http.HandleFunc("/events_for_week", logRequest(eventsForWeekHandler))
	http.HandleFunc("/events_for_month", logRequest(eventsForMonthHandler))

	port := "8080"
	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	userID, date, title, err := parseEventParams(r)
	if err != nil {
		http.Error(w, `{"error": "Invalid input", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	event := Event{
		ID:     idCounter,
		UserID: userID,
		Date:   date,
		Title:  title,
	}
	events[idCounter] = event
	idCounter++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"result": "Event created", "http_code": http.StatusOK})
	fmt.Printf("Created event: %+v\n", event)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, `{"error": "Invalid event ID", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	userID, date, title, err := parseEventParams(r)
	if err != nil {
		http.Error(w, `{"error": "Invalid input", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	if _, ok := events[id]; !ok {
		http.Error(w, `{"error": "Event not found", "http_code": 503}`, http.StatusServiceUnavailable)
		return
	}

	events[id] = Event{
		ID:     id,
		UserID: userID,
		Date:   date,
		Title:  title,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"result": "Event updated", "http_code": http.StatusOK})
	fmt.Printf("Updated event: %+v\n", events[id])
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, `{"error": "Invalid event ID", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	if _, ok := events[id]; !ok {
		http.Error(w, `{"error": "Event not found", "http_code": 503}`, http.StatusServiceUnavailable)
		return
	}

	delete(events, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"result": "Event deleted", "http_code": http.StatusOK})
	fmt.Printf("Deleted event ID: %d\n", id)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	eventsForPeriodHandler(w, r, 24*time.Hour)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	eventsForPeriodHandler(w, r, 7*24*time.Hour)
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	eventsForPeriodHandler(w, r, 30*24*time.Hour)
}

func eventsForPeriodHandler(w http.ResponseWriter, r *http.Request, period time.Duration) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "Invalid request method", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, `{"error": "Missing date parameter", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid date format", "http_code": 400}`, http.StatusBadRequest)
		return
	}

	fmt.Printf("Checking events for period from %s for %v\n", date, period)

	var result []Event
	endDate := date.Add(period)

	for _, event := range events {
		fmt.Printf("Checking event: %+v\n", event)
		if event.Date.After(date.Add(-time.Second)) && event.Date.Before(endDate.Add(time.Second)) {
			result = append(result, event)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(result) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{"result": []Event{}, "http_code": http.StatusOK})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "http_code": http.StatusOK})
	}
}

func parseEventParams(r *http.Request) (int, time.Time, string, error) {
	err := r.ParseForm()
	if err != nil {
		return 0, time.Time{}, "", err
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return 0, time.Time{}, "", err
	}

	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return 0, time.Time{}, "", err
	}

	title := r.FormValue("title")
	return userID, date, title, nil
}
