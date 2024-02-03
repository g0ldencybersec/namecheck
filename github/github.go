package github

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func IsValid(username string) bool {
	// Check for starting/ending with a hyphen
	if strings.HasPrefix(username, "-") || strings.HasSuffix(username, "-") {
		return false
	}
	// Check for two consecutive hyphens
	if strings.Contains(username, "--") {
		return false
	}
	// Use regex to check for 3 to 39 alphanumeric characters or hyphens
	re := regexp.MustCompile(`^[a-zA-Z0-9\-]{3,39}$`)
	return re.MatchString(username)
}

func IsAvailable(username string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("https://github.com/%s", username))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusNotFound:
		return true, nil
	case http.StatusOK:
		return false, nil
	default:
		return false, errors.New("github: unexpected status code. unkown availability")
	}
}
