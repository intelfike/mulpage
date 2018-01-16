package types

// コンテントの定義
import "errors"

type Content struct {
	Before   Method
	After    Method
	Packages map[string]*Package
}

func (con *Content) Init() {
	con.Packages = map[string]*Package{}
}

func (c Content) SetPackage(key string, p *Package) {
	c.Packages[key] = p
}

func (con Content) Exec(info *PageInfo) (*Redirect, error) {
	// 前置処理
	if red, _ := con.Before.Exec(info); red != nil {
		return red, nil
	}
	// 実行
	pack, ok := con.Packages[info.Package]
	if !ok {
		return nil, errors.New(info.Package + ":パッケージは定義されていません")
	}
	if red, err := pack.Exec(info); red != nil {
		return red, err
	}
	// 後置処理
	red, _ := con.After.Exec(info)
	return red, nil
}
