package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("chat.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	wordCount := wordCount(file)

	csv := ""
	for word, count := range wordCount {
		csv += fmt.Sprintf("%s,%d\n", word, count)
	}

	fmt.Println(csv)

	file, err = os.Create("word_count.csv")

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	file.WriteString(csv)

}

func wordCount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		// word = strings.ReplaceAll(word, ".", "")

		isTimestamp := strings.Contains(word, "]") || strings.Contains(word, "[")

		if isTimestamp == false {
			counts[word]++
		}
	}
	return counts
}
