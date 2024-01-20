package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// main is the entry point of the program.
func main() {
	// Create new CLI app
	app := &cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that checks whether a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			// Get port from command line argument
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			// Check the status of the domain
			status := Check(c.String("domain"), port)
			// Print the status
			fmt.Println(status)
			return nil
		},
	}

	// Run the CLI app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
