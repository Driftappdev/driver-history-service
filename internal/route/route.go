package route

import (
	"net/http"

	"dift_backend_driver/driver-history-service/internal/handler"
)

func Register(mux *http.ServeMux, h *handler.HistoryHandler) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/api/v1/driver/earnings", h.GetEarnings)
	mux.HandleFunc("/api/v1/driver/history", h.GetTripHistory)
}
