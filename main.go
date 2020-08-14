package main

import (
	"log"

	"github.com/mrcampbell/projects/go/cli-util/internal/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
