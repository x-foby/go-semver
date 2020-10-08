package semver

import (
	"fmt"
	"strings"
)

type Version struct {
	Major      uint
	Minor      uint
	Patch      uint
	PreRelease []string
	BuildData  []string
}

func Parse(src string) (*Version, error) {
	return newParser().parse(src)
}

func (v *Version) String() string {
	version := fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.PreRelease != nil {
		version += "-" + strings.Join(v.PreRelease, ".")
	}
	if v.BuildData != nil {
		version += "+" + strings.Join(v.BuildData, ".")
	}
	return version
}

func (v *Version) Less(target *Version) bool {
	if target == nil {
		return false
	}
	if v.Major != target.Major {
		return v.Major < target.Major
	}
	if v.Minor != target.Minor {
		return v.Minor < target.Minor
	}
	if v.Patch != target.Patch {
		return v.Patch < target.Patch
	}
	// TODO: pre-release compare
	return false
}

func (v *Version) LessOrEqual(target *Version) bool {
	if target == nil {
		return false
	}
	if v.Major != target.Major {
		return v.Major < target.Major
	}
	if v.Minor != target.Minor {
		return v.Minor < target.Minor
	}
	if v.Patch != target.Patch {
		return v.Patch < target.Patch
	}
	// TODO: pre-release compare
	return true
}

func (v *Version) Greater(target *Version) bool {
	if target == nil {
		return false
	}
	if v.Major != target.Major {
		return v.Major > target.Major
	}
	if v.Minor != target.Minor {
		return v.Minor > target.Minor
	}
	if v.Patch != target.Patch {
		return v.Patch > target.Patch
	}
	// TODO: pre-release compare
	return false
}

func (v *Version) GreaterOrEqual(target *Version) bool {
	if target == nil {
		return false
	}
	if v.Major != target.Major {
		return v.Major > target.Major
	}
	if v.Minor != target.Minor {
		return v.Minor > target.Minor
	}
	if v.Patch != target.Patch {
		return v.Patch > target.Patch
	}
	// TODO: pre-release compare
	return true
}

func (v *Version) Equals(target *Version) bool {
	if target == nil {
		return false
	}
	if v.Major != target.Major {
		return false
	}
	if v.Minor != target.Minor {
		return false
	}
	if v.Patch != target.Patch {
		return false
	}
	// TODO: pre-release compare
	return true
}
