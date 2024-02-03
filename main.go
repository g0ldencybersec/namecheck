package main

import (
	"fmt"
	"os"

	"github.com/g0ldencybersec/namecheck/github"
	"github.com/g0ldencybersec/namecheck/reddit"
)

func main() {
	username := os.Args[1]

	// Check GitHub Availability
	fmt.Println(github.IsValid(username))
	if github.IsValid(username) {
		ghResult, err := github.IsAvailable(username)
		if err != nil {
			fmt.Println()
		}
		fmt.Println(ghResult)
	}

	// Check Reddit Availability
	fmt.Println(reddit.IsValid(username))
	if reddit.IsValid(username) {
		redditResult, err := reddit.IsAvailable(username)
		if err != nil {
			fmt.Println()
		}
		fmt.Println(redditResult)
	}
}
