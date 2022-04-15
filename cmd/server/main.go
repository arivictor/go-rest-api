package main

import (
	"fmt"
	"log"
)

// Responsible for instantiation of startup of application
func Run() error {
	fmt.Println("Hello World")
	return nil
}

func main() {
	err := Run()
	if err != nil {
		log.Fatalln(err)
	}
}
