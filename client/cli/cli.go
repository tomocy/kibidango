package cli

import (
	"os"
	"runtime"

	"github.com/tomocy/kibidango/engine/booter"
	containerPkg "github.com/tomocy/kibidango/engine/container"
	"github.com/tomocy/kibidango/engine/creater"
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
			Name:   "start",
			Action: start,
		},
	}
}

func (c *CLI) Run(args []string) error {
	return c.app.Run(args)
}

func create(ctx *cli.Context) error {
	ctner := container()
	return ctner.Create()
}

func start(ctx *cli.Context) error {
	ctner := container()
	return ctner.Boot()
}

func container() *containerPkg.Container {
	creater := creater.ForOS(runtime.GOOS, os.Stdin, os.Stdout, os.Stderr)
	bter := booter.ForOS(runtime.GOOS, "/root/container")
	return containerPkg.New(creater, bter)
}
