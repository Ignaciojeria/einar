package in

import "context"

type EinarGenerate func(ctx context.Context, project string, componentKind string, componentName string) error
