package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/CloudyKit/jet"
)

var views = jet.NewHTMLSet("./views")

func main() {
	// remove in production
	views.SetDevelopmentMode(true)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view, err := views.GetTemplate("index.jet")
		if err != nil {
			log.Println("Unexpected template err:", err.Error())
		}
		view.Execute(w, nil, nil)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	} else if !strings.HasPrefix(":", port) {
		port = ":" + port
	}

	log.Println("Serving on " + port)
	http.ListenAndServe(port, nil)
}
