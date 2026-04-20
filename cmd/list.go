package cmd

import (
	"fmt"

	"github.com/matheusbrantdev/ghswitch/internal/profile"
	"github.com/matheusbrantdev/ghswitch/internal/ui"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved profiles",
	RunE:  runList,
}

func runList(_ *cobra.Command, _ []string) error {
	profiles, err := profile.Load()
	if err != nil {
		return err
	}

	if len(profiles) == 0 {
		fmt.Println("No profiles saved.")
		return nil
	}

	for _, p := range profiles {
		fmt.Printf("%s\n", ui.Name.Render(p.Name))
		fmt.Printf("  %s %s %s\n", ui.Muted.Render("git:"), p.GitName, ui.Email.Render("<"+p.GitEmail+">"))
		fmt.Printf("  %s %s\n\n", ui.Muted.Render("key:"), ui.Key.Render(p.SSHKey))
	}
	return nil
}
