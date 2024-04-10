package main

import (
	"fmt"

	"github.com/PhilAltech/go-wireless/hostapd"
)

func main() {
	wc, err := hostapd.NewClient("wlp1s0")
	if err != nil {
		panic(err)
	}
	defer wc.Close()

	status, err := wc.Status()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", status)
}
