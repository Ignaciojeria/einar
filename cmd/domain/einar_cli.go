package domain

type EinarCli struct {
	Version              string                `json:"version"`
	Project              string                `json:"project"`
	Installations        []Installation        `json:"installations"`
	InstallationCommands []InstallationCommand `json:"installation_commands"`
}

type Installation struct {
	Name      string   `json:"name"`
	Libraries []string `json:"libraries"`
}

func (c EinarCli) IsInstalled(component string) bool {
	for _, installation := range c.Installations {
		if installation.Name == component {
			return true
		}
	}
	return false
}
