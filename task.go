package harvest

import (
	"fmt"
	"time"
)

type TasksResponse struct {
	PagedResponse
	Tasks []*Task `json:"tasks"`
}

type TaskStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	BillableByDefault bool      `json:"billable_by_default"`
	Deactivated       bool      `json:"deactivated"`
	DefaultHourlyRate float64   `json:"default_hourly_rate"`
	IsDefault         bool      `json:"is_default"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (a *API) GetTask(taskID int64, args Arguments) (task *Task, err error) {
	task = &Task{}
	path := fmt.Sprintf("/tasks/%v", taskID)
	err = a.Get(path, args, task)
	return task, err
}

func (a *API) GetTasks(args Arguments) (tasks []*Task, err error) {
	tasks = make([]*Task, 0)
	tasksResponse := TasksResponse{}
	err = a.GetPaginated("/tasks", args, &tasksResponse, func() {
		for _, t := range tasksResponse.Tasks {
			tasks = append(tasks, t)
		}
		tasksResponse.Tasks = make([]*Task, 0)
	})
	return tasks, err
}

func (a *API) CreateTask(t *Task, args Arguments) (task *Task, err error) {
	task = &Task{}
	err = a.Post("/tasks", args, t, task)
	return task, err
}

func (a *API) UpdateTask(t *Task, args Arguments) (task *Task, err error) {
	task = &Task{}
	path := fmt.Sprintf("/tasks/%v", t.ID)
	err = a.Patch(path, args, t, task)
	return task, err
}

func (a *API) DeleteTask(taskID int64, args Arguments) error {
	path := fmt.Sprintf("/tasks/%v", taskID)
	err := a.Delete(path, args)
	return err
}
