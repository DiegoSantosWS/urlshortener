package controller

import (
	"html/template"
)

var (
	ModelosHome        = template.Must(template.ParseFiles("views/home.html"))
	ModelosRedirection = template.Must(template.ParseFiles("views/r.html"))
)
