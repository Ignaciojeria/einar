package nats_publihser

import (
	"context"
	"encoding/json"

	"github.com/Ignaciojeria/einar/app/domain"
	einar "github.com/Ignaciojeria/einar/app/shared/archetype/nats"
)

var PublishInstruction = func(ctx context.Context, e domain.Instruction) error {
	b, _ := json.Marshal(e)
	return einar.Conn.Publish(einar.EinarTopic, b)
}
