// メソッドを階層的に格納するための構造体。
package types

import (
	"errors"
	"fmt"

	"os"
)

// パッケージの定義
type Package struct {
	Depth     int
	Index     int
	Key       string
	Kind      string
	ChildKind string
	Name      string
	Before    Method
	After     Method
	Children  map[string]*Package
	Articles  map[string]*Article
}

// 構造体を初期化する
func (pack *Package) Init(name string) {
	pack.Name = name
	pack.Children = map[string]*Package{}
	pack.Articles = map[string]*Article{}
}

// 子パッケージの定義を実行する
func (pack *Package) FallDown(pi interface{Define(*Package)}) {
	pi.Define(pack)
}

// 新しい子パッケージを作成し、追加する
// SetChildのラッパ
func (pack *Package) NewChild(key string) *Package {
	child := &Package{}
	pack.SetChild(key, child)
	return child
}
// 子パッケージを追加する
func (pack *Package) SetChild(key string, child *Package) {
	if pack.Children == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init(\"web\")")
		os.Exit(1)
	}
	if _, ok := pack.Children[key]; ok {
		fmt.Println(key, ":重複したキーによる子パッケージは登録できません。\n違う名前で登録してください。")
	}

	child.Index = len(pack.Children)
	child.Key = key
	child.Depth = pack.Depth + 1
	if pack.ChildKind != "" {
		child.Kind = pack.ChildKind
	}
	pack.Children[key] = child
}
// 実行可能なメソッドを追加する
func (pack *Package) SetMethod(key, name string, m Method) {
	if pack.Articles == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init(\"web\")")
		os.Exit(1)
	}
	if _, ok := pack.Articles[key]; ok {
		fmt.Println(key, ":重複したキーによるメソッドは登録できません。\n違う名前で登録してください。")
	}
	atc := &Article{len(pack.Articles), key, name, m}
	pack.Articles[key] = atc
}

// リクエストに対してプログラムを実行する
// Beforeでリダイレクト→即座にリダイレクトするかどうか
// 階層が2の場合、Before1 => Before2 => Exec2 => After2 => After1 という順番で実行する
func (pack *Package) Exec(tpl *TplData, info PageInfo) (*Redirect, error) {
	if len(info.ExecPath) <= pack.Depth {
		return nil, errors.New("URLのパスが長すぎます")
	}
	// 前置処理
	if red, _ := pack.Before.Exec(tpl, info); red != nil {
		return red, nil
	}
	key := info.ExecPath[pack.Depth]
	// 実行
	atc, ok := pack.Articles[key]
	if !ok {
		// 子パッケージへの伝播
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

// キーを追加した順番でソートし、配列で返す
func (pack *Package) ChildrenArray() []*Package {
	arr := make([]*Package, len(pack.Children))
	for key, val := range pack.Children {
		val.Key = key
		arr[val.Index] = val
	}
	return arr
}
// キーを追加した順番でソートし、配列で返す
func (pack *Package) ArticlesArray() []*Article {
	arr := make([]*Article, len(pack.Articles))
	for key, val := range pack.Articles {
		val.Key = key
		arr[val.Index] = val
	}
	return arr
}
