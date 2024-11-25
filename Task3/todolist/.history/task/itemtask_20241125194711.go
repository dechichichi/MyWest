package task

import (
	"errors"
	"strconv"
	"todolist/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建任务
func CreateItem(Id string, Data *model.Data) error {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return errors.New("username already exists")
	}
	//查找是否存在相同标题的任务
	var item model.Data
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
func UpdateItem(Id string, Data *model.Data, statues string) error {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return errors.New("username already exists")
	}
	//更新任务
	statuesInt, err := strconv.Atoi(statues)
	if err != nil {
		return err // 处理转换错误
	}
	if err := db.Model(&Data).Updates(model.Data{Status: statuesInt}).Error; err != nil {
		return err
	}

}

// 根据关键词查找任务
func FindItem(Id string, key string) ([]model.Data, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return []model.Data{}, errors.New("username already exists")
	}
}

// 获取已完成的任务列表
func GetCompletedItemList(Id string) ([]model.Data, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return []model.Data{}, errors.New("username already exists")
	}
}

// 获取未完成的任务列表
func GetUncompletedItemList(Id string) ([]model.Data, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return []model.Data{}, errors.New("username already exists")
	}
}

// 获取全部任务列表
func GetAllItemList(Id string) ([]model.Data, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return []model.Data{}, errors.New("username already exists")
	}
}

// 删除任务
func DeleteItem(Id string, Data *model.Data) error {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return errors.New("username already exists")
	}
}

// 精准查找任务
func FindItemByTitle(Id string, title string) (model.Data, error) {
	db, err := gorm.Open(mysql.Open("root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return model.Data{}, err
	}
	//查找用户是否存在
	var user model.User
	if err := db.Where("id = ?", Id).First(&user).Error; err == nil {
		return model.Data{}, errors.New("username already exists")
	}
}
