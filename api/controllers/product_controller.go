package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/shaderboi/koffie-backend/api/db"
	"github.com/shaderboi/koffie-backend/api/products"
	"time"
)

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

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        productType,
				Description: "Create a new product",
				Args: graphql.FieldConfigArgument{
					"item_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"discount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					now := time.Now()

					disc := p.Args["discount"].(int)

					product := products.Product{
						Id:        p.Args["item_id"].(string),
						Name:      p.Args["name"].(string),
						Desc:      p.Args["description"].(string),
						Price:     p.Args["price"].(int),
						Image:     p.Args["image"].(string),
						Discount:  &disc,
						CreatedAt: &now,
					}

					create := conn.Create(&product)

					return product, create.Error
				},
			},
		},
	})

	rootQuery := graphql.ObjectConfig{Name: "Root", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery),
		Mutation: mutationType}
	schema, err := graphql.NewSchema(schemaConfig)

	return schema, err
}
