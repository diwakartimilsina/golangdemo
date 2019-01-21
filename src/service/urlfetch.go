package service

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func FetchURL(cmd string, args []string) {
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
	fmt.Println(string(cmdOut))
}