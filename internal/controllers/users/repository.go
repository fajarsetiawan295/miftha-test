package users

import (
	"test-agit/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(e *CreateRequest) (*model.UserModel, error)
	List() (*[]model.UserModel, error)
	Update(e *UpdateRequest) (*UpdateRequest, error)
	Find(e *Filter) (*model.UserModel, error)
	Destroy(e *Filter) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(e *CreateRequest) (*model.UserModel, error) {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var data model.UserModel
	data.User = e.User
	errMaster := tx.Debug().Create(&data).Error
	if errMaster != nil {
		tx.Rollback()
		return nil, errMaster
	}

	tx.Commit()
	return &data, nil
}

func (r *repository) List() (*[]model.UserModel, error) {

	var data []model.UserModel
	errMaster := r.db.Debug().Find(&data).Error
	if errMaster != nil {
		return nil, errMaster
	}

	return &data, nil
}

func (r *repository) Update(e *UpdateRequest) (*UpdateRequest, error) {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	errMaster := tx.Debug().Table("users").Where("id = ?", e.Id).Updates(model.User{
		Name:    e.Name,
		Address: e.Address,
	}).Error
	if errMaster != nil {
		tx.Rollback()
		return nil, errMaster
	}

	tx.Commit()
	return e, nil
}

func (r *repository) Find(e *Filter) (*model.UserModel, error) {

	var data model.UserModel
	err := r.db.Debug().Table("users").Where("id = ?", e.Id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) Destroy(e *Filter) error {
	err := r.db.Debug().Unscoped().Table("users").Where("id = ?", e.Id).Delete(model.UserModel{}).Error
	if err != nil {
		return err
	}
	return nil
}
