package main

import (
	// import martini and martini renderer
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic() // enable martini
	
	// static options 
	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))
	
	// use render settings
	m.Use(render.Renderer(render.Options{
		Directory:  "assets/html/", // path to templates
		Layout:     "layout", 
		Delims:     render.Delims{"{{", "}}"}, 
		Extensions: []string{".html"},
		Charset:    "UTF-8",                            
		IndentJSON: true,     // use js
		IndentXML: 	true,	  // use xml
	}))

	m.Get("/", func (r render.Render){
		// index func
		r.HTML(200, "index", "") // render index page
	})
	
	// you can also process other pages 
	/* m.Get("/:page", func (r render.Render, params martini.Params){
		r.HTML(200, params["page"], "")
	}) */
	
	m.NotFound(func() string {
		// 404 error func
		return "404: Ooops"
	})
	
	m.Run()
}