package service

import (
	"fmt"
	"github.com/Ephmeral/TodoList/model"
	"github.com/Ephmeral/TodoList/serializer"
	"time"
)

// CreateTaskService 创建任务的服务
type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0未完成，1已完成
}

type ShowTaskService struct {
}

type ShowAllTaskService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

type DeleteService struct {
}

type UpdateService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

// Create 创建一条新的备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, id)
	fmt.Println("create task service is :", service)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "备忘录创建失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "备忘录创建成功",
	}
}

// ShowTask 展示备忘录tid的一条记录
func (service *ShowTaskService) ShowTask(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "备忘录查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "备忘录查询成功",
		Data:   serializer.BuildTask(task),
	}
}

func (service *ShowAllTaskService) ShowAll(uid uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if service.Limit == 0 {
		service.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", uid).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

func (service *DeleteService) Delete(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "备忘录不存在，请检查备忘录ID",
		}
	}
	err = model.DB.Delete(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据库发生错误，删除失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "删除备忘录成功",
	}
}

func (service *UpdateService) Update(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.Model(model.Task{}).Where("id = ?", tid).First(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "更新备忘录失败，未找到对应的备忘录",
		}
	}
	task.Status = service.Status
	task.Content = service.Content
	task.Title = service.Title
	err = model.DB.Save(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "更新备忘录失败，数据库出现错误",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "更新备忘录成功",
	}
}
