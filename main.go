package main

import (
	"github.com/sieradzkiM/blockchain-golang/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
