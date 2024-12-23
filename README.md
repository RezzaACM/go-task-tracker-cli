# Task Tracker CLI

A command-line interface application for tracking tasks and managing productivity built with Go. Sample solution for the [task-tracer](https://roadmap.sh/projects/task-tracker) challange from [roadmap.sh](https://roadmap.sh/projects/task-tracker)

## Features

- Create and manage tasks
- Track task status
- Command-line based interface
- Easy to use and lightweight

## Installation

```bash
go get github.com:RezzaACM/go-task-tracker-cli.git
```

## Usage

```bash
# Start the application
task-tracker

# Common commands
task-tracker -add "New task" -description "This is a new task"
task-tracker -list todo
task-tracker -update "Updated task" -descripiton "This is a updated task" [task-id]
task-tracker -status completed [task-id]
```

## Development

1. Clone the repository
2. Install dependencies
3. Build the project:

```bash
go build -o task-tracer-cli ./cmd/task-tracer-cli
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
