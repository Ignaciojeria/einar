package subscription

import (
	"context"

	"cloud.google.com/go/pubsub"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type config struct {
	serviceName string
}

// A ReceiveOption is used to customize spans started by WrapReceiveHandler.
type ReceiveOption func(cfg *config)

// WithServiceName sets the service name tag for traces started by WrapReceiveHandler.
func WithServiceName(serviceName string) ReceiveOption {
	return func(cfg *config) {
		cfg.serviceName = serviceName
	}
}

// WrapReceiveHandler returns a receive handler that wraps the supplied handler,
// extracts any tracing metadata attached to the received message, and starts a
// receive span.
// WrapReceiveHandler returns a receive handler that wraps the supplied handler,
// extracts any tracing metadata attached to the received message, and starts a
// receive span.
func Middleware(subscriptionName string, f func(context.Context, *pubsub.Message), opts ...ReceiveOption) func(context.Context, *pubsub.Message) {

	var cfg config
	for _, opt := range opts {
		opt(&cfg)
	}
	//log.Debug("contrib/cloud.google.com/go/pubsub.v1: Wrapping Receive Handler: %#v", cfg)
	return func(ctx context.Context, msg *pubsub.Message) {
		parentSpanCtx, _ := tracer.Extract(tracer.TextMapCarrier(msg.Attributes))
		opts := []ddtrace.StartSpanOption{
			tracer.ResourceName(subscriptionName),
			tracer.SpanType(ext.SpanTypeMessageConsumer),
			tracer.Tag("message_size", len(msg.Data)),
			tracer.Tag("num_attributes", len(msg.Attributes)),
			//tracer.Tag("ordering_key", msg.OrderingKey),
			tracer.Tag("message_id", msg.ID),
			tracer.Tag("publish_time", msg.PublishTime.String()),
			tracer.ChildOf(parentSpanCtx),
		}
		if cfg.serviceName != "" {
			opts = append(opts, tracer.ServiceName(cfg.serviceName))
		}
		span, ctx := tracer.StartSpanFromContext(ctx, "pubsub.receive", opts...)
		if msg.DeliveryAttempt != nil {
			span.SetTag("delivery_attempt", *msg.DeliveryAttempt)
		}
		defer span.Finish()
		f(ctx, msg)
	}
}
