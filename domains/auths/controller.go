package auths

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var service = NewAuthService()

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (a *authController) Login(c *gin.Context) {
	/* Get params body */
	logrus.Info("Login")

	/** Get login request**/
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	/** Call service **/
	data, err := service.Login(request)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Login",
		"data":    data,
	})
}

func (a *authController) Register(c *gin.Context) {
	logrus.Info("Register")

	/** Get register request**/
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(422, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	/** Call service **/
	data, err := service.Register(request)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Register Success",
		"data":    data,
	})
}
