package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_REGISTER_USER = "failed register user"
	MESSAGE_FAILED_GET_USER      = "failed get user"
	MESSAGE_FAILED_LOGIN         = "failed login"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "success register user"
	MESSAGE_SUCCESS_GET_USER      = "success get user"
	MESSAGE_SUCCESS_LOGIN         = "success login"
)

var (
	ErrRoleNotAllowed        = errors.New("denied access for \"%v\" role")
	ErrGetUserById           = errors.New("failed to get user by id")
	ErrCredentialsNotMatched = errors.New("credentials not matched")
	ErrUsernameAlreadyExists = errors.New("username already exist")
	ErrCreateUser            = errors.New("failed to create user")
)

type (
	UserRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}
)
