//go:generate ${GO_PATH}/bin/goversioninfo -manifest=../../resource/goversioninfo.exe.manifest -64 -product-version=${VER} -file-version=${VER}

package main

import (
	"datasets_cli/v2/datasets"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "DEBUG: Starting main()")
	datasets.Execute()
	os.Exit(0) // Exit with status code 0
	fmt.Fprintln(os.Stderr, "DEBUG: This should never print")
}
