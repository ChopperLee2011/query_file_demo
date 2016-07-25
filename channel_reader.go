package query_file_demo

import (
	"bufio"
	"os"
	"regexp"
	"sync"
)

func ChannelReader(path string) string {
	f, err := os.Open(path)
	CheckError(err)
	defer f.Close()
	jobs := make(chan string)
	results := make(chan string)
	wg := new(sync.WaitGroup)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go grepLine(jobs, results, wg)
	}

	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	return <-results
}

func grepLine(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	re := regexp.MustCompile(`\bslowpoke\b`)
	for j := range jobs {
		if re.MatchString(j) {
			results <- j
		}
	}
}
