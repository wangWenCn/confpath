package conf

import (
	"os"
	"path/filepath"
	"strings"
)

// 查找项目根目录
func FindGoModPath() (string, bool) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", false
	}
	absDir = strings.ReplaceAll(absDir, `\`, `/`)
	var rootPath string
	tempPath := absDir
	hasGoMod := false
	for {
		_, err = os.Stat(filepath.Join(tempPath, "go.mod"))
		if err == nil {
			rootPath = tempPath
			hasGoMod = true
			break
		}
		if tempPath == filepath.Dir(tempPath) {
			break
		}
		tempPath = filepath.Dir(tempPath)
		if tempPath == string(filepath.Separator) {
			break
		}
	}
	if hasGoMod {
		return rootPath, true
	}
	return "", false
}
