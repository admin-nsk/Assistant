package version

import (
	"strings"
	"sync"
)

const (
	BuildTypeProduction  BuildType = "production"
	BuildTypeDevelopment BuildType = "development"
	BuildTypeStage       BuildType = "stage"
)

type BuildType = string

var (
	Version   = "0.0.0"              //nolint:gochecknoglobals
	Commit    = ""                   //nolint:gochecknoglobals
	Build     = BuildTypeDevelopment //nolint:gochecknoglobals
	AppName   = "note_service"       //nolint:gochecknoglobals
	EnvPrefix = "NOTE_SERVICE"       //nolint:gochecknoglobals
)

func IsDevelopment() bool {
	return strings.ToLower(Build) == BuildTypeDevelopment
}

//nolint:gochecknoglobals
var (
	once        sync.Once
	fullVersion string
)

func FullVersion() string {
	once.Do(func() {
		fversion := &strings.Builder{}
		fversion.WriteString(Version)

		if strings.ToLower(Build) == BuildTypeDevelopment {
			fversion.WriteString("dev")
		} else if strings.ToLower(Build) == BuildTypeStage {
			fversion.WriteString("stage")
		}

		if Commit != "" {
			fversion.WriteString("/")
			fversion.WriteString(Commit)
		}
		fullVersion = fversion.String()
	})
	return fullVersion
}
