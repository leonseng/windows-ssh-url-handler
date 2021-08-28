package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	ssh_url := os.Args[1]
	target := ssh_url[6:] // strip ssh:// prefix

	host := target
	port := "22"
	port_match := regexp.MustCompile(`:\d+$`)

	// port provided in URL
	if port_match.MatchString(target) {
		url_match := regexp.MustCompile(`^(?P<host>.+?):(?P<port>\d+)$`)
		match := url_match.FindStringSubmatch(target)
		result := make(map[string]string)
		for i, name := range url_match.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		host = result["host"]
		port = result["port"]
	}

	fmt.Printf("ssh -p %s %s", port, strings.TrimRight(host, "/"))
}
