package out

import "context"

type Shutdown func(ctx context.Context) error
