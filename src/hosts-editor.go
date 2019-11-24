package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

//TODO handle lines with empty space
func parseEntry(ent string) entry {
	ip := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	spaces := regexp.MustCompile(`\s*`)
	submatchall := ip.FindAllString(ent, -1)
	ent = ip.ReplaceAllString(ent, "")
	ent = spaces.ReplaceAllString(ent, "")
	val := entry{ipAddress: submatchall[0], hostname: ent}
	fmt.Println(val.ipAddress)
	fmt.Println(val.hostname)
	return val
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	f, err := os.Open("E:/dev/go-lang/hosts-editior/src/hosts")
	check(err)
	defer f.Close()

	var entries []entry

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "#") && scanner.Text() != "" {
			//fmt.Println(scanner.Text())
			parseEntry(scanner.Text())
			item := entry{"", ""}
			entries = append(entries, item)
		}

	}
}
