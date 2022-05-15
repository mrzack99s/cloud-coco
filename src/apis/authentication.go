package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/cloud-coco/src/authentication"
	"github.com/mrzack99s/cloud-coco/src/services"
)

type AuthenticationController struct {
	router gin.IRouter
}

// Headers godoc
// @Summary Revoke token
// @Description Revoke token
// @security ApiKeyAuth
// @ID get-credential
// @Accept   json
// @Tags	Authentication
// @Produce  json
// @Success 200 {object} models.Users
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /authentication/get-credential [get]
func (ctl *AuthenticationController) getCredential(c *gin.Context) {
	access_token, err := authentication.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := services.GetCredential(access_token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)

}

// Headers godoc
// @Summary Check credential
// @Description Check credential for get access
// @ID check-credential
// @Accept   json
// @Tags	Authentication
// @Produce  json
// @Param params body authentication.CredentialParams true "Parameters"
// @Success 200 {object} authentication.TokenResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /authentication/check-credential [post]
func (ctl *AuthenticationController) checkCredential(c *gin.Context) {
	params := authentication.CredentialParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	r, err := services.CheckCredential(params)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)

}

// Headers godoc
// @Summary Revoke token
// @Description Revoke token
// @security ApiKeyAuth
// @ID revoke-token
// @Accept   json
// @Tags	Authentication
// @Produce  json
// @Success 200 {string} string "Revoked"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /authentication/revoke-token [get]
func (ctl *AuthenticationController) revokeToken(c *gin.Context) {

	access_token, err := authentication.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = services.RevokeCredential(access_token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "revoked")

}

func NewAuthenticationController(router gin.IRouter) *AuthenticationController {
	s := &AuthenticationController{
		router: router,
	}
	s.register()
	return s
}

func (ctl *AuthenticationController) register() {
	api := ctl.router.Group("/authentication")

	api.GET("get-credential", ctl.getCredential)
	api.POST("check-credential", ctl.checkCredential)
	api.GET("revoke-token", authentication.TokenMiddleware, ctl.revokeToken)
}
