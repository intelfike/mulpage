// 末端パッケージ
// ファイル関連の入出力を共通化するためのパッケージ
package file

import (
	"encoding/json"
	"io/ioutil"
)

// ファイルを読み取る
func Read(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	return string(b), err
}

func ReadJson(filename string) (interface{}, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	i := new(interface{})
	err = json.Unmarshal(b, i)
	return *i, err
}

func Write(filename, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0777)
}
func WriteByte(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0777)
}

func WriteJson(filename string, i interface{}, indent string) error {
	data, err := json.MarshalIndent(i, "", indent)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0777)
}
