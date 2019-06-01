package usecases

import (
	"github.com/gin-gonic/gin"
)

type APIController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
