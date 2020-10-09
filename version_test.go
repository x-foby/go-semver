package semver

import (
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"v0", "v0.0.0"},
		{"0", "v0.0.0"},
		{"v0.1", "v0.1.0"},
		{"0.1", "v0.1.0"},
		{"v0.1.2", "v0.1.2"},
		{"0.1.2", "v0.1.2"},
		{"v0.1.2-alpha.beta", "v0.1.2-alpha.beta"},
		{"0.1.2-alpha.beta", "v0.1.2-alpha.beta"},
		{"v0.1.2-alpha.beta+a-b.c", "v0.1.2-alpha.beta+a-b.c"},
		{"0.1.2-alpha.beta+a-b.c", "v0.1.2-alpha.beta+a-b.c"},
		{"v0.1.2+a-b.c", "v0.1.2+a-b.c"},
		{"0.1.2+a-b.c", "v0.1.2+a-b.c"},

		{"v0.01.0", "v0.1.0"},
		{"v0.01.00", "v0.1.0"},
	}

	for _, c := range cases {
		v, err := Parse(c.in)
		if err != nil {
			t.Errorf("Parse(%s) = (nil, %s), want (%s, nil)", c.in, err, c.want)
			continue
		}
		if got := v.String(); got != c.want {
			t.Errorf("Parse(%s) = (%s, nil), want (%s, nil)", c.in, got, c.want)
		}
	}
}

// nolint:dupl
func TestLess(t *testing.T) {
	compare(
		t,
		func(a, b *Version) bool { return a.Less(b) },
		"Less",
		[]compareCase{
			{"0.1.0", "1.0.1", true},
			{"0.0.1", "0.1.0", true},
			{"0.1.0", "0.1.1", true},
			{"1.0.0-alpha", "1.0.0-alpha.1", true},
			{"1.0.0-alpha.1", "1.0.0-alpha.beta", true},
			{"1.0.0-alpha.beta", "1.0.0-beta", true},
			{"1.0.0-beta", "1.0.0-beta.2", true},
			{"1.0.0-beta.2", "1.0.0-beta.11", true},
			{"1.0.0-beta.11", "1.0.0-rc.1", true},
			{"1.0.0-rc.1", "1.0.0", true},

			{"1.0.0", "0.0.1", false},
			{"1.0.0", "1.0.0", false},
			{"1.1.0", "1.0.0", false},
			{"1.1.0", "1.1.0", false},
			{"1.1.1", "1.1.0", false},
			{"1.1.1", "1.1.1", false},
			{"1.0.0-alpha.1", "1.0.0-alpha", false},
			{"1.0.0-alpha.beta", "1.0.0-alpha.1", false},
			{"1.0.0-beta", "1.0.0-alpha.beta", false},
			{"1.0.0-beta.2", "1.0.0-beta", false},
			{"1.0.0-beta.11", "1.0.0-beta.2", false},
			{"1.0.0-rc.1", "1.0.0-beta.11", false},
			{"1.0.0", "1.0.0-rc.1", false},
		},
	)
}

func TestLessOrEqual(t *testing.T) {
	compare(
		t,
		func(a, b *Version) bool { return a.LessOrEqual(b) },
		"LessOrEqual",
		[]compareCase{
			{"0.1.0", "1.0.1", true},
			{"0.1.0", "0.1.0", true},
			{"0.0.1", "0.1.0", true},
			{"0.0.1", "0.0.1", true},
			{"0.1.0", "0.1.1", true},
			{"0.1.0", "0.1.0", true},
			{"1.0.0-alpha", "1.0.0-alpha.1", true},
			{"1.0.0-alpha.1", "1.0.0-alpha.1", true},
			{"1.0.0-alpha.1+a", "1.0.0-alpha.1", true},
			{"1.0.0-alpha.1", "1.0.0-alpha.beta", true},
			{"1.0.0-alpha.beta", "1.0.0-beta", true},
			{"1.0.0-beta", "1.0.0-beta.2", true},
			{"1.0.0-beta.2", "1.0.0-beta.11", true},
			{"1.0.0-beta.11", "1.0.0-rc.1", true},
			{"1.0.0-rc.1", "1.0.0", true},

			{"1.0.0", "0.0.1", false},
			{"1.1.0", "1.0.0", false},
			{"1.1.1", "1.1.0", false},
			{"1.0.0-alpha.1", "1.0.0-alpha", false},
			{"1.0.0-alpha.beta", "1.0.0-alpha.1", false},
			{"1.0.0-beta", "1.0.0-alpha.beta", false},
			{"1.0.0-beta.2", "1.0.0-beta", false},
			{"1.0.0-beta.11", "1.0.0-beta.2", false},
			{"1.0.0-rc.1", "1.0.0-beta.11", false},
			{"1.0.0", "1.0.0-rc.1", false},
		},
	)
}

