package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-sharex/config"
	"go-sharex/er"
	"log"
	"net/http"
	"path/filepath"
)

func Upload(ctx *gin.Context) {
	if file, err := ctx.FormFile("upload"); err == nil {
		newName := uuid.New().String() + filepath.Ext(file.Filename) // Why UUID? Cuz always unique
		if err = ctx.SaveUploadedFile(file, "./uploads/" + newName); err == nil {
			log.Printf("%s has uploaded %s\n", config.Get().Keys[ctx.Query("key")], newName)
			ctx.String(http.StatusOK, config.Get().BaseURL + "/" + newName)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": er.UploadFailed})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{"error" : er.UploadNoFileProvided})
	}
}