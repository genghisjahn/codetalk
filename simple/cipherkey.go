package main

import (
	"math/rand"
	"time"
)

var cKey map[string]string

func init() {
	cKey = make(map[string]string)
	//simpleNumber()
	//simpleTranspose()
	randomTranspose()
}
func simpleNumber() {
	cKey["A"] = "01"
	cKey["B"] = "02"
	cKey["C"] = "03"
	cKey["D"] = "04"
	cKey["E"] = "05"
	cKey["F"] = "06"
	cKey["G"] = "07"
	cKey["H"] = "08"
	cKey["I"] = "09"
	cKey["J"] = "10"
	cKey["K"] = "11"
	cKey["L"] = "12"
	cKey["M"] = "13"
	cKey["N"] = "14"
	cKey["O"] = "15"
	cKey["P"] = "16"
	cKey["Q"] = "17"
	cKey["R"] = "18"
	cKey["S"] = "19"
	cKey["T"] = "20"
	cKey["U"] = "21"
	cKey["V"] = "22"
	cKey["W"] = "23"
	cKey["X"] = "24"
	cKey["Y"] = "25"
	cKey["Z"] = "26"
}

func simpleTranspose() {
	cKey["A"] = "B"
	cKey["B"] = "C"
	cKey["C"] = "D"
	cKey["D"] = "E"
	cKey["E"] = "F"
	cKey["F"] = "G"
	cKey["G"] = "H"
	cKey["H"] = "I"
	cKey["I"] = "J"
	cKey["J"] = "K"
	cKey["K"] = "L"
	cKey["L"] = "M"
	cKey["M"] = "N"
	cKey["N"] = "O"
	cKey["O"] = "P"
	cKey["P"] = "Q"
	cKey["Q"] = "R"
	cKey["R"] = "S"
	cKey["S"] = "T"
	cKey["T"] = "U"
	cKey["U"] = "V"
	cKey["V"] = "W"
	cKey["W"] = "X"
	cKey["X"] = "Y"
	cKey["Y"] = "K"
	cKey["Z"] = "A"
}

func randomTranspose() {
	rand.Seed(time.Now().UTC().UnixNano())
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	mixed := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := range mixed {
		j := rand.Intn(i + 1)
		mixed[i], mixed[j] = mixed[j], mixed[i]
	}
	for k, v := range letters {
		cKey[v] = mixed[k]
	}
}
