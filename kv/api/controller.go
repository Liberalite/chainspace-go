package api

import (
	"net/http"

	"chainspace.io/prototype/kv"
	"chainspace.io/prototype/sbac"
	"github.com/gin-gonic/gin"
)

// Controller is the Key-Value controller
type controller struct {
	service *service
}

// Controller is the Key-Value controller
type Controller interface {
	GetByLabel(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new kv.Controller
func New(kvStore kv.Service, sbac *sbac.Service) Controller {
	return &controller{newService(kvStore, sbac)}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.GET("/api/kv/:label", controller.GetByLabel)
	router.GET("/api/kv/:label/version-id", controller.GetVersionIDByLabel)
}

// GetByLabel Retrieves a key by its value
// @Summary Retrieve a key by its label
// @Description get string by label
// @ID getbyLabel
// @Accept  json
// @Produce  json
// @Tags kv
// @Param   label      path   string     true  "label"
// @Success 200 {object} api.ObjectResponse
// @Failure 400 {object} api.Error
// @Failure 404 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /api/kv/{label} [get]
func (controller *controller) GetByLabel(c *gin.Context) {
	label := c.Param("label")
	obj, status, err := controller.service.Get(label)
	if err != nil {
		c.JSON(status, Error{err.Error()})
		return
	}
	c.JSON(http.StatusOK, ObjectResponse{Object: obj})
}

// GetVersionIDByLabel Retrieves a version ID by its value
// @Summary Retrieve a version ID by its label
// @Description get version ID by label
// @ID getVersionIDbyLabel
// @Accept  json
// @Produce  json
// @Tags kv
// @Param   label      path   string     true  "label"
// @Success 200 {object} api.VersionIDResponse
// @Failure 400 {object} api.Error
// @Failure 404 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /api/kv/{label}/version-id [get]
func (controller *controller) GetVersionIDByLabel(c *gin.Context) {
	label := c.Param("label")
	versionID, status, err := controller.service.GetVersionID(label)
	if err != nil {
		c.JSON(status, Error{err.Error()})
		return
	}
	c.JSON(http.StatusOK, VersionIDResponse{VersionID: versionID})
}