package repo

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
		return t.db.Create(&data).Error
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	err := t.db.Unscoped().Find(&teachers).Error
	return teachers, err
}

func (t TeacherRepo) Update(id uint, name string) error {
	return t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name).Error
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	if result := t.db.Where("id=?",id).Delete(&teacher);result.Error != nil{
		return fmt.Errorf("Error DELETE TEACHER")
	}
	return nil
}