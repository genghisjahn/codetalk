package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var ckey = map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25}
var outkey = map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z"}

func main() {
	fmt.Println("Enter message:")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	msg := s.Text()
	files, err := ioutil.ReadDir("/Users/jonwear/gocode/src/github.com/genghisjahn/codetalk/makeotp/pads")
	if err != nil {
		fmt.Println(err)
		return
	}
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
	fmt.Println("3rd, now let's pull a One Time Pad and Encrypt")
	s.Scan()
	fmt.Println("Let's use: ", files[0].Name())
	path := fmt.Sprintf("/Users/jonwear/gocode/src/github.com/genghisjahn/codetalk/makeotp/pads/%s", files[0].Name())
	s.Scan()
	padData, padErr := ioutil.ReadFile(path)
	if padErr != nil {
		fmt.Println(padErr)
		return
	}
	padStr := string(padData)
	padStr = strings.Replace(padStr, "\n", "", -1)
	padStr = strings.Replace(padStr, " ", "", -1)
	outputmsg := ""
	for k, v := range msg {
		key := string(padStr[k])
		keyVal := ckey[key]
		msgVal := ckey[string(v)]
		output := (keyVal + msgVal) % 26
		if (k)%5 == 0 && k != 0 {
			if (k)%25 == 0 {
				outputmsg += "\n"
			} else {
				outputmsg += " "
			}
		}
		outputmsg += outkey[output]

	}
	filename := fmt.Sprintf("encrypted-%s.msg", time.Now().Format("20060102150405"))

	fmt.Println(outputmsg)
	createfile(filename, outputmsg)

}
func createfile(path, msg string) {
	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("output/%s", path)); err != nil {
		panic(err)
	}
	defer file.Close()
	_, errWrite := file.WriteString(msg)
	if errWrite != nil {
		fmt.Println(errWrite)
	}

}
