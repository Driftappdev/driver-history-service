package natsadmin

import (
	"context"
	"encoding/json"
	"strings"

	"dift_backend_driver/driver-history-service/internal/service"

	"github.com/nats-io/nats.go"
)

type AdminConsumer struct {
	svc *service.HistoryService
}

func NewAdminConsumer(svc *service.HistoryService) *AdminConsumer { return &AdminConsumer{svc: svc} }

func (c *AdminConsumer) Subscribe(ctx context.Context, nc *nats.Conn, subject string) error {
	_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		_ = c.handle(msg.Data)
	})
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = nc.Drain()
	}()
	return nil
}

func (c *AdminConsumer) handle(raw []byte) error {
	var cmd struct {
		Action  string         `json:"action"`
		Payload map[string]any `json:"payload"`
	}
	if err := json.Unmarshal(raw, &cmd); err != nil {
		return err
	}
	switch strings.TrimSpace(cmd.Action) {
	case "driver.history.prewarm":
		driverID, _ := cmd.Payload["driver_id"].(string)
		if driverID != "" {
			_ = c.svc.GetTripHistory(driverID)
			_ = c.svc.GetEarnings(driverID)
		}
	}
	return nil
}
