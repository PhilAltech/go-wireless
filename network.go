package wireless

import (
	"net"
)

// This file contains components from github.com/brlbil/wpaclient
//
// Copyright (c) 2017 Birol Bilgin
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// nd/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// UT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// NewNamedNetwork will create a new network with the given parameters
func NewNamedNetwork(name, ssid, psk string) Network {
	return Network{IDStr: name, SSID: ssid, PSK: psk}
}

// NewNetwork will create a new network with the given parameters
func NewNetwork(ssid, psk string) Network {
	return NewNamedNetwork(ssid, ssid, psk)
}

// NewDisabledNetwork will create a new disabled network with the given parameters
func NewDisabledNetwork(ssid, psk string) Network {
	n := NewNamedNetwork(ssid, ssid, psk)
	n.Flags = append(n.Flags, "DISABLED")
	return n
}

// Network represents a known network
type Network struct {
	ID       int      `json:"id"`
	IDStr    string   `json:"id_str"`
	KeyMgmt  string   `json:"key_mgmt"`
	SSID     string   `json:"ssid"`
	BSSID    string   `json:"bssid"`
	ScanSSID bool     `json:"scan_ssid"`
	PSK      string   `json:"psk"`
	Flags    []string `json:"flags"`
}

// Connector is interfce than can connect a network
type Connector interface {
	Connect(Network) error
}

type Connectable interface {
	SetCmds() []string
}

// IsDisabled will return true if the network is disabled
func (net Network) IsDisabled() bool {
	for _, f := range net.Flags {
		if f == "DISABLED" {
			return true
		}
	}
	return false
}

// Disable or enabled the network
func (net Network) Disable(on bool) {
	var idx int
	var found bool
	for i, f := range net.Flags {
		if f == "DISABLED" {
			found = true
			idx = i
			break
		}
	}

	if on && !found {
		net.Flags = append(net.Flags, "DISABLED")
		return
	}

	net.Flags = append(net.Flags[:idx], net.Flags[idx:]...)
}

// Connect will connect the network to the given connector
func (net Network) Connect(cl Connector) error {
	return cl.Connect(net)
}

// SetCmds will generate the set_network commands to run to set this network up in
func (net Network) SetCmds() [][]string {
	cmds := [][]string{}
	cmds = append(cmds, []string{CmdSetNetwork, itoa(net.ID), "ssid", quote(net.SSID)})
	if net.IDStr != "" {
		cmds = append(cmds, []string{CmdSetNetwork, itoa(net.ID), "id_str", quote(net.IDStr)})
	}

	if net.ScanSSID {
		cmds = append(cmds, []string{CmdSetNetwork, itoa(net.ID), "scan_ssid", "1'"})
	}

	for _, f := range net.Flags {
		switch f {
		case "DISABLED":
			cmds = append(cmds, []string{CmdSetNetwork, itoa(net.ID), "disabled", "1"})
		}
	}

	if net.PSK == "" {
		cmds = append(cmds, []string{CmdSetNetwork, itoa(net.ID), "key_mgmt", "None"})
	} else {
		cmds = append(cmds, []string{CmdSetNetwork, itoa(net.ID), "psk", quote(net.PSK)})
	}

	return cmds
}

// AP represents an access point seen by the scan networks command
type AP struct {
	ID        int              `json:"id"`
	Freq      int              `json:"freq"`
	RSSI      int              `json:"rssi"`
	BSSID     net.HardwareAddr `json:"bssid"`
	SSID      string           `json:"ssid"`
	ESSID     string           `json:"essid"`
	Flags     []string         `json:"flags"`
	Signal    int              `json:"signal"`
	Frequency int              `json:"frequency"`
}
