package taskhandler

import (
	"context"
	"encoding/json"
	"todolist/model"

	"github.com/cloudwego/hertz/pkg/app"
)

func List(ctx context.Context, c *app.RequestContext) {

}

func CreateTask(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.ID = c.Query("id")
	var data model.Data
	err := json.Unmarshal([]byte(c.Query("data")), &data)
	if err != nil {
		// 处理错误
	}
	user.Data = data

}

func UpdateTask(ctx context.Context, c *app.RequestContext) {

}

func DeleteTask(ctx context.Context, c *app.RequestContext) {

}

func GetTask(ctx context.Context, c *app.RequestContext) {

}

func GetTasks(ctx context.Context, c *app.RequestContext) {

}
