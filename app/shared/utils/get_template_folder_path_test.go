package utils

import "testing"

func TestGetLatestTag(t *testing.T) {
	// Este test asume que la función simplemente extrae el último segmento de una ruta dada,
	// tratándolo como el tag, sin validar su formato como un tag de versión específico.
	tests := []struct {
		name    string
		path    string
		wantTag string
		wantErr bool
	}{
		{
			name:    "simple path with version tag",
			path:    "github.com/username/repository/v1.0.0",
			wantTag: "v1.0.0",
			wantErr: false,
		},
		{
			name:    "windows style path with version tag",
			path:    `C:\path\to\repository\v1.2.3`,
			wantTag: "v1.2.3",
			wantErr: false,
		},
		{
			name:    "unix style path with version tag",
			path:    "/path/to/repository/v2.3.4",
			wantTag: "v2.3.4",
			wantErr: false,
		},
		{
			name:    "path without version tag",
			path:    "github.com/username/repository",
			wantTag: "repository",
			wantErr: false,
		},

		{
			name:    "path without version tag",
			path:    "C:\\Users\\ignac\\go\\bin\\github.com\\Ignaciojeria\\template-tutorial\\v1.0.0",
			wantTag: "v1.0.0",
			wantErr: false,
		},
		{
			name:    "empty path",
			path:    "",
			wantTag: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTag, err := GetLatestTag(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTag != tt.wantTag {
				t.Errorf("GetLatestTag() gotTag = %v, want %v", gotTag, tt.wantTag)
			}
		})
	}
}
