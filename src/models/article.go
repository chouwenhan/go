package models

import (
    "github.com/graphql-go/graphql"
)
var Articles []Article 

type Article struct {
    ID    string   `json:"_id"`
    Name  string  `json:"name"`
    Describle  string  `json:"info,omitempty"`
    Note string `json:"note"`
}

type CreateArticle struct {
    Name  string  `json:"name"`
    Describle  string  `json:"info,omitempty"`
    Note string `json:"note"`
}

var ArticleType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Article",
        Fields: graphql.Fields{
            "_id": &graphql.Field{
                Type: graphql.String,
            },
            "name": &graphql.Field{
                Type: graphql.String,
            },
            "describle": &graphql.Field{
                Type: graphql.String,
            },
            "note": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

// var ArticleListType = graphql.NewList(ArticleType)
