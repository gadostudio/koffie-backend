package controllers

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/products"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	conn, err := db.Connection()

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data products.Product

	if err := json.Unmarshal(reqBody, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now()

	data.CreatedAt = &now

	if err := conn.Create(&data); err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetProduct() (graphql.Schema, error) {

	conn, _ := db.Connection()

	var productType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"item_id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"discount": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

	fields := graphql.Fields{
		"product": &graphql.Field{
			Type:        productType,
			Description: "Get product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["item_id"].(string)

				if ok {
					var _products []products.Product

					conn.Find(&_products, products.Product{Id: id})

					return _products, nil
				}

				return nil, nil
			},
		},
		"products": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Get product list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var _products []products.Product

				conn.Find(&_products)

				return _products, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "Root", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	return schema, err
}
