// リクエスト情報から取り出したデータ
// リクエストによって定義される
// ページ関連のファイルを取得するための情報
package types

import (
	"path/filepath"
	"strings"
)

type PageInfo struct {
	Contents string // dev/contents/*
	Package  string // dev/contents/*/page/X
	Method   string // dev/contents/*/page/X/X.go.Method
}

// 初期化処理
func (info *PageInfo) Init(contents, packageName, method string) {
	*info = PageInfo{
		Contents: contents,
		Package:  packageName,
		Method:   method,
	}
}

// ピリオド区切りでメソッド名のフルパスを表示
func (pi *PageInfo) FullMethod() string {
	return strings.Join([]string{pi.Contents, pi.Package, pi.Method}, ".")
}

// 相対パスを取得する
func (pi *PageInfo) ContentsPath() string {
	path := filepath.Join("contents", pi.Contents)
	return path
}
func (pi *PageInfo) PackagePath() string {
	path := filepath.Join(pi.ContentsPath(), "page", pi.Package)
	return path
}
