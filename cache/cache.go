package cache

import (
	"strconv"
	"time"

	"github.com/netopssh/agent-tools/models"
)

// Cache ..
type Cache struct {
	EnclosureDeviceId    int
	MegaCliLogicalDisks  map[string]models.MegaCliLogicalDisk
	MegaCliPhysicalDisks map[string]models.MegaCliPhysicalDisk
}

// New ..
func New() *Cache {
	c := &Cache{}
	c.PopulateEnclosureDeviceId()
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
func (c *Cache) PopulateEnclosureDeviceId() {
	c.EnclosureDeviceId, _ = GetMegaCliEnclosureDeviceId()
}

// PopulateCache ..
func (c *Cache) PopulateCache() {
	// can this be done in the constructor ?
	c.MegaCliLogicalDisks = make(map[string]models.MegaCliLogicalDisk)
	c.MegaCliPhysicalDisks = make(map[string]models.MegaCliPhysicalDisk)

	// Logical Disks
	logicalDiskLocation := "0"
	disk := GetMegaCliLogicalDisk(logicalDiskLocation)
	c.MegaCliLogicalDisks[logicalDiskLocation] = disk

	// Physical Disks
	i := 0
	for i < 24 {
		physicalDiskLocation := "[" + strconv.Itoa(c.EnclosureDeviceId) + ":" + strconv.Itoa(i) + "]"
		disk := GetMegaCliPhysicalDisk(c.EnclosureDeviceId, i)
		c.MegaCliPhysicalDisks[physicalDiskLocation] = disk
		i += 1
	}

}
