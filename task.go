package harvest

import (
	"fmt"
	"time"
)

type TaskResponse struct {
	Task *Task `json:"task"`
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
	taskResponse := TaskResponse{}
	path := fmt.Sprintf("/tasks/%v", taskID)
	err = a.Get(path, args, &taskResponse)
	return taskResponse.Task, err
}

func (a *API) GetTasks(args Arguments) (tasks []*Task, err error) {
	tasksResponse := make([]*TaskResponse, 0)
	path := fmt.Sprintf("/tasks")
	err = a.Get(path, args, &tasksResponse)
	for _, tr := range tasksResponse {
		tasks = append(tasks, tr.Task)
	}
	return tasks, err
}
