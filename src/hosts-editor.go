/*
Proposed Commandline flags to be used:
-h -help prints out a list of arguments and what they do
-l -list lists out all the host entries currently in the host file
-a -append [ip] [hostname] Example: -a 127.0.0.1 mattou07.local
-d -delete -del [hostname] Delete an entry in the hosts file
*/
package main

import (
	"bufio"
	"flag"
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

func parseEntry(ent string) entry {
	ip := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	spaces := regexp.MustCompile(`\s*`)
	matchedIP := ip.FindAllString(ent, -1)
	ent = ip.ReplaceAllString(ent, "")
	ent = spaces.ReplaceAllString(ent, "")
	val := entry{ipAddress: matchedIP[0], hostname: ent}
	return val
}

func list(entries []entry) {
	fmt.Println("You have ", len(entries), " entries in your hosts file")
	for _, ent := range entries {
		fmt.Println("IP Address:", ent.ipAddress, "|", "Hostname:", ent.hostname)
	}
}

/*
Since we don't know what order the user will provide the ipaddress and hostname
The arguments in this function may not be accurate
parseEntry function will figure out which variables are the hostname and the ip address and place it into a struct
*/
func add(hostname string, ipAddress string) {
	item := parseEntry(hostname + " " + ipAddress)
	fmt.Println("You entered", item.hostname)
	fmt.Println("You entered", item.ipAddress)
	entry := "\n" + item.ipAddress + "    " + item.hostname
	f, err := os.OpenFile("D:/dev/go-lang/hosts-editor/src/hosts", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	f.WriteString(entry)
	f.Close()
}

func main() {
	//Our commandline flags
	var listArg bool
	var delArg string
	var appendArg string
	flag.BoolVar(&listArg, "list", false, "Lists out all entries in the host file")
	flag.BoolVar(&listArg, "l", false, "Lists out all entries in the host file")
	flag.StringVar(&delArg, "delete", "", "Specifiy a host to be removed from the hosts file")
	flag.StringVar(&delArg, "d", "", "Specifiy a host to be removed from the hosts file")
	flag.StringVar(&appendArg, "append", "", "Specifiy a host and ip to be added to the hosts file")
	flag.StringVar(&appendArg, "a", "", "Specifiy a host and ip to be added to the hosts file")

	flag.Parse()

	// fmt.Println("List:", listArg)
	// fmt.Println("Delete:", delArg)
	// fmt.Println("Append:", appendArg)
	// fmt.Println("Tail:", flag.Args())

	f, err := os.Open("D:/dev/go-lang/hosts-editor/src/hosts")
	check(err)
	defer f.Close()

	var entries []entry

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "#") && scanner.Text() != "" {
			//fmt.Println(scanner.Text())

			item := parseEntry(scanner.Text())
			entries = append(entries, item)
		}

	}
	f.Close()
	if listArg {
		list(entries)
	}
	if appendArg != "" {
		add(appendArg, flag.Args()[0])
	}
}
