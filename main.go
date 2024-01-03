package main

import (
	"fmt"
	"os"

	"github.com/zawazawa0316/designDocConverter/cmd/ddc"
)

func main() {
	cli := &ddc.CLI{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	if err := cli.Run(os.Args); err != nil {
		fmt.Fprintln(cli.Stderr, "Error:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
