package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("User exist!!")
	ErrorUserNotExist    = errors.New("User is not exist!!")
	ErrorInvalidPassword = errors.New("Password error!!")
	ErrorInvalidID       = errors.New("Invalid ID!!")
)
