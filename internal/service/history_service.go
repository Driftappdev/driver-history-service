package service

import (
	"time"

	"dift_backend_driver/driver-history-service/internal/dto"
)

type HistoryService struct{}

func NewHistoryService() *HistoryService { return &HistoryService{} }

func (s *HistoryService) GetEarnings(driverID string) dto.DriverEarningsResponse {
	return dto.DriverEarningsResponse{
		DriverID:      driverID,
		Today:         2450.30,
		Week:          12840.20,
		Month:         48690.55,
		Incentives:    1820.00,
		CommissionDue: 3180.40,
		Trend: []dto.DriverEarningsPoint{
			{Label: "Mon", Amount: 1420},
			{Label: "Tue", Amount: 1680},
			{Label: "Wed", Amount: 1510},
			{Label: "Thu", Amount: 1970},
			{Label: "Fri", Amount: 2210},
			{Label: "Sat", Amount: 2450},
			{Label: "Sun", Amount: 2100},
		},
	}
}

func (s *HistoryService) GetTripHistory(driverID string) dto.DriverTripHistoryResponse {
	return dto.DriverTripHistoryResponse{
		DriverID: driverID,
		Items: []dto.DriverTripHistoryItem{
			{RouteID: "route-240401", PassengerName: "Narin P.", Pickup: "Asok, Bangkok", Dropoff: "Silom, Bangkok", Status: "completed", Fare: 168.00, DistanceKm: 8.4, CompletedAt: time.Now().Add(-2 * time.Hour).Format(time.RFC3339)},
			{RouteID: "route-240400", PassengerName: "Mina T.", Pickup: "Ekkamai, Bangkok", Dropoff: "Rama 9, Bangkok", Status: "completed", Fare: 142.50, DistanceKm: 6.9, CompletedAt: time.Now().Add(-4 * time.Hour).Format(time.RFC3339)},
			{RouteID: "route-240399", PassengerName: "Krit C.", Pickup: "Lat Phrao, Bangkok", Dropoff: "Mo Chit, Bangkok", Status: "cancelled", Fare: 0, DistanceKm: 0, CompletedAt: time.Now().Add(-7 * time.Hour).Format(time.RFC3339)},
		},
	}
}
