package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate is a function that returns a cli.App
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application new command line"
	app.Usage = "Get information about application"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IP address",
			Flags:  flags,
			Action: GetIP,
		},
		{
			Name:   "server",
			Usage:  "Get name server",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func GetIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
