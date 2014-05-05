package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	myfile "github.com/tyokoyama/aetesthelper/file"
)

func main() {
	const GOPATH = "GOPATH"
	var (
		info     []string
		path     string
		err      error
		sdk_path = flag.String("sdk_path", "", "Appengine SDK PATH")
	)

	// 対象ディレクトリの取得
	flag.Parse()

	if *sdk_path == "" {
		flag.PrintDefaults()
	}

	if _, err := os.Stat(*sdk_path); err != nil {
		log.Fatalln("Appengine SDK PATH not found.")
	}

	// GAE/GのプロジェクトはSDKディレクトリ以下にあることが前提。
	arg := flag.Arg(0)
	target := fmt.Sprintf("%s/%s", *sdk_path, arg)

	// GOPATH変数のチェック
	original_gopath := os.Getenv(GOPATH)

	// 一時ディレクトリ作成して、GOPATH設定
	if path, err = ioutil.TempDir("", "aetesthelper"); err != nil {
		panic(err)
	}
	testDir := fmt.Sprintf("%s/src", path)

	if err = os.Mkdir(testDir, 0755); err != nil {
		panic(err)
	}

	if err = os.Setenv(GOPATH, path); err != nil {
		panic(err)
	}

	// ディレクトリの検索
	info, err = myfile.SearchDirectory(target)
	if err != nil {
		panic(err)
	}

	// サブパッケージのソースコードのコピー
	for i := 0; i < len(info); i++ {
		myfile.FileCopy(fmt.Sprintf("%s/%s", target, info[i]), fmt.Sprintf("%s/%s", testDir, info[i]))
	}

	// aetest
	err = os.Chdir(*sdk_path)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(info); i++ {
		var targetInfo os.FileInfo
		targetPackage := fmt.Sprintf("%s/%s", target, info[i])
		targetInfo, err = os.Stat(targetPackage)
		if err != nil {
			panic(err)
		}

		if targetInfo.IsDir() {
			fmt.Println(info[i])
			cmd := exec.Command("./goapp", "test", fmt.Sprintf("./%s/%s", arg, info[i]))

			cmd.Env = os.Environ()
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			if err = cmd.Start(); err != nil {
				panic(err)
			}

			// if err = cmd.Wait(); err != nil {
			// 	panic(err)
			// }
			cmd.Wait()
		}
	}

	// テスト終了後に一時ファイルを削除
	if err = os.RemoveAll(path); err != nil {
		panic(err)
	}

	// 環境変数を元に戻す。
	if err = os.Setenv(GOPATH, original_gopath); err != nil {
		panic(err)
	}
}
