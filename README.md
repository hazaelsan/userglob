# userglob
User homedir globbing for Go.

## Usage

To perform globbing for an arbitrary user `user`:

```go
import "github.com/hazaelsan/userglob"

if expanded, err := userglob.Glob("~user/some/path"); err == nil {
	fmt.Printf("Expanded path is %v", expanded)
}
```

The following two are similar, but using `userglob` is more resilient:

```go
// $HOME may be manipulated or even missing
expanded := path.Join(os.Getenv("HOME"), "/some/path")

if expanded, err := userglob.Glob("~/some/path"); err == nil {
	fmt.Printf("Expanded path is %v", path)
}

```
