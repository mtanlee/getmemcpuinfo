package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	run "github.com/getmemcpuinfo/info"
)

func CreateConf(c *cli.Context) error {
	if len(c.Args()) != 0 {
		log.Fatalf("The `create` command takes no arguments. See '%s create --help'.", c.App.Name)
	}
	run.CreateConf()
	return nil
}
