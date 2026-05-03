package handler

import (
	"net/http"

	response "github.com/PlatformCore/engine-core/core/contracts/response"

	"dift_backend_driver/driver-history-service/internal/service"
)

type HistoryHandler struct{ svc *service.HistoryService }

func NewHistoryHandler(svc *service.HistoryService) *HistoryHandler { return &HistoryHandler{svc: svc} }

func (h *HistoryHandler) GetEarnings(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: "driver_id required"}})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetEarnings(driverID)})
}

func (h *HistoryHandler) GetTripHistory(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{Error: &response.AppError{Code: "bad_request", Message: "driver_id required"}})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetTripHistory(driverID)})
}
