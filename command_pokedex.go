package main

import "fmt"

func commandPokedex(conf *config, unused...string) error {
	
	if len(unused) > 0 {
		fmt.Println("The pokedex command doesn't take arguments.")
	}

	if len(conf.pokedex) < 1 {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key := range conf.pokedex {
		fmt.Printf("- %s\n", key)
	}

	return nil
}