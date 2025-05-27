package datasets

import (
	"fmt"
	"os"

	"github.com/gosuri/uiprogress"
)

var (
	progress *uiprogress.Progress
)

func Execute() {
	fmt.Fprintln(os.Stderr, "DEBUG: About to call os.Exit()")
	os.Exit(0)
	fmt.Fprintln(os.Stderr, "DEBUG: Code should never reach here")
}
