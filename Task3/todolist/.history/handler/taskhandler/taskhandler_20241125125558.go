package taskhandler

import (
	"context"
	"todolist/model"

	"github.com/cloudwego/hertz/pkg/app"
)

func List(ctx context.Context, c *app.RequestContext) {

}

func CreateTask(ctx context.Context, c *app.RequestContext) {
	var task model.Data
	task.Title = c.Query("title")
	task.Content = c.Query("content")

}

func UpdateTask(ctx context.Context, c *app.RequestContext) {

}

func DeleteTask(ctx context.Context, c *app.RequestContext) {

}

func GetTask(ctx context.Context, c *app.RequestContext) {

}

func GetTasks(ctx context.Context, c *app.RequestContext) {

}
