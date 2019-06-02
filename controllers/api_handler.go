package controllers

import (
	"github.com/developer-kikikaikai/api_server_simulator/usecases"
	"github.com/gin-gonic/gin"
)

func GETEndpointsHandler(c *gin.Context) {
	usecases.GetAPIController().Get(c)
}

func POSTEndpointsHandler(c *gin.Context) {
	usecases.GetAPIController().Post(c)
}

func PUTEndpointsHandler(c *gin.Context) {
	usecases.GetAPIController().Put(c)
}

func DELETEEndpointsHandler(c *gin.Context) {
	usecases.GetAPIController().Delete(c)
}
