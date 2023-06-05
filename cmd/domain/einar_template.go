package domain

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

type EinarTemplate struct {
	BaseTemplate         BaseTemplate          `json:"base_template"`
	InstallationsBase    []InstallationsBase   `json:"installations_base"`
	InstallationCommands []InstallationCommand `json:"installation_commands"`
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

type ComponentGenerationCommand struct {
	Name string
}

type ComponentFile struct {
	SourceFile     string          `json:"source_file"`
	DestinationDir string          `json:"destination_dir"`
	DependsOn      string          `json:"depends_on"`
	ReplaceHolders []ReplaceHolder `json:"replace_holders"`
}

type ReplaceHolder struct{
	
}
