package controllers

import (
	"errors"
	"strings"

	output "github.com/Goldwin/ies-pik-cms/internal/out/auth"
	"github.com/Goldwin/ies-pik-cms/pkg/auth"
	"github.com/Goldwin/ies-pik-cms/pkg/auth/dto"
	"github.com/Goldwin/ies-pik-cms/pkg/common/commands"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authComponent       auth.AuthComponent
	authOutputComponent output.AuthOutputComponent
}

func InitializeAuthController(r *gin.Engine, authComponent auth.AuthComponent, authOutputComponent output.AuthOutputComponent) {
	authController := authController{
		authComponent:       authComponent,
		authOutputComponent: authOutputComponent,
	}

	authGroup := r.Group("auth")
	authGroup.GET("", authController.auth, func(ctx *gin.Context) {
		result, ok := ctx.Get("auth_data")
		if !ok {
			ctx.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": result,
		})
	})
	authGroup.POST("registration", authController.auth, authController.completeRegistration)
	authGroup.POST("otp", authController.otp)
	authGroup.POST("otp/signin", authController.otpSignIn)
}

func (a *authController) completeRegistration(c *gin.Context) {
	var input dto.CompleteRegistrationInput
	authRaw, ok := c.Get("auth_data")

	if !ok {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	c.BindJSON(&input)
	authData, ok := authRaw.(dto.AuthData)

	if !ok {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Malformed Token",
		})
		return
	}

	input.Email = authData.Email

	output := &outputDecorator[dto.AuthData]{
		output: a.authOutputComponent.RegistrationOutput(),
		errFunction: func(err commands.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.AuthData) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}
	a.authComponent.CompleteRegistration(c, input, output)
}

func (a *authController) auth(c *gin.Context) {
	var input dto.AuthInput
	header := c.GetHeader("Authorization")
	token, err := extractBearerToken(header)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": gin.H{
				"code":    1,
				"message": "Unauthorized",
			},
		})
		return
	}

	input.Token = token

	output := &outputDecorator[dto.AuthData]{
		output: a.authOutputComponent.AuthOutput(),
		errFunction: func(err commands.AppErrorDetail) {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.AuthData) {
			c.Set("auth_data", result)
		},
	}
	a.authComponent.Auth(c, input, output)
}

func (a *authController) otp(c *gin.Context) {
	var input dto.OtpInput
	c.BindJSON(&input)
	a.authComponent.GenerateOtp(c, input, a.authOutputComponent.OTPOutput())
	c.JSON(204, gin.H{})
}

func (a *authController) otpSignIn(c *gin.Context) {
	var input dto.SignInInput
	c.BindJSON(&input)
	input.Method = "otp"
	output := &outputDecorator[dto.SignInResult]{
		output: a.authOutputComponent.SignInOutput(),
		errFunction: func(err commands.AppErrorDetail) {
			c.JSON(400, gin.H{
				"error": err,
			})
		},
		successFunc: func(result dto.SignInResult) {
			c.JSON(200, gin.H{
				"data": result,
			})
		},
	}
	a.authComponent.SignIn(c, input, output)
}

func extractBearerToken(bearer string) (string, error) {
	if bearer == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(bearer, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
