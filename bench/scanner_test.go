package query_file_demo_test

import (
	"fmt"
	"testing"

	. "github.com/chopperlee2011/query_file_demo"
)

func BenchmarkScanner(b *testing.B) {
	var result string
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		result = Scanner("pokemon.csv")
	}
	fmt.Println(result)
}
