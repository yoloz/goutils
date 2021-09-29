//在net/http包中动态文件的路由和静态文件的路由是分开的，
//动态文件使用http.HandleFunc进行设置，静态文件就需要使用到http.FileServer
package control

import (
	"net/http"
)

func StartHttpServer(srv http.Server) error {

	http.Handle("/", http.FileServer(http.Dir("static")))
	// web for develop
	// http.Handle("/", http.FileServer(http.Dir("web")))

	http.Handle("/add", new(AddControl))
	http.Handle("/delete", new(DeleteControl))
	http.Handle("/update", new(UpdateControl))
	http.Handle("/query", new(QueryControl))

	// for test
	// return srv.ListenAndServe()
	return srv.ListenAndServeTLS("cert.pem", "key.pem")
}
