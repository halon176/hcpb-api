package api

import (
	"encoding/json"
	control "hcpb-api/controllers"
	s "hcpb-api/schemas"
	"log/slog"
	"math"
	"net/http"
	"strconv"
)

func getLastCallsHandlerNew(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	if pageStr == "" && pageSizeStr == "" {
		jsonString, err := control.GetLastCalls(ctx)
		if err != nil {
			slog.Error("Error getting last calls", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		writeJson(w, http.StatusOK, jsonString)
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	params := s.PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}

	totalItems, err := control.GetTotalCallsCount(ctx)
	if err != nil {
		slog.Error("Error getting total calls count", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := control.GetPaginatedCalls(ctx, params)
	if err != nil {
		slog.Error("Error getting paginated calls", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	var callsArray any
	if err := json.Unmarshal(data, &callsArray); err != nil {
		slog.Error("Error unmarshaling calls data", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := s.PaginatedCallsResponse{
		Data:       callsArray,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		slog.Error("Error marshaling paginated response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, jsonResponse)
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

func getExcludedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonString, err := control.GetExcluded(ctx)
	if err != nil {
		slog.Error("Error getting excluded", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeJson(w, http.StatusOK, jsonString)

}

func insertExcludedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	item := r.PathValue("item")
	err := control.InsertExcluded(ctx, item)
	if err != nil {
		slog.Error("Error inserting excluded", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteExcludedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	item := r.PathValue("item")
	err := control.DeleteExcluded(ctx, item)
	if err != nil {
		slog.Error("Error deleting excluded", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
