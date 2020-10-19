package main

import (
	"io"
	"net/http"
	
	"github.com/ramizkhan99/meetingsAPI/src/app"
	"github.com/ramizkhan99/meetingsAPI/src/controllers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})
	http.HandleFunc("/meetings/", controllers.MeetingHandler)

	app.Run("8080")
}