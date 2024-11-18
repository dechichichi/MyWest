package task

import "todolist/model"

//增
func Add(username string, password string, email string) (*model.User, error) {

}

//删
func Delete(username string, password string) error {

}

//改
func ModifyName(username string, password string, newname string) (*model.User, error) {

}

func ModifyPassword(username string, password string, newpassword string) (*model.User, error) {

}

//查
func Ask(username string) (*model.User, error) {

}
