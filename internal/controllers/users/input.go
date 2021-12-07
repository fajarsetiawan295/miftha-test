package users

import "test-agit/internal/model"

//CREATE
type CreateRequest struct {
	model.User
}
type UpdateRequest struct {
	Id uint `json:"id" validate:"required"`
	model.User
}
type DeleteRequest struct {
	model.User
}
type Filter struct {
	Id      uint  `json:"id"`
	Name    int64 `json:"name"`
	Address int64 `json:"address"`
}
