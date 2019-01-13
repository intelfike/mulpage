// 末端パッケージ
// ファイル関連の入出力を共通化するためのパッケージ
package file

import (
	"io/ioutil"
)

// ファイルを読み取る
func Read(fileName string) (string, err) {
	b, err := ioutil.ReadFile(fileName)
	return string(b), err
}
