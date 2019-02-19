package main

import (
    "github.com/gin-gonic/gin"
    "src/controllers"
    "src/models"
    "fmt"
)

// middleware
func CORSMiddleware(c *gin.Context){
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

    c.Next()
}

func main() {
    models.CreateDatabase("article")
    // gin start
    router := gin.Default()
    router.Use(CORSMiddleware)
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
        fmt.Println("123")
        graphql.Any("/article", controllers.ArticleHandler())
    }

    router.Run() // listen and serve on 0.0.0.0:8080
}