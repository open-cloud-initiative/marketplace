package main

import (
	"github.com/open-cloud-initiative/marketplace/cmd/cli/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
