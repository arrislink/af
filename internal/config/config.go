package config

import (
	"os"
	"path/filepath"
)

func AIDir() string {
	return filepath.Join(os.Getenv("PWD"), ".ai")
}

func ContextDir() string {
	return filepath.Join(os.Getenv("PWD"), "context")
}

func TasksDir() string {
	return filepath.Join(os.Getenv("PWD"), "tasks")
}
