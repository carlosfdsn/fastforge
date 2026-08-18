package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
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

	folders := []string{
		"app",
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
		"requirements.txt",
		".gitignore",
		".env",
		".env.example",
		"README.md",
	}

	for _, folder := range folders {
		path := filepath.Join(f.Name, folder)

		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	for _, file := range files {
		path := filepath.Join(f.Name, file)

		var content []byte

		switch file {
		case "app/main.py":
			template, err := templatesFS.ReadFile("templates/main.py")
			if err != nil {
				fmt.Println("Error reading template:", err)
				return
			}

			content = template

		case ".gitignore":
			template, err := templatesFS.ReadFile("templates/gitignore.txt")
			if err != nil {
				fmt.Println("Error reading template:", err)
				return
			}

			content = template

		case "requirements.txt":
			template, err := templatesFS.ReadFile("templates/requirements.txt")
			if err != nil {
				fmt.Println("Error reading template:", err)
				return
			}

			content = template
		}

		err := os.WriteFile(path, content, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(path)
	}

	fmt.Printf("\nProject created: { Name: %s }\n", f.Name)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}