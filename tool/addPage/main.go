package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	filemod "github.com/intelfike/module/file"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("addPage CONTENTS PAGE...")
		os.Exit(1)
	}
	// コンテンツがあるかどうか
	con := args[0]
	gopath, _ := os.LookupEnv("GOPATH")
	fullpath := filepath.Join(gopath, "src/github.com/intelfike/mulpage/contents", con)
	if !filemod.ExistsDir(fullpath) {
		fmt.Println(fullpath)
		fmt.Println("CONTENTSが見つかりませんでした")
		os.Exit(1)
	}

	// ページを作成する
	for _, v := range args[1:] {
		newPagePath := filepath.Join(fullpath, "page", v)
		if filemod.ExistsDir(newPagePath) {
			fmt.Println(v, ":パッケージは既に存在するためスキップします")
			continue
		}
		// 新規ディレクトリ作成
		if err := os.Mkdir(newPagePath, 0777); err != nil {
			fmt.Println(err)
			continue
		}
		if err := os.Mkdir(filepath.Join(newPagePath, "template"), 0777); err != nil {
			fmt.Println(err)
			continue
		}

		// ファイル作成
		// tplファイル
		ftpl, err := os.Create(filepath.Join(newPagePath, "template/Index.tpl"))
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer ftpl.Close()
		forigin, err := os.Open("template/Index.tpl")
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer forigin.Close()
		io.Copy(ftpl, forigin)

		// goファイル
		fgo, err := os.Create(filepath.Join(newPagePath, v+".go"))
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer fgo.Close()
		tplfiles, err := template.ParseFiles("template/tpl.go")
		if err != nil {
			fmt.Println(err)
			continue
		}
		tplfiles.Execute(fgo, struct{ Package string }{v})
	}
}
