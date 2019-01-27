package controllers

import (
    // "math/rand"
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
                    Type: graphql.Int,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                id, ok := p.Args["id"].(string)
                if ok {
                    // Find product
                    for _, article := range models.Articles {
                        if string(article.ID) == id {
                            return article, nil
                        }
                    }
                }
                return nil, nil
            },
        },
        // "article_list": &graphql.Field{
        //     Type:        models.ArticleListType,
        //     Description: "Get article list",
        //     Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        //         id, ok := p.Args["id"].(int)
        //         if ok {
        //             // Find product
        //             for _, article := range models.Articles {
        //                 if int(article.ID) == id {
        //                     return article, nil
        //                 }
        //             }
        //         }
        //         return nil, nil
        //     },
        // }
    },
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Mutation",
    Fields: graphql.Fields{
        /* Create new product item
        http://localhost:8080/product?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
        */
        "article": &graphql.Field{
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


