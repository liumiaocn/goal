package main

import "net/http"
import "controller"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/charts.html", controller.Charts)
	http.HandleFunc("/blank-page.html", controller.Blank)
	http.HandleFunc("/bootstrap-elements.html", controller.Elements)
	http.HandleFunc("/bootstrap-grid.html", controller.Grid)
	http.HandleFunc("/forms.html", controller.Forms)
	http.HandleFunc("/tables.html", controller.Tables)
	http.HandleFunc("/typography.html", controller.Graphy)
	http.ListenAndServe(":8080", nil)
}
