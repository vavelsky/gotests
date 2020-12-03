package concurrency

import "time"

// WebsiteChecker type for check the websites
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites function verifies if the url is website
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	time.Sleep(2 * time.Second)

	return results
}
