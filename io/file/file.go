// 末端パッケージ
package file

import (
	"io/ioutil"
)

func Read(fileName string) (string, err) {
	b, err := ioutil.ReadFile(fileName)
	return string(b), err
}
