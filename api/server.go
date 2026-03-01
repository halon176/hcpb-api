package api

import (
	"hcpb-api/configs"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Init() {
	r := http.NewServeMux()

	r.HandleFunc("POST /calls", insertCallHandlerDriver)
	r.HandleFunc("GET /calls", getLastCallsHandlerNew)
	r.HandleFunc("GET /statistics", getStatisticsHandler)
	r.HandleFunc("GET /excluded", getExcludedHandler)
	r.HandleFunc("POST /excluded/{item}", insertExcludedHandler)
	r.HandleFunc("DELETE /excluded/{item}", deleteExcludedHandler)

	handler := enableCORS(Logging(checkAPIKey(r)))

	server := http.Server{
		Addr:    ":" + configs.SERVICE_PORT,
		Handler: otelhttp.NewHandler(handler, "hcpb-api"),
	}
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server error", "error", err)
	}
}

func writeJson(w http.ResponseWriter, code int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response", "error", err)
	}
}
