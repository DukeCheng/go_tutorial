package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range messages {
		fmt.Println(msg)
	}
	fmt.Println(messages)

	i := 8
	iPtr := &i
	j := *iPtr
	fmt.Println(iPtr)
	fmt.Println(j)
	j = 10
	fmt.Println(j)
	fmt.Println(*iPtr)
	fmt.Println(&j)
	fmt.Println("hello, world", i)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
