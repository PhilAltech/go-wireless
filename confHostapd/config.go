package confHostapd

type Config struct {
	Ctrl_interface string
	Bridge 			 string
	Interface string
	SSID string
	Password string
	IgnoreBroadcastSSID string
	Ieee80211n string
	Hw_mode string
	Country_code string
	Channel string

	Additionals map[string]string
}

func NewDefaultConfig() *Config {
	return &Config{
		Ctrl_interface: "/var/run/hostapd",
		Bridge: "br0",
		Interface: "wlan0",
		SSID: "MyAP",
		IgnoreBroadcastSSID: "0",
		Ieee80211n: "1",
		Hw_mode: "g",
		Country_code: "US",
		Channel: "6",
		Additionals: make(map[string]string),
	}
}