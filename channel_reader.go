package query_file_demo

import (
	"bufio"
	"os"
	"regexp"
)

func ChannelReader(path string) []string {
	works := 10
	f, err := os.Open(path)
	CheckError(err)
	defer f.Close()
	jobs := make(chan string)
	results := make(chan string, 10)
	complete := make(chan bool)
	var matchLines []string

	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	for i := 0; i <= works; i++ {
		go grepLine(jobs, results, complete)
	}

	for i := 0; i < works; i++ {
		<-complete
	}

	close(results)
	for v := range results {
		matchLines = append(matchLines, v)
	}
	return matchLines
}

func grepLine(jobs <-chan string, results chan<- string, complete chan bool) {
	re := regexp.MustCompile(`\bslowpoke\b`)
	for j := range jobs {
		if re.MatchString(j) {
			results <- j
		}
	}
	complete <- true
}
