package update

import (
	"os"
	"path/filepath"
)

func CreateTmpDir() (string, error) {
	loc := filepath.Join(os.TempDir(), "newyear")
	if _, err := os.Stat(loc); err == nil {
		err := os.RemoveAll(loc)
		if err != nil {
			return "", err
		}
	}
	err := os.Mkdir(loc, 0755)
	if err != nil {
		return "", err
	}
	return loc, nil
}
