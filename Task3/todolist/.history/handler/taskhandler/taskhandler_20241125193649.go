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

}

func CreateTask(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.Query("id")
	var data model.Data
	err := json.Unmarshal([]byte(c.Query("data")), &data)
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
	user.ID = c.Query("id")
	var data model.Data
	title := c.Query("title")
	data, err := task.FindItemByTitle(user.ID, title)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid data",
			"code":    http.StatusBadRequest,
		})
	}
	newstatues := c.Query("status")
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
	user.ID = c.Query("id")
	var title string
	title = c.Query("title")
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
		"data":    data,
	})
}

func GetTasksToDone(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.Query("id")
	way := c.Query("way")
	var data []model.Data
	if way == "done" {
		data = task.GetCompletedItemList(user.ID)
	} else if way == "todo" {
		data = task.GetUncompletedItemList(user.ID)
	} else if way == "all" {
		data = task.GetAllItemList(user.ID)
	} else {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "invalid way",
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

func GetTasksToKey(ctx context.Context, c *app.RequestContext) {

}
