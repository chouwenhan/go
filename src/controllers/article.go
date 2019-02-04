package controllers

import (
    // "math/rand"
    // "encoding/json"
    // "fmt"
    "src/models"
    "github.com/graphql-go/graphql"
    "github.com/gin-gonic/gin"
    "github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        /* Get (read) single article by id
           http://localhost:8080/product?query={product(id:1){name,info,price}}
        */
        "article": &graphql.Field{
            Type:        models.ArticleType,
            Description: "Get article by id",
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                id, ok := p.Args["id"].(string)
                if ok {
                    // Find product
                    db := models.ConnDB("article")
                    result, err := models.ReadDocument(db, id)
                    if err != nil {
                        panic(err)
                    }
                    return result, nil
                }
                return ok, nil
            },
        },
        "article_list": &graphql.Field{
            Type:    graphql.NewList(models.ArticleType),
            Description: "Get article list",
            Args: graphql.FieldConfigArgument{
                "selector": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                selector_json := map[string]string{
                                      "type": "article",
                                 }
                db := models.ConnDB("article")
                results := models.Articles{}
                err := models.Find(db, &results, selector_json)
                if err != nil {
                    panic(err)
                }
                return results.Docs, nil
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
        "createArticle": &graphql.Field{
            Type:        models.ArticleType,
            Description: "Create new article",
            Args: graphql.FieldConfigArgument{
                "name": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "describle": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
                "note": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                article := models.CreateArticle{
                    Type:  "article",
                    Name:  params.Args["name"].(string),
                    Describle:  params.Args["describle"].(string),
                    Note: params.Args["note"].(string),
                }
                db := models.ConnDB("article")
                id := models.CreateDocument(db, article)
                result, err := models.ReadDocument(db, id)
                if err != nil {
                    panic(err)
                }
                return result, nil
            },
        },
        "updateArticle": &graphql.Field{
            Type:        models.ArticleType,
            Description: "Update article",
            Args: graphql.FieldConfigArgument{
                "_id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "name": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
                "describle": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
                "note": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                id, _ := params.Args["_id"].(string)
                db := models.ConnDB("article")
                result, err := models.ReadDocument(db, id)
                if err != nil {
                    panic(err)
                }
                if params.Args["name"] != nil {
                    result.Name = params.Args["name"].(string)
                }
                if params.Args["describle"] != nil {
                    result.Describle = params.Args["describle"].(string)
                }
                if params.Args["note"] != nil {
                    result.Note = params.Args["note"].(string)
                }
                id = models.UpdateDocument(db, result, id, result.Rev)
                result, err = models.ReadDocument(db, id)
                if err != nil {
                    panic(err)
                }
                return result, nil
            },
        },
    },
})

func ArticleHandler() gin.HandlerFunc{
    h := handler.New(&handler.Config{
        Schema:   &schema,
        Pretty:   true,
        GraphiQL: true,
    })

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

var schema, _ = graphql.NewSchema(
    graphql.SchemaConfig{
        Query:    queryType,
        Mutation: mutationType,
    },
)


