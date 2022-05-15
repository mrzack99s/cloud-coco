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

type serviceVersionsController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create ServiceVersions
// @Description Create ServiceVersions
// @ID create-service-versions
// @Accept   json
// @security ApiKeyAuth
// @Tags	ServiceVersions
// @Produce  json
// @Param params body models.ServiceVersions true "Parameters"
// @Success 200 {object} models.ServiceVersions
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/create [post]
func (ctl *serviceVersionsController) create(c *gin.Context) {
	params := models.ServiceVersions{}
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
// @Summary Update ServiceVersions
// @Description Update ServiceVersions
// @ID update-service-versions
// @Accept   json
// @security ApiKeyAuth
// @Tags	ServiceVersions
// @Produce  json
// @Param params body models.ServiceVersions true "Parameters"
// @Success 200 {object} models.ServiceVersions
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/update [post]
func (ctl *serviceVersionsController) update(c *gin.Context) {
	params := models.ServiceVersions{}
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
// @Summary Soft Delete ServiceVersions
// @Description Soft Delete ServiceVersions
// @ID soft-delete-service-versions
// @Tags	ServiceVersions
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Services UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/soft-delete/{uuid} [delete]
func (ctl *serviceVersionsController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDeleteServiceVersions(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete ServiceVersions
// @Description Hard Delete ServiceVersions
// @ID hard-delete-service-versions
// @Tags	ServiceVersions
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Services UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/hard-delete/{uuid} [delete]
func (ctl *serviceVersionsController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDeleteServiceVersions(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get ServiceVersions by uuid
// @Description Get ServiceVersions by uuid
// @ID get-service-versions-by-uuid
// @Tags	ServiceVersions
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.ServiceVersions
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/get/{uuid} [get]
func (ctl *serviceVersionsController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.ServiceVersions{}
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
// @Summary Get ServiceVersions by sid
// @Description Get ServiceVersions by sid
// @ID get-service-versions-by-sid
// @Tags	ServiceVersions
// @Produce  json
// @security ApiKeyAuth
// @Param sid query int true "ServiceID"
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/get-by-sid [get]
func (ctl *serviceVersionsController) getBySid(c *gin.Context) {
	sid, _ := strconv.Atoi(c.Query("sid"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	rCound, o, err := services.GetByServicesVersionBySID(uint(sid), page_no, limit)
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

// Headers godoc
// @Summary Get ServiceVersions by offset
// @Description Get ServiceVersions by offset
// @ID get-service-versions-by-offset
// @Tags	ServiceVersions
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /service-versions/get-by-offset [get]
func (ctl *serviceVersionsController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.ServiceVersions{}
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

func NewServiceVersionsController(router gin.IRouter) *serviceVersionsController {
	s := &serviceVersionsController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *serviceVersionsController) register() {
	api := ctl.router.Group("/service-versions")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
	api.GET("get-by-sid", ctl.getBySid)

}
