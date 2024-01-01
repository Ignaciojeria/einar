package in

import "context"

type EinarInit func(ctx context.Context, templateFilePath string, project string) error
