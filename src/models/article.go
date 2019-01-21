package models

import (
    "github.com/graphql-go/graphql"
)
var Articles []Article 

type Article struct {
    ID    int64   `json:"id"`
    Name  string  `json:"name"`
    Describle  string  `json:"info,omitempty"`
}

var ArticleType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Article",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "name": &graphql.Field{
                Type: graphql.String,
            },
            "describle": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

func InitArticlesData(p *[]Article) {
    Article1 := Article{ID: 1, Name: "Chicha Morada", Describle: "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)"}
    *p = append(*p, Article1)
}