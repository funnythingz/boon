package routes

import (
    "github.com/go-martini/martini"
    //"github.com/martini-contrib/render"
)

func Router() {

    m := martini.Classic()

    m.Use(martini.Static("assets"))

    m.Get("/", func() string {
        return "Hello world!"
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
