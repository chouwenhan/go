package controllers

import (
    "math/rand"
    "time"
    "src/models"

    "github.com/graphql-go/graphql"
    "github.com/gin-gonic/gin"
    "github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(
    graphql.ObjectConfig{
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
                    id, ok := p.Args["id"].(int)
                    if ok {
                        // Find product
                        for _, article := range models.Articles {
                            if int(article.ID) == id {
                                return article, nil
                            }
                        }
                    }
                    return nil, nil
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
            Type:        models.ArticleType,
            Description: "Create new article",
            Args: graphql.FieldConfigArgument{
                "name": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "describle": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                rand.Seed(time.Now().UnixNano())
                article := models.Article{
                    ID:    int64(rand.Intn(100000)), // generate random ID
                    Name:  params.Args["name"].(string),
                    Describle:  params.Args["describle"].(string),
                }
                models.Articles = append(models.Articles, article)
                return article, nil
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


