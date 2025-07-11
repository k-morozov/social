// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Code HTTP status code
	Code int `json:"code"`

	// Error Error message
	Error string `json:"error"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	// Email User email address
	Email openapi_types.Email `json:"email"`

	// Password User password
	Password string `json:"password"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	// Token JWT authentication token
	Token string `json:"token"`
	User  User   `json:"user"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	// Email User email address
	Email openapi_types.Email `json:"email"`

	// Name User full name
	Name string `json:"name"`

	// Password User password
	Password string `json:"password"`
}

// RegisterResponse defines model for RegisterResponse.
type RegisterResponse struct {
	// Message Success message
	Message string `json:"message"`
	User    User   `json:"user"`
}

// User defines model for User.
type User struct {
	// CreatedAt User creation timestamp
	CreatedAt time.Time `json:"created_at"`

	// Email User email address
	Email openapi_types.Email `json:"email"`

	// Id Unique user identifier
	Id int64 `json:"id"`

	// Name User full name
	Name string `json:"name"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = LoginRequest

// PostUserRegisterJSONRequestBody defines body for PostUserRegister for application/json ContentType.
type PostUserRegisterJSONRequestBody = RegisterRequest
