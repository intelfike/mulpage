// メソッドによって定義されるレスポンス情報
package types

import "path/filepath"

type ResponseData struct {
	// デフォルトでMethod名と同じ、変更可能
	Template string // dev/contents/*/page/X/template/*.tpl
	Layout   string

	// テンプレートファイルを追加するため
	TemplateFiles []string
	// 主にテンプレートに渡すデータたち
	AssignData map[string]interface{}
}

func (rd *ResponseData) Init() {
	*rd = ResponseData{
		Template:      "",
		Layout:        "default",
		TemplateFiles: []string{},
		AssignData:    map[string]interface{}{},
	}
}

func (rd *ResponseData) SetPageInfo(info *PageInfo) {
	rd.Template = info.Method
	rd.Assign("Title", info.Package+" | "+info.Contents)
}

// データを追加
func (rd *ResponseData) AddTpl(fileNames ...string) {
	rd.TemplateFiles = append(rd.TemplateFiles, fileNames...)
}
func (rd *ResponseData) Assign(key string, value interface{}) {
	rd.AssignData[key] = value
}

// 相対パスを取得する
func (rd *ResponseData) PageTemplatePath(info *PageInfo) string {
	path := filepath.Join(info.PackagePath(), "template", rd.Template+".tpl")
	return path
}

func (rd *ResponseData) LayoutPath(info *PageInfo) string {
	path := filepath.Join(info.ContentsPath(), "template", "layout", rd.Layout+".tpl")
	return path
}
