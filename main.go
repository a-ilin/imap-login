package main

import (
	"fmt"
	"math"
	"os"

	"github.com/urfave/cli/v2"

	log "github.com/sirupsen/logrus"
)

var (
	// Specified via LDFLAGS in Makefile.
	AppVersion string

	verbose bool
	config  ImapConfig
)

func main() {
	// Start application.
	app := &cli.App{
		Usage: "Test login to IMAP server",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"vv"},
				Usage:       "Verbose logging",
				Destination: &verbose,
			},
			&cli.StringFlag{
				Name:        "server",
				Aliases:     []string{"s"},
				EnvVars:     []string{"IMAP_SERVER"},
				Usage:       "Server name or IP",
				Destination: &config.Server,
				Required:    true,
			},
			&cli.UintFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				EnvVars:     []string{"IMAP_PORT"},
				Usage:       "Port",
				DefaultText: "143/993",
				Action: func(ctx *cli.Context, v uint) error {
					if v > math.MaxUint16 {
						return fmt.Errorf("IMAP port is incorrect: %d", v)
					}
					config.Port = uint16(v)
					return nil
				},
			},
			&cli.BoolFlag{
				Name:        "tls",
				Aliases:     []string{"t"},
				EnvVars:     []string{"IMAP_TLS"},
				Usage:       "Use TLS",
				Destination: &config.TLS,
			},
			&cli.StringFlag{
				Name:        "username",
				Aliases:     []string{"u"},
				EnvVars:     []string{"IMAP_USERNAME"},
				Usage:       "Username",
				Destination: &config.User,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"P"},
				EnvVars:     []string{"IMAP_PASSWORD"},
				Usage:       "Password",
				Destination: &config.Password,
				Required:    true,
			},
		},
		Before: func(ctx *cli.Context) error {
			// Set log level.
			if verbose {
				log.SetLevel(log.TraceLevel)
			} else {
				log.SetLevel(log.InfoLevel)
			}

			// Set port.
			if config.Port == 0 {
				if config.TLS {
					config.Port = 993
				} else {
					config.Port = 143
				}
			}

			return nil
		},
		Action: func(ctx *cli.Context) error {
			cli := &ImapClient{
				Config: config,
			}

			if err := cli.Login(); err != nil {
				return err
			}

			if err := cli.Logout(); err != nil {
				return err
			}

			fmt.Println("Success")

			return nil
		},
		HideHelpCommand: true,
		Version:         AppVersion,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