// nolint:dupl
func TestGreater(t *testing.T) {
	compare(t,
		func(a, b *Version) bool { return a.Greater(b) },
		"Greater",
		[]compareCase{
			{"0.1.0", "1.0.1", false},
			{"0.1.0", "0.1.0", false},
			{"0.0.1", "0.1.0", false},
			{"0.0.1", "0.0.1", false},
			{"0.1.0", "0.1.1", false},
			{"0.1.0", "0.1.0", false},
			{"1.0.0-alpha.1", "1.0.0-alpha", true},
			{"1.0.0-alpha.beta", "1.0.0-alpha.1", true},
			{"1.0.0-beta", "1.0.0-alpha.beta", true},
			{"1.0.0-beta.2", "1.0.0-beta", true},
			{"1.0.0-beta.11", "1.0.0-beta.2", true},
			{"1.0.0-rc.1", "1.0.0-beta.11", true},
			{"1.0.0", "1.0.0-rc.1", true},

			{"1.0.0", "0.0.1", true},
			{"1.1.0", "1.0.0", true},
			{"1.1.1", "1.1.0", true},
			{"1.0.0-alpha", "1.0.0-alpha.1", false},
			{"1.0.0-alpha.1", "1.0.0-alpha.beta", false},
			{"1.0.0-alpha.beta", "1.0.0-beta", false},
			{"1.0.0-beta", "1.0.0-beta.2", false},
			{"1.0.0-beta.2", "1.0.0-beta.11", false},
			{"1.0.0-beta.11", "1.0.0-rc.1", false},
			{"1.0.0-rc.1", "1.0.0", false},
		},
	)
}

func TestGreaterOrEqual(t *testing.T) {
	compare(t,
		func(a, b *Version) bool { return a.GreaterOrEqual(b) },
		"GreaterOrEqual",
		[]compareCase{
			{"0.1.0", "1.0.1", false},
			{"0.0.1", "0.1.0", false},
			{"0.1.0", "0.1.1", false},
			{"1.0.0-alpha.1", "1.0.0-alpha", true},
			{"1.0.0-alpha.1", "1.0.0-alpha.1", true},
			{"1.0.0-alpha.1+a", "1.0.0-alpha.1", true},
			{"1.0.0-alpha.beta", "1.0.0-alpha.1", true},
			{"1.0.0-beta", "1.0.0-alpha.beta", true},
			{"1.0.0-beta.2", "1.0.0-beta", true},
			{"1.0.0-beta.11", "1.0.0-beta.2", true},
			{"1.0.0-rc.1", "1.0.0-beta.11", true},
			{"1.0.0", "1.0.0-rc.1", true},

			{"1.0.0", "0.0.1", true},
			{"1.0.0", "1.0.0", true},
			{"1.1.0", "1.0.0", true},
			{"1.1.0", "1.1.0", true},
			{"1.1.1", "1.1.0", true},
			{"1.1.1", "1.1.1", true},
			{"1.0.0-alpha", "1.0.0-alpha.1", false},
			{"1.0.0-alpha.1", "1.0.0-alpha.beta", false},
			{"1.0.0-alpha.beta", "1.0.0-beta", false},
			{"1.0.0-beta", "1.0.0-beta.2", false},
			{"1.0.0-beta.2", "1.0.0-beta.11", false},
			{"1.0.0-beta.11", "1.0.0-rc.1", false},
			{"1.0.0-rc.1", "1.0.0", false},
		},
	)
}

func TestEquals(t *testing.T) {
	compare(t,
		func(a, b *Version) bool { return a.Equals(b) },
		"Equals",
		[]compareCase{
			{"0.0.1", "0.0.2", false},
			{"0.1.0", "0.2.0", false},
			{"1.0.0", "2.0.0", false},

			{"1.0.0", "1.0.0", true},
			{"0.1.0", "0.1.0", true},
			{"0.1.1", "0.1.1", true},
			{"0.1.1+a", "0.1.1+b", true},
		},
	)
}

type compareCase struct {
	a, b string
	want bool
}

func compare(t *testing.T, cb func(a, b *Version) bool, funcName string, cases []compareCase) {
	for _, c := range cases {
		a, err := Parse(c.a)
		if err != nil {
			t.Errorf("Parse(%s) = (nil, %s), want (not nil, nil)", c.a, err)
			continue
		}
		b, err := Parse(c.b)
		if err != nil {
			t.Errorf("Parse(%s) = (nil, %s), want (not nil, nil)", c.b, err)
			continue
		}
		if got := cb(a, b); got != c.want {
			t.Errorf("(%s).%s(%s) = (%v), want (%v)", a.String(), funcName, b.String(), got, c.want)
		}
	}
}
