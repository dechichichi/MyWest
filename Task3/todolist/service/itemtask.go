package service

import (
	"errors"
	"strconv"
	"todolist/database"
)

var user database.User

func FindUser(Id string) error {
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return errors.New("username does not exist")
	}
	return nil
}

// 创建任务
func CreateItem(Id string, Data *database.Data) error {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//查找是否存在相同标题的任务
	var item database.Data
	if err := db.Where("title = ?", Data.Title).First(&item).Error; err == nil {
		return errors.New("task already exists")
	}
	//创建任务
	if err := db.Create(&Data).Error; err != nil {
		return err
	}
	return nil
}

// 更新任务
func UpdateItem(Id string, Data *database.Data, statues string) error {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//更新任务
	statuesInt, err := strconv.Atoi(statues)
	if err != nil {
		return err // 处理转换错误
	}
	if err := db.Model(&Data).Updates(database.Data{Status: statuesInt}).Error; err != nil {
		return err
	}
	return nil
}

// 根据关键词查找任务
func FindItem(Id string, key string) ([]database.Data, error) {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//查找任务
	var items []database.Data
	if err := db.Where("title LIKE ?", "%"+key+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// 获取已完成的任务列表
func GetCompletedItemList(Id string) ([]database.Data, error) {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//查找任务
	var items []database.Data
	if err := db.Where("status = ?", 1).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// 获取未完成的任务列表
func GetUncompletedItemList(Id string) ([]database.Data, error) {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//查找任务
	var items []database.Data
	if err := db.Where("status = ?", 0).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// 获取全部任务列表
func GetAllItemList(Id string) ([]database.Data, error) {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//查找任务
	var items []database.Data
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// 删除任务
func DeleteItem(Id string, Data *database.Data) error {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//删除任务
	if err := db.Delete(&Data).Error; err != nil {
		return err
	}
	return nil
}

// 精准查找任务
func FindItemByTitle(Id string, title string) (database.Data, error) {
	//查找用户是否存在
	err := FindUser(Id)
	if err != nil {
		panic(err)
	}
	//查找任务
	var item database.Data
	if err := db.Where("title = ?", title).First(&item).Error; err != nil {
		return database.Data{}, err
	}
	return item, nil
}
