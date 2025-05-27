//go:generate ${GO_PATH}/bin/goversioninfo -manifest=../../resource/goversioninfo.exe.manifest -64 -product-version=${VER} -file-version=${VER}

package main

import (
	"fmt"
	"os"

	"github.com/gosuri/uiprogress"
)

var (
	progress *uiprogress.Progress
)

func main() {
	fmt.Fprintln(os.Stderr, "DEBUG: Exiting the program immediately")
	os.Exit(0) // Hangs here on Apple silicon macOS >=12 on go>=1.21 with gosuri/uilive v0.0.4
}
