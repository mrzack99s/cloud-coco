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

type servicesController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create Services
// @Description Create Services
// @ID create-services
// @Accept   json
// @security ApiKeyAuth
// @Tags	Services
// @Produce  json
// @Param params body models.Services true "Parameters"
// @Success 200 {object} models.Services
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /services/create [post]
func (ctl *servicesController) create(c *gin.Context) {
	params := models.Services{}
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
// @Summary Update Services
// @Description Update Services
// @ID update-services
// @Accept   json
// @security ApiKeyAuth
// @Tags	Services
// @Produce  json
// @Param params body models.Services true "Parameters"
// @Success 200 {object} models.Services
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /services/update [post]
func (ctl *servicesController) update(c *gin.Context) {
	params := models.Services{}
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
// @Summary Soft Delete Services
// @Description Soft Delete Services
// @ID soft-delete-services
// @Tags	Services
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Services UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /services/soft-delete/{uuid} [delete]
func (ctl *servicesController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDeleteServices(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete Services
// @Description Hard Delete Services
// @ID hard-delete-services
// @Tags	Services
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Services UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /services/hard-delete/{uuid} [delete]
func (ctl *servicesController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDeleteServices(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get Services by uuid
// @Description Get Services by uuid
// @ID get-services-by-uuid
// @Tags	Services
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Services
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /services/get/{uuid} [get]
func (ctl *servicesController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.Services{}
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
// @Summary Get Services by offset
// @Description Get Services by offset
// @ID get-services-by-offset
// @Tags	Services
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /services/get-by-offset [get]
func (ctl *servicesController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.Services{}
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

func NewServicesController(router gin.IRouter) *servicesController {
	s := &servicesController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *servicesController) register() {
	api := ctl.router.Group("/services")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
