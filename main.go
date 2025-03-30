package main

import (
	"fmt"
	"github.com/mreza-mohammadzadeh/createapp/core"
	"log"
	"os"
	"os/exec"
)

// FolderStructure defines a hierarchical structure that allows unlimited subdirectories.
var FolderStructure = map[string]interface{}{
	"delivery": map[string]interface{}{
		"httpserver": map[string]interface{}{
			"login.go":   core.SampleLoginHandlerCode(),
			"route.go":   core.SampleRouteCode(),
			"handler.go": core.SampleHandlerCode(),
		},
		"grpcserver": map[string]interface{}{},
	},
	"cmd": map[string]interface{}{
		"build.go": core.SampleBuildCode(),
	},
	"migrator": map[string]interface{}{
		"migrations": map[string]interface{}{
			"1702146900_add_users_table.sql": core.SampleMigrationCode(),
		},
		"provider.go": core.SampleProviderCode(),
	},
	"entity": map[string]interface{}{
		"user.go": core.SampleUserEntityCode(),
	},
	"param": map[string]interface{}{
		"get_by_id.go": core.SampleParamsCode(),
	},
	"repository": map[string]interface{}{
		"db.go":        core.SampleDBRepositoryCode(),
		"get_by_id.go": core.SampleRepositoryCode(),
	},
	"service": map[string]interface{}{
		"service.go":   core.SampleServiceCode(),
		"get_by_id.go": core.SampleMethodServiceCode(),
	},
}

func main() {
	if len(os.Args) < 4 || os.Args[1] != "create" || os.Args[2] != "app" {
		log.Fatal("Usage: create app <AppName> [Directory]")
	}

	appName := os.Args[3]
	directory := "."

	if len(os.Args) > 4 {
		directory = os.Args[4]
	}

	if core.IsAppNameDuplicate(appName) {
		log.Fatalf("❌ Error: Project '%s' already exists. Please choose a different name.", appName)
	}

	finalPath := fmt.Sprintf("%s/%s", directory, appName)

	if err := core.GenerateProjectStructure(finalPath, FolderStructure); err != nil {
		log.Fatalf("Error generating project: %v", err)
	}

	fmt.Println("✅ Project structure created successfully!")
	if err := installCLI(); err != nil {
		log.Fatalf("Error installing CLI: %v", err)
	}
}

// installCLI builds and installs the CLI tool globally.
func installCLI() error {
	if err := exec.Command("go", "build", "-o", "create").Run(); err != nil {
		return fmt.Errorf("failed to build CLI: %w", err)
	}

	if err := exec.Command("sudo", "mv", "create", "/usr/local/bin/").Run(); err != nil {
		return fmt.Errorf("failed to move CLI binary: %w", err)
	}

	if err := exec.Command("chmod", "+x", "/usr/local/bin/create").Run(); err != nil {
		return fmt.Errorf("failed to set executable permission: %w", err)
	}

	fmt.Println("✅ CLI installed successfully! Use 'create app <AppName>' anywhere.")
	return nil
}
