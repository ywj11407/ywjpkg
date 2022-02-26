// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pathutil

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Clean cleans up given path and returns a relative path that goes straight down.
func Clean(p string) string {
	return strings.Trim(path.Clean("/"+p), "/")
}

func WorkDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(fmt.Sprintf("Read workdir fail: %s", err.Error()))
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func FileDir(path string, userId string) (string, bool) {
	folderPath := filepath.Join(path, userId)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
		return folderPath, true
	}
	return folderPath, false
}

func CreateDateDir(Path string) string {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}
