// ページ関連のファイルを取得するための情報
// ページとは、パッケージとメソッドによって定義される
package types

import (
	"path/filepath"
	"strings"
)

type PageInfo struct {
	Contents string // dev/contents/*
	Package  string // dev/contents/*/page/X
	Method   string // dev/contents/*/page/X/X.go.Method
	// デフォルトでMethod名と同じ、変更可能
	Template string // dev/contents/*/page/X/template/*.tpl
	Layout   string

	// テンプレートファイルを追加するため
	TemplateFiles []string
	// 主にテンプレートに渡すデータたち
	AssignData map[string]interface{}
}

func (info *PageInfo) Init(contents, packageName, method string) {
	*info = PageInfo{
		Contents:      contents,
		Package:       packageName,
		Method:        method,
		Template:      method,
		Layout:        "default",
		TemplateFiles: []string{},
		AssignData: map[string]interface{}{
			"Title": packageName + " - " + contents,
		},
	}
}

// データを追加
func (pi *PageInfo) AddTpl(fileNames ...string) {
	pi.TemplateFiles = append(pi.TemplateFiles, fileNames...)
}
func (pi *PageInfo) Assign(key string, value interface{}) {
	pi.AssignData[key] = value
}

// 相対パスを取得する
func (pi *PageInfo) ContentsPath() string {
	path := filepath.Join("contents", pi.Contents)
	return path
}
func (pi *PageInfo) PagePath() string {
	path := filepath.Join(pi.ContentsPath(), "page", pi.Package)
	return path
}
func (pi *PageInfo) PageTemplatePath() string {
	path := filepath.Join(pi.PagePath(), "template", pi.Template+".tpl")
	return path
}

func (pi *PageInfo) LayoutPath() string {
	path := filepath.Join(pi.ContentsPath(), "template", "layout", pi.Layout+".tpl")
	return path
}

// ピリオド区切りでメソッド名のフルパスを表示
func (pi *PageInfo) FullMethod() string {
	return strings.Join([]string{pi.Contents, pi.Package, pi.Method}, ".")
}
