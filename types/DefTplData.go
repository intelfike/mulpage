// メソッドによって定義されるテンプレート関連情報
package types

type DefTplData struct {
	// テンプレートファイルを追加するため
	TemplateFiles []string
	// 主にテンプレートに渡すデータたち
	AssignData map[string]interface{}
}

func (rd *DefTplData) Init() {
	*rd = DefTplData{
		TemplateFiles: []string{},
		AssignData:    map[string]interface{}{},
	}
}

// データを追加
func (rd *DefTplData) AddTpl(fileNames ...string) {
	rd.TemplateFiles = append(rd.TemplateFiles, fileNames...)
}
func (rd *DefTplData) Assign(key string, value interface{}) {
	rd.AssignData[key] = value
}
