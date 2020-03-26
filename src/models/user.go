package models

import (
	"fmt"
	"strconv"
	"time"
	"bytes"
)

type Event struct {
	TimeStamp int64
	Lat       float64
	Lon       float64
}

var UserList map[string][]Event

func init() {
	UserList = make(map[string][]Event)
}

func GetEvents(pnumber string) ([]Event, error) {
	if eventslist, ok := UserList[pnumber]; ok {
		return eventslist, nil
	}
	return nil, fmt.Errorf("No Events Registered in number: %s", pnumber)
}

//adiciona um local quer exista ou não o numero dentro do mapa de utilizadores
func AddEvent(pnumber string, lat string, lon string) error {
	floatLat,err := strconv.ParseFloat(lat, 32)
	if err != nil{
		return fmt.Errorf("Latitude conversion Error")	
	}
	floatLon,err := strconv.ParseFloat(lon, 32)
	if err != nil{
		return fmt.Errorf("Longitude conversion Error")
	}
	newEvent := Event{time.Now().Unix(), floatLat, floatLon}
	UserList[pnumber] = append(UserList[pnumber], newEvent)
	return nil
}

func (e *Event) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"ts\" : %d,", e.TimeStamp))
	buffer.WriteString(fmt.Sprintf("\"lat\" : \"%f\",", e.Lat))
	buffer.WriteString(fmt.Sprintf("\"lon\" : %f", e.Lon))
	buffer.WriteString(fmt.Sprintf("}"))
	return buffer.Bytes(), nil
}
