package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type entry struct {
	ipAddress string
	hostname  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./hosts")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "#") {
			fmt.Println(scanner.Text())
		}

	}

}
