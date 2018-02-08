package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/netopssh/megacli-agent/cache"
)

const agentVersion = "0.1.151228"
const agentPort = "3214"

var ctrl *Controller

func main() {

	fmt.Println("megacli-agent " + agentVersion + " : initialising cache..")
	c := cache.New()
	ctrl = New(c)
	c.Run()

	e := echo.New()

	// Routes
	e.GET("/megacli/version", ctrl.VersionHandle)
	e.GET("/megacli/logicaldisk/summary", ctrl.MegaCliLogicalDiskSummaryHandle)
	e.GET("/megacli/logicaldisk/state/:disk", ctrl.MegaCliLogicalDiskStateHandle)
	e.GET("/megacli/physicaldisk/summary", ctrl.MegaCliPhysicalDiskSummaryHandle)
	e.GET("/megacli/physicaldisk/state/:disk", ctrl.MegaCliPhysicalDiskStateHandle)

	// Start server
	fmt.Println("megacli-agent " + agentVersion + " : starting server on :" + agentPort)
	e.Start(":" + agentPort)
}
