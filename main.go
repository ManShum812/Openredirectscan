package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	fmt.Println("Open Redirect Vulnerability Scanner")
	fmt.Println("\\__(-_-)__/")

	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	sc := bufio.NewScanner(os.Stdin)

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				if isRedirectToEvil(url) {
					fmt.Printf("%sOpen Redirect Found:%s %s -> evil.com\n", colorRed, colorReset, url)
				} else {
					fmt.Printf("%sNot Vulnerable:%s %s\n", colorGreen, colorReset, url)
				}
			}
		}()
	}

	for sc.Scan() {
		url := sc.Text()
		jobs <- url
	}

	close(jobs)
	wg.Wait()
}

// Function to check if the URL redirects to "evil.com"
func isRedirectToEvil(url string) bool {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Do not follow redirects
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check if the response status code is a redirect (3xx)
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		if location != "" && containsEvil(location) {
			return true
		}
	}
	return false
}

// Function to check if the URL contains "evil.com"
func containsEvil(url string) bool {
	return containsDomain(url, "evil.com")
}

// Function to check if the URL contains a specific domain
func containsDomain(url, domain string) bool {
	return len(url) >= len(domain) && url[len(url)-len(domain):] == domain
}
