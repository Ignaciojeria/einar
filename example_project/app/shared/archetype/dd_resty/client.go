package resty

import (
	"time"

	"github.com/go-resty/resty/v2"

	"net/http"

	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

var Client *resty.Client = newRestyClientWithResourceName()

func newRestyClientWithResourceName() *resty.Client {
	client := resty.New()
	// Set the custom transport for the Resty client
	client.SetTransport(createTracedTransport())
	// Set timeout to 10 seconds
	client.SetTimeout(10 * time.Second)
	return client
}

func createTracedTransport() http.RoundTripper {
	// Get the Resty client's default transport
	transport := http.DefaultTransport
	// Wrap the transport with Datadog's tracing transport
	tracedTransport := ddhttp.WrapRoundTripper(
		transport,
		ddhttp.RTWithResourceNamer(func(req *http.Request) string {
			return req.URL.Path
		}),
	)
	return tracedTransport
}
