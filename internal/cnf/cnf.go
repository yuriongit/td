/*
Package cnf provides helpers for working
with td's configuration directory.
*/
package cnf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

/*
ChdirToConfigDir changes the current directory
to td's config directory.
*/
func ChdirToConfigDir() (err error) {
	if err := constructConfigPath(); err != nil {
		return err
	}
	if err := os.Chdir(configPath); err != nil {
		return errors.New("config directory not found: run td init")
	}

	return nil
}

/*
GetEnvVariable retrieves an environment
variable from the config directory. Must
call be in config directory first.
*/
func GetEnvVariable(varName string) (
	variable string,
	err error,
) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if currentPath != configPath {
		return "", errors.New("current directory is not %s")
	}

	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("")
		return "", fmt.Errorf("Environment variables load error: %w\nMaybe run 'td init'?", err)
	}

	// Retrieves DB_URI environment variable.
	return os.Getenv(varName), nil
}

var configPath string

func constructConfigPath() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Failed to get home directory: %w", err)
	}

	configPath = filepath.Join(homeDir, ".td")

	return nil
}
