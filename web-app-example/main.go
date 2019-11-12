// go run main.go
// localhost:8080
// localhost:8080/?name=your-name-here

package main

import (
	"fmt"           // to format the output text
	"html/template" // interact with the HTML file
	"net/http"      // to access the core Go HTTP functionality
	"time"          // to work with date and time
)

// Welcome is a struct that we'll use to hold info to display in the HTML file
type Welcome struct {
	Name string
	Time string
}

// App entrypoint
func main() {
	// Instantiate Welcome struct object with default info
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	// Tell Go the relative path to find the HTML file to parse
	// template.Must() handles errors and halts on fatal errors
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	// To get the CSS that the HTML needs, create a handle that looks in the static directory.
	// Go then uses the "/static/" as a URL that the HTML can look into to get CSS and other files.
	http.Handle("/static/", // final URL can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	// Go looks in the relative static directory first, then matches it to a
	// URL of our choice as shown in http.Handle("/static/").
	// In HTML: <link rel="stylesheet" href="/static/stylesheets/...">
	// Note: the final URL can be anything, so long as we're consistent.

	// URL path "/" and a function (response writer, HTTP request):
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// if have name from http request r -> /?name=your-name-here -> name
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name // set struct object to hold name from URL
		}
		// if have error
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			// show error message
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start web server, set the port to listen to 8080 (assumes localhost if unspecified).
	// Print any errors from starting the webserver using fmt
	fmt.Println(http.ListenAndServe(":8080", nil))
}
