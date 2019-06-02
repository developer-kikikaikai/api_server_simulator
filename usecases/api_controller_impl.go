package usecases

import (
	"net/http"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
	"github.com/developer-kikikaikai/api_server_simulator/entities"
	"github.com/gin-gonic/gin"
)

type APIControllerImpl struct {
}

func (this *APIControllerImpl) Get(c *gin.Context) {
	entities := entities.GetAll()
	c.JSON(http.StatusOK, entities)
}

func (this *APIControllerImpl) Post(c *gin.Context) {
	req_body := new(datatypes.Endpoint)
	if err := c.Bind(req_body); err != nil {
		c.JSON(http.StatusBadRequest, datatypes.ParseFailed)
		return
	}

	if err := entities.Add(req_body); err != nil {
		c.JSON(http.StatusBadRequest, datatypes.FailedToAdd)
		return
	}

	c.JSON(http.StatusOK, datatypes.PostSuccess)
}

func (this *APIControllerImpl) Put(c *gin.Context) {
	req_body := new(datatypes.Endpoint)
	if err := c.Bind(req_body); err != nil {
		c.JSON(http.StatusBadRequest, datatypes.ParseFailed)
		return
	}

	if err := entities.Update(req_body); err != nil {
		c.JSON(http.StatusBadRequest, datatypes.FailedToUpdate)
		return
	}

	c.JSON(http.StatusOK, datatypes.PutSuccess)
}

func (this *APIControllerImpl) Delete(c *gin.Context) {
	req_body := new(datatypes.ExistingEndpoint)
	if err := c.Bind(req_body); err != nil {
		c.JSON(http.StatusBadRequest, datatypes.ParseFailed)
		return
	}

	if err := entities.Delete(req_body.Endpoint, req_body.Method); err != nil {
		c.JSON(http.StatusBadRequest, datatypes.FailedToDelete)
		return
	}

	c.JSON(http.StatusOK, datatypes.DeleteSuccess)
}
