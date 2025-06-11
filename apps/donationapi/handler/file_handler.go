package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"sharedinfra/cipher"
	"sharedinfra/fileio"
)

const (
	GroupPathFile = "/file"
)

// func encodeRot128(w http.ResponseWriter, r *http.Request) {
func encodeRot128(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}
	file, _ := formFile.Open()

	readBuffer := make([]byte, 1024)
	rot128Buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	writer, _ := cipher.NewRot128Writer(rot128Buffer)
	for {
		bytesRead, err := file.Read(readBuffer)
		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		_, _ = writer.Write(readBuffer[:bytesRead])
	}

	c.Data(http.StatusOK, "text/plain", rot128Buffer.Bytes())
}

// func decodeRot128(w http.ResponseWriter, r *http.Request) {
func decodeRot128(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}
	file, _ := formFile.Open()

	decoded, err := fileio.DecodeRot128(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}
	defer file.Close()

	data := make([]byte, 0, 2048)
	readBuffer := make([]byte, 1024)

	for {
		bytesRead, err := decoded.Read(readBuffer)
		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		data = append(data, readBuffer[:bytesRead]...)
	}

	c.Data(http.StatusOK, "application/octet-stream", data)
}

//func FileRouter() http.Handler {
//	mux := http.NewServeMux()
//	mux.HandleFunc("POST /rot128/encode", encodeRot128)
//	mux.HandleFunc("POST /rot128/decode", decodeRot128)
//
//	return mux
//}

func FileRouter(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group(GroupPathFile)
	rot128 := group.Group("/rot128")
	{
		rot128.POST("/encode", encodeRot128)
		rot128.POST("/decode", decodeRot128)
	}

	return group
}
