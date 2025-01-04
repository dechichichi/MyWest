package taskhandler

import (
	"context"
	"encoding/json"
	"net/http"
	"todolist/model"
	"todolist/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func List(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.PostForm("id")
	data, err := task.GetAllItemList(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
		"data":    data,
	})
}

func CreateTask(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.PostForm("id")
	var data model.Data
	err := json.Unmarshal([]byte(c.PostForm("data")), &data)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
	}
	task.CreateItem(user.ID, &data)
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}

func UpdateTask(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.PostForm("id")
	var data model.Data
	title := c.PostForm("title")
	data, err := task.FindItemByTitle(user.ID, title)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
	}
	newstatues := c.PostForm("status")
	if len(newstatues) == 0 {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid status",
			"code":    http.StatusBadRequest,
		})
	}
	task.UpdateItem(user.ID, &data, newstatues)
}

func DeleteTask(ctx context.Context, c *app.RequestContext) {
	var user model.User
	var title string
	user.ID = c.PostForm("id")
	title = c.PostForm("title")
	data, err := task.FindItemByTitle(user.ID, title)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
	}
	task.DeleteItem(user.ID, &data)
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}

func GetTasksToDone(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.PostForm("id")
	way := c.PostForm("way")
	var data []model.Data
	var err error
	if way == "done" {
		data, err = task.GetCompletedItemList(user.ID)
	} else if way == "todo" {
		data, err = task.GetUncompletedItemList(user.ID)
	} else if way == "all" {
		data, err = task.GetAllItemList(user.ID)
	} else {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid way",
			"code":    http.StatusBadRequest,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
		"data":    data,
	})
}

func GetTasksByKey(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.PostForm("id")
	key := c.PostForm("key")
	if len(key) == 0 {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid key",
			"code":    http.StatusBadRequest,
		})
		return
	}
	data, err := task.FindItem(user.ID, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
		"data":    data,
	})
}
