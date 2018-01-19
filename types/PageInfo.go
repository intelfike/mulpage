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
	Method   string // dev/contents/*/page/X/X.go.Method
}

// 初期化処理
func (info *PageInfo) Init(r *http.Request) error {
	*info = PageInfo{"homepage", "top", "Index"} // デフォルト値
	// パスを取得
	path := httpio.ReadPath(r)
	if path != nil && !strings.HasPrefix(path[0], "_") {
		return errors.New("Method呼び出しではありません")
	}
	// ページ情報を定義
	info.Contents = httpio.ReadContents(r)
	if len(path) >= 1 {
		info.Package = strings.TrimPrefix(path[0], "_")
	}
	if len(path) >= 2 {
		info.Method = path[1]
	}
	return nil
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
