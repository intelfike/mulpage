// ページ関連のファイルを取得するための情報
package types

import (
	"path/filepath"
)

type PageInfo struct {
	Contents string // dev/contents/*
	Package  string // dev/contents/*/page/X
	Method   string // dev/contents/*/page/X/X.go.Method
	// デフォルトでMethod名と同じ、変更可能
	Template string // dev/contents/*/page/X/template/*.tpl
	Layout   string

	// 主にテンプレートに渡すデータたち
	AssignData map[string]string
}

func DefaultPageInfo(contents, packageName, method string) *PageInfo {
	info := &PageInfo{
		Contents: contents,
		Package:  packageName,
		Method:   method,
		Template: method,
		Layout:   "default",
		AssignData: map[string]string{
			"Title": packageName + " - " + contents,
		},
	}
	return info
}
func (pi *PageInfo) Assign(key, value string) {
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
