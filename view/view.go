package view

import (
    "github.com/martini-contrib/render"
    "../util"
)

func AboutRender(r render.Render) {
    r.HTML(200, "about", nil)
}

func BoonList() string {
    return "BoonList"
}

func ShowBoon() string {
    return "ShowBoon"
}

func NewBoon(r render.Render) {
    r.HTML(200, "boon/new", map[string]interface{}{
        "Options": util.CreateTimeOptions(),
    })
}

func PostNewBoon() {
}

func EditBoon() string {
    return "EditBoon"
}

func DeleteBoon() string {
    return "DeleteBoon"
}
