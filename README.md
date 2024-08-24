# Project RixLog

Webpage with artciles and pages.

Stack:
.Go
.Htmx
.Sqlite3

## Getting Started

. make build
. ./bin/rixlog

## MakeFile

Run all make commands with clean tests
```bash
make all build
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

Live reload the application
```bash
make watch
```

Run the test suite
```bash
make test
```

Clean up binary from the last build
```bash
make clean
```
