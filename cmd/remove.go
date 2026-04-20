package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/matheusbrantdev/ghswitch/internal/profile"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a saved profile",
	RunE:  runRemove,
}

func runRemove(_ *cobra.Command, _ []string) error {
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
		Title("Select profile to remove").
		Options(opts...).
		Value(&name).
		Run(); err != nil {
		return err
	}

	updated := make([]profile.Profile, 0, len(profiles)-1)
	for _, p := range profiles {
		if p.Name != name {
			updated = append(updated, p)
		}
	}

	if err := profile.Save(updated); err != nil {
		return err
	}

	fmt.Printf("Profile %q removed.\n", name)
	return nil
}
