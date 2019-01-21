package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sync"
    "time"
    "os/exec"
    "strconv"

)

var wg sync.WaitGroup

func main() {
	
	fmt.Println("Started at: ")
	fmt.Println(time.Now())
	
	counter := 0
	
	file, err := os.Open("/tmp/zipcodes.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	url := scanner.Text()
	
		fmt.Println("Counter " + strconv.Itoa(counter))
	    
	    wg.Add(1)
	    
		go FetchURL(url)
		counter++
		
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    wg.Wait()

	fmt.Println("done")
	fmt.Println(time.Now())
}

func FetchURL(url string) {
	var (
		err    error
		cmdOut []byte
	)


	cmd := "curl"
	
	fmt.Println(time.Now())

	cmdOut, err = exec.Command(cmd, "http://www.google.com/" + url).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(string(cmdOut))
	
	wg.Done()
}