package api

import (
	"github.com/graphql-go/handler"
	"github.com/shaderboi/koffie-backend/api/controllers"
	"net/http"
	"os"
)

func Routes() {

	schema, _ := controllers.GetProduct()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
