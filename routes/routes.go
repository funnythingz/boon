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
        header := HeaderModel{"Boon", "日程を調整するおっ(^^"}
        r.HTML(200, "index", header)
    })

    m.Get("/about", AboutRender)

    m.Group("/boon", func(r martini.Router) {
        r.Get("/admin", BoonList)
        r.Get("/entry/:id", ShowBoon)
        r.Get("/new", NewBoon)
        r.Post("/new", PostNewBoon)
        r.Put("/edit/:id", EditBoon)
        r.Delete("/delete/:id", DeleteBoon)
    })

    m.Run()

}

func AboutRender(r render.Render) {
    header := HeaderModel{"About", "日程をえらべるお(^^"}
    r.HTML(200, "about", header)
}

func BoonList() string {
    return "BoonList"
}

func ShowBoon() string {
    return "ShowBoon"
}

func NewBoon(r render.Render) {
    header := HeaderModel{"日程をえらブーン", "さぁえらブーンだ！"}
    r.HTML(200, "boon/new", header)
}

func PostNewBoon() {
}

func EditBoon() string {
    return "EditBoon"
}

func DeleteBoon() string {
    return "DeleteBoon"
}
