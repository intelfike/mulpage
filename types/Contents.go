package types

import "errors"

// コンテンツの定義
type Contents struct {
	Title    string
	Before   Method
	After    Method
	Contents map[string]*Content
}

func (cs *Contents) Init() {
	cs.Contents = map[string]*Content{}
}

func (cs Contents) SetContent(key string, c *Content) {
	cs.Contents[key] = c
}

// 指定された関数を実行する
func (cs Contents) Exec(info *PageInfo) (*Redirect, error) {
	if red, _ := cs.Before.Exec(info); red != nil {
		return red, nil
	}
	con, ok := cs.Contents[info.Contents]
	if !ok {
		return nil, errors.New(info.Contents + ":コンテンツは定義されていません")
	}
	if red, err := con.Exec(info); red != nil || err != nil {
		return red, err
	}
	red, _ := cs.After.Exec(info)
	return red, nil
}
