package main

import (
	"net/http"

	"github.com/Isabelle-Dev/isabelle-graphql/schema"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema.Root,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", h)
	http.ListenAndServe(":4000", nil)
}
