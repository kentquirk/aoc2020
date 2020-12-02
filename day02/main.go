package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./inputsample.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	day02a(lines)
	day02b(lines)
}
