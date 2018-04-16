package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	//	"gopkg.in/yaml.v2"

	//	"github.com/lxc/lxd/shared"
	//	"github.com/lxc/lxd/shared/api"
	cli "github.com/lxc/lxd/shared/cmd"
	"github.com/lxc/lxd/shared/i18n"
	//	"github.com/lxc/lxd/shared/log15"
	//	"github.com/lxc/lxd/shared/logging"
)

type cmdTop struct {
	global   *cmdGlobal
	flagType []string
}

func (c *cmdTop) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = i18n.G("top [<remote>:]")
	cmd.Short = i18n.G("Shows an htop like interface for LXD")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Shows an htop like interface for LX

Useful for monitoring your LXD instances, to understand which container

is taking how much resources on the host server`))
	cmd.Example = cli.FormatSection("", i18n.G(
		`lxc top`))
	cmd.Hidden = false

	cmdTop := cmdAliasAdd{global: c.global}
	cmd.AddCommand(cmdTop.Command())

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdTop) Run(cmd *cobra.Command, args []string) error {
	conf := c.global.conf

	var err error
	var remote string

	// Sanity checks
	exit, err := c.global.CheckArgs(cmd, args, 0, 1)
	if exit {
		return err
	}

	if len(args) == 0 {
		remote, _, err = conf.ParseRemote("")
		if err != nil {
			return err
		}
	} else {
		remote, _, err = conf.ParseRemote(args[0])
		if err != nil {
			return err
		}
	}

	d, err := conf.GetContainerServer(remote)
	if err != nil {
		return err
	}

	listener, err := d.GetEvents()
	if err != nil {
		return err
	}

	handler := func(message interface{}) {
		// Special handling for logging only output

		render, err := json.Marshal(&message)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n\n", render)
	}

	_, err = listener.AddHandler(c.flagType, handler)
	if err != nil {
		return err
	}

	return listener.Wait()
}
