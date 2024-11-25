package systemx

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MatchParentDir returns target's directory's full path,
// returning error if `dir`'s parent dir names don't match `target`
func MatchParentDir(dir string, target string) (string, error) {
	var currentDir string
	var file string
	for {
		currentDir = filepath.Dir(dir)
		file = filepath.Base(dir)

		// Match target directory
		if file == target {
			return dir, nil
		}

		// Reach the top of directory
		if currentDir == dir {
			return "", fmt.Errorf(
				"diretory `%s` doesn't match `%s`", dir, target)
		}

		dir = currentDir
	}
}

// FindFilesBySuffix walks dir and return all files with specific suffix.
func FindFilesBySuffix(dir string, suffix string) ([]string, []string) {
	var trimSuffixRelResult []string
	var trimSuffixBaseResult []string
	if err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			ext := filepath.Ext(path)
			if ext == suffix {
				relName, err := filepath.Rel(dir, path)
				if err != nil {
					panic(err)
				}
				trimSuffixRel := strings.TrimSuffix(relName, ext)
				trimSuffixRelResult = append(trimSuffixRelResult, trimSuffixRel)

				baseName := filepath.Base(path)
				trimSuffixBase := strings.TrimSuffix(baseName, ext)
				trimSuffixBaseResult = append(trimSuffixBaseResult, trimSuffixBase)
			}
			return nil
		}); err != nil {
		panic(err)
	}
	return trimSuffixRelResult, trimSuffixBaseResult
}

// Exists returns if exists a dir or file
func Exists(s string) bool {
	_, err := os.Stat(s) // os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
