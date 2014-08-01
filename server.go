package main

import "github.com/go-martini/martini"

func main() {
    routes()
}

func routes() {

    m := martini.Classic()

    m.Use(martini.Static("assets"))

    m.Get("/", func() string {
        return "Hello world!"
    })

    m.Get("/ids/:name", func(params martini.Params) string {
        return "Hello " + params["name"]
    })

    m.Group("/users", func(router martini.Router) {
        router.Get("/:id", GetUsers)
        router.Post("/new", NewUser)
        router.Put("/update/:id", UpdateUser)
        router.Delete("/delete/:id", DeleteUser)
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
