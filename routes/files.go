package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sharex/config"
	"go-sharex/er"
	"go-sharex/utils"
	"net/http"
	"os"
)

const discordAgent = "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)"

func Files(ctx *gin.Context) {
	if fName := ctx.Param("file"); len(fName) > 0 {
		fPath := fmt.Sprintf("./uploads/%s", fName)
		if _, err := os.Stat(fPath); err == nil {
			if agent := ctx.Request.UserAgent(); agent == discordAgent && (utils.GetFileType(fPath) == "image" || utils.GetFileType(fPath) == "video") { // Discord embed
				data := gin.H{
					"url": fmt.Sprintf("%s/%s", config.Get().BaseURL, fName),
					"file": fmt.Sprintf("/%s", fName),
					"color": config.Get().EmbedColor,
				}
				ctx.HTML(http.StatusOK, fmt.Sprintf("%s.tmpl", utils.GetFileType(fPath)), data)
			} else { // regular browser / not image / video
				ctx.File(fmt.Sprintf("./uploads/%s", fName))
			}
		} else {
			ctx.JSON(http.StatusOK, er.FilesNotFound)
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{"error": er.FilesNoFile})
	}
}