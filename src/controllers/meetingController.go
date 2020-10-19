package controllers

import (
	"time"
	"strings"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	
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
				var meetings []models.Meeting

				opts := options.Find()
				
				skip, noSkip := strconv.Atoi(r.URL.Query()["skip"][0])
				limit, noLimit := strconv.Atoi(r.URL.Query()["limit"][0])

				if noSkip != nil {
					skip = 0
				}

				if noLimit != nil {
					limit = 0
				}
				
				opts.SetLimit((int64)(limit))
				opts.SetSkip((int64)(skip))
				cur, err := collection.Find(context.TODO(), bson.D{}, opts)

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

		filter := bson.M{"start": bson.M{"$gte": meeting.StartAt}, "end": bson.M{"$lte": meeting.EndAt}}
		
		cur, e := collection.Find(context.TODO(), filter)

		if e != nil {
			log.Fatal(e)
			return
		}

		defer cur.Close(context.Background())
		for cur.Next(context.Background()) {
			var prevMeetings 	[]models.Meeting
			var prevMeeting 	models.Meeting

			e := cur.Decode(&prevMeeting)


			if e != nil {
				log.Fatal(e)
				return
			}

			prevMeetings = append(prevMeetings, prevMeeting)
		}

		_, err := collection.InsertOne(context.TODO(), meeting)

		if err != nil {
			log.Fatal(json.NewEncoder(w).Encode(err))
			return
		}
		
		json.NewEncoder(w).Encode(meeting)
	}	
}
