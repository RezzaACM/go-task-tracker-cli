package cli

import (
	"flag"
	"fmt"
	"strconv"

	models "github.com/yourusername/task-tracer-cli/internal/models/task"
	service "github.com/yourusername/task-tracer-cli/internal/service/task"
)

var taskFile = "tasks.json"

// Run the cli application.
//
// This function parses the flags and executes the appropriate method
// on the TaskManager instance.
//
// The available flags are:
//
//	-add <task>    Add a new task
//	-update <id>   Update a task by ID
//	-delete <id>   Delete a task by ID
//	-status <id>   Set task status (todo, in-progress, done)
//	-list          List tasks by status (todo, in-progress, done)
func Run() {
	tm := &service.TaskManager{}
	if err := tm.Load(taskFile); err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	add := flag.String("add", "", "Add a new task")
	description := flag.String("description", "", "Add a description to the task")
	update := flag.String("update", "", "Update an existing task")
	delete := flag.Int("delete", 0, "Delete a task by ID")
	status := flag.String("status", "", "Set task status (TODO, IN_PROGRESS, DONE)")
	list := flag.String("list", "", "List tasks by status (TODO, IN_PROGRESS, DONE)")

	flag.Parse()

	switch {
	case *list != "":
		fmt.Println("Listing tasks by status:", *list)
		status := models.Status(*list)
		tasks := tm.List(status)
		for _, task := range tasks {
			fmt.Printf(" ID: %d\n Title: %s\n Status: %s\n Description: %s\n Created At: %s\n Updated At: %s\n\n", task.ID, task.Title, task.Status, task.Description, task.CreatedAt, task.UpdatedAt)
		}
	case *add != "":
		fmt.Println("Adding new task:", *add)
		tm.Add(models.Task{Title: *add, Description: *description})
	case *status != "":
		fmt.Printf("Setting task status for ID %s to %s\n", flag.Arg(0), *status)
		id, _ := strconv.Atoi(flag.Arg(0))
		tm.SetStatus(id, models.Status(*status))
	case *update != "":
		fmt.Printf("Updating task with ID %s\n", flag.Arg(0))
		id, _ := strconv.Atoi(flag.Arg(0))
		tm.Update(id, models.Task{Title: *update, Description: *description})
	case *delete != 0:
		fmt.Printf("Deleting task with ID %d\n", *delete)
		tm.Delete(*delete)
	default:
		fmt.Println("Usage: task-tracer-cli [options]")
		fmt.Println("Options:")
		fmt.Println("  -add <task>    Add a new task")
		fmt.Println("  -update <id>   Update a task by ID")
		fmt.Println("  -delete <id>   Delete a task by ID")
		fmt.Println("  -status <id>   Set task status (todo, in-progress, done)")
		fmt.Println("  -list          List tasks by status (todo, in-progress, done)")
	}

	if err := tm.Save(taskFile); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}
