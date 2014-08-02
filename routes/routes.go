package routes

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

type HeaderModel struct {
    Title string
    Description string
}

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

    m.Get("/about", func(r render.Render) {
        header := HeaderModel{"About", "ブーンについて"}
        r.HTML(200, "about", header)
    })

    m.Group("/boon", func(r martini.Router) {
        r.Get("/admin", BoonList)
        r.Get("/:id", ShowBoon)
        r.Post("/new", NewBoon)
        r.Put("/edit/:id", EditBoon)
        r.Delete("/delete/:id", DeleteBoon)
    })

    m.Run()

}

func BoonList() string {
    return "BoonList"
}

func ShowBoon() string {
    return "ShowBoon"
}

func NewBoon() string {
    return "NewBoon"
}

func EditBoon() string {
    return "EditBoon"
}

func DeleteBoon() string {
    return "DeleteBoon"
}
