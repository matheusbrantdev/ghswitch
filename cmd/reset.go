package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/matheusbrantdev/ghswitch/internal/profile"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete all saved profiles",
	RunE:  runReset,
}

func runReset(_ *cobra.Command, _ []string) error {
	var confirm bool

	if err := huh.NewConfirm().
		Title("Delete all profiles?").
		Affirmative("No").
		Negative("Yes").
		Value(&confirm).
		Run(); err != nil {
		return err
	}

	if confirm {
		fmt.Println("Aborted.")
		return nil
	}

	if err := profile.Save(nil); err != nil {
		return err
	}

	home, _ := os.UserHomeDir()
	os.Remove(filepath.Join(home, ".ghswitch", "active"))

	fmt.Println("All profiles deleted.")
	return nil
}
