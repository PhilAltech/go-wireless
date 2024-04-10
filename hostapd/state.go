package hostapd

import "strings"

// State represents the current status of HOSTAPD
type State struct {
	State   string `json:"state"`
	Phy     string `json:"phy"`
	Frequ   string `json:"freq"`
	Channel string `json:"channel"`
}

// NewState will return the state of the HOSTAPD when given the raw output
func NewState(data string) State {
	s := State{}
	for _, l := range strings.Split(data, "\n") {
		bits := strings.Split(strings.TrimSpace(l), "=")
		if len(bits) < 2 {
			continue
		}

		switch bits[0] {
		case "state":
			s.State = bits[1]
		case "phy":
			s.Phy = bits[1]
		case "freq":
			s.Frequ = bits[1]
		case "channel":
			s.Channel= bits[1]
		}
	}

	return s
}
