// Meeting scheduler REST API built with Gin.
//
//	POST /schedule                     find a free slot for all participants and book it
//	GET  /users/:userId/calendar?start=&end=  list a user's events (RFC3339 times)
package main

import (
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// --------- Models ---------

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	UserID    string    `json:"userId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type MeetingRequest struct {
	ParticipantIds  []string `json:"participantIds"`
	DurationMinutes int      `json:"durationMinutes"`
	TimeRange       struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"timeRange"`
}

type MeetingResponse struct {
	MeetingId      string   `json:"meetingId"`
	Title          string   `json:"title"`
	ParticipantIds []string `json:"participantIds"`
	StartTime      string   `json:"startTime"`
	EndTime        string   `json:"endTime"`
}

// --------- In-memory DB ---------

var users = []User{
	{"u1", "Revti"},
	{"u2", "Sejal"},
	{"u3", "Tanvi"},
}

// mu protects events, because Gin runs each request in its own goroutine
var mu sync.Mutex

var events = []Event{
	{"e1", "Existing Meeting", "u1", parseTime("2025-08-02T14:00:00Z"), parseTime("2025-08-02T15:00:00Z")},
	{"e2", "Existing Meeting", "u2", parseTime("2025-08-03T09:00:00Z"), parseTime("2025-08-03T10:00:00Z")},
	{"e3", "Existing Meeting", "u3", parseTime("2025-08-04T13:00:00Z"), parseTime("2025-08-04T14:00:00Z")},
}

// --------- Utils ---------

// parseTime is only used for the hard-coded sample data above
func parseTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	return t
}

func userExists(id string) bool {
	for _, u := range users {
		if u.ID == id {
			return true
		}
	}
	return false
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// --------- API ---------

func main() {
	router := gin.Default()

	router.POST("/schedule", scheduleMeeting)
	router.GET("/users/:userId/calendar", getUserCalendar)

	router.Run(":8080")
}

// --------- Handlers ---------

func getUserCalendar(c *gin.Context) {
	userId := c.Param("userId")
	startStr := c.Query("start")
	endStr := c.Query("end")
	start, err := time.Parse(time.RFC3339, startStr)
	end, err2 := time.Parse(time.RFC3339, endStr)
	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	mu.Lock()
	defer mu.Unlock()

	userEvents := []Event{}
	for _, e := range events {
		if e.UserID == userId && e.StartTime.Before(end) && e.EndTime.After(start) {
			userEvents = append(userEvents, e)
		}
	}
	c.JSON(http.StatusOK, userEvents)
}

func scheduleMeeting(c *gin.Context) {
	var req MeetingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if len(req.ParticipantIds) == 0 || req.DurationMinutes <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "participantIds and a positive durationMinutes are required"})
		return
	}
	for _, id := range req.ParticipantIds {
		if !userExists(id) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Unknown participant: " + id})
			return
		}
	}
	start, err := time.Parse(time.RFC3339, req.TimeRange.Start)
	end, err2 := time.Parse(time.RFC3339, req.TimeRange.End)
	if err != nil || err2 != nil || !end.After(start) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timeRange: use RFC3339 and end must be after start"})
		return
	}
	duration := time.Duration(req.DurationMinutes) * time.Minute

	mu.Lock()
	defer mu.Unlock()

	slots := findCommonSlots(req.ParticipantIds, start, end, duration)
	if len(slots) == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "No available time slot found for all participants."})
		return
	}

	bestSlot := scoreAndSelectSlot(slots, req.ParticipantIds)
	meetingId := uuid.New().String()
	for _, userId := range req.ParticipantIds {
		events = append(events, Event{
			ID:        uuid.New().String(),
			Title:     "New Meeting",
			UserID:    userId,
			StartTime: bestSlot.Start,
			EndTime:   bestSlot.End,
		})
	}
	resp := MeetingResponse{
		MeetingId:      meetingId,
		Title:          "New Meeting",
		ParticipantIds: req.ParticipantIds,
		StartTime:      bestSlot.Start.Format(time.RFC3339),
		EndTime:        bestSlot.End.Format(time.RFC3339),
	}
	c.JSON(http.StatusCreated, resp)
}

// --------- Slot Discovery and Scoring ---------

type TimeSlot struct {
	Start time.Time
	End   time.Time
}

func findCommonSlots(participantIds []string, periodStart, periodEnd time.Time, duration time.Duration) []TimeSlot {
	busyIntervals := map[string][]TimeSlot{}
	for _, id := range participantIds {
		bi := []TimeSlot{}
		for _, ev := range events {
			if ev.UserID == id {
				bi = append(bi, TimeSlot{ev.StartTime, ev.EndTime})
			}
		}
		busyIntervals[id] = bi
	}

	slots := []TimeSlot{}
	date := periodStart
	for date.Add(duration).Before(periodEnd) || date.Add(duration).Equal(periodEnd) {
		possible := true
		slot := TimeSlot{date, date.Add(duration)}
		for _, id := range participantIds {
			for _, bi := range busyIntervals[id] {
				if slot.Start.Before(bi.End) && slot.End.After(bi.Start) {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			slots = append(slots, slot)
		}
		date = date.Add(15 * time.Minute)
	}
	return slots
}

func scoreAndSelectSlot(slots []TimeSlot, participantIds []string) TimeSlot {
	type ScoredSlot struct {
		Slot  TimeSlot
		Score int
	}
	var scored []ScoredSlot
	for _, slot := range slots {
		score := 0
		hour := slot.Start.Hour()
		// Lower score is better.
		// Prefer earlier slots
		score += hour * 2
		// Penalty for slots outside working hours (9 to 17)
		if hour < 9 || hour > 17 {
			score += 20
		}
		// Buffer check and gap minimization
		for _, id := range participantIds {
			for _, e := range events {
				if e.UserID == id {
					// Buffer: penalize if less than 15 min from an adjacent event
					if abs(int(slot.Start.Sub(e.EndTime).Minutes())) < 15 ||
						abs(int(e.StartTime.Sub(slot.End).Minutes())) < 15 {
						score += 10
					}
					// Awkward gap penalty
					gap := int(slot.Start.Sub(e.EndTime).Minutes())
					if gap > 0 && gap < 30 {
						score += 5
					}
				}
			}
		}
		scored = append(scored, ScoredSlot{slot, score})
	}
	// SliceStable keeps the earliest slot first when scores are equal
	sort.SliceStable(scored, func(i, j int) bool { return scored[i].Score < scored[j].Score })
	return scored[0].Slot
}
