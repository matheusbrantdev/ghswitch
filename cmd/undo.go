package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matheusbrantdev/ghswitch/internal/backup"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Restore SSH and Git config to pre-ghswitch state",
	RunE:  runUndo,
}

func runUndo(_ *cobra.Command, _ []string) error {
	if !backup.Exists() {
		fmt.Println("No backup found.")
		return nil
	}

	if err := backup.Restore(); err != nil {
		return err
	}

	home, _ := os.UserHomeDir()
	os.Remove(filepath.Join(home, ".ghswitch", "active"))

	fmt.Println("Restored to original state.")
	return nil
}
