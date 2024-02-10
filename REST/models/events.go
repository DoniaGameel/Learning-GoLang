package models

import "time"

// if a field can't be null ==> use struct tag `binding:"required"`
type Event struct{
	ID				int
	Name 			string		`binding:"required"`
	Description		string		`binding:"required"`
	Location		string		`binding:"required"`
	DateTime		time.Time	`binding:"required"`
	UserID			int
}

var events = []Event{} //var events []Event = []Event()

func (e Event) Save(){
	// later save to the database
	events = append (events, e)
}

func GetAllEvents() []Event{
	return events
}