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

var __archetype_subscription_stop bool = false

type __archetype_subscription_struct struct {
	subscriptionName string
}

func __archetype_subscription_constructor(
	r subscription.Receive,
	subscriptionName string) (__archetype_subscription_struct, error) {
	if __archetype_subscription_stop {
		return __archetype_subscription_struct{}, nil
	}
	s := __archetype_subscription_struct{
		subscriptionName: subscriptionName,
	}
	ctx := context.Background()
	if err := r(ctx, subscription.Middleware(subscriptionName, s.receive)); err != nil {
		log.
			Error().
			Err(err).
			Str(constants.SUBSCRIPTION_NAME, subscriptionName).
			Msg(constants.SUSBCRIPTION_SIGNAL_BROKEN)
		time.Sleep(10 * time.Second)
		go __archetype_subscription_constructor(r, subscriptionName)
		return s, err
	}
	return s, nil
}

func init() {
	const subscription_name = "INSERT YOUR SUBSCRIPTION NAME"
	container.InjectComponent(func() error {
		subscription_setup := archetype.Client.Subscription(subscription_name)
		subscription_setup.ReceiveSettings.Synchronous = true
		subscription_setup.ReceiveSettings.NumGoroutines = 1
		subscription_setup.ReceiveSettings.MaxOutstandingMessages = 1
		__archetype_subscription_constructor(subscription_setup.Receive, subscription_name)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
		Paralel:      true,
	})
}

func (s __archetype_subscription_struct) receive(ctx context.Context, m *pubsub.Message) {
	s.processMessage(ctx, m)
}

func (s __archetype_subscription_struct) processMessage(ctx context.Context, m *pubsub.Message) error {

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
