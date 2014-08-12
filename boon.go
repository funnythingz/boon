package main

import (
	"github.com/codegangsta/negroni"
	render "github.com/unrolled/render"
	_ "log"
	"net/http"
)

func main() {

	r := render.New(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl"},
		Charset:    "utf-8",
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "about", nil)
	})

	mux.HandleFunc("/new", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "new", map[string]interface{}{
			"Options": CreateTimeOptions(),
		})
	})

	n := negroni.Classic()
	n.UseHandler(mux)
    //n.UseHandler(http.NotFoundHandler())

	n.Run(":3000")

}
