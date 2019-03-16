package mulpage

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"

	mhttpio "github.com/intelfike/mulpage/io/http"
	"github.com/intelfike/mulpage/types"
)

func All(contents types.Definer, asLocalhost string, port string) {
	Handle(contents, asLocalhost)
	Listen(port)
}

func Handle(contents types.Definer, asLocalhost string) {
	app := types.NewPackage()

	pack := contents.Define(app)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		info, err := types.NewPageInfo(r)
		if err != nil {
			// localhostの場合のアクセス先
			if info.Contents == "localhost" {
				info.Contents = asLocalhost
			}

			// メソッドがなければファイルを書き出す
			err := mhttpio.WriteFile(w, filepath.Join("contents", info.Contents, "public", r.URL.Path))
			if err != nil {
				err := mhttpio.WriteFile(w, "public"+r.URL.Path)
				if err != nil {
					mhttpio.Write(w, err.Error())
					http.NotFound(w, r)
					return
				}
			}
			return
		}

		// localhostの場合のアクセス先
		if info.Contents == "localhost" {
			info.Contents = asLocalhost
		}

		// テンプレートのデフォルト値を設定
		tpl := types.NewTplData()
		tpl.SetLayoutPath(filepath.Join("contents", info.Contents, "layout"))
		tpl.SetPagePath(filepath.Join("contents", info.Contents, "page", info.Group))
		tpl.SetPartsPath(filepath.Join("contents", info.Contents, "parts"))
		tpl.SetPage(info.Page + ".tpl")
		tpl.SetLayout("default.tpl")
		tpl.Assign("Rand", strconv.Itoa(rand.Int()))
		tpl.Assign("Info", info)

		// 関数実行
		redirect, err := pack.Exec(info, tpl)
		if tpl.GetDebugPrintText() != "" {
			mhttpio.Write(w, tpl.GetDebugPrintText())
			return
		}
		// mhttpio.Write(w, fmt.Sprintln(pack))
		if err != nil {
			result := "エラー:" + fmt.Sprint(err) + "\n"
			result += "エラーメッセージを記録して管理者に報告してください。"
			mhttpio.Write(w, result)
			return
		}

		if redirect == nil {
			// テンプレートを実行して送信
			err = tpl.Write(w)
		}

		// 必要ならリダイレクトする
		if redirect != nil {
			redirect.Exec(w, r)
		}
	})
}

func Listen(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
