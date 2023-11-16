package in

import "context"

type EinarInstall func(ctx context.Context, project, commandName string) error
