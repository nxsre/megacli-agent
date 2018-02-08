package cache

import (
	"strconv"
	"time"

	"github.com/soopsio/agent-tools/models"
)

// Cache ..
type Cache struct {
	MegaCliEnclosureInfo models.MegaCliEnclosureInfo
	MegaCliLogicalDisks  map[string]models.MegaCliLogicalDisk
	MegaCliPhysicalDisks map[string]models.MegaCliPhysicalDisk
}

// New ..
func New() *Cache {
	c := &Cache{}
	c.PopulateCache()
	return c
}

// Run ..
func (c *Cache) Run() {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			// purge the cache
			*c = Cache{}
			// populate the cache
			c.PopulateCache()
		}
	}()
}

// EnclosureDeviceId
func (c *Cache) PopulateEnclosureInfo() {
	c.MegaCliEnclosureInfo = models.MegaCliEnclosureInfo{}
	enclosure := GetMegaCliEnclosureDeviceId()
	c.MegaCliEnclosureInfo = enclosure
}

// PopulateCache ..
func (c *Cache) PopulateCache() {
	c.PopulateEnclosureInfo()
	// can this be done in the constructor ?
	c.MegaCliLogicalDisks = make(map[string]models.MegaCliLogicalDisk)
	c.MegaCliPhysicalDisks = make(map[string]models.MegaCliPhysicalDisk)

	// Logical Disks
	logicalDiskLocation := "0"
	disk := GetMegaCliLogicalDisk(logicalDiskLocation)
	c.MegaCliLogicalDisks[logicalDiskLocation] = disk

	// Physical Disks
	i := 0
	enclosureDeviceId, _ := strconv.Atoi(c.MegaCliEnclosureInfo.EnclosureDeviceId)
	numOfDevices, _ := strconv.Atoi(c.MegaCliEnclosureInfo.NumberOfPhysicalDrives)
	for i < numOfDevices {
		physicalDiskLocation := "[" + strconv.Itoa(enclosureDeviceId) + ":" + strconv.Itoa(i) + "]"
		disk := GetMegaCliPhysicalDisk(enclosureDeviceId, i)
		c.MegaCliPhysicalDisks[physicalDiskLocation] = disk
		i += 1
	}

}
