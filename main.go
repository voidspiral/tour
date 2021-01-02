package main

import (
	"log"
	"tour/cmd"
)

func main() {
	err := cmd.Exercute()
	if err != nil {
		log.Fatalf("cmd.Exercute err: %v", err)
	}
}