package main

import (
	"fmt"
	"service"
)

func main() {
	cmd := "curl"
	args := []string{"google.com"}

	for i := 0; i < 5; i++ {
		go service.FetchURL(cmd, args)
	}
	fmt.Scanln()
	fmt.Println("done")
}