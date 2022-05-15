package apis

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/cloud-coco/src/services"
	"github.com/mrzack99s/cloud-coco/src/types"
)

type setupController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Setup system
// @Description Setup system
// @ID setup-system
// @Accept   json
// @security ApiKeyAuth
// @Tags	Setup
// @Produce  json
// @Param params body types.SetupParams true "Parameters"
// @Success 200 {object} models.Users
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /setup [post]
func (ctl *setupController) create(c *gin.Context) {
	params := types.SetupParams{}
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

func NewSetupController(router gin.IRouter) *setupController {
	s := &setupController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *setupController) register() {
	api := ctl.router.Group("/users")

	api.POST("create", ctl.create)
}
