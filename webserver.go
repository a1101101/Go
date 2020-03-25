//Handlerの処理とそれを呼び出す文字列を紐づけて、Listen状態にする
package main

import (
//	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/<path>")))
	http.ListenAndServe(":8080", nil)
}

//以下コメント...

//Handle関数:
//HandleはHandler型のオブジェクトと、それを呼び出す文字列を紐づける

//Handler型:
//インターフェース型で、ServeHTTPをメソッドとして持つ
//ServeHTTPの実装はレシーバーの違いにより2種類ある

//ServeHTTP:
//レシーバーとしてHandlerFuncをとる場合と、ServeMuxをとる場合とがある
//ServeHTTPをメソッドとしてとる型と、Serve

//HandlerFunc:
//func handlerfunc(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hey %s", r.URL.Path[1:])
//
//}
//関数fの引数が適切なら、型変換HandlerFunc(f)はHandler型

//ListenAndServe:
//ListenAndServe はServeを呼び出す
//ListenAndServeの第2引数はHandlerで、nilの場合はDefaultServeMuxが呼び出される

//レシーバーの機能:
//ある型データをレシーバーとして受け取る関数は、その型データのメソッドのように記述できる
//たとえばservemuxをレシーバとして受け取るfunc (mux *ServeMux) Handleにおいては、mux.HandleというふうにHandleを呼び出せる

//引数に取った関数の定義:
//引数に関数を受け取る場合、その関数の処理も続けて書くことができる
//たとえば、func(mux *ServeMux)HandleFunc(pattern String, handler func(ResponseWriter, *Request))は、
//mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
	//処理内容
//	})
//というようにfuncの処理を引数宣言部に書ける

//リクエストを取得するにはどうしたらよいか?:
//URLはtype Request に含まれる
//type HandlerFunc func(w, r)からr.URLを呼び出せる
//HandlerFuncはfunc(w, r)という型の関数
//HandlerFunc fの処理を定義して、ServeHttpのレシーバに渡すとfが呼び出される
