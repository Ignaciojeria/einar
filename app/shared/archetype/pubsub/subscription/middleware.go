package subscription

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func Middleware(
	subscriptionName string,
	f func(ctx context.Context, msg *pubsub.Message)) func(context.Context, *pubsub.Message) {
	return func(ctx context.Context, m *pubsub.Message) {
		//PUT YOUR MIDDLEWARE LOGIC
		f(ctx, m)
	}
}
