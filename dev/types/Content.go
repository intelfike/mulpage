package types

// コンテントの定義
import "errors"

type Content struct {
	Before   Method
	After    Method
	Packages map[string]Package
}

func (con *Content) Init() {
	con.Packages = map[string]Package{}
}

func (con Content) ExecBefore(info *PageInfo) (*Redirect, error) {
	if con.Before == nil {
		return nil, nil
	}
	return con.Before.Exec(info)
}

func (con Content) ExecAfter(info *PageInfo) (*Redirect, error) {
	if con.After == nil {
		return nil, nil
	}
	return con.After.Exec(info)
}

func (con Content) Exec(info *PageInfo) (*Redirect, error) {
	// 前置処理
	if red, err := con.ExecBefore(info); red != nil || err != nil {
		return red, err
	}
	// 実行
	pack, ok := con.Packages[info.Package]
	if !ok {
		return nil, errors.New(info.Package + ":パッケージは定義されていません")
	}
	if red, err := pack.Exec(info); red != nil || err != nil {
		return red, err
	}
	// 後置処理
	red, err := con.ExecBefore(info)
	return red, err
}

func (c Content) SetPackage(key string, p Package) {
	c.Packages[key] = p
}
