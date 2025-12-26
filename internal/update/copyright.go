package update

import (
	"os"
	"strings"

	"go.mattglei.ch/newyear/internal/api"
)

func Copyright(repo api.Repo) (bool, error) {
	err := os.Chdir(repo.Name)
	if err != nil {
		return false, err
	}

	fsObjects, err := os.ReadDir(".")
	if err != nil {
		return false, err
	}
	updated := false
	for _, fsObject := range fsObjects {
		name := fsObject.Name()
		if !fsObject.IsDir() && strings.Contains(strings.ToLower(name), "license") {
			b, err := os.ReadFile(name)
			if err != nil {
				return false, err
			}

			content := string(b)
			patchedFile := strings.ReplaceAll(content, os.Args[1], os.Args[2])
			if content != patchedFile {
				updated = true
				err = os.WriteFile(name, []byte(patchedFile), fsObject.Type().Perm())
				if err != nil {
					return false, err
				}
			}
		}
	}
	return updated, nil
}
