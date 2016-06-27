package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	run "github.com/getmemcpuinfo/info"
	"time"
)

func GetMemCpuInfo(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Fatalf("the `GetMemCpuInfo` command takes no arguments. See '%s GetMemInfo --help'.", c.App.Name)

	}
	address := c.String("addr")
	if address == "" {
		log.Fatal("missing mandatory --addr flag ")
	}
	heartbeat, err := time.ParseDuration(c.String("heart"))
	if err != nil {
		log.Fatalf("invalid --heart: %v", err)
	}
	if heartbeat < 1*time.Second {

		log.Fatal("--heart should be at least one second")
	}

	for {
		log.Infof(address)
		run.GetMemCpuInfo(address)
		time.Sleep(heartbeat)
	}

}
