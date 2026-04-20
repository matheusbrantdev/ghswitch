package profile

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type store struct {
	Profiles []Profile `yaml:"profiles"`
}

func storePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".ghswitch", "profiles.yml"), nil
}

func Load() ([]Profile, error) {
	path, err := storePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return []Profile{}, nil
	}
	if err != nil {
		return nil, err
	}

	var s store
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return s.Profiles, nil
}

func Save(profiles []Profile) error {
	path, err := storePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	data, err := yaml.Marshal(store{Profiles: profiles})
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func ActiveName() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath.Join(home, ".ghswitch", "active"))
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(data), err
}

func SetActive(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(home, ".ghswitch", "active"), []byte(name), 0600)
}

func Add(p Profile) error {
	profiles, err := Load()
	if err != nil {
		return err
	}
	profiles = append(profiles, p)
	return Save(profiles)
}
