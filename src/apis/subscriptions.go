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

type subscriptionsController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Create Subscriptions
// @Description Create Subscriptions
// @ID create-subscriptions
// @Accept   json
// @security ApiKeyAuth
// @Tags	Subscriptions
// @Produce  json
// @Param params body models.Subscriptions true "Parameters"
// @Success 200 {object} models.Subscriptions
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subscriptions/create [post]
func (ctl *subscriptionsController) create(c *gin.Context) {
	params := models.Subscriptions{}
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
// @Summary Update Subscriptions
// @Description Update Subscriptions
// @ID update-subscriptions
// @Accept   json
// @security ApiKeyAuth
// @Tags	Subscriptions
// @Produce  json
// @Param params body models.Subscriptions true "Parameters"
// @Success 200 {object} models.Subscriptions
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subscriptions/update [post]
func (ctl *subscriptionsController) update(c *gin.Context) {
	params := models.Subscriptions{}
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
// @Summary Soft Delete Subscriptions
// @Description Soft Delete Subscriptions
// @ID soft-delete-subscriptions
// @Tags	Subscriptions
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Subscriptions UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subscriptions/soft-delete/{uuid} [delete]
func (ctl *subscriptionsController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDeleteSubscriptions(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete Subscriptions
// @Description Hard Delete Subscriptions
// @ID hard-delete-subscriptions
// @Tags	Subscriptions
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Subscriptions UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subscriptions/hard-delete/{uuid} [delete]
func (ctl *subscriptionsController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDeleteSubscriptions(uuid); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get Subscriptions by uuid
// @Description Get Subscriptions by uuid
// @ID get-subscriptions-by-uuid
// @Tags	Subscriptions
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Subscriptions
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subscriptions/get/{uuid} [get]
func (ctl *subscriptionsController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.Subscriptions{}
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
// @Summary Get Subscriptions by offset
// @Description Get Subscriptions by offset
// @ID get-subscriptions-by-offset
// @Tags	Subscriptions
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subscriptions/get-by-offset [get]
func (ctl *subscriptionsController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.Subscriptions{}
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

func NewSubscriptionsController(router gin.IRouter) *subscriptionsController {
	s := &subscriptionsController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *subscriptionsController) register() {
	api := ctl.router.Group("/subscriptions")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
