package update

import (
	"os"
	"os/exec"
	"slices"

	"go.mattglei.ch/newyear/internal/api"
)

func Clone(repo api.Repo) error {
	source := repo.URL + ".git"

	// switch to ssh if flag "--ssh" is passed in
	if slices.Contains(os.Args, "--ssh") {
		source = "git@github.com:" + repo.NameWithOwner
	}

	err := exec.Command("git", "clone", source).Run()
	if err != nil {
		return err
	}
	return nil
}
