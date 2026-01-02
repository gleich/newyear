package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go.mattglei.ch/newyear/internal/api"
	"go.mattglei.ch/newyear/internal/out"
	"go.mattglei.ch/newyear/internal/update"
	"go.mattglei.ch/timber"
)

func main() {
	timber.Timezone(time.Local)
	timber.TimeFormat("03:04:05")

	pat := out.Ask("What is your PAT (personal access token)?")
	if pat == "" || !strings.HasPrefix(pat, "ghp_") {
		timber.FatalMsg("Please enter a valid personal access token")
	}

	client := api.Client(pat)

	repos, err := api.Repos(client)
	if err != nil {
		timber.Fatal(err, "Failed to load repos")
	}

	tmpDir, err := update.CreateTmpDir()
	if err != nil {
		timber.Fatal(err, "Failed to create temp directory for cloning")
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		timber.Fatal(err, "Failed to change directory to temporary directory for cloning")
	}

	updates := 0
	for i, repo := range repos {
		if repo.IsArchived || repo.IsDisabled || repo.IsEmpty || repo.IsFork || repo.IsMirror {
			continue
		}
		err = update.Clone(repo)
		if err != nil {
			timber.Fatal(err, "Failed to clone", repo.NameWithOwner)
		}
		timber.Done("Cloned", repo.NameWithOwner, fmt.Sprintf("(%v/%v)", i+1, len(repos)))

		updated, err := update.Copyright(repo)
		if err != nil {
			timber.Fatal(err, "Failed to update copyright for", repo.NameWithOwner)
		}

		if updated {
			updates++
			err = update.Commit(repo)
			if err != nil {
				timber.Fatal(err, "Failed to commit & push changes for", repo.NameWithOwner)
			}
		}

		err = os.Chdir("..")
		if err != nil {
			timber.Fatal(err, "Failed to change directory up out of the repository")
		}
	}

	fmt.Println()
	timber.Done("Updated", updates, "repositories from", os.Args[1], "to", os.Args[2])

	err = os.RemoveAll(tmpDir)
	if err != nil {
		timber.Fatal(
			err,
			"Failed to remove temporary directory. Please remove",
			tmpDir,
			"manually.",
		)
	}
}
