package file

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// ファイルのコピー
func FileCopy(src, dst string) (err error) {
	var info os.FileInfo
	var txt []byte

	info, err = os.Stat(src)
	if info.IsDir() {
		if err = os.Mkdir(dst, 0755); err != nil {
			return
		}
	} else {
		txt, err = ioutil.ReadFile(src)
		if err != nil {
			return
		}

		err = ioutil.WriteFile(dst, txt, 0755)
		if err != nil {
			return
		}
	}

	return nil
}

// ディレクトリの検索
func SearchDirectory(target string) (info []string, err error) {
	var fileinfo []os.FileInfo

	info = make([]string, 0)

	fileinfo, err = ioutil.ReadDir(target)
	for i := 0; i < len(fileinfo); i++ {
		if !strings.HasPrefix(fileinfo[i].Name(), ".") {
			if fileinfo[i].IsDir() {

				subDir := fmt.Sprintf("%s/%s", target, fileinfo[i].Name())
				subinfo, subErr := SearchDirectory(subDir)
				if err != nil {
					err = subErr
					return
				}

				info = append(info, fileinfo[i].Name())
				for j := 0; j < len(subinfo); j++ {
					info = append(info, fmt.Sprintf("%s/%s", fileinfo[i].Name(), subinfo[j]))
				}
			} else {
				info = append(info, fmt.Sprintf("%s", fileinfo[i].Name()))
			}
		}
	}

	return
}

