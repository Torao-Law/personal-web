package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//for route path public
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	//for routing url
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/blog", blog).Methods("GET")
	r.HandleFunc("/form-blog", formBlog).Methods("GET")
	r.HandleFunc("/add-blog", addBlog).Methods("POST")
	r.HandleFunc("/contact", contact).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset-utp8")
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html charset-utp8")
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("views/blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}

func formBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html charset-utp8")
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("views/add-blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html charset-utp8")
	w.WriteHeader(http.StatusOK)

	temp, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	temp.Execute(w, nil)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title : " + r.PostForm.Get("inputTitle")) // get value berdasarkan dari tag input name
	fmt.Println("Content : " + r.PostForm.Get("inputContent"))

	http.Redirect(w, r, "/blog", http.StatusMovedPermanently) // untuk mengarahkan kemana nanti yg akan dituju
}
