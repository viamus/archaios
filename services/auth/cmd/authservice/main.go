package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/viamus/archaios/libs/serviceutils/envy"
)

func main() {

	envy.LoadDotEnvIfDev()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world! 👋🌍")
	})

	port := envy.Get("PORT")
	log.Printf("🚀 Server running at http://localhost%s/hello", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
