package routes

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "../view"
)

func Router() {

    m := martini.Classic()

    m.Use(martini.Static("assets"))
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

    m.Get("/about", view.AboutRender)

    m.Group("/boon", func(r martini.Router) {
        r.Get("/admin", view.BoonList)
        r.Get("/entry/:id", view.ShowBoon)
        r.Get("/new", view.NewBoon)
        r.Post("/new", view.PostNewBoon)
        r.Put("/edit/:id", view.EditBoon)
        r.Delete("/delete/:id", view.DeleteBoon)
    })

    m.Run()

}
