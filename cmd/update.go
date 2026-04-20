package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/matheusbrantdev/ghswitch/internal/profile"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a saved profile",
	RunE:  runUpdate,
}

func runUpdate(_ *cobra.Command, _ []string) error {
	profiles, err := profile.Load()
	if err != nil {
		return err
	}

	if len(profiles) == 0 {
		fmt.Println("No profiles saved.")
		return nil
	}

	var name string

	opts := make([]huh.Option[string], len(profiles))
	for i, p := range profiles {
		opts[i] = huh.NewOption(p.Name, p.Name)
	}

	if err := huh.NewSelect[string]().
		Title("Select profile to update").
		Options(opts...).
		Value(&name).
		Run(); err != nil {
		return err
	}

	var current profile.Profile
	for _, p := range profiles {
		if p.Name == name {
			current = p
			break
		}
	}

	keys, err := listSSHKeys()
	if err != nil {
		return err
	}

	keyOptions := make([]huh.Option[string], len(keys))
	for i, k := range keys {
		keyOptions[i] = huh.NewOption(filepath.Base(k), k)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Profile name").Value(&current.Name),
			huh.NewInput().Title("Git name").Value(&current.GitName),
			huh.NewInput().Title("Git email").Value(&current.GitEmail),
			huh.NewSelect[string]().Title("SSH key").Options(keyOptions...).Value(&current.SSHKey),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	updated := make([]profile.Profile, len(profiles))
	for i, p := range profiles {
		if p.Name == name {
			updated[i] = current
		} else {
			updated[i] = p
		}
	}

	if err := profile.Save(updated); err != nil {
		return err
	}

	fmt.Printf("Profile %q updated.\n", current.Name)
	return nil
}
