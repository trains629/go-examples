package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git.trains629.com/trains629/go-examples/day1"
)

func main() {
	b, err := ioutil.ReadFile("../index.xml")

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	out, err := day1.Xml2Json(&b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ioutil.WriteFile("./index.json", *out, 0600)
}
