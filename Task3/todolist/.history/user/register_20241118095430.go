package user

import "todolist/task"

func Register(username, password string) error {
	// TODO: implement the registration logic
	task.Add(username, password)
	return nil
}
