package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	filename := fmt.Sprintf("useonce-%s.otp", time.Now().Format("20060102150405"))
	otp := [999]int{}
	for k := range otp {
		otp[k] = rand.Intn(26) + 1
	}
	makeotp(filename, otp)
}

func makeotp(path string, otp [999]int) {
	var file *os.File
	var err error
	if file, err = os.Create(path); err != nil {
		panic(err)
	}
	defer file.Close()

	for _, elem := range otp {
		v := strconv.Itoa(elem)
		fmt.Println(":", v)
		_, err := file.WriteString(v + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
