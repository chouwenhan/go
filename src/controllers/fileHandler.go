package controllers
import (
    "github.com/gin-gonic/gin"
    "src/models"
    // "fmt"
)


// type File struct {
//     doc_id string
//     url string
// }

func FileHandler(c *gin.Context) {
    doc_id := c.Param("doc_id")
    db := models.ConnDB("article")
    result, _ := models.ReadDocument(db, doc_id)
    file, _ := c.FormFile("file")
    fileName := file.Filename
    fileType := file.Header["Content-Type"][0]
    file_open, _ := file.Open()
    save_result, err := db.SaveAttachment(doc_id, result.Rev, fileName, fileType, file_open)
    if err != nil {
        panic(err)
    }
    c.JSON(200, gin.H{"result": save_result})
}