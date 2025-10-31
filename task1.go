package main

import (
    "fmt"
    "strings"
    "regexp"
)

func counter(words string) map[string]int {
    words = strings.ToLower(words)

    re := regexp.MustCompile(`[^a-zA-Z\s]`)
    words = re.ReplaceAllString(words, "")

    wordList := strings.Fields(words)
    count := make(map[string]int)

    for _, word := range wordList {
        count[word]++
    }

    return count
}

func palindrome(words string) bool {
    words = strings.ToLower(words)

    re := regexp.MustCompile(`[^a-zA-Z]`)
    words = re.ReplaceAllString(words, "")

    reversed := ""
    for i := len(words) - 1; i >= 0; i-- {
        reversed += string(words[i])
    }

    return words == reversed
}


func main() {
    fmt.Println(counter("Go go GO, let's go!"))
	fmt.Println(palindrome("A man, a plan, a canal, Panama"))
}
