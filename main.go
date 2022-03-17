package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := run()
	log.Printf("Exit with err: %v", err)
	os.Exit(1)
}
func run() error {
	return fmt.Errorf("no implementation, run function test cases to generate figures - see `makefile`")
}
