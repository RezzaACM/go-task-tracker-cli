# Task Tracker CLI

A command-line interface application for tracking tasks and managing productivity built with Go. Sample solution for the [task-tracer](https://roadmap.sh/projects/task-tracker) challange from [roadmap.sh](https://roadmap.sh/projects/task-tracker)

## Features

- Create and manage tasks
- Track task status
- Command-line based interface
- Easy to use and lightweight

## How to run

```bash
git clone git@github.com:RezzaACM/go-task-tracker-cli.git
cd go-task-tracker-cli
```

## Usage

```bash
# Build the application
go build -o task-tracer-cli ./cmd/task-tracer-cli

# Start the application
./task-tracer-cli

# Common commands
./task-tracer-cli -add "New task" -description "This is a new task"
./task-tracer-cli -list todo
./task-tracer-cli -update "Updated task" -descripiton "This is a updated task" [task-id]
./task-tracer-cli -status completed [task-id]
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
