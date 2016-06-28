package cli

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/getmemcpuinfo/version"
	"os"
	"path"
)

var (
	commands = []cli.Command{
		{
			Name:      "GetMemCpuInfo",
			ShortName: "g",
			Usage:     "Run a TransferTool",
			Flags:     []cli.Flag{sAddr, sHeart},
			Action:    GetMemCpuInfo,
		},
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "Create a getinfconf",
			Action:    CreateConf,
		},
	}

	sAddr = cli.StringFlag{
		Name:  "addr",
		Usage: "It is a www address",
	}
	sHeart = cli.StringFlag{
		Name:  "heart",
		Usage: "It is a heartbeat time",
	}
)

func Run() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "A transfer message"
	app.Version = version.VERSION

	app.Author = "mtanlee"
	app.Email = "mtanlee07@gmail.com"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},

		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (optios: debug, info, warn, error, fatal, panic)"),
		},
	}

	//logs
	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level)

		if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
