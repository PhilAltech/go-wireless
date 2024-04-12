package confHostapd

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	// See if we can open a file that not exists
	_, err := Open("/tmp/does_not_exist")
	if err == nil {
		t.Fatalf("Expected error opening non-existing file")
	}
	
	// Create a temporary file for testing
	tempFilePath := "/tmp/test_hostapd.conf"
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFilePath)

	// Write some test data to the temporary file
	testData := "interface=wlan0\nssid=TestNetwork #with Comment\nAdditionalData=SomeValue\n"
	_, err = tempFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write test data to temporary file: %v", err)
	}
	tempFile.Close()

	// Open the temporary file
	file, err := Open(tempFilePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	// Verify that the path is set correctly
	if file.Path() != tempFilePath {
		t.Errorf("Expected path %s, got %s", tempFilePath, file.Path())
	}

	// Verify that the config is parsed correctly
	 if file.HostapdConf.Interface != "wlan0" {
		t.Errorf("Expected interface wlan0, got %s", file.HostapdConf.Interface)
	}

	if file.HostapdConf.SSID != "TestNetwork" {
		t.Errorf("Expected SSID TestNetwork, got %s", file.HostapdConf.SSID)
	}

	if file.HostapdConf.Additionals["AdditionalData"] != "SomeValue" {
		t.Errorf("Expected AdditionalData SomeValue, got %s", file.HostapdConf.Additionals["AdditionalData"])
	}


	file.HostapdConf.Interface = "wlan1"
	file.Save()

	file, err = Open(tempFilePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	if file.HostapdConf.Interface != "wlan1" {
		t.Errorf("Expected interface wlan1, got %s", file.HostapdConf.Interface)
	}
}

