/*
Package cnf is for initializing td's
This package prepares the needed
environment variables for the application to
use. Also, this package directly enables global
usage of td.
*/
package cnf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yuriongit/td/internal/domain"
)

var (
	envExampleValue = fmt.Sprintf("YOUR_%s", domain.DBConnStringVarName)
	envExample      = fmt.Sprintf("%s=%s\n", domain.DBConnStringVarName, envExampleValue)
)

/*
InitConfig initializes the td configuration directory and .env file.
*/
func InitConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Failed to get home directory: %w", err)
	}

	tdDir := filepath.Join(homeDir, ".td")

	// Ensure ~/.td exists and log status
	if _, err := os.Stat(tdDir); err == nil {
		fmt.Println("✓ Found existing ~/.td directory")
	} else {
		if err := os.MkdirAll(tdDir, 0700); err != nil {
			return fmt.Errorf("Failed to create ~/.td directory: %w", err)
		}
		fmt.Println("✓ Created ~/.td directory")
	}

	// Move into ~/.td early so subsequent file operations are relative
	if err := ChdirToConfigDir(); err != nil {
		return err
	}

	return setupEnvFile()
}

// setupEnvFile handles creation, loading, and verification of the local .env file.
func setupEnvFile() error {
	// os.O_EXCL creates the file atomically or fails if it already exists
	f, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)

	switch {
	case os.IsExist(err):
		fmt.Println("✓ Found existing .env file in ~/.td")
		return nil
	case err != nil:
		return fmt.Errorf("Failed to check or create .env file: %w", err)
	}

	_, writeErr := f.WriteString(envExample)
	if err := f.Close(); err != nil {
		return fmt.Errorf("Failed to close .env file: %w", err)
	}
	if writeErr != nil {
		return fmt.Errorf("Failed to write .env file: %w", writeErr)
	}
	fmt.Println("✓ Created .env file in ~/.td")

	// Load relative .env file
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("✗ Failed to read .env file: %w", err)
	}

	dbConnVar := os.Getenv(domain.DBConnStringVarName)
	if strings.Contains(dbConnVar, envExampleValue) {
		fmt.Printf(
			"◌ Please update %q in ~/.td/.env with your own actual values\n",
			domain.DBConnStringVarName,
		)
	} else {
		fmt.Println("✓ td already initialized")
	}

	return nil
}
