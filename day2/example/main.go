package main

import (
	"fmt"

	"git.trains629.com/trains629/go-examples/day2"
	reg "golang.org/x/sys/windows/registry"
)

func main() {

	s, err := day2.ReadStringKey(reg.CURRENT_USER, `Software\Google\Chrome\BLBeacon`, "version")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q\n", s)
}
