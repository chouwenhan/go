package models

import (
    "github.com/graphql-go/graphql"
)
type Articles struct {
    Docs []Article `json:"docs"`
}

type Article struct {
    ID    string  `json:"_id"`
    Rev   string  `json:"_rev"`
    Type  string  `json:"type"`
    Title  string  `json:"title"`
    Content  string  `json:"info,omitempty"`
    Tags string  `json:"tags"`
    Attachments map[string]Attachment  `json:"_attachments,omitempty"`
}

type Attachment struct {
    ContentType string `json:"content_type"`
    Revpos int `json:"revpos"`
    Digest string `json:"digest"`
    Length int `json:"length"`
    Stub bool `json:"stub"`
}

type CreateArticle struct {
    Type  string   `json:"type"`
    Title  string  `json:"title"`
    Content  string  `json:"info,omitempty"`
    Tags string `json:"tags"`
}

type UpdateArticle struct {
    Type  string   `json:"type"`
    Title  string  `json:"title"`
    Content  string  `json:"info,omitempty"`
    Tags string `json:"tags"`
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
            "tags": &graphql.Field{
                Type: graphql.String,
            },
            "message": &graphql.Field{
                Type: graphql.String,
            },
            "_attachments": &graphql.Field{
                Type:  JSON,
            },
        },
    },
)

// var ArticleListType = graphql.NewList(ArticleType)
