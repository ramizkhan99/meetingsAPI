# meetingsAPI

An API to schedule meetings written in Go with Mongodb

## Endpoints
#### POST /meetings
Route to add a meeting<br />
Body:
```
{
  title: string,
  participants: [{
    name: string,
    email: string,
    rsvp: string
  }],
  start: time,
  end: time
}
```

#### GET /meetings
Route to get all meetings

#### GET /meetings/<meetingID>
Route to get a meeting with an objectID of `meetingID`
