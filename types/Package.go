// メソッドを階層的に格納するための構造体。
// 構造体内の変数を直接弄くった場合は、動作を保証できません。
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
	Kind      string // 表示用の種別
	ChildKind string
	Name      string // 表示用のラベル
	Before    Method
	After     Method
	Children  map[string]Executable
}

type Executable interface {
	Exec(PageInfo, *TplData) (*Redirect, error)
}

func NewPackage(name, childKind string) *Package {
	pack := new(Package)
	pack.Init("web")
	pack.Kind = "root"
	pack.ChildKind = childKind
	return pack
}

// 構造体を初期化する
func (pack *Package) Init(name string) {
	pack.Name = name
	pack.Children = map[string]Executable{}
}

type Definer interface {
	Define(*Package)
}

// 子パッケージの定義を実行する
func (pack *Package) FallDown(pis ...Definer) {
	for _, pi := range pis {
		pi.Define(pack)
	}
}

// 新しい子パッケージを作成し、追加する
// SetChildのラッパ
func (pack *Package) NewChild(key, name string) *Package {
	if _, ok := pack.Children[key]; ok {
		fmt.Println(key, ":重複したキーによる子パッケージは登録できません。\n違う名前で登録してください。")
		return nil
	}
	child := NewPackage(name, "")
	child.Index = len(pack.Children)
	child.Key = key
	child.Depth = pack.Depth + 1
	if pack.ChildKind != "" {
		child.Kind = pack.ChildKind
	}

	pack.SetChild(key, child)
	return child
}

// 子パッケージを追加する
func (pack *Package) SetChild(key string, child *Package) {
	pack.setExecutable(key, child)
}

// 実行可能なメソッドを追加する
func (pack *Package) SetMethod(key, name string, m Method) {
	atc := &Page{len(pack.Children), key, name, m}
	pack.setExecutable(key, atc)
}

// 子を追加する
func (pack *Package) setExecutable(key string, child Executable) {
	if pack.Children == nil {
		fmt.Println("types.Packageは必ずInit()してください\n例: pack.Init(\"web\")")
		os.Exit(1)
	}
	if _, ok := pack.Children[key]; ok {
		fmt.Println(key, ":重複したキーによるメソッドは登録できません。\n違う名前で登録してください。")
	}

	pack.Children[key] = child
}

// リクエストに対してプログラムを実行する
// Beforeでリダイレクト→即座にリダイレクトするかどうか
// 階層が2の場合、Before1 => Before2 => Exec2 => After2 => After1 という順番で実行する
func (pack *Package) Exec(info PageInfo, tpl *TplData) (*Redirect, error) {
	if len(info.ExecPath) <= pack.Depth {
		return nil, errors.New("URLのパスが長すぎます")
	}
	// 前置処理
	if red, _ := pack.Before.Exec(info, tpl); red != nil {
		return red, nil
	}
	key := info.ExecPath[pack.Depth]
	// 実行
	atc, ok := pack.Children[key]
	if !ok {
		// 子パッケージへの伝播
		child, ok := pack.Children[key]
		if !ok {
			return nil, errors.New("エラー:" + pack.Kind + "において、" + key + "は定義されていません")
		}
		if red, err := child.Exec(info, tpl); red != nil {
			return red, err
		}
	} else {
		if red, err := atc.Exec(info, tpl); red != nil {
			return red, err
		}
	}
	// 後置処理
	red, _ := pack.After.Exec(info, tpl)
	return red, nil
}

// キーを追加した順番でソートし、配列で返す
func (pack *Package) ChildrenArray() []Executable {
	arr := make([]Executable, len(pack.Children))
	for key, val := range pack.Children {
		switch v := val.(type) {
		case *Page:
			v.Key = key
			arr[v.Index] = v
		case *Package:
			v.Key = key
			arr[v.Index] = v
		}
	}
	return arr
}
