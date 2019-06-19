package cli

import "github.com/urfave/cli"

const (
	name = "kibidango"
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
	c.initBasic()
}

func (c *CLI) initBasic() {
	c.app.Name = name
}

func (c *CLI) Run(args []string) error {
	return c.app.Run(args)
}
