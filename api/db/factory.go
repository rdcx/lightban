package db

import (
	"lightban/api/model"
	"lightban/util"

	"github.com/brianvoe/gofakeit/v6"
)

func (db *DB) CreateRandomUser() *model.User {
	name := util.RandomStringWithPrefix("user-")

	user, err := db.CreateUser(name, gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 32))

	if err != nil {
		panic(err)
	}

	return user
}

func (db *DB) CreateRandomProject() *model.Project {
	name := util.RandomStringWithPrefix("project-")

	proj := &model.Project{
		User: db.CreateRandomUser(),
		Name: name,
	}

	err := db.CreateProject(proj)

	if err != nil {
		panic(err)
	}

	return proj
}

func (db *DB) CreateRandomList() *model.List {
	name := util.RandomStringWithPrefix("list-")

	list := &model.List{
		Name:    name,
		Project: db.CreateRandomProject(),
	}

	err := db.CreateList(list)

	if err != nil {
		panic(err)
	}

	return list
}

func (db *DB) CreateRandomTask() *model.Task {
	name := util.RandomStringWithPrefix("task-")

	task := &model.Task{
		Name: name,
		List: db.CreateRandomList(),
	}

	err := db.CreateTask(task)

	if err != nil {
		panic(err)
	}

	return task
}
