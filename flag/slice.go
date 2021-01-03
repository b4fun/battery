package flag

import (
	"flag"
	"strings"
)

// RepeatedStringSlice parses repeated string slice flags.
//
// var sf RepeatedStringSlice
// fs.Var(&sf, "name", "character names")
//
// in cli:
//   a.out --name foo --name bar
type RepeatedStringSlice []string

func (s *RepeatedStringSlice) String() string {
	if s == nil || len(*s) == 0 {
		return "<empty>"
	}

	return strings.Join([]string(*s), ",")
}

// Set - set value
func (s *RepeatedStringSlice) Set(value string) error {
	*s = append(*s, value)

	return nil
}

// CommaSeparatedStringSlice parses comman separted string flag.
//
// var sf CommaSeparatedStringSlice
// fs.Var(&sf, "name", "character names")
//
// in cli:
//   a.out --name foo,bar
type CommaSeparatedStringSlice []string

func (s *CommaSeparatedStringSlice) String() string {
	if s == nil || len(*s) == 0 {
		return "<empty>"
	}

	return strings.Join([]string(*s), ",")
}

// Set - set value
func (s *CommaSeparatedStringSlice) Set(value string) error {
	// NOTE: we keep *unfiltered* value
	parts := strings.Split(value, ",")

	// NOTE: we keep the lastest value only
	*s = append([]string{}, parts...)

	return nil
}

var (
	_ flag.Value = (*RepeatedStringSlice)(nil)
	_ flag.Value = (*CommaSeparatedStringSlice)(nil)
)
