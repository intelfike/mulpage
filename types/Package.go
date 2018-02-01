package types

import (
	"errors"
	"fmt"

	"os"
)

// パッケージの定義
type Package struct {
	Index    int
	Key      string
	depth    int
	Kind     string
	Name     string
	Before   Method
	After    Method
	Children map[string]*Package
	Articles map[string]*Article
}

type PackageIfc interface {
	Define(*Package)
}

func (pack *Package) Init(kind, name string) {
	pack.Kind = kind
	pack.Name = name
	pack.Children = map[string]*Package{}
	pack.Articles = map[string]*Article{}
}

func (pack *Package) FallDown(pi PackageIfc) {
	pi.Define(pack)
}

func (pack *Package) NewChild(key string) *Package {
	child := &Package{}
	pack.SetChild(key, child)
	return child
}
func (pack *Package) SetChild(key string, child *Package) {
	if pack.Children == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init()")
		os.Exit(1)
	}
	if _, ok := pack.Children[key]; ok {
		fmt.Println(key, ":メソッドは既に定義されています。")
	}

	child.Index = len(pack.Children)
	child.Key = key
	child.depth = pack.depth + 1
	pack.Children[key] = child
}
func (pack *Package) SetMethod(key, name string, m Method) {
	if pack.Articles == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init()")
		os.Exit(1)
	}
	if _, ok := pack.Articles[key]; ok {
		fmt.Println(key, ":メソッドは既に定義されています。")
	}
	atc := &Article{len(pack.Articles), key, name, m}
	pack.Articles[key] = atc
}

func (pack *Package) Exec(tpl *TplData, info PageInfo) (*Redirect, error) {
	if len(info.ExecPath) <= pack.depth {
		return nil, errors.New("パスが長すぎます")
	}
	// 前置処理
	if red, _ := pack.Before.Exec(tpl, info); red != nil {
		return red, nil
	}
	key := info.ExecPath[pack.depth]
	// 実行
	atc, ok := pack.Articles[key]
	if !ok {
		// 伝播
		child, ok := pack.Children[key]
		if !ok {
			return nil, errors.New("エラー:" + pack.Kind + "において、" + key + "は定義されていません")
		}
		if red, err := child.Exec(tpl, info); red != nil {
			return red, err
		}
	} else {
		if red, err := atc.Exec(tpl, info); red != nil {
			return red, err
		}
	}
	// 後置処理
	red, _ := pack.After.Exec(tpl, info)
	return red, nil
}

func (pack *Package) ChildrenArray() []*Package {
	arr := make([]*Package, len(pack.Children))
	for key, val := range pack.Children {
		val.Key = key
		arr[val.Index] = val
	}
	return arr
}
func (pack *Package) ArticlesArray() []*Article {
	arr := make([]*Article, len(pack.Articles))
	for key, val := range pack.Articles {
		val.Key = key
		arr[val.Index] = val
	}
	return arr
}
