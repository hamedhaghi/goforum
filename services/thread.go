package services

import (
	"github.com/hamedhaghi/goforum/models"
	"gorm.io/gorm"
)

type ThreadServiceInterface interface {
	FindAll() (*[]models.Thread, error)
	FindByID(id uint) (*models.Thread, error)
	Create(t *models.Thread) error
	Update(t *models.Thread) error
	Delete(id uint) error
}

type threadService struct {
	storage *gorm.DB
}

func NewThreadService(storage *gorm.DB) ThreadServiceInterface {
	return &threadService{
		storage: storage,
	}
}

func (ts *threadService) FindAll() (*[]models.Thread, error) {
	var threads *[]models.Thread
	r := ts.storage.Find(&threads)
	return threads, r.Error
}

func (ts *threadService) FindByID(id uint) (*models.Thread, error) {
	var thread *models.Thread
	r := ts.storage.First(&thread, id)
	return thread, r.Error
}

func (ts *threadService) Create(t *models.Thread) error {
	r := ts.storage.Create(&t)
	return r.Error
}

func (ts *threadService) Update(t *models.Thread) error {
	r := ts.storage.Save(&t)
	return r.Error
}

func (ts *threadService) Delete(id uint) error {
	r := ts.storage.Delete(&models.Thread{}, id)
	return r.Error
}
