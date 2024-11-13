package main

import (
	"embed"
	"fmt"
	"goopher/internal"
	"os"
	"strings"
	"text/template"
)

//go:embed internal/templates/*.tmpl
var templatesFS embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goval setup <project> or make:<type> name")
		return
	}

	command := os.Args[1]

	if command == "setup" {
		if err := internal.SetupProject(os.Args[2]); err != nil {
			fmt.Println("Setup failed:", err)
		}

		return
	}

	if !strings.HasPrefix(command, "make:") {
		fmt.Println("No command found. Usage: make:<type> Name")
		return
	}

	commandType := strings.TrimPrefix(command, "make:")

	if len(os.Args) < 3 {
		fmt.Printf("Provide a name for %s\n", commandType)
		return
	}

	name := os.Args[2]

	switch commandType {
	case "controller":
		createController(name)
	default:
		fmt.Printf("Unknown type: %s\n", commandType)
	}
}

func createController(controllerName string) {
	fmt.Printf("Creating controller: %s\n", controllerName)

	// TODO: Put those base paths in config
	dirName := "controllers"
	fileName := fmt.Sprintf("%s/%s.go", dirName, strings.ToLower(controllerName))

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory", err)
			return
		}
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	tmpl, err := template.ParseFS(templatesFS, "internal/templates/controller.tmpl")

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(file, struct{ ControllerName string }{ControllerName: controllerName})
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("%s created!\n", fileName)
}
