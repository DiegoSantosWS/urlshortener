package controller

import (
	"html/template"
)

var (
	//ModelosHome teste teste teste
	ModelosHome = template.Must(template.ParseFiles("views/home.html"))
	//ModelosRedirection teste teste teste
	ModelosRedirection = template.Must(template.ParseFiles("views/r.html"))
	//ModelosAnalytics teste teste teste
	ModelosAnalytics = template.Must(template.ParseFiles("views/analytics-wd.html"))
)
