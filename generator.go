package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//go:embed templates/*
var templatesFS embed.FS

type FastForge struct {
	Name string
}

func (f FastForge) CreateProject() {

	if fileExists(f.Name) {
		fmt.Printf("Error: project '%s' already exists.\n", f.Name)
		return
	}

	if err := os.MkdirAll(f.Name, 0755); err != nil {
		errorMessage(fmt.Sprintf("Failed to create project directory: %v", err))
		return
	}

	fmt.Println()
	fmt.Printf("%sCreating %s%s\n", cyan, f.Name, reset)
	fmt.Println()

	fmt.Printf("%sSetup%s\n", gray, " ──────────────────────────")
	if !isUVInstalled() {
		infoMessage("uv not found. Installing...")

		if err := installUV(); err != nil {
			errorMessage(fmt.Sprintf("Failed to install uv: %v", err))
			return
		}

		successMessage("uv installed ")
	} else {
		successMessage("uv found")
	}

	if err := runCommand(f.Name, "uv", "init"); err != nil {
		errorMessage(fmt.Sprintf("Failed to initialize uv project: %v", err))
		return
	}

	successMessage("Python project initialized\n")

	dependencies := []string {
		"fastapi",
		"sqlalchemy",
		"pydantic",
		"pydantic-settings",
		"alembic",
		"psycopg[binary]",
		"pyjwt",
		"pwdlib[argon2]",
	}

	args := append([]string{"add"}, dependencies...)

	if err := runCommand(f.Name, "uv", args...); err != nil {
		errorMessage(fmt.Sprintf("Failed to install dependencies: %v", err))
		return
	}

	successMessage("\nDependencies installed\n")

	fmt.Printf("%sStructure%s\n", gray, " ──────────────────────────")

	folders := []string{
		"app",
		"app/api",
		"app/api/routes",
		"app/core",
		"app/db",
		"app/models",
		"app/schemas",
		"app/services",
		"app/repositories",
		"app/utils",
		"tests",
	}

	for _, folder := range folders {
		path := filepath.Join(f.Name, folder)

		if err := os.MkdirAll(path, 0755); err != nil {
			errorMessage(fmt.Sprintf("Failed to create folder %s: %v", folder, err))
			return
		}
	}

	successMessage("FastAPI architecture")
	successMessage("Application structure")
	successMessage("Test structure")

	fmt.Printf("\n%sFiles%s\n", gray, " ──────────────────────────")

	files := []string{
		"app/__init__.py",
		"app/main.py",
		"app/api/__init__.py",
		"app/api/routes/__init__.py",
		"app/core/__init__.py",
		"app/core/config.py",
		"app/core/security.py",
		"app/db/database.py",
		"app/models/__init__.py",
		"app/schemas/__init__.py",
		"app/services/__init__.py",
		"app/repositories/__init__.py",
		"app/utils/__init__.py",
		"tests/__init__.py",
		".gitignore",
		".env",
		".env.example",
		"README.md",
	}

	for _, file := range files {
		if err := createFile(f.Name, file); err != nil {
			errorMessage(fmt.Sprintf("Failed to create %s: %v", file, err))
			return
		}
	}

	successMessage(fmt.Sprintf("%d files created", len(files)))

	fmt.Println()
	fmt.Printf("%s─────────────────────────────────────────────%s\n", gray, reset)

	fmt.Printf("\n%sSuccess!%s\n\n", green, reset)

	fmt.Printf("%sProject:%s %s\n", gray, reset, f.Name)
	fmt.Printf("%sFiles:%s   %d\n", gray, reset, len(files))
	fmt.Printf("%sPackages:%s %d\n", gray, reset, len(dependencies))

	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println()
	fmt.Printf("  %scd %s%s\n", cyan, f.Name, reset)
	fmt.Printf("  %suv run fastapi dev app/main.py%s\n", cyan, reset)
	fmt.Println()

	if err := runCommand(f.Name, ".venv/Scripts/activate.ps1"); err != nil {

	}
}

func createFile(projectName string, file string) error {
	path := filepath.Join(projectName, file)

	var content []byte

	switch file {
	case "app/main.py":
		template, err := templatesFS.ReadFile("templates/main.py")
		if err != nil {
			return err
		}

		content = template

	case ".gitignore":
		template, err := templatesFS.ReadFile("templates/gitignore.txt")
		if err != nil {
			return err
		}

		content = template

	default:
		content = []byte{}
	}

	return os.WriteFile(path, content, 0644)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isUVInstalled() bool {
	_, err := exec.LookPath("uv")
	return err == nil
}

func installUV() error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("automatic uv installation is currently supported only on Windows")
	}

	cmd := exec.Command(
		"powershell",
		"-ExecutionPolicy",
		"ByPass",
		"-c",
		"irm https://astral.sh/uv/install.ps1 | iex",
	)

	cmd.Stdout = nil
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func successMessage(message string) {
	fmt.Printf("  %s✓%s %s\n", green, reset, message)
}

func errorMessage(message string) {
	fmt.Printf("  %s✗%s %s\n", red, reset, message)
}

func infoMessage(message string) {
	fmt.Printf("  %s•%s %s\n", yellow, reset, message)
}