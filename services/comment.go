package services

import (
	"github.com/hamedhaghi/goforum/models"
	"gorm.io/gorm"
)

type CommentServiceInterface interface {
	FindAll() (*[]models.Comment, error)
	FindByID(id uint) (*models.Comment, error)
	Create(c *models.Comment) error
	Update(c *models.Comment) error
	Delete(id uint) error
}

type commentService struct {
	storage *gorm.DB
}

func NewCommentService(storage *gorm.DB) CommentServiceInterface {
	return &commentService{
		storage: storage,
	}
}

func (cs *commentService) FindAll() (*[]models.Comment, error) {
	var comments *[]models.Comment
	r := cs.storage.Find(&comments)
	return comments, r.Error
}

func (cs *commentService) FindByID(id uint) (*models.Comment, error) {
	var comment *models.Comment
	r := cs.storage.First(&comment, id)
	return comment, r.Error
}

func (cs *commentService) Create(c *models.Comment) error {
	r := cs.storage.Create(&c)
	return r.Error
}

func (cs *commentService) Update(c *models.Comment) error {
	r := cs.storage.Save(&c)
	return r.Error
}

func (cs *commentService) Delete(id uint) error {
	r := cs.storage.Delete(&models.Comment{}, id)
	return r.Error
}
