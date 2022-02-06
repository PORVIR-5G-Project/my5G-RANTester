package main

import (
	"fmt"
	"my5G-RANTester/config"
	"my5G-RANTester/internal/templates"

	// "fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

const version = "1.0"

func init() {

	cfg, err := config.GetConfig()
	if err != nil {
		//return nil
		log.Fatal("Error in get configuration")
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	fmt.Println(cfg.Logs.Level)
	log.SetLevel(log.Level(cfg.Logs.Level))
	spew.Config.Indent = "\t"

	log.Info("my5G-RANTester version " + version)
}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "ue",
				Aliases: []string{"ue"},
				Usage:   "Testing an ue attached with configuration",
				Action: func(c *cli.Context) error {
					name := "Testing an ue attached with configuration"
					cfg := config.Data

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", 1)
					log.Info("[TESTER][GNB] Control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] Data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestAttachUeWithConfiguration()
					return nil
				},
			},
			{
				Name:    "gnb",
				Aliases: []string{"gnb"},
				Usage:   "Testing an gnb attached with configuration",
				Action: func(c *cli.Context) error {
					name := "Testing an gnb attached with configuration"
					cfg := config.Data

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][GNB] Number of GNBs: ", 1)
					log.Info("[TESTER][GNB] Control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] Data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestAttachGnbWithConfiguration()
					return nil
				},
			},
			{
				Name:    "load-test",
				Aliases: []string{"load-test"},
				Usage: "\nLoad endurance stress tests.\n" +
					"Example for testing multiple UEs: load-test -n 5 \n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-ues", Value: 1, Aliases: []string{"n"}},
				},
				Action: func(c *cli.Context) error {
					var numUes int
					name := "Testing registration of multiple UEs"
					cfg := config.Data

					if c.IsSet("number-of-ues") {
						numUes = c.Int("number-of-ues")
					} else {
						log.Info(c.Command.Usage)
						return nil
					}

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of UEs: ", numUes)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					templates.TestMultiUesInQueue(numUes)

					return nil
				},
			},
			{
				Name:    "amf-requests",
				Aliases: []string{"amf-requests"},
				Usage: "\nTesting AMF requests for the specified time in milliseconds\n" +
					"Example for testing multiple requests per second: amf-requests -n 5 \n" +
					"Example for testing multiple requests per specifying time: amf-requests -n 5 -t 200\n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-requests", Value: 1, Aliases: []string{"n"}},
					&cli.IntFlag{Name: "time", Value: 1000, Aliases: []string{"t"}},
				},
				Action: func(c *cli.Context) error {
					var numRqs int
					var time int
					name := "Testing AMF requests for the specified time"
					cfg := config.Data

					numRqs = c.Int("number-of-requests")
					time = c.Int("time")

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of Requests per second: ", numRqs)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					log.Warn("[TESTER][GNB] AMF Requests per Time:", templates.TestRqsPerTime(numRqs, time))
					return nil
				},
			},
			{
				Name:    "amf-requests-loop",
				Aliases: []string{"amf-requests"},
				Usage: "\nTesting AMF requests in interval\n" +
					"Example for testing multiple requests in 20 seconds: amf-requests-loop -t 20\n",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "number-of-requests", Value: 1, Aliases: []string{"n"}},
					&cli.IntFlag{Name: "time", Value: 1, Aliases: []string{"t"}},
				},
				Action: func(c *cli.Context) error {
					var time int
					var numRqs int

					name := "Testing AMF requests for the specified time"
					cfg := config.Data

					numRqs = c.Int("number-of-requests")
					time = c.Int("time")

					log.Info("---------------------------------------")
					log.Info("[TESTER] Starting test function: ", name)
					log.Info("[TESTER][UE] Number of Requests per second: ", numRqs)
					log.Info("[TESTER][GNB] gNodeB control interface IP/Port: ", cfg.GNodeB.ControlIF.Ip, "/", cfg.GNodeB.ControlIF.Port)
					log.Info("[TESTER][GNB] gNodeB data interface IP/Port: ", cfg.GNodeB.DataIF.Ip, "/", cfg.GNodeB.DataIF.Port)
					log.Info("[TESTER][AMF] AMF IP/Port: ", cfg.AMF.Ip, "/", cfg.AMF.Port)
					log.Info("---------------------------------------")
					log.Warn("[TESTER][GNB] AMF Requests GLOBAL per Time:", templates.TestRqsLoop(numRqs, time))
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
