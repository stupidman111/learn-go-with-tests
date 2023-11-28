package main
 


type WebsiteChecker func(string) bool


func CheckWebsites(w WebsiteChecker, urls []string) map[string]bool {
	urlMapping := make(map[string]bool)
	for _, url := range urls {
		urlMapping[url] = w(url)
	}

	return urlMapping
}