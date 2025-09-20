package main
import (
	"fmt"
	"io"
	"net/http"
	"sync"
)
type HomePageSize struct {
	URL  string
	Size int
}
func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize, len(urls))
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				return
			}
			defer res.Body.Close()
			bs, err := io.ReadAll(res.Body)
			if err != nil {
				return
			}
			results <- HomePageSize{URL: url, Size: len(bs)}
		}(url)
	}
	wg.Wait()
	close(results)
	var biggest HomePageSize
	for result := range results {
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Printf("The biggest home page: %s, size = %d bytes\n", biggest.URL, biggest.Size)
}
