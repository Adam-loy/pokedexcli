package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.pritnln("Welcome to the FTPS ")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Println("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
