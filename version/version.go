package version

import "fmt"

var (
	version = "dev"
	commit  = "init"
)

func Text() string {
	return fmt.Sprintf("yae version %s (%s)", version, commit)
}
