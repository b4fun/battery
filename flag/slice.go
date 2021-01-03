package flag

import (
	"flag"
	"strings"
)

// RepeatedStringSlice implements repeated string slice flag.
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

var (
	_ flag.Value = (*RepeatedStringSlice)(nil)
)
