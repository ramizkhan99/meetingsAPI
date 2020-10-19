package controllers

import (
	"time"
	"strings"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"github.com/ramizkhan99/meetingsAPI/src/models"
	"github.com/ramizkhan99/meetingsAPI/src/app"
)

var collection = app.GetClient().Database("meetingsApi").Collection("Meetings")

// MeetingHandler : Function to handle meeting requests
func MeetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.Method, r.URL, r.URL.Query())

	switch r.Method {
	case "GET":
		queries := r.URL.Query()
		if _, ok := queries["start"]; ok {
			var meetings []models.Meeting

			filter := bson.M{"start": bson.M{"$gte": queries["start"][0]}, "end": bson.M{"$lte": queries["end"][0]}}

			cur, err := collection.Find(context.TODO(), filter)

			if err != nil {
				log.Fatal(err)
				return
			}

			defer cur.Close(context.Background())
			for cur.Next(context.Background()) {
				var meeting models.Meeting

				err := cur.Decode(&meeting)
				if err != nil {
					log.Fatal(err)
					return
				}

				meetings = append(meetings, meeting)
			}

			if err := cur.Err(); err != nil {
				log.Fatal(err)
				return
			}

			if _, ok := queries["participant"]; ok {
				var meetingsWithParticipant []models.Meeting
				for i := range meetings {
					meeting := meetings[i]

					for j := range meeting.Participants {
						if queries["participant"][0] == meeting.Participants[j].Email {
							meetingsWithParticipant = append(meetingsWithParticipant, meeting)
						}
					}
				}
			}

			json.NewEncoder(w).Encode(meetings)
		} else {
			var meeting models.Meeting
			
			urlArray := strings.Split(r.URL.Path, "/")

			if urlArray[1] == "meetings" {
				log.Println(len(urlArray))
				var meetings []models.Meeting
				cur, err := collection.Find(context.TODO(), bson.D{})

				if err != nil {
					log.Fatal(err)
					return
				}

				defer cur.Close(context.Background())
				for cur.Next(context.Background()) {
					var meeting models.Meeting

					err := cur.Decode(&meeting)
					if err != nil {
						log.Fatal(err)
						return
					}

					meetings = append(meetings, meeting)
				}

				if err := cur.Err(); err != nil {
					log.Fatal(err)
					return
				}

				json.NewEncoder(w).Encode(meetings)
				return
			}
			
			meetingID, _ := primitive.ObjectIDFromHex(urlArray[len(urlArray)-1])
			
			filter := bson.M{"_id": meetingID}
			err := collection.FindOne(context.TODO(), filter).Decode(&meeting)
		
			if err != nil {
				log.Fatal(err)
				return
			}
		
			json.NewEncoder(w).Encode(meeting)
		}

	case "POST":
		log.Println("Here")
		var meeting models.Meeting
		_ = json.NewDecoder(r.Body).Decode(&meeting)

		meeting.CreatedAt = time.Now()
		_, err := collection.InsertOne(context.TODO(), meeting)

		if err != nil {
			log.Fatal(json.NewEncoder(w).Encode(err))
			return
		}
		
		json.NewEncoder(w).Encode(meeting)
	}	
}