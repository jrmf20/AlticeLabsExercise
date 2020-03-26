package models

import (
	"fmt"
	"testing"
)

//TODO: Add Nonsense test
//Check Given Values
func TestGetEvents(t *testing.T) {
	phoneNumber := "961234567"
	UserList = map[string][]Event{
		"961234567": {{TimeStamp: 202003251000, Lat: 34.176253176, Lon: -8.4232526},
			{TimeStamp: 202003251030, Lat: 34.3325, Lon: -8.7232526},
			{TimeStamp: 202003251040, Lat: 33.176256, Lon: -8.3232526},
			{TimeStamp: 202003251100, Lat: 34.576253176, Lon: -8.5232526},
		}}

	_, err := GetEvents(phoneNumber)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

//TODO: Test time stamp in added Event
// Check recieved Values
// Add nonsense test
func TestPost(t *testing.T) {
	phoneNumber := "961234567"
	lat := "0.001"
	lon := "-0.001"
	UserList = make(map[string][]Event)

	err := AddEvent(phoneNumber, lat, lon)
	if err != nil {
		fmt.Println("No event added")
		t.Fail()
	}
}
