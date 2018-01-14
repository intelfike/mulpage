// UNIX哲学に出来る限り従う方向で！
// isearとホームページを作りたい
// lazyの構成を参考に
package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	httpio "github.com/intelfike/mulpage/dev/io/http"
	htmlproc "github.com/intelfike/mulpage/dev/proc/policy/html"
	pathmod "github.com/intelfike/mulpage/dev/proc/policy/path"
)

var port = flag.String("http", ":80", "HTTP port number.")

func init() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// パスを取得
		path := httpio.ReadPath(r)
		path = pathmod.ParseURIPath(path)
		if len(path) == 0 || !strings.HasPrefix(path[0], "_") {
			err := httpio.WriteFile(w, "../public"+r.URL.Path)
			if err == nil {
				return
			}
			return
		}
		if len(path) == 1 {
			http.Redirect(w, r, path[0], 307)
			return
		}
		// HTMLを生成
		contents := httpio.ReadContents(r)
		pageName := pathmod.AvailePageName(path[0])
		redirect, err := htmlproc.Write(w, contents, pageName, path[1])
		if err != nil {
			result := "エラー:" + fmt.Sprint(err) + "\n"
			result += "エラーメッセージを記録して管理者に報告してください。"
			httpio.Write(w, result)
			return
		}
		if redirect != nil {
			redirect.Exec(w, r)
		}
	})
}

func main() {
	fmt.Printf("Start HTTP Server localhost%s\n", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
