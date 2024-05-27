package api

import (
	"encoding/json"
	control "hcpb-api/controllers"
	s "hcpb-api/schemas"
	"log/slog"

	"net/http"
)

func getLastCallsHandlerNew(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonString, err := control.GetLastCalls(ctx)
	if err != nil {
		slog.Error("Error getting last calls", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, jsonString)

}

func getStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonString, err := control.GetStatistics(ctx)
	if err != nil {
		slog.Error("Error getting statistics", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, jsonString)

}

func insertCallHandlerDriver(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var call s.Call
	err := json.NewDecoder(r.Body).Decode(&call)
	if err != nil {
		slog.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = control.InsertCall(ctx, call)
	if err != nil {
		slog.Error("Error inserting call", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
