package resolvers

import (
	"finnbank/services/graphql/graphql_config/entities"
	"finnbank/services/graphql/types"
	// "math/rand"
	// "time"
	"github.com/graphql-go/graphql"
)

type Product = types.Product

var productType = entities.GetProductType()

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single product by id
			   http://localhost:8080/product?query={product(id:1){name,info,price}}
			*/
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// id, ok := p.Args["id"].(int)
					// if ok {
					// 	// Find product
						
					// 	// Call gRPC product service

					// 	// for _, product := range products {
					// 	// 	if int(product.ID) == id {
					// 	// 		return product, nil
					// 	// 	}
					// 	// }
					// }
					// return nil, nil
				},
			},
			/* Get (read) product list
			   http://localhost:8080/product?query={list{id,name,info,price}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					
					// Call gRPC product service

					// return products, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new product item
		http://localhost:8080/product?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
		*/
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// rand.Seed(time.Now().UnixNano())
				// product := Product{
				// 	ID:    int64(rand.Intn(100000)), // generate random ID
				// 	Name:  params.Args["name"].(string),
				// 	Info:  params.Args["info"].(string),
				// 	Price: params.Args["price"].(float64),
				// }
				// products = append(products, product)
				// return product, nil

				// Call gRPC product service

			},
		},

		/* Update product by id
		   http://localhost:8080/product?query=mutation+_{update(id:1,price:3.95){id,name,info,price}}
		*/
		"update": &graphql.Field{
			Type:        productType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// id, _ := params.Args["id"].(int)
				// name, nameOk := params.Args["name"].(string)
				// info, infoOk := params.Args["info"].(string)
				// price, priceOk := params.Args["price"].(float64)
				// product := Product{}
				
				
				// Call gRPC product service

				// for i, p := range products {
				// 	if int64(id) == p.ID {
				// 		if nameOk {
				// 			products[i].Name = name
				// 		}
				// 		if infoOk {
				// 			products[i].Info = info
				// 		}
				// 		if priceOk {
				// 			products[i].Price = price
				// 		}
				// 		product = products[i]
				// 		break
				// 	}
				// }
				// return product, nil
			},
		},

		/* Delete product by id
		   http://localhost:8080/product?query=mutation+_{delete(id:1){id,name,info,price}}
		*/
		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// id, _ := params.Args["id"].(int)
				// product := Product{}
				// for i, p := range products {
				// 	if int64(id) == p.ID {
				// 		product = products[i]
				// 		// Remove from product list
				// 		products = append(products[:i], products[i+1:]...)
				// 	}
				// }

				return product, nil
			},
		},
	},
})
