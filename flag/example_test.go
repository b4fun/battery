package flag_test

import (
	"flag"
	"fmt"

	bflag "github.com/b4fun/battery/flag"
)

func ExampleRepeatedStringSlice() {
	fs := flag.NewFlagSet("example", flag.ExitOnError)

	var names bflag.RepeatedStringSlice
	fs.Var(&names, "name", "character names")
	fs.Parse([]string{"-name", "alice", "-name", "bob"})
	fmt.Printf("character names: %q", names)
}
