package subscription

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Receive func(ctx context.Context, f func(context.Context, *pubsub.Message)) error
