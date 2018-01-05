package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type TaskStatus int

const (
	TaskStatus_Todo    TaskStatus = 0
	TaskStatus_Done    TaskStatus = 1
	TaskStatus_Deleted TaskStatus = 2
)

var TaskStatusName = map[int]string{
	0: "TODO",
	1: "DONE",
	2: "DELETED",
}

type Task struct {
	ID          int        `yaml:"id"`
	Status      TaskStatus `yaml:"status"`
	Title       string     `yaml:"title"`
	Detail      string     `yaml:"detail"`
	Tags        []string   `yaml:"tags"`
	Deadline    string     `yaml:"deadline"`
	CompletedAt string     `yaml:"completed_at"`
}

type Tasks struct {
	Tasks []Task `yaml:"tasks"`
}

func newTask(file string) *Task {

	t := new(Task)
	t.ID = issueTaskID(file)
	t.Status = TaskStatus_Todo

	return t
}

func loadTasks(file string) (*Tasks, error) {

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var t Tasks
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Task) save(file string) error {

	ts, err := loadTasks(file)
	if err != nil {
		return err
	}

	ts.Tasks = append(ts.Tasks, *t)

	return ts.save(file)
}

func (t *Tasks) save(file string) error {

	d, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, d, 0644)
	if err != nil {
		return err
	}

	return nil
}

func issueTaskID(file string) int {

	ts, err := loadTasks(file)
	if err != nil {
		return 0
	}

	id := 0
	for _, t := range ts.Tasks {
		if t.ID > id {
			id = t.ID
		}
	}
	id++

	return id
}

func (t *Task) interactiveEdit() error {

	fmt.Printf("title:[%s] ", t.Title)
	fmt.Scanln(&t.Title)

	fmt.Printf("Detail:[%s] ", t.Detail)
	fmt.Scanln(&t.Detail)

	fmt.Printf("Tags:%v ", t.Tags)
	var tags = ""
	fmt.Scanln(&tags)
	if tags != "" {
		t.Tags = strings.Split(tags, ",")
	}

	return nil
}
