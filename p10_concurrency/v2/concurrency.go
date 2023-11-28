package main
 
type WebsietChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsietChecker, urls[]string) map[string]bool {
	resluts := make(map[string]bool)
	resultChannel := make(chan result)

	//这里resultChannel是一个缓冲队列

	for _, url := range urls {
		go func (u string)  {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		reslut := <- resultChannel
		resluts[reslut.string] = reslut.bool
	}

	return resluts
}