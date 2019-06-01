package controllers

import (
	"io"

	"github.com/gin-gonic/gin"
)

func GETEndpointsHandler(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}

func POSTEndpointsHandler(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}
func PUTEndpointsHandler(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}
func DELETEEndpointsHandler(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}
