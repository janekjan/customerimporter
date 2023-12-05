// This file contatins a simple CLI program that uses
// the customerimporter package. It accepts filepath
// as a command line argument and returns the results
// to the terminal as CSV, which can be piped to a file.
package main

import (
	"fmt"
	"log"
	"github.com/janekjan/customerimporter/customerimporter"
	"os"
)

func main() {
	// Retrieve the CSV filepath from CLI args
	args := os.Args
	if len(args) < 2 {
		log.Fatal("No file path to the CSV file given")
	}

	filepath := args[1]
	domains, err := customerimporter.CustomerCSVToDomainCount(filepath, log.Default())
	if err != nil {
		log.Fatal(err)
	}
	
	for _, domain := range domains {
		fmt.Printf(domain.String())
	}
}


	
