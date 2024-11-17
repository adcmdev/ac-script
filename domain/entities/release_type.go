package entities

type TulReleaseType int

const (
	Debug TulReleaseType = iota
	Bugfix
	Minor
	Major
)

func TulReleaseTypeFromString(arg string) TulReleaseType {
	switch arg {
	case "debug":
		return Debug
	case "bugfix":
		return Bugfix
	case "minor":
		return Minor
	case "major":
		return Major
	}

	return Bugfix
}
