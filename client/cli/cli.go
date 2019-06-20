package cli

import (
	"os"
	"runtime"

	"github.com/tomocy/kibidango/engine/container"
	"github.com/tomocy/kibidango/engine/creater"
	"github.com/tomocy/kibidango/engine/initializer"
	"github.com/urfave/cli"
)

const (
	name    = "kibidango"
	usage   = "a client for linux container runtime"
	version = "0.0.1"
)

func New() *CLI {
	c := new(CLI)
	c.init()
	return c
}

type CLI struct {
	app *cli.App
}

func (c *CLI) init() {
	c.app = cli.NewApp()
	c.initBasic()
	c.initCommands()
}

func (c *CLI) initBasic() {
	c.app.Name = name
	c.app.Usage = usage
	c.app.Version = version
}

func (c *CLI) initCommands() {
	c.app.Commands = []cli.Command{
		{
			Name:   "create",
			Action: create,
		},
		{
			Name:   "init",
			Action: initialize,
		},
	}
}

func (c *CLI) Run(args []string) error {
	return c.app.Run(args)
}

func create(ctx *cli.Context) error {
	ctner := new(container.Container)
	creater := creater.ForOS(runtime.GOOS, os.Stdin, os.Stdout, os.Stderr)
	return ctner.Create(creater, "init")
}

func initialize(ctx *cli.Context) error {
	ctner := new(container.Container)
	initer := initializer.ForOS(runtime.GOOS, "/root/container")
	return ctner.Init(initer)
}
