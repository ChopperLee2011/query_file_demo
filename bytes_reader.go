package query_file_demo

import (
	"log"
	"os"
)

func bytes_reader() {
	f, err := os.Open("./bench/pokemon.csv")
	CheckError(err)
	readBuf := make([]byte, 1024)
	data, err := f.Read(readBuf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read info is %v", data)
}
