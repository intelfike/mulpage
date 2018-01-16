package types

import "errors"

// コンテンツの定義
type Contents struct {
	Before   Method
	After    Method
	Contents map[string]Content
}

func (cs *Contents) Init() {
	cs.Contents = map[string]Content{}
}

func (cs Contents) SetContent(key string, c Content) {
	cs.Contents[key] = c
}

func (cs Contents) ExecBefore(info *PageInfo) *Redirect {
	if cs.Before == nil {
		return nil
	}
	return cs.Before(info)
}

func (cs Contents) ExecAfter(info *PageInfo) *Redirect {
	if cs.After == nil {
		return nil
	}
	return cs.After(info)
}

// 指定された関数を実行する
func (cs Contents) Exec(info *PageInfo) (*Redirect, error) {
	if red := cs.ExecBefore(info); red != nil {
		return red, nil
	}
	con, ok := cs.Contents[info.Contents]
	if !ok {
		return nil, errors.New(info.Contents + ":コンテンツは定義されていません")
	}
	if red, err := con.Exec(info); red != nil || err != nil {
		return red, err
	}
	if red := cs.ExecBefore(info); red != nil {
		return red, nil
	}
	return nil, nil
}
