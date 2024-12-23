package service

import (
	"os"
	"testing"

	models "github.com/yourusername/task-tracer-cli/internal/models/task"
)

func TestTaskManager_Add(t *testing.T) {
	tm := &TaskManager{}
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	tm.Add(task)
	if len(tm.Tasks) != 1 {
		t.Errorf("Expected 1 task, but got %d", len(tm.Tasks))
	}

	addedTask := tm.Tasks[0]
	if addedTask.ID != 1 {
		t.Errorf("Expected ID 1, got %d", addedTask.ID)
	}
	if addedTask.Title != task.Title {
		t.Errorf("Expected title %s, got %s", task.Title, addedTask.Title)
	}
	if addedTask.Status != models.Todo {
		t.Errorf("Expected status %s, got %s", models.Todo, addedTask.Status)
	}
}

func TestTaskManager_Update(t *testing.T) {
	tm := &TaskManager{}
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	tm.Add(task)
	updatedTask := models.Task{
		Title: "Updated Task",
	}
	tm.Update(1, updatedTask)

	if len(tm.Tasks) != 1 {
		t.Errorf("Expected 1 task, but got %d", len(tm.Tasks))
	}

	updatedTask = tm.Tasks[0]
	if updatedTask.Title != "Updated Task" {
		t.Errorf("Expected title 'Updated Task', got %s", updatedTask.Title)
	}
	if updatedTask.Description != "This is a test task" {
		t.Errorf("Expected description 'This is a test task', got %s", updatedTask.Description)
	}
}

func TestTaskManager_Delete(t *testing.T) {
	tm := &TaskManager{}
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	tm.Add(task)
	tm.Delete(1)
	if len(tm.Tasks) != 0 {
		t.Errorf("Expected 0 tasks, but got %d", len(tm.Tasks))
	}
}

func TestTaskManager_SetStatus(t *testing.T) {
	tm := &TaskManager{}
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	tm.Add(task)
	tm.SetStatus(1, models.Completed)
	if len(tm.Tasks) != 1 {
		t.Errorf("Expected 1 task, but got %d", len(tm.Tasks))
	}

	updatedTask := tm.Tasks[0]
	if updatedTask.Status != models.Completed {
		t.Errorf("Expected status %s, got %s", models.Completed, updatedTask.Status)
	}
}

func TestTaskManager_List(t *testing.T) {
	tm := &TaskManager{}
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	tm.Add(task)

	taskCompleted := models.Task{
		Title:       "Completed Task",
		Description: "This is a completed task",
	}
	tm.Add(taskCompleted)
	tm.SetStatus(2, models.Completed)

	listTaskCompleted := tm.List(models.Completed)
	if len(listTaskCompleted) != 1 {
		t.Errorf("Expected 1 task, but got %d", len(listTaskCompleted))
	}

	tasks := tm.List(models.All)
	if len(tasks) != 2 {
		t.Errorf("Expected 1 task, but got %d", len(tasks))
	}
}

func TestTaskManager_SaveAndLoad(t *testing.T) {
	tm := &TaskManager{}
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	tm.Add(task)
	filename := "test_tasks.json"

	// Save the tasks
	err := tm.Save(filename)
	if err != nil {
		t.Errorf("Error saving tasks: %v", err)
	}

	// Load the tasks
	err = tm.Load(filename)
	if err != nil {
		t.Errorf("Error loading tasks: %v", err)
	}

	if len(tm.Tasks) != 1 {
		t.Errorf("Expected 1 task, but got %d", len(tm.Tasks))
	}

	// Clean up
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Error removing file: %v", err)
	}
}
