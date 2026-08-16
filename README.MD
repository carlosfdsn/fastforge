# FastForge

FastForge is a command-line tool written in Go for bootstrapping FastAPI projects. It automates the initial scaffolding steps that are usually done by hand: creating a directory structure, generating base files from templates, and reporting what was created.

## Status

FastForge is an early-stage project. The current version implements a minimal, linear scaffolding flow. It is not yet a full-featured scaffolding tool — see [Planned Features](#planned-features) for what is not implemented yet.

## What FastForge Does Today

When you run FastForge, it performs the following steps in order:

1. Clears the terminal.
2. Prints an ASCII logo.
3. Prompts the user for a project name via standard input.
4. Reads the project name from the terminal.
5. Creates a directory structure for the new project.
6. Creates the necessary project files.
7. Reads template files from a local `templates/` folder.
8. Writes the contents of those templates into the generated project.
9. Displays the list of files that were created.
10. Prints a confirmation message that the project was created.

Steps 5 through 10 are performed by a `FastForge` struct through its `CreateProject` method, which is invoked after the project name is captured from the terminal.

## How It Works (Current Implementation)

The entry point (`main.go`) does the following:

- Clears the terminal using `cmd /c cls` on Windows (detected via the `OS` environment variable) or `clear` on other systems.
- Prints a colored ASCII logo using raw ANSI escape codes (`cyan`, `gray`, `green`, `reset` constants are defined for this purpose).
- Reads a single line from standard input using `bufio.Reader` to capture the project name, trimming surrounding whitespace with `strings.TrimSpace`.
- Instantiates a `FastForge` struct with the provided name and calls `project.CreateProject()`, which is responsible for generating the project structure, copying template files, and reporting the result.

The template files used to generate new projects live in a `templates/` directory. This directory is embedded directly into the compiled binary at build time using Go's `embed` package (`//go:embed templates/*`), so the resulting `fastforge` executable is self-contained and does not require the `templates/` folder to be present at runtime.

## Installation

FastForge is a Go program. To build it from source, you need Go installed on your machine.

```bash
git clone https://github.com/carlosfdsn/fastforge.git
cd fastforge
go build -o fastforge .
```

This produces a `fastforge` binary in the project directory.

## Usage

Run the binary from a terminal:

```bash
./fastforge
```

You will be prompted for a project name:

```
? Project name: my-api
```

After entering a name, FastForge creates the project directory structure, writes the template files into it, and prints the list of files it created along with a confirmation message.

## Requirements

- Go (to build FastForge from source)
- A `templates/` directory present in the source tree at build time (it is embedded into the binary via `go:embed`, so it is not required afterwards to run the compiled executable)

## Planned Features

The following are ideas for future versions of FastForge and are **not implemented yet**:

- A more complete FastAPI scaffolding tool, beyond the current minimal flow
- Support for different project layouts or architectural styles (e.g. Clean Architecture)
- Command-line flags/arguments instead of only interactive prompts
- Configurable or selectable templates
- Cross-platform packaging and distribution
- Templates for the remaining generated files that are currently created empty (e.g. `app/core/config.py`, `app/core/security.py`, `app/db/database.py`, the generated project's `README.md`)

## Contributing

This project is under active development. Issues and pull requests are welcome, but keep in mind the codebase is still small and evolving, and interfaces may change without notice.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.