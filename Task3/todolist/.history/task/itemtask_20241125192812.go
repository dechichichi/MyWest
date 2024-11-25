package task

import "todolist/model"

//创建任务
func CreateItem(Id string, Data *model.Data) {

}

//更新任务
func UpdateItem(Id string, Data *model.Data, statues string) {

}

//根据关键词查找任务
func FindItem(Id string, key string) []model.Data {

}

//获取已完成的任务列表
func GetCompletedItemList(Id string) []model.Data {

}

//获取未完成的任务列表
func GetUncompletedItemList(Id string) []model.Data {

}

//获取全部任务列表
func GetAllItemList(Id string) []model.Data {

}

//删除任务
func DeleteItem(Id string, Data *model.Data) {

}

//精准查找任务
func FindItemByTitle(Id string, title string) []model.Data {

}
