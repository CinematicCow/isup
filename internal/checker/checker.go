package checker

import (
	"fmt"
	Http "net/http"
	"sync"

	"github.com/CinematicCow/isup/internal/http"
)

type WebsiteChecker struct{

}

func (w *WebsiteChecker) CheckWebsites(websites []string) map[string]string {
	var wg sync.WaitGroup
	results := make(chan string)
	resultMap := make(map[string]string)

	for _, website := range websites {
		wg.Add(1)
		go func(site string) {
			defer wg.Done()
			if w.isWebsiteUp(site) {
				results <- fmt.Sprintf("%s is up", site)
			} else {
				results <- fmt.Sprintf("%s is down", site)
			}
		}(website)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		resultMap[result] = result
	}

	return resultMap
}

func (w *WebsiteChecker) isWebsiteUp(website string) bool {
	resp, err := http.Get(fmt.Sprintf("https://www.%s.com", website))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == Http.StatusOK
}
