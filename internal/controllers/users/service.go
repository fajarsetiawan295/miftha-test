package users

import (
	"errors"
	"test-agit/internal/model"
)

type Service interface {
	Create(payload *CreateRequest) (*model.UserModel, error)
	Update(payload *UpdateRequest) (*UpdateRequest, error)
	Detail(e *Filter) (*model.UserModel, error)
	List() (*[]model.UserModel, error)
	Destroy(e *Filter) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(payload *CreateRequest) (*model.UserModel, error) {

	data, err := s.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (s *service) Update(payload *UpdateRequest) (*UpdateRequest, error) {

	_, err := s.repository.Find(&Filter{
		Id: payload.Id,
	})
	if err != nil {
		return nil, errors.New("data user tidak di temukan")
	}

	data, err := s.repository.Update(payload)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) Destroy(e *Filter) error {
	// cek data
	_, err := s.repository.Find(e)
	if err != nil {
		return errors.New("data user tidak di temukan")
	}
	// delete data
	errDelete := s.repository.Destroy(e)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
func (s *service) Detail(e *Filter) (*model.UserModel, error) {
	// cek data
	data, err := s.repository.Find(e)
	if err != nil {
		return nil, errors.New("data user tidak di temukan")
	}

	return data, nil
}
func (s *service) List() (*[]model.UserModel, error) {
	// delete data
	data, errDelete := s.repository.List()
	if errDelete != nil {
		return nil, errDelete
	}
	return data, nil
}
