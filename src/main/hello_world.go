package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := "curl"
	args := []string{"google.com"}

	for i := 0; i < 500; i++ {
		go fetchURL(cmd, args)
	}
	fmt.Scanln()
	fmt.Println("done")
}

func fetchURL(cmd string, args []string) {
	var (
		err    error
		cmdOut []byte
	)

	fmt.Println(time.Now())
	fmt.Println("Command: " + cmd)

	cmdOut, err = exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("I am here")
	fmt.Println(string(cmdOut))
}
