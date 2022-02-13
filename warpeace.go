package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func countWords(words []string) (morewords map[string]int) {
	var nim = len(words)
	morewords = make(map[string]int)
	for i := 0; i < nim; i++ {
		/*val, ok := morewords[words[i]]
		if ok == false && val == 0 {
			morewords[words[i]] = 1
		} else {
			morewords[words[i]] += 1
		}*/
		morewords[words[i]] += 1

	}
	return
}

func reportResults(words map[string]int) {
	for word, num := range words {
		fmt.Printf("%s : %d", word, num)
		fmt.Println("")
	}
}

func main() {
	var filename string
	fmt.Print("type the name of a file to read: ")
	fmt.Scan(&filename)
	byteArray, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	str := string(byteArray)
	reg, err3 := regexp.Compile("[^a-zA-Z0-9]+")
	if err3 != nil {
		log.Fatal(err3)
	}
	processedString := reg.ReplaceAllString(str, " ")
	forcount := strings.Split(processedString, " ")
	forpr := countWords(forcount)
	reportResults(forpr)
}
