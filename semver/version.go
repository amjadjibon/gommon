package semver

import (
	"github.com/Masterminds/semver/v3"
)

type Version struct {
	Major      uint64
	Minor      uint64
	Patch      uint64
	PreRelease string
	Metadata   string
}

func NewVersion(v string) (*Version, error) {
	version, err := semver.NewVersion(v)
	if err != nil {
		return nil, err
	}
	return &Version{
		Major:      version.Major(),
		Minor:      version.Minor(),
		Patch:      version.Patch(),
		PreRelease: version.Prerelease(),
		Metadata:   version.Metadata(),
	}, nil
}

func IsSemver(v string) bool {
	version, err := semver.NewVersion(v)
	if err != nil {
		return false
	}
	return version != nil
}

func Constraint(version string, constraint string) bool {
	v, err := semver.NewVersion(version)
	if err != nil {
		return false
	}
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false
	}
	return c.Check(v)
}

func Compare(v1 string, v2 string) int {
	version1, err := semver.NewVersion(v1)
	if err != nil {
		return -2
	}
	version2, err := semver.NewVersion(v2)
	if err != nil {
		return -2
	}
	return version1.Compare(version2)
}
