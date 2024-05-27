package api

import (
	"hcpb-api/configs"
	"net/http"
)

func Init() {
	r := http.NewServeMux()

	r.HandleFunc("POST /call", insertCallHandlerDriver)
	r.HandleFunc("GET /calls", getLastCallsHandlerNew)
	r.HandleFunc("GET /statistics", getStatisticsHandler)

	server := http.Server{
		Addr:    ":" + configs.SERVICE_PORT,
		Handler: enableCORS(Logging(checkAPIKey(r))),
	}
	server.ListenAndServe()
}

func writeJson(w http.ResponseWriter, code int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
