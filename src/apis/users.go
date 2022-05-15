package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/cloud-coco/src/authentication"
	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
	"github.com/mrzack99s/cloud-coco/src/services"
	"github.com/mrzack99s/cloud-coco/src/types"
	"github.com/mrzack99s/cloud-coco/src/utils"
)

type usersController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Reset password Users
// @Description Reset password Users
// @ID reset-passwd-users
// @Accept   json
// @security ApiKeyAuth
// @Tags	Users
// @Produce  json
// @Param params body types.UserResetPasswdParams true "Parameters"
// @Success 200 {object} models.Users
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/reset-password [post]
func (ctl *usersController) resetPassword(c *gin.Context) {
	params := types.UserResetPasswdParams{}
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

	r, err := services.ResetUserPassword(params.UUID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)

}

// Headers godoc
// @Summary Change password Users
// @Description Change password Users
// @ID change-passwd-users
// @Accept   json
// @security ApiKeyAuth
// @Tags	Users
// @Produce  json
// @Param params body types.UserChangePasswdParams true "Parameters"
// @Success 200 {string} string "changed"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/change-password [post]
func (ctl *usersController) changePasswd(c *gin.Context) {
	params := types.UserChangePasswdParams{}
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

	err := services.ChangeUserPassword(params)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(200, "changed")

}

// Headers godoc
// @Summary BYO password Users
// @Description BYO password Users
// @ID byo-passwd-users
// @Accept   json
// @security ApiKeyAuth
// @Tags	Users
// @Produce  json
// @Param params body types.UserChangePasswdParams true "Parameters"
// @Success 200 {string} string "ok"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/byo-password [post]
func (ctl *usersController) bringYourOwnPassword(c *gin.Context) {
	params := types.UserChangePasswdParams{}
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

	access_token, err := authentication.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if utils.RedisFindExistingKey(fmt.Sprintf("token:access/%s", access_token)) {
		tokenDetailStr, err := configures.CacheInstance().Get(fmt.Sprintf("token:access/%s", access_token)).Result()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		tokenDetail := authentication.TokenDetails{}
		json.Unmarshal([]byte(tokenDetailStr), &tokenDetail)

		err = services.ChangeUserPassword(params)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

	} else {
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.String(200, "ok")

}

// Headers godoc
// @Summary Create Users
// @Description Create Users
// @ID create-users
// @Accept   json
// @security ApiKeyAuth
// @Tags	Users
// @Produce  json
// @Param params body models.Users true "Parameters"
// @Success 200 {object} models.Users
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/create [post]
func (ctl *usersController) create(c *gin.Context) {
	params := models.Users{}
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

	if err := services.CreateUser(&params); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, params)

}

// Headers godoc
// @Summary Update Users
// @Description Update Users
// @ID update-users
// @Accept   json
// @security ApiKeyAuth
// @Tags	Users
// @Produce  json
// @Param params body models.Users true "Parameters"
// @Success 200 {object} models.Users
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/update [post]
func (ctl *usersController) update(c *gin.Context) {
	params := models.Users{}
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
// @Summary Soft Delete Users
// @Description Soft Delete Users
// @ID soft-delete-users
// @Tags	Users
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Users UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/soft-delete/{uuid} [delete]
func (ctl *usersController) softDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.SoftDelete(uuid, &models.Users{}); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Hard Delete Users
// @Description Hard Delete Users
// @ID hard-delete-users
// @Tags	Users
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "Users UUID"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/hard-delete/{uuid} [delete]
func (ctl *usersController) hardDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := services.HardDelete(uuid, &models.Users{}); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "deleted")
}

// Headers godoc
// @Summary Get Users by uuid
// @Description Get Users by uuid
// @ID get-users-by-uuid
// @Tags	Users
// @Produce  json
// @security ApiKeyAuth
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Users
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/get/{uuid} [get]
func (ctl *usersController) get(c *gin.Context) {
	uuid := c.Param("uuid")
	o := models.Users{}
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
// @Summary Get Users by offset
// @Description Get Users by offset
// @ID get-users-by-offset
// @Tags	Users
// @Produce  json
// @security ApiKeyAuth
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} types.ArrayResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/get-by-offset [get]
func (ctl *usersController) getByOffset(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page_no := (offset - 1) * limit

	o := []models.Users{}
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

func NewUsersController(router gin.IRouter) *usersController {
	s := &usersController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *usersController) register() {
	api := ctl.router.Group("/users")

	api.POST("create", ctl.create)
	api.POST("update", ctl.update)
	api.POST("reset-password", ctl.resetPassword)
	api.POST("change-password", ctl.changePasswd)
	api.POST("byo-password", ctl.bringYourOwnPassword)
	api.DELETE("soft-delete/:uuid", ctl.softDelete)
	api.DELETE("hard-update/:uuid", ctl.hardDelete)
	api.GET("get/:uuid", ctl.get)
	api.GET("get-by-offset", ctl.getByOffset)
}
