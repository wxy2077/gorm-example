package controller

import (
	"github.com/gin-gonic/gin"
	"gorm-example/global"
	"gorm-example/logic"
	"gorm-example/utils"
	"net/http"
)

func JoinFunc(c *gin.Context) {

	list, pagination := logic.NewUserLogic().UserDepList(global.DB.WithContext(c.Request.Context()), 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, res)
}

func PreloadFunc(c *gin.Context) {
	list, pagination := logic.NewUserLogic().PreloadUserDep(global.DB.WithContext(c.Request.Context()), 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, res)
}

func PreloadsFunc(c *gin.Context) {

	list, pagination := logic.NewUserLogic().PreloadUserDeps(global.DB.WithContext(c.Request.Context()), 1)

	res := &utils.OkWithPage{
		List:       list,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, res)
}
