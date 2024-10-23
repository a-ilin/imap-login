package main

import (
	"fmt"

	imap_client "github.com/emersion/go-imap/client"
	log "github.com/sirupsen/logrus"
)

type ImapConfig struct {
	Server   string
	Port     uint16
	User     string
	Password string
	TLS      bool
}

type ImapClient struct {
	Config ImapConfig

	c *imap_client.Client
}

func (cli *ImapClient) Login() error {
	var err error

	// Connect to server.
	log.Traceln("Connecting to", cli.Config.Server)

	if cli.Config.TLS {
		cli.c, err = imap_client.DialTLS(fmt.Sprintf("%s:%d", cli.Config.Server, cli.Config.Port), nil)
	} else {
		cli.c, err = imap_client.Dial(fmt.Sprintf("%s:%d", cli.Config.Server, cli.Config.Port))
	}

	if err != nil {
		return fmt.Errorf("cannot connect to: %s:%d: %w", cli.Config.Server, cli.Config.Port, err)
	}

	log.Traceln("Connected to", cli.Config.Server)

	// Login.
	if err := cli.c.Login(cli.Config.User, cli.Config.Password); err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	log.Traceln("Logged in as", cli.Config.User)

	return nil
}

func (cli *ImapClient) Logout() error {
	if cli.c != nil {
		if err := cli.c.Logout(); err != nil {
			return fmt.Errorf("logout failed: %w", err)
		}

		log.Traceln("Logged out")

		return nil
	}

	return nil
}
