package storage

import "time"

// interface for interacting with repositories
type Storage interface {

	//adds any events
	Add(events []Event) (err error)

	// update event data in db
	Update(event *Event) (err error)

	//Removal(repo *Repository) (err error)

	//return all events
	All() (events []Event, err error)

	//return the last added event
	Last() (event *Event, err error)
}

type Event struct {
	Id          int64
	Title       string
	Type        string
	Start       time.Time
	AtEnd       time.Time
	AtAddress   string
	Coordinates string
	Format      string
	Description string
	Status      string
	AtAddition  time.Time
}
