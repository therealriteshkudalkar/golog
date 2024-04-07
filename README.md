# golog - A library with a built-in log manager

With an intent to simplify logging and reuse the logging code, this library has a built-log manger
that works with Go's built-in log/slog package.

## Features

- Logging to the standard output
  - Log levels are colour coded (RED: Error, WARNING: ORANGE, BLUE: INFO, GREEN: DEBUG)
  - Log levels are specified along with the timestamp
- Logging to a file
  - Logs are written to a file specified while creating the logger
  - Log levels are specified along with the timestamp


## Installation

Install the package with the following command

```zsh
go get github.com/therealriteshkudalkar/golog
```

## Example

### 1. Creating a logger that logs into a file

### 2. Creating a logger that logs into standard output
