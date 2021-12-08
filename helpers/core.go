package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {
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

func MapSum(theMap map[int]int) int {
	sum := 0
	for _, x := range theMap {
		sum += x
	}
	return sum
}

func Max(list []int) int {
	max := 0
	for _, x := range list {
		if x > max {
			max = x
		}
	}
	return max
}

func Contains(str string, substr string) bool {
	for _, i := range substr {
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}

func ContainsExactly(str string, substr string) bool {
	if len(str) != len(substr) {
		return false
	}
	for _, i := range substr {
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}
