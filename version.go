package semver

import (
	"fmt"
	"strings"
)

// Version contains information about a major version, a minor version,
// a patch version, a set of pre-release identifiers and build data
type Version struct {
	Major      uint
	Minor      uint
	Patch      uint
	PreRelease []string
	BuildData  []string
}

// Parse returns new Version or error if parsing failed
func Parse(src string) (*Version, error) {
	return newParser().parse(src)
}

// String returns a string representation of the Version of the conforming Semantic Versioning Specification.
// See https://semver.org/#is-v123-a-semantic-version for an understanding of skipping leading v
func (v *Version) String() string {
	version := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.PreRelease != nil {
		version += "-" + strings.Join(v.PreRelease, ".")
	}
	if v.BuildData != nil {
		version += "+" + strings.Join(v.BuildData, ".")
	}
	return version
}

// Compare compares the Version to target and returns 1 if target is less or nil, -1 if target is greater and 0 if target is equivalent
func (v *Version) Compare(target *Version) int {
	if target == nil {
		return 1
	}
	if v.Major != target.Major {
		if v.Major > target.Major {
			return 1
		}
		return -1
	}
	if v.Minor != target.Minor {
		if v.Minor > target.Minor {
			return 1
		}
		return -1
	}
	if v.Patch != target.Patch {
		if v.Patch > target.Patch {
			return 1
		}
		return -1
	}
	return v.comparePreRelease(target)
}

// nolint:gocyclo
func (v *Version) comparePreRelease(target *Version) int {
	vLen := len(v.PreRelease)
	targetLen := len(target.PreRelease)
	if vLen == 0 && targetLen > 0 {
		return 1
	}
	if targetLen == 0 && vLen > 0 {
		return -1
	}

	maxLen := vLen
	if maxLen < targetLen {
		maxLen = targetLen
	}

	for i := 0; i < maxLen; i++ {
		if i > vLen-1 {
			return -1
		}
		if i > targetLen-1 {
			return 1
		}

		vIdent := v.PreRelease[i]
		targetIdent := target.PreRelease[i]

		var vIdentIsNumber, targetIdentIsNumber bool
		vIdentNumber, err := strToUint(vIdent, 0, 0)
		if err == nil {
			vIdentIsNumber = true
		}
		targetIdentNumber, err := strToUint(targetIdent, 0, 0)
		if err == nil {
			targetIdentIsNumber = true
		}

		if vIdentIsNumber != targetIdentIsNumber {
			if targetIdentIsNumber {
				return 1
			}
			return -1
		}
		if vIdentIsNumber {
			if vIdentNumber > targetIdentNumber {
				return 1
			} else if vIdentNumber < targetIdentNumber {
				return -1
			}
			return 0
		}
		if vIdent > targetIdent {
			return 1
		} else if vIdent < targetIdent {
			return -1
		}
	}

	return 0
}

func (v *Version) Less(target *Version) bool {
	return v.Compare(target) == -1
}

func (v *Version) LessOrEqual(target *Version) bool {
	cmp := v.Compare(target)
	return cmp == -1 || cmp == 0
}

func (v *Version) Greater(target *Version) bool {
	return v.Compare(target) == 1
}

func (v *Version) GreaterOrEqual(target *Version) bool {
	cmp := v.Compare(target)
	return cmp == 1 || cmp == 0
}

func (v *Version) Equals(target *Version) bool {
	return v.Compare(target) == 0
}
