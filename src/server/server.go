package server

import (
	"encoding/json"
	"models"
	"net/http"
	"fmt"
)

func AddLocal(w http.ResponseWriter, r *http.Request) {
	pnumber := r.PostFormValue("id")
	lat := r.PostFormValue("lat")
	lon := r.PostFormValue("lon")

	if pnumber == "" || lat == "" || lon == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Insuficient parameters"))
		return
	}
	if err := models.AddEvent(pnumber, lat, lon); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func GetLocals(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	pnumber := values.Get("id")

	if pnumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Id Number"))
		return
	}

	eventlist, err := models.GetEvents(pnumber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	jsonEventList, err := json.Marshal(eventlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	out := []byte(fmt.Sprintf("{locs: %s}",jsonEventList))
	w.Write(out)
}
