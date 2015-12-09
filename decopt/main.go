package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var ckey = map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}
var outkey = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I", 10: "J", 11: "K", 12: "L", 13: "M", 14: "N", 15: "O", 16: "P", 17: "Q", 18: "R", 19: "S", 20: "T", 21: "U", 22: "V", 23: "W", 24: "X", 25: "Y", 26: "Z"}

func main() {
	filesotp, err := ioutil.ReadDir("/Users/jonwear/gocode/src/github.com/genghisjahn/codetalk/makeotp/pads")
	if err != nil {
		fmt.Println(err)
		return
	}
	filesenc, errEnc := ioutil.ReadDir("/Users/jonwear/gocode/src/github.com/genghisjahn/codetalk/encotp/output")
	if errEnc != nil {
		fmt.Println(errEnc)
		return
	}

	pathotp := fmt.Sprintf("/Users/jonwear/gocode/src/github.com/genghisjahn/codetalk/makeotp/pads/%s", filesotp[0].Name())
	pathenc := fmt.Sprintf("/Users/jonwear/gocode/src/github.com/genghisjahn/codetalk/encotp/output/%s", filesenc[0].Name())
	otpstr := getotp(pathotp)
	encstr := getenc(pathenc)
	outputmsg := ""
	_, _, _ = otpstr, encstr, outputmsg
	for k, v := range encstr {
		_ = k
		t := ckey[string(v)]
		o := ckey[string(otpstr[k])]
		var d int
		d = (t - o) % 26
		if d < 0 {
			d = 26 + d
		}
		outputmsg += outkey[d]
	}
	fmt.Println("Decrypted message:")
	for k, v := range outputmsg {
		if (k)%5 == 0 {
			if (k)%25 == 0 {
				fmt.Printf("\n")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Printf("%s", string(v))
	}
	fmt.Printf("\n")
}

func getenc(path string) string {
	encData, encErr := ioutil.ReadFile(path)
	if encErr != nil {
		panic(encErr)
	}
	encStr := string(encData)
	encStr = strings.Replace(encStr, "\n", "", -1)
	encStr = strings.Replace(encStr, " ", "", -1)
	return encStr
}

func getotp(path string) string {

	padData, padErr := ioutil.ReadFile(path)
	if padErr != nil {
		panic(padErr)
	}
	padStr := string(padData)
	padStr = strings.Replace(padStr, "\n", "", -1)
	padStr = strings.Replace(padStr, " ", "", -1)
	return padStr
}
