// 目視・手動での管理のしやすさを重視して、高い一覧性を維持しつつ素早く目的のデータにたどり着くためのJSON形式でのDBライブラリ
package jsondb

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	fileio "github.com/intelfike/mulpage/io/file"
)

type JsonDB struct {
	path   string
	indent string
}

func NewJsonDB(path string) *JsonDB {
	j := new(JsonDB)
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	j.path = path
	return j
}

func (j *JsonDB) SetIndent(indent string) {
	j.indent = indent
}

func (j *JsonDB) GetList(direname string) []string {
	dirs, _ := ioutil.ReadDir(filepath.Join(j.path, direname)) // NewJsonDB() で保証されているためエラーは握り潰す
	ss := make([]string, 0, len(dirs))
	for _, file := range dirs {
		ss = append(ss, file.Name())
	}
	return ss
}

func (j *JsonDB) Get(file string) (interface{}, error) {
	i, err := fileio.ReadJson(filepath.Join(j.path, file))
	return i, err
}
func (j *JsonDB) Set(file string, data map[string]interface{}) error {
	path := filepath.Join(j.path, file)
	if _, err := os.Stat(filepath.Dir(path)); err != nil {
		os.MkdirAll(path, 0777)
	}
	return fileio.WriteJson(path, data, j.indent)
}
