package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

var config globalConfig

const (
	configFileName = "config.toml"
	todoFileName   = ".todo"
)

func list(c *cli.Context) error {

	NewTasksFromFile(getTodoFilePath())

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

func getConfigFilePath() string {

	filePath := ""
	isExist := false
	curDir, err := os.Getwd()
	if err == nil {
		filePath = filepath.Join(curDir, configFileName)
		_, err = os.Stat(filePath)
		if err == nil {
			isExist = true
		}
	}

	if !isExist {
		filePath = filepath.Join(os.Getenv("HOME"), ".config", "todo", configFileName)
	}

	return filePath
}

func getTodoFilePath() string {

	filePath := ""

	if config.TodoDir != "" {

		filePath = filepath.Join(config.TodoDir, todoFileName)

	} else {

		isExist := false
		curDir, err := os.Getwd()
		if err == nil {
			filePath = filepath.Join(curDir, todoFileName)
			_, err = os.Stat(filePath)
			if err == nil {
				isExist = true
			}
		}

		if !isExist {
			filePath = filepath.Join(os.Getenv("HOME"), todoFileName)
		}
	}

	return filePath
}

func main() {

	configFile := getConfigFilePath()
	_, err := os.Stat(configFile)
	if err != nil {

		config = getDefaultConfig()

	} else {

		_, err := toml.DecodeFile(getConfigFilePath(), &config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

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

	// TODO delete
	app.Action = list

	app.RunAndExitOnError()
}
