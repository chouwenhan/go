package controllers
import (
    "github.com/gin-gonic/gin"
    "src/models"
    "fmt"
)


func FileHandler(c *gin.Context) {
    doc_id := c.Param("doc_id")
    db := models.ConnDB("article")
    result, _ := models.ReadDocument(db, doc_id)
    form, _ := c.MultipartForm()
    fmt.Println(form)
    files := form.File["files"]
    fmt.Println(files)
    for _, file := range files {
        fileName := file.Filename
        fileType := file.Header["Content-Type"][0]
        file_open, _ := file.Open()
        save_result, err := db.SaveAttachment(doc_id, result.Rev, fileName, fileType, file_open)
        if err != nil {
            panic(err)
        }
        result, _ = models.ReadDocument(db, doc_id)
        fmt.Println(save_result)
    }
    c.JSON(200, gin.H{
        "message": "upload files successful",
    })
}