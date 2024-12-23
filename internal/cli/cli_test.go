package cli

import (
	"bytes"
	"flag"
	"io"
	"os"
	"testing"
)

// testSetup prepares the test environment
func testSetup(t *testing.T) (string, func()) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "tasks_test_*.json")
	if err != nil {
		t.Fatal(err)
	}

	// Save the original taskFile value
	originalTaskFile := taskFile
	// Set the taskFile to our temporary file
	taskFile = tmpfile.Name()

	// Save original stdout
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	// Return cleanup function
	return tmpfile.Name(), func() {
		// Restore original taskFile
		taskFile = originalTaskFile
		// Restore original stdout
		os.Stdout = oldStdout
		// Close and remove temporary file
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
}

// captureOutput captures stdout and returns it as a string
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// resetFlags resets all command line flags
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestRun_Add(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	tests := []struct {
		name         string
		args         []string
		wantContains string
	}{
		{
			name:         "add task",
			args:         []string{"-add", "Test Task", "-description", "Test Description"},
			wantContains: "Adding new task: Test Task",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = append([]string{"cmd"}, tt.args...)

			output := captureOutput(Run)

			if output != "" && !bytes.Contains([]byte(output), []byte(tt.wantContains)) {
				t.Errorf("Run() output = %v, want to contain %v", output, tt.wantContains)
			}
		})
	}
}

func TestRun_List(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	// First add a task
	resetFlags()
	os.Args = []string{"cmd", "-add", "Test Task"}
	Run()

	// Then test listing
	tests := []struct {
		name         string
		args         []string
		wantContains string
	}{
		{
			name:         "list all tasks",
			args:         []string{"-list", "all"},
			wantContains: "Test Task",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = append([]string{"cmd"}, tt.args...)

			output := captureOutput(Run)

			if output != "" && !bytes.Contains([]byte(output), []byte(tt.wantContains)) {
				t.Errorf("Run() output = %v, want to contain %v", output, tt.wantContains)
			}
		})
	}
}

func TestRun_Update(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	// First add a task
	resetFlags()
	os.Args = []string{"cmd", "-add", "Test Task"}
	Run()

	tests := []struct {
		name         string
		args         []string
		wantContains string
	}{
		{
			name:         "update task",
			args:         []string{"-update", "Updated Task", "1"},
			wantContains: "Updating task with ID 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = append([]string{"cmd"}, tt.args...)

			output := captureOutput(Run)

			if output != "" && !bytes.Contains([]byte(output), []byte(tt.wantContains)) {
				t.Errorf("Run() output = %v, want to contain %v", output, tt.wantContains)
			}
		})
	}
}

func TestRun_Delete(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	// First add a task
	resetFlags()
	os.Args = []string{"cmd", "-add", "Test Task"}
	Run()

	tests := []struct {
		name         string
		args         []string
		wantContains string
	}{
		{
			name:         "delete task",
			args:         []string{"-delete", "1"},
			wantContains: "Deleting task with ID 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = append([]string{"cmd"}, tt.args...)

			output := captureOutput(Run)

			if output != "" && !bytes.Contains([]byte(output), []byte(tt.wantContains)) {
				t.Errorf("Run() output = %v, want to contain %v", output, tt.wantContains)
			}
		})
	}
}

func TestRun_SetStatus(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	// First add a task
	resetFlags()
	os.Args = []string{"cmd", "-add", "Test Task"}
	Run()

	tests := []struct {
		name         string
		args         []string
		wantContains string
	}{
		{
			name:         "set status",
			args:         []string{"-status", "in-progress", "1"},
			wantContains: "Setting task status for ID 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = append([]string{"cmd"}, tt.args...)

			output := captureOutput(Run)

			if output != "" && !bytes.Contains([]byte(output), []byte(tt.wantContains)) {
				t.Errorf("Run() output = %v, want to contain %v", output, tt.wantContains)
			}
		})
	}
}

func TestRun_Default(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	tests := []struct {
		name         string
		args         []string
		wantContains string
	}{
		{
			name:         "show usage",
			args:         []string{},
			wantContains: "Usage: task-tracer-cli [options]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			os.Args = append([]string{"cmd"}, tt.args...)

			output := captureOutput(Run)

			if output != "" && !bytes.Contains([]byte(output), []byte(tt.wantContains)) {
				t.Errorf("Run() output = %v, want to contain %v", output, tt.wantContains)
			}
		})
	}
}
