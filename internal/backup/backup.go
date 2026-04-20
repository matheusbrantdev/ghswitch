package backup

import (
	"os"
	"path/filepath"
)

func dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	d := filepath.Join(home, ".ghswitch", "backup")
	return d, os.MkdirAll(d, 0700)
}

func Save() error {
	d, err := dir()
	if err != nil {
		return err
	}

	home, _ := os.UserHomeDir()

	for _, src := range []struct{ from, to string }{
		{filepath.Join(home, ".gitconfig"), filepath.Join(d, "gitconfig")},
		{filepath.Join(home, ".ssh", "config"), filepath.Join(d, "ssh_config")},
	} {
		data, err := os.ReadFile(src.from)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		if err := os.WriteFile(src.to, data, 0600); err != nil {
			return err
		}
	}
	return nil
}

func Restore() error {
	d, err := dir()
	if err != nil {
		return err
	}

	home, _ := os.UserHomeDir()

	for _, src := range []struct{ from, to string }{
		{filepath.Join(d, "gitconfig"), filepath.Join(home, ".gitconfig")},
		{filepath.Join(d, "ssh_config"), filepath.Join(home, ".ssh", "config")},
	} {
		data, err := os.ReadFile(src.from)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		if err := os.WriteFile(src.to, data, 0600); err != nil {
			return err
		}
	}
	return nil
}

func Exists() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	_, err = os.Stat(filepath.Join(home, ".ghswitch", "backup", "gitconfig"))
	return err == nil
}
