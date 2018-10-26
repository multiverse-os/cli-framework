package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli-framework"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name: "example",
		//Version: cli.Version{Major: 0, Minor: 1, Patch: 1},
		Usage: "make an explosive entrance",
		DefaultAction: func(c *cli.Context) error {
			fmt.Println("Example output for an action (or command)!")
			fmt.Println("version is: ", c.CLI.Version.ANSIFormattedString())
			return nil
		},
	})

	cmd.Run(os.Args)
}
