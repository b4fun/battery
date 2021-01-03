package flag

import (
	"flag"
	"testing"
)

func createFS(v flag.Value) *flag.FlagSet {
	fs := flag.NewFlagSet("test-fs", flag.ExitOnError)
	fs.Var(v, "name", "names")

	return fs
}

func TestRepeatedStringSlice(t *testing.T) {
	cases := []struct {
		input          []string
		expectedValue  []string
		expectedString string
		reason         string
	}{
		{
			input:          nil,
			expectedValue:  nil,
			expectedString: "<empty>",
			reason:         "empty value",
		},
		{
			input:          []string{"-name=foo"},
			expectedValue:  []string{"foo"},
			expectedString: "foo",
			reason:         "single value",
		},
		{
			input:          []string{"-name=foo", "-name=bar"},
			expectedValue:  []string{"foo", "bar"},
			expectedString: "foo,bar",
			reason:         "multiple values",
		},
		{
			input:          []string{"-name=foo", "-name=bar", "-name=foo"},
			expectedValue:  []string{"foo", "bar", "foo"},
			expectedString: "foo,bar,foo",
			reason:         "duplicated values",
		},
	}

	for _, c := range cases {
		t.Logf("running case: %s", c.reason)
		var v RepeatedStringSlice
		fs := createFS(&v)
		err := fs.Parse(c.input)
		if err != nil {
			t.Errorf("unexpected flag parse error: %s", err)
		}

		if actual := v.String(); actual != c.expectedString {
			t.Errorf("string output, expected=%s, got=%s", c.expectedString, actual)
		}
		if c.expectedValue != nil {
			actualValue := []string(v)
			if len(actualValue) != len(c.expectedValue) {
				t.Errorf("value, expected=%q, got=%q", c.expectedValue, actualValue)
			}
			for i, v := range actualValue {
				if v != c.expectedValue[i] {
					t.Errorf("value, expected=%q, got=%q", c.expectedValue, actualValue)
				}
			}
		}
	}
}
