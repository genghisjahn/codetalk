package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var ckey = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I", 10: "J", 11: "K", 12: "L", 13: "M", 14: "N", 15: "O", 16: "P", 17: "Q", 18: "R", 19: "S", 20: "T", 21: "U", 22: "V", 23: "W", 24: "X", 25: "Y", 26: "Z"}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	filename := fmt.Sprintf("useonce-%s.otp", time.Now().Format("20060102150405"))
	otp := [1000]int{}
	for k := range otp {
		otp[k] = rand.Intn(26) + 1
	}
	makeotp(filename, otp)
}

func makeotp(path string, otp [1000]int) {
	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("pads/%s", path)); err != nil {
		panic(err)
	}
	defer file.Close()
	var separator = ""
	fmt.Println(len(otp))
	for k, elem := range otp {
		if (k+1)%5 == 0 {
			separator = " "
		}
		if (k+1)%25 == 0 {
			separator = "\n"
		}
		v := ckey[elem]
		_, err := file.WriteString(v + separator)
		separator = ""
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
