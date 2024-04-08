package main

import (
	"hcpb-api/configs"
	"hcpb-api/db"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		statusColor := color.New(color.FgGreen).SprintFunc()
		methodColor := color.New(color.FgCyan).SprintFunc()
		pathColor := color.New(color.FgYellow).SprintFunc()

		log.Printf("%s %s %s %s", statusColor(wrapped.statusCode), methodColor(r.Method), pathColor(r.URL.Path), time.Since(start))
	})
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		expectedKey := configs.API_KEY
		if apiKey != expectedKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	timezone := "Europe/Rome"
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	time.Local = location

	db.Main()

	r := http.NewServeMux()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/call", insertCallHandlerDriver)
	r.HandleFunc("/calls", getLastCallsHandlerNew)
	r.HandleFunc("/statistics", getStatisticsHandler)

	// Chain middleware in the order you want them to execute
	server := http.Server{
		Addr:    ":7777",
		Handler: enableCORS(Logging(checkAPIKey(r))),
	}
	server.ListenAndServe()
}
