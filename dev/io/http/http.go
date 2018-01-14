// "github.com/intelfike/mulpage/dev/io/http"
package httpio

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ReadBody(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return string(b), err
}

// 戻り値はパス
func ReadPath(r *http.Request) []string {
	path := r.URL.EscapedPath()
	// パスを取得
	path = strings.Trim(path, "/")
	if len(path) == 0 {
		return nil
	}
	return strings.Split(path, "/")
}

// コンテンツという名の、アプリケーションの動作を決める大まかなグループを取得
// ドメインの一番左側を取得する
func ReadContents(r *http.Request) string {
	host := strings.Split(r.Host, ".")
	host[0] = "isear"
	return host[0]
}

func Write(w http.ResponseWriter, s string) error {
	_, err := w.Write([]byte(s))
	return err
}
func WriteFile(w http.ResponseWriter, fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(w, f)
	return nil
}
