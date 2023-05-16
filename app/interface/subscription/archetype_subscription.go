package subscription

import (
	"context"
	"encoding/json"

	"archetype/app/shared/archetype/container"
	archetype "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/pubsub/subscription"

	"archetype/app/shared/constants"

	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var archetype_subscription_stop bool = false

type archetype_subscription_struct struct {
	dependencyID     string
	subscriptionName string
}

func archetype_subscription_constructor(
	r subscription.Receive,
	subscriptionName string) (archetype_subscription_struct, error) {
	if archetype_subscription_stop {
		return archetype_subscription_struct{}, nil
	}
	s := archetype_subscription_struct{
		subscriptionName: subscriptionName,
	}
	ctx := context.Background()
	if err := r(ctx, s.receive); err != nil {
		log.
			Error().
			Err(err).
			Str(constants.SUBSCRIPTION_NAME, subscriptionName).
			Msg(constants.SUSBCRIPTION_SIGNAL_BROKEN)
		time.Sleep(10 * time.Second)
		go archetype_subscription_constructor(r, subscriptionName)
		return s, err
	}
	return s, nil
}

func init() {
	const archetype_subscription_name = "INSERT YOUR SUBSCRIPTION NAME"
	dependencyID := uuid.NewString()
	container.InjectComponent(func() error {
		archetype_subscription_pubsub := archetype.Client.Subscription(archetype_subscription_name)
		archetype_subscription_pubsub.ReceiveSettings.Synchronous = true
		archetype_subscription_pubsub.ReceiveSettings.NumGoroutines = 1
		archetype_subscription_pubsub.ReceiveSettings.MaxOutstandingMessages = 1
		s, _ := archetype_subscription_constructor(archetype_subscription_pubsub.Receive, archetype_subscription_name)
		s.dependencyID = dependencyID
		return nil
	}, container.InjectionProps{
		DependencyID: dependencyID,
		Paralel:      true,
	})
}

func (s archetype_subscription_struct) receive(ctx context.Context, m *pubsub.Message) {
	s.processMessage(ctx, m)
}

func (s archetype_subscription_struct) processMessage(ctx context.Context, m *pubsub.Message) error {

	var replace_by_your_model interface{}
	err := json.Unmarshal(m.Data, &replace_by_your_model)

	if err != nil {
		log.
			Error().
			Str(constants.SUBSCRIPTION_NAME, s.subscriptionName).
			Err(err).
			Msg("error unmarshaling m.Data")
		m.Ack()
		return err
	}
	m.Ack()
	return nil
}
