package internal

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

func SetupProject(projectName string) error {

	directories := []string{
		"controller",
		"routes",
		"middleware",
		"vendor",
	}

	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("Error creating directory %s: %w", dir, err)
		}
	}

	exampleRoutePath := filepath.Join("routes", "routes.go")

	if err := generateFileFromTemplate(exampleRoutePath, "templates/api.tmpl"); err != nil {
		return err
	}

	exampleControllerPath := filepath.Join("controller", "hello_world_controller.go")

	if err := generateFileFromTemplate(exampleControllerPath, "templates/hello_world_controller.tmpl"); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	_, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Error executing go mod init")
	}

	cmd2 := exec.Command("go", "mod", "tidy")
	_, err2 := cmd2.CombinedOutput()

	if err2 != nil {
		return fmt.Errorf("Error executing go mod tidy")
	}

	fmt.Println("Projektstruktur wurde erfolgreich eingerichtet.")
	return nil

}

func generateFileFromTemplate(filePath, templatePath string) error {
	tmpl, err := template.ParseFS(templatesFS, templatePath)
	if err != nil {
		return fmt.Errorf("Fehler beim Laden des Templates: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Fehler beim Erstellen der Datei %s: %w", filePath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, nil); err != nil {
		return fmt.Errorf("Fehler beim Schreiben der Datei %s: %w", filePath, err)
	}
	return nil
}
