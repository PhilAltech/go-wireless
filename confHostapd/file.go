package confHostapd

import (
	"os"
	"strings"
)

// DefaultPath is the default path for the HOSTAPD config file
const DefaultPath = "/etc/hostap/hostapd.conf"

// File represents the config file
type File struct {
	path        string
	HostapdConf Config
}

// Open will open and parse the config file at the given path
func Open(path string) (*File, error) {
	f := &File{path: path}
	return f, f.Load()
}

// Path will return the path of the file
func (f *File) Path() string {
	return f.path
}

// Load will parse the file contents
func (f *File) Load() error {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return err
	}

	f.HostapdConf.Additionals = make(map[string]string)

	for _, l := range strings.Split(string(data), "\n") {
		l = strings.TrimSpace(l)

		//Skip comments and empty lines
		if strings.HasPrefix(l, "#") || len(l) == 0 {
			continue
		}

		//No key value pair
		vs := strings.Split(l, "=")
		if len(vs) < 2 {
			continue
		}

		//Invalid key
		if strings.Contains(vs[0], "#") {
			continue
		}

		//Remove comments and trim spaces
		vs[1] = strings.Split(vs[1], "#")[0]
		vs[1] = strings.TrimSpace(vs[1])
		vs[0] = strings.TrimSpace(vs[0])

		if vs[0] == "" || vs[1] == "" {
			continue
		}

		switch vs[0] {
		case "ctrl_interface":
			f.HostapdConf.Ctrl_interface = vs[1]
		case "bridge":
			f.HostapdConf.Bridge = vs[1]
		case "interface":
			f.HostapdConf.Interface = vs[1]
		case "ssid":
			f.HostapdConf.SSID = vs[1]
		case "wpa_passphrase":
			f.HostapdConf.Password = vs[1]
		case "ignore_broadcast_ssid":
			f.HostapdConf.IgnoreBroadcastSSID = vs[1]
		case "ieee80211n":
			f.HostapdConf.Ieee80211n = vs[1]
		case "hw_mode":
			f.HostapdConf.Hw_mode = vs[1]
		case "country_code":
			f.HostapdConf.Country_code = vs[1]
		case "channel":
			f.HostapdConf.Channel = vs[1]
		default:
			f.HostapdConf.Additionals[vs[0]] = vs[1]
		}
	}
	
	return nil
}

// Save will save the config to the file path
func (f *File) Save() error {
	data := []string{}

	if f.HostapdConf.Interface != "" {
		data = append(data, "ctrl_interface=" + f.HostapdConf.Ctrl_interface)
	}
	if f.HostapdConf.Bridge != "" {
		data = append(data, "bridge=" + f.HostapdConf.Bridge)
	}
	if f.HostapdConf.Interface != "" {
		data = append(data, "interface=" + f.HostapdConf.Interface)
	}
	if f.HostapdConf.SSID != "" {
		data = append(data, "ssid=" + f.HostapdConf.SSID)
	}
	if f.HostapdConf.Password != "" {
		data = append(data, "wpa_passphrase=" + f.HostapdConf.Password)
	}
	if f.HostapdConf.IgnoreBroadcastSSID != "" {
		data = append(data, "ignore_broadcast_ssid=" + f.HostapdConf.IgnoreBroadcastSSID)
	}
	if f.HostapdConf.Ieee80211n != "" {
		data = append(data, "ieee80211n=" + f.HostapdConf.Ieee80211n)
	}
	if f.HostapdConf.Hw_mode != "" {
		data = append(data, "hw_mode=" + f.HostapdConf.Hw_mode)
	}
	if f.HostapdConf.Country_code != "" {
		data = append(data, "country_code=" + f.HostapdConf.Country_code)
	}
	if f.HostapdConf.Channel != "" {
		data = append(data, "channel=" + f.HostapdConf.Channel)
	}

	for k, v := range f.HostapdConf.Additionals {
		data = append(data, k+"="+v)
	}

	return os.WriteFile(f.path, []byte(strings.Join(data, "\n")), 0644)
}
