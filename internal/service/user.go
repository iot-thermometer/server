package service

type User interface {
}

type user struct {
}

func newUserService() User {
	return &user{}
}
