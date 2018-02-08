package cache

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"os"
	"path/filepath"

	"github.com/soopsio/agent-tools/models"
	"github.com/soopsio/agent-tools/scraper"
)

const (
	Megacli64 = "/opt/MegaRAID/MegaCli/MegaCli64"
)

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func ReleaseMegacli64() error {
	if PathExist(Megacli64) {
		return nil
	}
	dir := filepath.Dir(Megacli64)
	os.MkdirAll(dir, 0755)
	if err := RestoreAsset(dir, "MegaCli64"); err != nil {
		return err
	}
	return os.Chmod(Megacli64, 0777)
}

// GetMegaCliLogicalDisk ..
func GetMegaCliLogicalDisk(Id string) models.MegaCliLogicalDisk {
	var disk models.MegaCliLogicalDisk
	diskLocation := "0"
	if err := ReleaseMegacli64(); err != nil {
		fmt.Println("ReleaseMegacli64 failed: ", err)
		return disk
	}
	command := Megacli64 + " -LDInfo -L" + Id + " -a0"

	output := scraper.GetCommandOutput(command)
	output = scraper.RemoveLineFeed(output)
	lines := scraper.SplitLines(output)

	for _, line := range lines {
		kv := strings.Split(string(line), ":")
		value := strings.TrimSpace(kv[len(kv)-1])
		// replace this with switch/case
		if strings.HasPrefix(string(line), "Name ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.Name = value
		} else if strings.HasPrefix(string(line), "RAID Level ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.RaidLevel = value
		} else if strings.HasPrefix(string(line), "Size ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.Size = value
		} else if strings.HasPrefix(string(line), "Sector Size ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.SectorSize = value
		} else if strings.HasPrefix(string(line), "State ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.State = value
		} else if strings.HasPrefix(string(line), "Strip Size ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.StripSize = value
		} else if strings.HasPrefix(string(line), "Number Of Drives ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.NumberOfDrives = value
		}
	}

	return disk
}

// GetMegaCliEnclosureDeviceId
func GetMegaCliEnclosureDeviceId() models.MegaCliEnclosureInfo {
	enclosure := models.MegaCliEnclosureInfo{}
	if err := ReleaseMegacli64(); err != nil {
		fmt.Println("ReleaseMegacli64 failed: ", err)
		return disk
	}
	command := Megacli64 + " -EncInfo -a0"
	output := scraper.GetCommandOutput(command)
	output = scraper.RemoveLineFeed(output)
	lines := scraper.SplitLines(output)
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		kv := strings.Split(string(line), ":")
		value := strings.TrimSpace(kv[len(kv)-1])

		if strings.HasPrefix(string(line), "Device ID ") {
			enclosure.EnclosureDeviceId = value
		} else if strings.HasPrefix(string(line), "Number of Physical Drives  ") {
			enclosure.NumberOfPhysicalDrives = value
		} else {
			continue
		}

	}
	return enclosure
}

// GetMegaCliPhysicalDisk ..
func GetMegaCliPhysicalDisk(EnclosureDeviceId, i int) models.MegaCliPhysicalDisk {
	var disk models.MegaCliPhysicalDisk
	diskLocation := "[" + strconv.Itoa(EnclosureDeviceId) + ":" + strconv.Itoa(i) + "]"

	disk.EncDeviceId = EnclosureDeviceId
	disk.SlotNumber = i
	if err := ReleaseMegacli64(); err != nil {
		fmt.Println("ReleaseMegacli64 failed: ", err)
		return disk
	}
	command := Megacli64 + " -PDInfo -PhysDrv " + diskLocation + " -a0"

	output := scraper.GetCommandOutput(command)
	output = scraper.RemoveLineFeed(output)
	lines := scraper.SplitLines(output)

	for _, line := range lines {
		kv := strings.Split(string(line), ":")
		value := strings.TrimSpace(kv[len(kv)-1])

		// replace this with switch/case
		if strings.HasSuffix(string(line), " is not found.") {
			break
		} else if strings.HasPrefix(string(line), "WWN: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.Wwn = value
		} else if strings.HasPrefix(string(line), "Media Error Count: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.MedErrCount = value
		} else if strings.HasPrefix(string(line), "Other Error Count: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.OthErrCount = value
		} else if strings.HasPrefix(string(line), "Predictive Failure Count: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.PredictiveFailureCount = value
		} else if strings.HasPrefix(string(line), "Last Predictive Failure Event Seq Number: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.LastPredictiveFailureEventSeqNumber = value
		} else if strings.HasPrefix(string(line), "PD Type: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.PdType = value
		} else if strings.HasPrefix(string(line), "Raw Size: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.RawSize = value
		} else if strings.HasPrefix(string(line), "Firmware state: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.FirmwareState = value
		} else if strings.HasPrefix(string(line), "Inquiry Data: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.InquiryData = value
		} else if strings.HasPrefix(string(line), "Device Speed: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.DeviceSpeed = value
		} else if strings.HasPrefix(string(line), "Link Speed: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.LinkSpeed = value
		} else if strings.HasPrefix(string(line), "Media Type: ") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.MediaType = value
		} else if strings.HasPrefix(string(line), "Drive Temperature") {
			fmt.Println(diskLocation + " - " + string(line))
			disk.DriveTemp = value
		}
	}

	return disk
}
