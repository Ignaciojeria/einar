package echo_server

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type RoutingState struct {
	Context     map[string]string
	FlatContext string
}

func NewRoutingState(ctx echo.Context, m map[string]string) RoutingState {
	headerFlatContext := ctx.Request().Header.Get("FlatContext")

	routingState := RoutingState{
		Context: m,
	}

	// Flattening the Context map to a string in the format "key:value|key:value"
	var flatContextParts []string
	for k, v := range m {
		flatContextParts = append(flatContextParts, fmt.Sprintf("%s:%s", k, v))
	}
	flatContextString := strings.Join(flatContextParts, "|")

	if headerFlatContext != "" {
		// If FlatContext is present in the header, use the original format
		routingState.FlatContext = `{"FlatContext":"` + headerFlatContext + `"}`

		// Split the FlatContext by "|" to get key-value pairs
		pairs := strings.Split(headerFlatContext, "|")

		// Iterate over pairs and split by ":" to separate keys and values
		for _, pair := range pairs {
			kv := strings.Split(pair, ":")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}

		routingState.Context = m

	} else {
		// If FlatContext is not present in the header, use the flattened format of map m
		routingState.FlatContext = `{"FlatContext":"` + flatContextString + `"}`

	}

	return routingState
}
