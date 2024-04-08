package main

import (
	"encoding/json"
	"hcpb-api/db"
	"hcpb-api/models"

	"log"
	"net/http"
)



func getLastCallsHandlerNew(w http.ResponseWriter, r *http.Request) {
    jsonString, err := db.GetLastCallsDriver()
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Scrivi il JSON sulla risposta HTTP
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(jsonString))
}

func getStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	jsonString, err := db.GetStatistics()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonString))


}

func insertCallHandlerDriver(w http.ResponseWriter, r *http.Request) {
	var call models.Call
	err := json.NewDecoder(r.Body).Decode(&call)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.InsertCall(call)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}	



