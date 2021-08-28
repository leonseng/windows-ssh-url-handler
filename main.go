package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	input_ssh_url := os.Args[1]

	ssh_url, err := url.Parse(input_ssh_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL: %s\n", input_ssh_url)
		os.Exit(1)
	}

	if strings.ToLower(ssh_url.Scheme) != "ssh" {
		fmt.Fprintf(os.Stderr, "Invalid SSH URL: %s\n", input_ssh_url)
		os.Exit(1)
	}

	host := ssh_url.Hostname()
	if ssh_url.User.String() != "" {
		host = fmt.Sprintf("%s@%s", ssh_url.User.String(), host)
	}

	port := "22"
	if ssh_url.Port() != "" {
		port = ssh_url.Port()
	}

	fmt.Printf("ssh -p %s %s", port, host)
}
