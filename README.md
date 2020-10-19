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

#### GET /meetings/:meetingID
Route to get a meeting with an objectID of `meetingID`

#### GET /meetings/?limit=n&skip=m
Route to get `n` meetings after skipping `m` meetings. If not specified, the defaults are `limit = 10` and `skip = 0`

#### GET /meetings/?start=t1&end=t2
Route to get a list of all meetings scheduled between t1 and t2

#### GET /meetings/?start=t1&end=t2&participant=p
Route to get a list of all meetings scheduled between t1 and t2 of a certain participant p



### What is broken
Checking for overlap of meetings of a particular participant which have an `RSVP = Yes` is broken. The README would be updated if it is fixed.


### How to run
1. Clone the repo
2. `cd meetingsAPI`
3. Run `go build -o meetingsAPI`
4. Run `./meetingsAPI`

Run `go test -v` to run unit tests
