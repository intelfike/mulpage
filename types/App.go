package types

import "errors"

// コンテンツの定義
type App struct {
	Name     string
	Before   Method
	After    Method
	Contents map[string]*Content
}

func (cs *App) Init(name string, ContentList map[string]ContentIfc) {
	cs.Name = name
	cs.Contents = map[string]*Content{}
	for key, val := range ContentList {
		content := &Content{}
		cs.SetContent(key, content)
		val.Define(content)
	}
}

func (cs App) SetContent(key string, c *Content) {
	cs.Contents[key] = c
}

// 指定された関数を実行する
func (cs App) Exec(tpl *TplData, info PageInfo) (*Redirect, error) {
	if red, _ := cs.Before.Exec(tpl, info); red != nil {
		return red, nil
	}
	con, ok := cs.Contents[info.Contents]
	if !ok {
		return nil, errors.New(info.Contents + ":コンテンツは定義されていません")
	}
	if red, err := con.Exec(tpl, info); red != nil || err != nil {
		return red, err
	}
	red, _ := cs.After.Exec(tpl, info)
	return red, nil
}
