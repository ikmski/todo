package main

import (
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

func list(c *cli.Context) error {
	return nil
}

func add(c *cli.Context) error {
	return nil
}

func delete(c *cli.Context) error {
	return nil
}

func done(c *cli.Context) error {
	return nil
}

func undone(c *cli.Context) error {
	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "todo"
	app.Description = "command-line tool for TODO management"
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show task list",
			Action:  list,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add new task",
			Action:  add,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete task",
			Action:  delete,
		},
		{
			Name:   "done",
			Usage:  "make the task done",
			Action: done,
		},
		{
			Name:    "undone",
			Aliases: []string{"u"},
			Usage:   "undo the done task",
			Action:  undone,
		},
	}

	app.RunAndExitOnError()
}
