package dto

type DriverEarningsPoint struct {
	Label  string  `json:"label"`
	Amount float64 `json:"amount"`
}

type DriverEarningsResponse struct {
	DriverID      string                `json:"driver_id"`
	Today         float64               `json:"today"`
	Week          float64               `json:"week"`
	Month         float64               `json:"month"`
	Incentives    float64               `json:"incentives"`
	CommissionDue float64               `json:"commission_due"`
	Trend         []DriverEarningsPoint `json:"trend"`
}

type DriverTripHistoryItem struct {
	RouteID       string  `json:"route_id"`
	PassengerName string  `json:"passenger_name"`
	Pickup        string  `json:"pickup"`
	Dropoff       string  `json:"dropoff"`
	Status        string  `json:"status"`
	Fare          float64 `json:"fare"`
	DistanceKm    float64 `json:"distance_km"`
	CompletedAt   string  `json:"completed_at"`
}

type DriverTripHistoryResponse struct {
	DriverID string                  `json:"driver_id"`
	Items    []DriverTripHistoryItem `json:"items"`
}
