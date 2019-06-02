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
	c.JSON(http.StatusOK, entities.GetAll())
}

func (this *APIControllerImpl) Post(c *gin.Context) {
	req_body := new(datatypes.Endpoint)
	if err := c.Bind(req_body); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to parse request body")
		return
	}

	if err := entities.Add(req_body); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to add request Endpoint")
		return
	}

	c.JSON(http.StatusOK, "Success to add Endpoint")
}

func (this *APIControllerImpl) Put(c *gin.Context) {
	req_body := new(datatypes.Endpoint)
	if err := c.Bind(req_body); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to parse request body")
		return
	}

	if err := entities.Update(req_body); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to update request Endpoint")
		return
	}

	c.JSON(http.StatusOK, "Success to update Endpoint")
}

func (this *APIControllerImpl) Delete(c *gin.Context) {
	req_body := new(datatypes.ExistingEndpoint)
	if err := c.Bind(req_body); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to parse request body")
		return
	}

	if err := entities.Delete(req_body.Endpoint, req_body.Method); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to delete request Endpoint")
		return
	}

	c.JSON(http.StatusOK, "Success to delete Endpoint")
}
