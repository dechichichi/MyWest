package taskhandler

import (
	"context"
	"todolist/model"

	"github.com/cloudwego/hertz/pkg/app"
)

func List(c context.Context, ctx *app.RequestContext) {

}

func CreateTask(c context.Context, ctx *app.RequestContext) {
	var task model.Data
	task.Title = ctx.Query("title")
	task.Content = ctx.Query("content")

}

func UpdateTask(c context.Context, ctx *app.RequestContext) {

}

func DeleteTask(c context.Context, ctx *app.RequestContext) {

}

func GetTask(c context.Context, ctx *app.RequestContext) {

}

func GetTasks(c context.Context, ctx *app.RequestContext) {

}
