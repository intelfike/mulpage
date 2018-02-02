// リクエスト情報から取り出したデータ
// リクエストによって定義される
// ページ関連のファイルを取得するための情報
package types

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"

	httpio "github.com/intelfike/mulpage/io/http"
)

type PageInfo struct {
	Contents string // dev/contents/*
	Package  string // dev/contents/*/page/X
	Article  string // dev/contents/*/page/X/X.go.Function
	ExecPath []string
}

// 初期化処理
func (info *PageInfo) Init(r *http.Request) error {
	// コンテンツを取得
	info.ExecPath = []string{"www", "top", "Index"}
	info.ExecPath[0] = httpio.ReadContents(r)
	info.Contents = info.ExecPath[0]
	// パスを取得
	path := httpio.ReadPath(r)
	if path != nil && !strings.HasPrefix(path[0], "_") {
		return errors.New("Method呼び出しではありません")
	}
	// ページ情報を定義
	if len(path) >= 1 {
		info.ExecPath[1] = strings.TrimPrefix(path[0], "_")
	}
	if len(path) >= 2 {
		info.ExecPath[2] = path[1]
	}
	info.Package = info.ExecPath[1]
	info.Article = info.ExecPath[2]
	return nil
}

// ピリオド区切りでメソッド名のフルパスを表示
func (pi *PageInfo) FullMethod() string {
	return strings.Join([]string{pi.Contents, pi.Package, pi.Article}, ".")
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
