package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)


func main() {
	m := martini.Classic()

	m.Use( render.Renderer(render.Options{
		IndentJSON: true, // so we can read it..
	}))
