package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var fileNames map[string]struct{} = map[string]struct{}{"Mary": {}, "Vickie": {}, "Kelly": {}, "Catherine": {}, "Harold": {}, "Madison": {}, "Sydney": {}, "Rebecca": {}, "Abigail": {}, "Wendy": {}}

func main() {
    router := gin.Default()
    router.MaxMultipartMemory = 8 << 20  // 8 MiB
    router.POST("/newMaterial", newMaterial)
    router.POST("/delete", delete)
    router.POST("/newImage/:folderName/:size", newImage)

    // cd into ./images
    err := os.Chdir("images")
    if err != nil {
        fmt.Println("no images folder found")
        return
    }

    router.Run("localhost:8080")
}

func newMaterial(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "*")
    c.Header("Access-Control-Allow-Headers", "*")

    dst := "./images/"

    // // Single file
    // file, err := c.FormFile("imageFile")
    // if checkErr(err, "wrong key on form", c) {
    //     return
    // }

    // Get material name
    materialName, ok1 := c.GetPostForm("materialName")
    if ok1 != true {
        fmt.Println(materialName)
        fmt.Println("Error!!!111")
        c.String(http.StatusNotFound, "Name of material not found")
        return
    }
    fmt.Println(materialName)

    // Get file name
    // name := strings.Split(file.Filename, ".")
    // _, ok2 := fileNames[name[0]]
    // if ok2 != true {
    //     fmt.Println("Error!!!222")
    //     c.String(http.StatusNotFound, fmt.Sprintf("wrong naming of the file", file.Filename))
    //     return
    // }

    // mkdir for new material
    err := os.Mkdir(materialName + "/", os.ModePerm)
    if checkErr(err, "could not create directory", c) {
        fmt.Println("Error!!!444")
        return
    }

    fmt.Println(dst + materialName + "/")

    // Save the file to the current folder
    // c.SaveUploadedFile(file, "./" + materialName + "/" + file.Filename)
    // if checkErr(err, "could not save file", c) {
    //     fmt.Println("Error!!!555")
    //     return
    // }
    c.String(http.StatusOK, "Success")
}

func newImage(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "*")
    c.Header("Access-Control-Allow-Headers", "*")

    dst := "./images/"

    // Single file
    file, err := c.FormFile("imageFile")
    if checkErr(err, "wrong key on form", c) {
        return
    }

    // Get folder name
    materialName := c.Param("folderName")
    fmt.Println(materialName)

    // Get file name
    name := strings.Split(file.Filename, ".")
    extension := name[1]
    // _, ok2 := fileNames[name[0]]
    // if ok2 != true {
    //     fmt.Println("Error!!!222")
    //     c.String(http.StatusNotFound, fmt.Sprintf("wrong naming of the file", file.Filename))
    //     return
    // }

    // cd into ./images
    // err = os.Chdir("images")
    // if checkErr(err, "did not find image directory", c) {
    //     fmt.Println("Error!!!333")
    //     return
    // }

    // mkdir for new material
    // err = os.Mkdir(materialName + "/", os.ModePerm)
    // if checkErr(err, "could not create directory", c) {
    //     fmt.Println("Error!!!444")
    //     return
    // }

    fmt.Println(dst + materialName + "/")

    // Save the file to the current folder
    c.SaveUploadedFile(file, "./" + materialName + "/" + c.Param("size") + "." + extension)
    if checkErr(err, "could not save file", c) {
        fmt.Println("Error!!!555")
        return
    }
    c.String(http.StatusOK, "Success")
}

func delete(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "*")
    c.Header("Access-Control-Allow-Headers", "*")

    loc := c.PostForm("imageLoc")
    if loc == "" {
        fmt.Println("something went wrong 1")
        return
    }

    err := os.Remove(loc)
    if err != nil {
        fmt.Println(err)
        c.String(http.StatusNotFound, string(err.Error()))
        return
    }
}

func checkErr(err error, message string, c *gin.Context) bool {
    if err != nil {
        fmt.Println(err)
        c.String(http.StatusNotFound, message)
        return true
    }
    return false
}