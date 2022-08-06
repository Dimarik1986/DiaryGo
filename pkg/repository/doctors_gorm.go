package repository

import (
	"DiaryGo/models"
	"gorm.io/gorm"
)

type DoctorsActionDb struct {
	gorm gorm.DB
}

func NewDoctorsActionDb(gorm gorm.DB) *DoctorsActionDb {
	return &DoctorsActionDb{gorm: gorm}
}

func (d *DoctorsActionDb) UpdateDoctors(doctors *models.Doctor, docId int) *models.Doctor {
	d.gorm.Model(&models.Doctor{}).Where("id=?", docId).Updates(models.Doctor{
		Name:           doctors.Name,
		Surname:        doctors.Surname,
		Username:       doctors.Username,
		Password:       doctors.Password,
		Specialization: doctors.Specialization,
	})
	return doctors
}

func (d *DoctorsActionDb) DeleteDoctors(id int) int {
	d.gorm.Delete(&models.Doctor{}, id)
	return id
}

func (d *DoctorsActionDb) GetAllDoctors() []*models.Doctor {
	var users []*models.Doctor
	d.gorm.Raw("select * from doctors").Scan(&users)
	var docs []*models.Doctor
	docs = append(docs, users...)
	return docs
}
