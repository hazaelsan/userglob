package userglob

import (
	"fmt"
	"os/user"
	"path"
	"testing"
)

func TestExpandTilde(t *testing.T) {
	root, err := user.LookupId("0")
	if err != nil {
		t.Fatal(err)
	}
	tilde := fmt.Sprintf("~%v", root.Username)
	u, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		want string
		err  bool
	}{
		path.Join(tilde, "foo"): {want: path.Join(root.HomeDir, "foo")}, // ~root/path
		tilde:           {want: root.HomeDir}, // ~root
		"~/foo":         {want: path.Join(u.HomeDir, "foo")},
		"~":             {want: u.HomeDir},
		"/not/a/glob":   {want: "/not/a/glob"},
		"~1invalid":     {err: true},
		"~1invalid/foo": {err: true},
		"":              {},
	}
	for glob, tt := range tests {
		got, err := Glob(glob)
		if err != nil {
			if !tt.err {
				t.Errorf("Glob(%v) error = %v", glob, err)
			}
			continue
		}
		if got != tt.want {
			t.Errorf("Glob(%v) = %v, want %v", glob, got, tt.want)
		}
	}
}
