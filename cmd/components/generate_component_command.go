package components

func InstallCommand(project string, componentKind string, componentName string) error {
	return addComponentInsideCli(project, componentKind, componentName)
}
