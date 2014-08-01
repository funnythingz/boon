package main

import "github.com/go-martini/martini"

func main() {
    app()
}

func app() {

    m := martini.Classic()

    m.Get("/", func() string {
        return "Hello world!"
    })

    m.Get("/:name", func(params martini.Params) string {
        return "Hello " + params["name"]
    })

    m.Run()

}
