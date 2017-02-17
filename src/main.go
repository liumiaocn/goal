package main

import "net/http"
import "controller"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/report.html", controller.Report)
	http.HandleFunc("/monitor.html", controller.Monitor)
	http.HandleFunc("/audit.html", controller.Audit)
	http.HandleFunc("/package.html", controller.Package)
	http.HandleFunc("/deploy.html", controller.Deploy)
	http.HandleFunc("/analyze.html", controller.Analyze)
	http.HandleFunc("/build.html", controller.Build)
	http.HandleFunc("/test.html", controller.Test)
	http.HandleFunc("/docker.html", controller.Docker)
	http.HandleFunc("/gitlab.html", controller.Gitlab)
	http.HandleFunc("/jenkins.html", controller.Jenkins)
	http.HandleFunc("/sonarqube.html", controller.Sonarqube)
	http.HandleFunc("/redmine.html", controller.Redmine)
	http.HandleFunc("/flow.html", controller.Flow)
	http.HandleFunc("/feedback.html", controller.Feedback)
	http.HandleFunc("/culture.html", controller.Culture)
	http.ListenAndServe(":8080", nil)
}
