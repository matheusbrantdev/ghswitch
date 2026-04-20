package git

import (
	"fmt"
	"os/exec"
)

func SetGlobal(name, email string) error {
	for _, args := range [][]string{
		{"config", "--global", "user.name", name},
		{"config", "--global", "user.email", email},
	} {
		if out, err := exec.Command("git", args...).CombinedOutput(); err != nil {
			return fmt.Errorf("git %v: %s", args, out)
		}
	}
	return nil
}
