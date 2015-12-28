package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// VersionHandle ..
func (ctrl *Controller) VersionHandle(c *echo.Context) error {
	return c.JSON(http.StatusOK, agentVersion)
}

// MegaCliPhysicalDiskSummaryHandle ..
func (ctrl *Controller) MegaCliLogicalDiskSummaryHandle(c *echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.Cache.MegaCliLogicalDisks)
}

// MegaCliPhysicalDiskStateHandle ..
func (ctrl *Controller) MegaCliLogicalDiskStateHandle(c *echo.Context) error {
	diskLocation := c.Param("disk")
	return c.JSON(http.StatusOK, ctrl.Cache.MegaCliLogicalDisks[diskLocation].State)
}

// MegaCliPhysicalDiskSummaryHandle ..
func (ctrl *Controller) MegaCliPhysicalDiskSummaryHandle(c *echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.Cache.MegaCliPhysicalDisks)
}

// MegaCliPhysicalDiskStateHandle ..
func (ctrl *Controller) MegaCliPhysicalDiskStateHandle(c *echo.Context) error {
	diskLocation := "[" + strings.Replace(c.Param("disk"), "-", ":", 1) + "]"
	return c.JSON(http.StatusOK, ctrl.Cache.MegaCliPhysicalDisks[diskLocation].FirmwareState)
}
