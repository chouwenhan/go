package main

import (
    "github.com/gin-gonic/gin"
    "src/controllers"
    "src/models"
)

func main() {

    router := gin.Default()
    // test
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    // controllers start
    language :=router.Group("/language")
    {
        language.GET("",controllers.LanguageTest)
    }

    // Primary data initialization
    graphql := router.Group("/graphql")
    {
        models.InitArticlesData(&models.Articles)
        graphql.Any("/article", controllers.ArticleHandler())
    }

    router.Run() // listen and serve on 0.0.0.0:8080
}