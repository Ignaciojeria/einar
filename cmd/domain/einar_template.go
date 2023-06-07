package domain

type EinarTemplate struct {
	BaseTemplate         BaseTemplate          `json:"base_template"`
	InstallationsBase    []InstallationsBase   `json:"installations_base"`
	InstallationCommands []InstallationCommand `json:"installation_commands"`
	ComponentCommands    []ComponentCommands   `json:"component_commands"`
}

type BaseTemplate struct {
	Description string   `json:"description"`
	Folders     []Folder `json:"folders"`
	Files       []File   `json:"files"`
}

type Folder struct {
	SourceDir      string `json:"source_dir"`
	DestinationDir string `json:"destination_dir"`
}

type File struct {
	SourceFile      string `json:"source_file"`
	DestinationFile string `json:"destination_file"`
}

type InstallationCommand struct {
	Name           string   `json:"name"`
	SourceDir      string   `json:"source_dir"`
	DestinationDir string   `json:"destination_dir"`
	Command        string   `json:"command"`
	Libraries      []string `json:"libraries"`
}

type InstallationsBase struct {
	Name    string `json:"name"`
	Library string `json:"library"`
}

type ComponentCommands struct {
	Kind           string          `json:"kind"`
	Name           string          `json:"name"`
	ComponentFiles []ComponentFile `json:"files"`
}

type ComponentFile struct {
	SourceFile     string          `json:"source_file"`
	DestinationDir string          `json:"destination_dir"`
	IocDiscovery   bool            `json:"ioc_discovery"`
	Port           Port            `json:"port"`
	ReplaceHolders []ReplaceHolder `json:"replace_holders"`
}

type Port struct {
	SourceFile     string `json:"source_file"`
	DestinationDir string `json:"destination_dir"`
}

type ReplaceHolder struct {
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	AppendAtStart string `json:"append_at_start"`
	AppendAtEnd   string `json:"append_at_end"`
}
