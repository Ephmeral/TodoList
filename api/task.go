package api

import (
	"fmt"
	"github.com/Ephmeral/TodoList/pkg/util"
	"github.com/Ephmeral/TodoList/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createService service.CreateTaskService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))

	//fmt.Println("认证token成功")
	if err := c.ShouldBind(&createService); err == nil {
		fmt.Printf("createService info:%#v\n", createService)
		res := createService.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ShowTask(c *gin.Context) {
	var showService service.ShowTaskService

	if err := c.ShouldBind(&showService); err == nil {
		fmt.Printf("showService info:%#v\n", showService)
		res := showService.ShowTask(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func ShowAllTasks(c *gin.Context) {
	var showallService service.ShowAllTaskService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))

	//fmt.Println("认证token成功")
	if err := c.ShouldBind(&showallService); err == nil {
		fmt.Printf("createService info:%#v\n", showallService)
		res := showallService.ShowAll(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteTask(c *gin.Context) {
	var deleteService service.DeleteService

	res := deleteService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTask(c *gin.Context) {
	var updateService service.UpdateService

	if err := c.ShouldBind(&updateService); err == nil {
		fmt.Printf("createService info:%#v\n", updateService)
		res := updateService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
