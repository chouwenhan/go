package models

import (
    "github.com/graphql-go/graphql"
)
type Articles struct {
    Docs []Article `json:"docs"`
}

type Article struct {
    ID    string   `json:"_id"`
    Rev   string   `json:"_rev"`
    Type  string   `json:"type"`
    Title  string  `json:"title"`
    Content  string  `json:"info,omitempty"`
    Note string `json:"note"`
}

type CreateArticle struct {
    Type  string   `json:"type"`
    Title  string  `json:"title"`
    Content  string  `json:"info,omitempty"`
    Note string `json:"note"`
}

type UpdateArticle struct {
    Type  string   `json:"type"`
    Title  string  `json:"title"`
    Content  string  `json:"info,omitempty"`
    Note string `json:"note"`
}

var ArticleType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Article",
        Fields: graphql.Fields{
            "_id": &graphql.Field{
                Type: graphql.String,
            },
            "type": &graphql.Field{
                Type: graphql.String,
            },
            "title": &graphql.Field{
                Type: graphql.String,
            },
            "content": &graphql.Field{
                Type: graphql.String,
            },
            "note": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

// var ArticleListType = graphql.NewList(ArticleType)
