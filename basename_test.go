package main

import (
	"fmt"
	"testing"
)

func TestBasename(t *testing.T) {
	cases := []struct {
		param string
		want  string
	}{
		{param: "a", want: "a"},
		{param: "a.go", want: "a"},
		{param: "a/b/c.go", want: "c"},
		{param: "a/b.c.go", want: "b.c"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("for %q", c.param), func(t *testing.T) {
			got := basename(c.param)

			if got != c.want {
				t.Errorf("got %q want %q", got, c.want)
			}
		})
	}
}
