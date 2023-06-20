package export

import (
	"context"
	"dagger.io/dagger"
)
const output_folder = "output/"
func Output(ctx context.Context, container *dagger.Container) error {
	// get reference to build output directory in container
	output := container.Directory(output_folder)
	// write contents of container build/ directory to the host
	_, err := output.Export(ctx, output_folder)
	if err != nil {
		return err
	}
	return nil
}