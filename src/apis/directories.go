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

type directoriesController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create Directories
// @Description Create Directories
// @ID create-directories
// @Accept   json
// @security ApiKeyAuth
// @Tags	Directories
// @Produce  json
// @Param params body models.Directories true "Parameters"
// @Success 200 {object} models.Directories
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /directories/create [post]
func (ctl *directoriesController) create(c *gin.Context) {
	params := models.Directories{}
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
// @Summary Update Directories
// @Description Update Directories
// @ID update-directories
// @Accept   json
// @security ApiKeyAuth
// @Tags	Directories
// @Produce  json
// @Param params body models.Directories true "Parameters"
// @Success 200 {object} models.Directories
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /directories/update [post]
func (ctl *directoriesController) update(c *gin.Context) {
	params := models.Directories{}
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
// @Summary Soft Delete Directories
// @Description Soft Delete Directories
// @ID soft-delete-directories
// @Tags	Directories
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Directories UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /directories/soft-delete/{uuid} [delete]
func (ctl *directoriesController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDeleteDirectories(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete Directories
// @Description Hard Delete Directories
// @ID hard-delete-directories
// @Tags	Directories
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Directories UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /directories/hard-delete/{uuid} [delete]
func (ctl *directoriesController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDeleteDirectories(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get Directories by uuid
// @Description Get Directories by uuid
// @ID get-directories-by-uuid
// @Tags	Directories
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Directories
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /directories/get/{uuid} [get]
func (ctl *directoriesController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.Directories{}
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
// @Summary Get Directories by offset
// @Description Get Directories by offset
// @ID get-directories-by-offset
// @Tags	Directories
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /directories/get-by-offset [get]
func (ctl *directoriesController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.Directories{}
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

func NewDirectoriesController(router gin.IRouter) *directoriesController {
	s := &directoriesController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *directoriesController) register() {
	api := ctl.router.Group("/directories")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
