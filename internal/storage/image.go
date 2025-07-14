package storage

import (
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
    "fmt"
    "math/rand"
    "time"
)

func SaveMultipleFiles(c *gin.Context, field string, folder string) ([]string, error) {
    c.Request.ParseMultipartForm(100 << 20)
    form := c.Request.MultipartForm
    files := form.File[field]

    var paths []string
    basePath := "./public/image/" + folder

    os.MkdirAll(basePath, os.ModePerm)

    for _, file := range files {
        name := generateFileName(file.Filename)
        path := filepath.Join(basePath, name)
        c.SaveUploadedFile(file, path)
        paths = append(paths, filepath.ToSlash(filepath.Join("image", folder, name)))
    }

    return paths, nil
}

func SaveFile(c *gin.Context, field string, folder string) ([]string, error) {
    c.Request.ParseMultipartForm(100 << 20)
    form := c.Request.MultipartForm
    files := form.File[field]
    var paths []string
    basePath := "./public/file/" + folder
    os.MkdirAll(basePath, os.ModePerm)
    for _, file := range files {
        name := generateFileName(file.Filename)
        path := filepath.Join(basePath, name)
        c.SaveUploadedFile(file, path)
        paths = append(paths, filepath.ToSlash(filepath.Join("file", folder, name)))
    }
    return paths, nil
}

func generateFileName(original string) string {
    ext := filepath.Ext(original)
    return fmt.Sprintf("img_%d_%d%s", time.Now().UnixNano(), rand.Intn(9999), ext)
}
