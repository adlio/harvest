package harvest

import "testing"

func testTask(t *testing.T) *Task {
	a := testAPI()
	taskResponse := mockResponse("tasks", "task-example.json")
	a.BaseURL = taskResponse.URL
	task, err := a.GetTask(2086199, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if task.Name != "Admin" {
		t.Errorf("Incorrect Task Name '%s'", task.Name)
	}
	if task.ID != 2086199 {
		t.Errorf("Incorrect Task ID '%v'", task.ID)
	}
	return task
}

func TestGetTask(t *testing.T) {
	task := testTask(t)
	if task == nil {
		t.Fatal("testTask() returned nil instead of task")
	}
	if task.ID != 2086199 {
		t.Errorf("Incorrect task ID '%v'", task.ID)
	}
}

func TestGetTasks(t *testing.T) {
	a := testAPI()
	taskResponse := mockResponse("tasks", "tasks-example.json")
	a.BaseURL = taskResponse.URL
	tasks, err := a.GetTasks(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Errorf("Incorrect number of tasks '%v'", len(tasks))
	}
	if tasks[0].Name != "Admin" {
		t.Errorf("Incorrect Task Name '%s'", tasks[0].Name)
	}
	if tasks[1].ID != 2086200 {
		t.Errorf("Incorrect Task ID '%v'", tasks[1].ID)
	}
}
