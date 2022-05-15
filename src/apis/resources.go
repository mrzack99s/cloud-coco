package apis

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/cloud-coco/src/models"
	"github.com/mrzack99s/cloud-coco/src/services"
	"github.com/mrzack99s/cloud-coco/src/types"
)

type resourceController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create Resources
// @Description Create Resources
// @ID create-resources
// @Accept   json
// @security ApiKeyAuth
// @Tags	Resources
// @Produce  json
// @Param params body models.Resources true "Parameters"
// @Success 200 {object} models.Resources
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resources/create [post]
func (ctl *resourceController) create(c *gin.Context) {
	params := models.Resources{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		_, err := govalidator.ValidateStruct(params)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	if err := services.Create(&params); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, params)

}

// Headers godoc
// @Summary Update Resources
// @Description Update Resources
// @ID update-resources
// @Accept   json
// @security ApiKeyAuth
// @Tags	Resources
// @Produce  json
// @Param params body models.Resources true "Parameters"
// @Success 200 {object} models.Resources
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resources/update [post]
func (ctl *resourceController) update(c *gin.Context) {
	params := models.Resources{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		_, err := govalidator.ValidateStruct(params)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	if err := services.Update(&params); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, params)

}

// Headers godoc
// @Summary Soft Delete Resources
// @Description Soft Delete Resources
// @ID soft-delete-resources
// @Tags	Resources
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Resources UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resources/soft-delete/{uuid} [delete]
func (ctl *resourceController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDelete(uuid, &models.Resources{}); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete Resources
// @Description Hard Delete Resources
// @ID hard-delete-resources
// @Tags	Resources
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Resources UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resources/hard-delete/{uuid} [delete]
func (ctl *resourceController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDelete(uuid, &models.Resources{}); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get Resources by uuid
// @Description Get Resources by uuid
// @ID get-resources-by-uuid
// @Tags	Resources
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Resources
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resources/get/{uuid} [get]
func (ctl *resourceController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.Resources{}
	err := services.GetByUUID(uuid, &o)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, o)
}

// Headers godoc
// @Summary Get Resources by offset
// @Description Get Resources by offset
// @ID get-resources-by-offset
// @Tags	Resources
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resources/get-by-offset [get]
func (ctl *resourceController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.Resources{}
	err := services.GetByOffset(page_no, limit, &o)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	rCound, err := services.CountAll(&o)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.ArrayResponse{
		RecordCount: rCound,
		Records:     o,
	})
}

func NewResourceController(router gin.IRouter) *resourceController {
	s := &resourceController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *resourceController) register() {
	api := ctl.router.Group("/resources")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
