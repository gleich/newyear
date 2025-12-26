package update

import (
	"fmt"
	"os"
	"os/exec"

	"go.mattglei.ch/newyear/internal/api"
	"go.mattglei.ch/timber"
)

func Commit(repo api.Repo) error {
	err := exec.Command("git", "add", ".").Run()
	if err != nil {
		return err
	}

	err = exec.Command("git", "commit", "-m", fmt.Sprintf("%v -> %v", os.Args[1], os.Args[2])).Run()
	if err != nil {
		return err
	}

	err = exec.Command("git", "push").Run()
	if err != nil {
		return err
	}

	timber.Info("Updated", repo.NameWithOwner, "to new year")
	return nil
}
