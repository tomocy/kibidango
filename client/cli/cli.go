package cli

import "github.com/urfave/cli"

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
}

func (c *CLI) initBasic() {
	c.app.Name = name
	c.app.Usage = usage
	c.app.Version = version
}

func (c *CLI) Run(args []string) error {
	return c.app.Run(args)
}
