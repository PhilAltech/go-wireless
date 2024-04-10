package main

import (
	"fmt"

	wireless "github.com/PhilAltech/go-wireless/wpa"
)

func main() {
	//iface, ok := wireless.DefaultInterface()
	//if !ok {
	//	panic("no wifi cards on the system")
	//}
	

	ifaces := wireless.InterfacesFromWPARunDir("/run/wpa_supplicant")
	if len(ifaces) == 0 {
		panic("no wifi cards on the system")
	}
	iface := ifaces[0]

	fmt.Printf("Using interface: %s\n", iface)
	wc, err := wireless.NewClient(iface)
	if err != nil {
		panic(err)
	}
	defer wc.Close()

	aps, err := wc.Scan()
	if err != nil {
		panic(err)
	}

	for _, ap := range aps {
		fmt.Println(ap.SSID)
	}
}
