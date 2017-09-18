package harvest

import (
	"fmt"
	"time"
)

type TasksResponse struct {
	Tasks []*Task `json:"tasks"`
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
	tasksResponse := TasksResponse{}
	path := fmt.Sprintf("/tasks")
	err = a.Get(path, args, &tasksResponse)
	return tasksResponse.Tasks, err
}
