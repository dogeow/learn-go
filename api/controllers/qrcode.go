package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

func RenderQRExport(c *gin.Context) {
	img, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	images := base64.StdEncoding.EncodeToString(img)
	if err != nil {
		fmt.Print(err)
	}

	c.HTML(http.StatusOK, "qr.html", gin.H{
		"images": images,
	})
}
