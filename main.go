package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

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
	todoFileName   = "todo.yaml"
)

func list(c *cli.Context) error {

	ts, err := loadTasks(getTodoFilePath())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, t := range ts.Tasks {
		if t.Status == TaskStatusTodo {
			fmt.Printf("\u2610 %03d: %s\n  %s\n\n", t.ID, t.Title, t.Detail)
		} else if t.Status == TaskStatusDone {
			fmt.Printf("\u2611 %03d: %s\n  %s\n\n", t.ID, t.Title, t.Detail)
		}
	}

	return nil
}

func add(c *cli.Context) error {

	file := getTodoFilePath()

	ts, err := loadTasks(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	t := ts.newTask()
	save, err := ts.interactiveEdit(t.ID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if save {
		err = ts.save(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}

func edit(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}

	targetID, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	file := getTodoFilePath()

	ts, err := loadTasks(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	save, err := ts.interactiveEdit(targetID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if save {
		err = ts.save(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}

func delete(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}

	targetID, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	file := getTodoFilePath()

	ts, err := loadTasks(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	save, err := ts.interactiveDelete(targetID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if save {
		err = ts.save(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}

func done(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}

	targetID, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	file := getTodoFilePath()

	ts, err := loadTasks(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	save, err := ts.interactiveDone(targetID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if save {
		err = ts.save(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return nil
}

func undone(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("argument is required")
	}

	targetID, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	file := getTodoFilePath()

	ts, err := loadTasks(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	save, err := ts.interactiveUndone(targetID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if save {
		err = ts.save(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

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

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return filePath
	}

	if config.TodoDir != "" {

		dir := strings.Replace(config.TodoDir, "~", user.HomeDir, -1)
		filePath = filepath.Join(dir, todoFileName)

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
		err = config.save(configFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

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
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "edit task",
			Action:  edit,
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
