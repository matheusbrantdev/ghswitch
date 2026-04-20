package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/matheusbrantdev/ghswitch/internal/profile"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Register a new GitHub profile",
	RunE:  runAdd,
}

func runAdd(_ *cobra.Command, _ []string) error {
	var (
		name     string
		gitName  string
		gitEmail string
		sshKey   string
	)

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
			huh.NewInput().Title("Profile name").Value(&name),
			huh.NewInput().Title("Git name").Value(&gitName),
			huh.NewInput().Title("Git email").Value(&gitEmail),
			huh.NewSelect[string]().Title("SSH key").Options(keyOptions...).Value(&sshKey),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	p := profile.Profile{
		Name:     name,
		GitName:  gitName,
		GitEmail: gitEmail,
		SSHKey:   sshKey,
	}

	if err := profile.Add(p); err != nil {
		return err
	}

	fmt.Printf("Profile %q saved\n", name)
	return nil
}

func listSSHKeys() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(filepath.Join(home, ".ssh"))
	if err != nil {
		return nil, err
	}

	var keys []string
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || filepath.Ext(name) == ".pub" || !strings.HasPrefix(name, "id") {
			continue
		}
		keys = append(keys, filepath.Join(home, ".ssh", name))
	}
	return keys, nil
}
