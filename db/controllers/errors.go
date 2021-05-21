package controllers

import "errors"

var (
	ErrParameterNotFound = errors.New("controller: missing parameter")
	ErrEmptyUsers        = errors.New("controller: empty users table")
)

var (
	StrInsertUserFail = "Fail when try to insert a new user on db"
	StrGetUsersFail   = "Fail when try to get users from db"
	StrGetUserFail    = "Fail when try to get user by ID from db"
)
