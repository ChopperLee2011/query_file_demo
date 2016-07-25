package query_file_demo

import (
	"bufio"
	"os"
	"regexp"
)

func Scanner(path string) string {
	f, err := os.Open(path)
	CheckError(err)
	defer f.Close()

	var result string
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`\bslowpoke\b`)
	for scanner.Scan() {
		s := scanner.Text()
		if re.MatchString(s) {
			result = s
		}
	}
	return result
}
