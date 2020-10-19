package main

import (
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/ramizkhan99/meetingsAPI/src/controllers"
)

func TestGetMeetings(t *testing.T) {
	req, err := http.NewRequest("GET", "/meetings", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.MeetingHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unit test failed with: %v | %v", status, rr.Code)
	}
}