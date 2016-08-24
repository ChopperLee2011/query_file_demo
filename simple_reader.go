package query_file_demo

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func SimpleReader(path string) []string {
	f, err := ioutil.ReadFile(path)
	CheckError(err)
	lines := strings.Split(string(f), "\n")
	re := regexp.MustCompile(`\bslowpoke\b`)
	var result []string

	for _, line := range lines {
		if re.MatchString(line) {
			result = append(result, line)
		}
	}
	return result
}
