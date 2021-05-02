package fs

import (
	"os"
)

func GetTile(path string) {
	if FileExists(path) {

	}
}

func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
