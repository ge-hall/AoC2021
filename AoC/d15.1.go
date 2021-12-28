package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D15_1() {
	dat, err := ioutil.ReadFile("inputd15")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")

	for _, line := range data {
		fmt.Printf("%s\n", line)

	}

}
