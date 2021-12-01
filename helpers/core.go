package helpers

import (
	"os"
	"log"
	"bufio"
)

func OpenFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func Sum(list []int) int {
	sum := 0
	for _, x := range list {
		sum += x
	}
	return sum
}