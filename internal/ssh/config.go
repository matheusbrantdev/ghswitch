package ssh

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var githubHostRe = regexp.MustCompile(`(?m)^Host github\.com\b[^\n]*\n(?:[ \t]+[^\n]*\n)*`)

func SetGitHubKey(keyPath string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, ".ssh", "config")
	block := fmt.Sprintf("Host github.com\n  HostName github.com\n  User git\n  IdentityFile %s\n  IdentitiesOnly yes\n", keyPath)

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return os.WriteFile(path, []byte(block), 0600)
	}
	if err != nil {
		return err
	}

	current := string(data)
	var updated string
	if githubHostRe.MatchString(current) {
		updated = githubHostRe.ReplaceAllString(current, block)
	} else {
		if !strings.HasSuffix(current, "\n") {
			current += "\n"
		}
		updated = current + block
	}

	return os.WriteFile(path, []byte(updated), 0600)
}
