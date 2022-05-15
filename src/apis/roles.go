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

type rolesController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create Roles
// @Description Create Roles
// @ID create-roles
// @Accept   json
// @security ApiKeyAuth
// @Tags	Roles
// @Produce  json
// @Param params body models.Roles true "Parameters"
// @Success 200 {object} models.Roles
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/create [post]
func (ctl *rolesController) create(c *gin.Context) {
	params := models.Roles{}
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
// @Summary Update Roles
// @Description Update Roles
// @ID update-roles
// @Accept   json
// @security ApiKeyAuth
// @Tags	Roles
// @Produce  json
// @Param params body models.Roles true "Parameters"
// @Success 200 {object} models.Roles
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/update [post]
func (ctl *rolesController) update(c *gin.Context) {
	params := models.Roles{}
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
// @Summary Soft Delete Roles
// @Description Soft Delete Roles
// @ID soft-delete-roles
// @Tags	Roles
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Roles UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/soft-delete/{uuid} [delete]
func (ctl *rolesController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDelete(uuid, &models.Roles{}); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete Roles
// @Description Hard Delete Roles
// @ID hard-delete-roles
// @Tags	Roles
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Roles UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/hard-delete/{uuid} [delete]
func (ctl *rolesController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDelete(uuid, &models.Roles{}); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get Roles by uuid
// @Description Get Roles by uuid
// @ID get-roles-by-uuid
// @Tags	Roles
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Roles
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/get/{uuid} [get]
func (ctl *rolesController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.Roles{}
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
// @Summary Get Roles by offset
// @Description Get Roles by offset
// @ID get-roles-by-offset
// @Tags	Roles
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/get-by-offset [get]
func (ctl *rolesController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.Roles{}
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

func NewRolesController(router gin.IRouter) *rolesController {
	s := &rolesController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *rolesController) register() {
	api := ctl.router.Group("/roles")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
