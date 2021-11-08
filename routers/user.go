package routers

import "github.com/gin-gonic/gin"

func UserLog(e *gin.Engine) {
	e.GET("/signIn", signInHandler)
	e.PUT("/signUp", signUpHandler)
	e.PUT("/changePwd", changePwdHandler)
	e.PUT("/changeName", changeNameHandler)
}

func signInHandler(c *gin.Context) {}

func signUpHandler(c *gin.Context) {}

func changePwdHandler(c *gin.Context) {}

func changeNameHandler(c *gin.Context) {}
