package services

import (
	"DiaryGo/models"
	"DiaryGo/pkg/repository"
)

type DoctorsActionService struct {
	repo repository.DoctorsAction
}

func NewDoctorsActionService(repo repository.DoctorsAction) *DoctorsActionService {
	return &DoctorsActionService{repo: repo}
}
func (s *DoctorsActionService) UpdateDoctors(doctors *models.Doctor, doctorsId int) *models.Doctor {
	return s.repo.UpdateDoctors(doctors, doctorsId)
}
func (s *DoctorsActionService) DeleteDoctors(id int) int {
	return s.repo.DeleteDoctors(id)
}

func (s *DoctorsActionService) GetAllDoctors() []*models.Doctor {
	return s.repo.GetAllDoctors()
}
