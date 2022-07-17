package controllers

import (
	"fmt"
	"net/http"
	
	"dev-platform/models"

	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context){

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	userModel := models.User{Email: email, Password: password}

	userInfo, err := userModel.UserGetByEmail(userModel.Email)

	fmt.Println("userInfo", userInfo , err)

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
		"uid": userInfo.Id,
	})
}