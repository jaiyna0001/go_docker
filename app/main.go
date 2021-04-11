// このソースファイルは「main」パッケージに属すことを指定しています。
// また、Go言語のソースファイルは必ず「package」文で始まります。
package main

// 一つの import ステートメントで複数のパッケージをインポートしています。
import (

	// フォーマットI/Oを扱うパッケージです。
	// C言語のprintfおよびscanfと似た関数を持ちます。

	"encoding/json"

	// ロギングを行うシンプルなパッケージです。
	// 出力をフォーマットするメソッドを持つLogger型が定義されています。
	"log"

	// HTTPを扱うパッケージです。
	// HTTPクライアントとHTTPサーバーを実装するために必要な機能が提供されています
	"net/http"
)

type param struct {
	Key   string
	Value string
}

// 「func」が関数の宣言で、ここから「main」関数の記述が始まります。
// ※関数宣言のような波括弧「{}」は、開始波括弧「{」をfuncと同じ行に書く必要があります。
func main() {
	// 第一引数にURL、第二引数にハンドラを渡し、DefaultServeMuxに登録する。
	// つまり、「/」というURLがリクエストされた際に、handlerが起動する。
	//   ・URL            ：HTTPリクエストのパターン
	//   ・ハンドラ        ：リクエストに対する応答
	//   ・DefaultServeMux：パターンにマッチしたリクエストに対して、そのパターンを持つhandlerを返却
	http.HandleFunc("/sample/", handle)
	http.HandleFunc("/sample/get/", get_handle)
	http.HandleFunc("/sample/post/", post_handle)

	// ポート番号3000で待ち受けるHTTPサーバを起動します。
	log.Fatal("[Fatal] [main] ", http.ListenAndServe(":3000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		v := r.URL.Query()
		if v == nil {
			return
		}

		var para []param
		for key, vs := range v {
			para = append(para, param{Key: key, Value: vs[0]})
		}
		json.NewEncoder(w).Encode(para)
		// res, _ := json.Marshal(para)
		// fmt.Fprintf(w, string(res))
	} else if r.Method == "POST" {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, "Request Error", http.StatusBadRequest)
			return
		}

		mf := r.MultipartForm

		var para []param
		for key, vs := range mf.Value {
			para = append(para, param{Key: key, Value: vs[0]})
		}
		json.NewEncoder(w).Encode(para)
		// res, _ := json.Marshal(para)
		// fmt.Fprintf(w, string(res))
	} else {
		http.Error(w, "Request Error", http.StatusBadRequest)
		return
	}
}

func get_handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Request Error", http.StatusBadRequest)
		return
	}
	v := r.URL.Query()
	if v == nil {
		return
	}
	var para []param
	for key, vs := range v {
		para = append(para, param{Key: key, Value: vs[0]})
	}
	json.NewEncoder(w).Encode(para)
	// res, _ := json.Marshal(para)
	// fmt.Fprintf(w, string(res))
}

func post_handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Request Error", http.StatusBadRequest)

		return
	}
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Request Error", http.StatusBadRequest)
		return
	}

	mf := r.MultipartForm

	var para []param
	for key, vs := range mf.Value {
		para = append(para, param{Key: key, Value: vs[0]})
	}
	json.NewEncoder(w).Encode(para)
	// res, _ := json.Marshal(para)
	// fmt.Fprintf(w, string(res))
}
