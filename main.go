package main

import (
	"os"

	"github.com/condezero/go-cryptocurrency-scratch/client"
)

func main() {
	defer os.Exit(0)

	cmd := client.CommandLine{}
	cmd.Run()

}
