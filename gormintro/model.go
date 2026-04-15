package gormintro

import (
  	"gorm.io/gorm"
)

type Student struct {
	// ID			int			`gorm:"primaryKey"`	Эти все поля включены в поле gorm.Model(id, create, update, delete) с аказанием и имени поля и его типа
	// CreatedAt	time.Time						Эти все поля включены в поле gorm.Model(id, create, update, delete) с аказанием и имени поля и его типа
	// UpdatedAt	time.Time						Эти все поля включены в поле gorm.Model(id, create, update, delete) с аказанием и имени поля и его типа
	// DeletedAt	gorm.DeletedAt `gorm:"index"`	Эти все поля включены в поле gorm.Model(id, create, update, delete) с аказанием и имени поля и его типа
	gorm.Model
	Name		string
	Age			int
}

type Group struct {
	gorm.Model
	Name	string
}