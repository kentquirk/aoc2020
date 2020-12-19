package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func day18b(lines []string) (int, error) {
	total := 0
	for _, line := range lines {
		calc := &AoCCalc{Buffer: line}
		calc.Init()
		calc.Expression.Init(line)
		if err := calc.Parse(); err != nil {
			return 0, err
		}
		calc.Execute()
		result := calc.Evaluate()
		total += result
	}
	return total, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	result, err := day18b(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
