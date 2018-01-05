package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type TaskStatus int

const (
	TaskStatusTodo TaskStatus = 0
	TaskStatusDone TaskStatus = 1
)

type Task struct {
	ID     int        `yaml:"id"`
	Title  string     `yaml:"title"`
	Detail string     `yaml:"detail"`
	Status TaskStatus `yaml:"status"`
}

type Tasks struct {
	Tasks []Task `yaml:"tasks"`
}

func loadTasks(file string) (*Tasks, error) {

	_, err := os.Stat(file)
	if err != nil {
		err = createTodoFile(file)
		if err != nil {
			return nil, err
		}
	}

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

func createTodoFile(file string) error {

	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func (ts *Tasks) newTask() *Task {

	t := new(Task)
	t.ID = ts.issueTaskID()
	t.Status = TaskStatusTodo

	ts.Tasks = append(ts.Tasks, *t)

	return t
}

func (ts *Tasks) getTaskIndex(id int) (int, error) {

	idx := -1
	for i, t := range ts.Tasks {
		if t.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return idx, fmt.Errorf("task not found")
	}

	return idx, nil
}

func (ts *Tasks) save(file string) error {

	d, err := yaml.Marshal(ts)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, d, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (ts *Tasks) issueTaskID() int {

	id := 0
	for _, t := range ts.Tasks {
		if t.ID > id {
			id = t.ID
		}
	}
	id++

	return id
}

func (ts *Tasks) interactiveEdit(id int) (bool, error) {

	tasks := ts.Tasks

	idx, err := ts.getTaskIndex(id)
	if err != nil {
		return false, err
	}

	edit := false

	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("title:[%s] ", tasks[idx].Title)
	s.Scan()
	if s.Text() != "" {
		tasks[idx].Title = s.Text()
		edit = true
	}

	fmt.Printf("Detail:[%s] ", tasks[idx].Detail)
	s.Scan()
	if s.Text() != "" {
		tasks[idx].Detail = s.Text()
		edit = true
	}

	return edit, nil
}

func (ts *Tasks) interactiveDelete(id int) (bool, error) {

	tasks := ts.Tasks

	idx, err := ts.getTaskIndex(id)
	if err != nil {
		return false, err
	}

	s := ""
	fmt.Printf("Delete task %03d: %s ? (y/N) ", tasks[idx].ID, tasks[idx].Title)
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {

		result := []Task{}
		for _, t := range ts.Tasks {
			if t.ID != id {
				result = append(result, t)
			}
		}

		ts.Tasks = result
		return true, nil
	}

	return false, nil
}

func (ts *Tasks) interactiveDone(id int) (bool, error) {

	tasks := ts.Tasks

	idx, err := ts.getTaskIndex(id)
	if err != nil {
		return false, err
	}

	if tasks[idx].Status == TaskStatusDone {
		return false, fmt.Errorf("task status is already DONE")
	}

	s := ""
	fmt.Printf("Done task %03d: %s ? (y/N) ", tasks[idx].ID, tasks[idx].Title)
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {

		tasks[idx].Status = TaskStatusDone
		return true, nil
	}

	return false, nil
}

func (ts *Tasks) interactiveUndone(id int) (bool, error) {

	tasks := ts.Tasks

	idx, err := ts.getTaskIndex(id)
	if err != nil {
		return false, err
	}

	if tasks[idx].Status == TaskStatusTodo {
		return false, fmt.Errorf("task status is already TODO")
	}

	s := ""
	fmt.Printf("Undone task %03d: %s ? (y/N) ", tasks[idx].ID, tasks[idx].Title)
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {

		tasks[idx].Status = TaskStatusTodo
		return true, nil
	}

	return false, nil
}
