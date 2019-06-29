package main

import (
	"net/http"
	"html/template"
)

// page handlers
func errorPage(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("home.html"))
	tmpl.Execute(res, nil)
}
func homePage(res http.ResponseWriter, req *http.Request) {
	// database stuff
	data := struct{head string; sidebar string; footer string;}{"", "", ""}
	tmpl, err := template.ParseFiles("home.html")
	tmpl.Execute(res, data)
	if err != nil {
        http.Redirect(res, req, "/error", 500)
    }
}
func postPage(res http.ResponseWriter, req *http.Request) {
	// database stuff
	data := struct{head string; sidebar string; footer string;}{"", "", ""}
	tmpl, err := template.ParseFiles("post.html")
	tmpl.Execute(res, data)
	if err != nil {
        http.Redirect(res, req, "/error", 500)
    }
}

// router
func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/error", errorPage)
    http.HandleFunc("/post/", postPage)
    http.ListenAndServe(":80", nil) // set listen port
}