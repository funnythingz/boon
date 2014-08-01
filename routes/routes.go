package routes

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

type Header struct {
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

    m.Get("/", func(r render.Render) {
        header := Header{"Index", "Description"}
        r.HTML(200, "index", header)
    })

    m.Get("/ids/:name", func(params martini.Params) string {
        return "Hello " + params["name"]
    })

    m.Group("/users", func(r martini.Router) {
        r.Get("/:id", GetUsers)
        r.Post("/new", NewUser)
        r.Put("/update/:id", UpdateUser)
        r.Delete("/delete/:id", DeleteUser)
    })

    m.Run()

}

func GetUsers() string {
    return "get Users"
}

func NewUser() {
}

func UpdateUser() {
}

func DeleteUser() {
}
