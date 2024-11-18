package user

import "todolist/task"

func Register(username, password, email string) error {
	// TODO: implement the registration logic
	task.Add(username, password, email)
	return nil
}
