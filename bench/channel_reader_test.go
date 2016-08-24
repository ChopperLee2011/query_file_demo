package query_file_demo_test

import (
	"fmt"
	"testing"

	. "github.com/chopperlee2011/query_file_demo"
)

func BenchmarkChannelReader(b *testing.B) {
	var result []string
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		result = ChannelReader("pokemon.csv")
	}
	fmt.Printf("grep line: %v\n", result)
}
