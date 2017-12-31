// Package userglob provides user homedir globbing expansions.
package userglob

import (
	"os/user"
	"regexp"
)

var globRE = regexp.MustCompile(`^~(.*?)(/.+)?$`)

// Glob expands ~user/dir paths.
func Glob(path string) (string, error) {
	if m := globRE.FindStringSubmatch(path); m != nil {
		var u *user.User
		var err error
		if m[1] == "" {
			u, err = user.Current()
		} else {
			u, err = user.Lookup(m[1])
		}
		if err != nil {
			return "", err
		}
		path = u.HomeDir + m[2]
	}
	return path, nil
}
