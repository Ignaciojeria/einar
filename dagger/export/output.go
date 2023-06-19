package export

import (
	"context"
	"github.com/Ignaciojeria/einar/dagger/constants"
	"dagger.io/dagger"
)


func Output(ctx context.Context, container *dagger.Container) error {
	// get reference to build output directory in container
	output := container.Directory(constants.PATH)
	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, constants.PATH)
	if err != nil {
		return err
	}
	return nil
}