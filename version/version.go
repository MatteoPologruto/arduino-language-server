package version

import (
	"github.com/arduino/arduino-cli/i18n"
)

var (
	defaultVersionString = "0.0.0-git"
	versionString        = ""
	commit               = ""
	status               = ""
	date                 = ""
	tr                   = i18n.Tr
)

// Info FIXMEDOC
type Info struct {
	Application   string `json:"Application"`
	VersionString string `json:"VersionString"`
	Commit        string `json:"Commit"`
	Status        string `json:"Status"`
	Date          string `json:"Date"`
}

// NewInfo FIXMEDOC
func NewInfo(application string) *Info {
	return &Info{
		Application:   application,
		VersionString: versionString,
		Commit:        commit,
		Status:        status,
		Date:          date,
	}
}

func (i *Info) String() string {
	return tr("%[1]s %[2]s Version: %[3]s Commit: %[4]s Date: %[5]s", i.Application, i.Status, i.VersionString, i.Commit, i.Date)
}

//nolint:gochecknoinits
func init() {
	if versionString == "" {
		versionString = defaultVersionString
	}
}
