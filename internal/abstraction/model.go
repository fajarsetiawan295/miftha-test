package abstraction

import (
	util "test-agit/helpers/utils"
	"time"

	"gorm.io/gorm"
)

type Id struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type IdAI struct {
	ID uint `gorm:"primaryKey;autoIncrement;" json:"id"`
}

type Entity struct {
	Id
	Filter
}

type EntityAI struct {
	IdAI
	Filter
}

type Filter struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = *util.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *util.DateTodayLocal()
	return
}

// checking function
func (m *EntityAI) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = *util.DateTodayLocal()
	return
}
func (m *EntityAI) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *util.DateTodayLocal()
	return
}
