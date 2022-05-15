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

type resourcePoolsController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create ResourcePools
// @Description Create ResourcePools
// @ID create-resource-pools
// @Accept   json
// @security ApiKeyAuth
// @Tags	ResourcePools
// @Produce  json
// @Param params body models.ResourcePools true "Parameters"
// @Success 200 {object} models.ResourcePools
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resource-pools/create [post]
func (ctl *resourcePoolsController) create(c *gin.Context) {
	params := models.ResourcePools{}
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
// @Summary Update ResourcePools
// @Description Update ResourcePools
// @ID update-resource-pools
// @Accept   json
// @security ApiKeyAuth
// @Tags	ResourcePools
// @Produce  json
// @Param params body models.ResourcePools true "Parameters"
// @Success 200 {object} models.ResourcePools
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resource-pools/update [post]
func (ctl *resourcePoolsController) update(c *gin.Context) {
	params := models.ResourcePools{}
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
// @Summary Soft Delete ResourcePools
// @Description Soft Delete ResourcePools
// @ID soft-delete-resource-pools
// @Tags	ResourcePools
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "ResourcePools UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resource-pools/soft-delete/{uuid} [delete]
func (ctl *resourcePoolsController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDeleteResourcePools(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete ResourcePools
// @Description Hard Delete ResourcePools
// @ID hard-delete-resource-pools
// @Tags	ResourcePools
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "ResourcePools UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resource-pools/hard-delete/{uuid} [delete]
func (ctl *resourcePoolsController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDeleteResourcePools(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get ResourcePools by uuid
// @Description Get ResourcePools by uuid
// @ID get-resource-pools-by-uuid
// @Tags	ResourcePools
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.ResourcePools
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resource-pools/get/{uuid} [get]
func (ctl *resourcePoolsController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.ResourcePools{}
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
// @Summary Get ResourcePools by offset
// @Description Get ResourcePools by offset
// @ID get-resource-pools-by-offset
// @Tags	ResourcePools
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resource-pools/get-by-offset [get]
func (ctl *resourcePoolsController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.ResourcePools{}
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

func NewResourcePoolsController(router gin.IRouter) *resourcePoolsController {
	s := &resourcePoolsController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *resourcePoolsController) register() {
	api := ctl.router.Group("/resource-pools")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
