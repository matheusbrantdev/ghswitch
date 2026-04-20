package cmd

import (
	"fmt"

	"github.com/matheusbrantdev/ghswitch/internal/git"
	"github.com/matheusbrantdev/ghswitch/internal/profile"
	internalssh "github.com/matheusbrantdev/ghswitch/internal/ssh"
	"github.com/matheusbrantdev/ghswitch/internal/ui"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [profile]",
	Short: "Switch to a profile",
	Args:  cobra.ExactArgs(1),
	RunE:  runUse,
}

func runUse(_ *cobra.Command, args []string) error {
	name := args[0]

	profiles, err := profile.Load()
	if err != nil {
		return err
	}

	var p *profile.Profile
	for i := range profiles {
		if profiles[i].Name == name {
			p = &profiles[i]
			break
		}
	}

	if p == nil {
		return fmt.Errorf("profile %q not found", name)
	}

	if err := internalssh.SetGitHubKey(p.SSHKey); err != nil {
		return fmt.Errorf("updating SSH config: %w", err)
	}

	if err := git.SetGlobal(p.GitName, p.GitEmail); err != nil {
		return fmt.Errorf("updating git config: %w", err)
	}

	if err := profile.SetActive(name); err != nil {
		return err
	}

	fmt.Printf("Switched to %s\n", ui.Name.Render(name))
	fmt.Printf("  %s %s %s\n", ui.Muted.Render("git:"), p.GitName, ui.Email.Render("<"+p.GitEmail+">"))
	fmt.Printf("  %s %s\n", ui.Muted.Render("key:"), ui.Key.Render(p.SSHKey))
	return nil
}
