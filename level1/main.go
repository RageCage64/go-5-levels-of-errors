package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("not_real.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
