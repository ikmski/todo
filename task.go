package main

import "fmt"

type TaskStatus int32

const (
	TaskStatus_Todo    TaskStatus = 0
	TaskStatus_Done    TaskStatus = 1
	TaskStatus_Deleted TaskStatus = 2
)

var TaskStatusName = map[int32]string{
	0: "TODO",
	1: "DONE",
	2: "DELETED",
}

type Task struct {
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

func NewTasksFromFile(file string) *Tasks {

	fmt.Printf("%v\n", file)

	return nil
}

func (t *Tasks) WriteFile(file string) {

}
