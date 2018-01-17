package types

// コンテントの定義
import "errors"

type Content struct {
	Name     string
	Before   Method
	After    Method
	Packages map[string]*Package
}

type ContentIfc interface {
	Define(*Content)
}

func (con *Content) Init(name string, PackageList map[string]PackageIfc) {
	con.Name = name
	con.Packages = map[string]*Package{}
	for key, val := range PackageList {
		p := &Package{}
		con.SetPackage(key, p)
		val.Define(p)
	}
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
