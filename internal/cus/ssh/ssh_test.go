package ssh

import "testing"

func TestCmd(t *testing.T) {
	cases := []struct {
		Name     string
		input    string
		expected string
	}{
		{"echo", "echo CIALLO | tr '[A-Z]' '[a-z]'", "ciallo\n"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if re := Cmd(c.input); re != c.expected {
				t.Fatalf("expected %s, but got %s", c.expected, re)
			}
		})
	}
}
