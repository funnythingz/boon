package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

type Options struct {
    Options []string
}

func main() {

    m := martini.Classic()

    m.Use(render.Renderer(render.Options{
        Directory: "templates",
        Layout: "layout",
        Extensions: []string{".tmpl"},
        Charset: "utf-8",
    }))

    m.NotFound(func (r render.Render){
        r.Redirect("/")
    })

    m.Get("/", func(r render.Render) {
        r.HTML(200, "index", nil)
    })

    m.Get("/about", func(r render.Render) {
        r.HTML(200, "about", nil)
    })

    m.Get("/new", func(r render.Render) {
        r.HTML(200, "new", &Options{CreateTimeOptions()})
    })

    m.Run()

}
