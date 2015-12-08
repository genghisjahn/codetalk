package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter message:")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	msg := s.Text()
	fmt.Println("Okay, let's go through the steps:")
	s.Scan()
	fmt.Println("1st, let's make it all upper-case:")
	s.Scan()
	msg = strings.ToUpper(msg)
	fmt.Println(msg)
	fmt.Println("")
	s.Scan()
	fmt.Println("2nd, let's get rid of the spaces:")
	s.Scan()
	msg = strings.Replace(msg, " ", "", -1)
	fmt.Println(msg)
	s.Scan()
	fmt.Println("3rd, let's split it into groups of 5 letters:")
	s.Scan()
	words := []string{}
	cwords := []string{}
	_ = cwords
	word := ""
	for _, v := range msg {
		if len(word) == 5 {
			words = append(words, word)
			word = ""
		}
		word += string(v)
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	for _, w := range words {
		fmt.Printf("%s ", w)
	}
	s.Scan()
	fmt.Println("")
	fmt.Println("4th, let's encrypt it:")
	s.Scan()
	fmt.Println("")
	for _, w := range words {
		for _, letter := range w {
			fmt.Printf("%s", string(letter))
		}
		fmt.Printf("%s", " ")
	}
	fmt.Println("")

	for _, w := range words {
		for _, letter := range w {
			fmt.Printf("%s", encryptWord(string(letter)))
		}
		fmt.Printf("%s", " ")
	}
	fmt.Println("")
}

func encryptWord(w string) string {
	c, _ := cKey[w]
	return c
}
