package main

import (
	"log"
	"net/http"

	"github.com/russmiles/MyAwesomeGoService/HelloCourse"
)

func main() {
	HelloCourse.AddServices()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
