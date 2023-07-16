package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arcaptcha/arcaptcha-go"
)

func abc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	website := arcaptcha.NewWebsite("1pr8g5l057", "66sw65442gwh0wummxvk")
	challenge_id := r.FormValue("arcaptcha-token")
	result, err := website.Verify(challenge_id)
	if err != nil {
		// throw Error
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./Html-frontend.html")

	case "POST":
		if !result.Success {
			http.ServeFile(w, r, "./Html-frontend.html")
		}
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "parseForm() err: %v\n", err)
			return
		}

		fmt.Println(w, "post form website r.postform = %v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Fprintf(w, "username=%s\n", username)
		fmt.Fprintf(w, "password=%s\n", password)

	default:
		fmt.Fprint(w, "only get and post")
	}
}

func main() {
	http.HandleFunc("/", abc)

	fmt.Printf("starting server got testing\n")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		log.Fatal(err)

	}

}
