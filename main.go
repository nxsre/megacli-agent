package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"
	"github.com/soopsio/megacli-agent/cache"
)

const agentVersion = "0.1.151228"

var host *string
var port *string

var ctrl *Controller

func main() {
	host = flag.String("host", "0.0.0.0", "listen host")
	port = flag.String("port", "2024", "listen port")
	flag.Parse()
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
	fmt.Println("megacli-agent " + agentVersion + " : starting server on :" + *host + ":" + *port)
	e.Start(*host + ":" + *port)
}
