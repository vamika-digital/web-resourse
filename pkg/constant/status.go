package constant

import "strings"

type Status int8

const (
	Disable Status = 0
	Enable  Status = 1
	Draft   Status = 2
)

func StatusFromString(statusStr string) Status {
	if strings.ToLower(statusStr) == "enable" {
		return Enable
	} else if strings.ToLower(statusStr) == "disable" {
		return Disable
	} else if strings.ToLower(statusStr) == "draft" {
		return Draft
	} else {
		return Status(-1)
	}
}

func (s Status) AsText() string {
	switch s {
	case Disable:
		return "Disable"
	case Enable:
		return "Enable"
	case Draft:
		return "Draft"
	default:
		return "Unknown"
	}
}

func (s Status) IsValid() bool {
	return !(s < 0 || s > 2)
}
