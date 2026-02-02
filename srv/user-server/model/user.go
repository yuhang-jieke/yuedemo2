package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(30);comment:名称"`
	Age     int    `gorm:"type:int(11);comment:年龄"`
	Address string `gorm:"type:varchar(30);comment:地址"`
}

func (u *User) Register(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) FindName(db *gorm.DB, name string) error {
	return db.Where("name=?", name).Find(&u).Error
}

type TitleContent struct {
	gorm.Model
	Content string `gorm:"type:varchar(50);comment:内容"`
}

func (c *TitleContent) CreateTitle(db *gorm.DB) error {
	return db.Create(&c).Error
}

type ImgContent struct {
	gorm.Model
	Content string `gorm:"type:varchar(100);comment:内容"`
}
