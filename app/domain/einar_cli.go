package domain

type EinarCli struct {
	Version       string         `json:"version"`
	Project       string         `json:"project"`
	Template 	  Template 		 `json:"template"`
	Installations []Installation `json:"installations"`
	Components    []Component    `json:"components"`
}

type Template struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Installation struct {
	Name      string   `json:"name"`
	Libraries []string `json:"libraries"`
}

type Component struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func (c EinarCli) IsInstalled(component string) bool {
	for _, installation := range c.Installations {
		if installation.Name == component {
			return true
		}
	}
	return false
}
