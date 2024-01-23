package services

import (
	"github.com/hamedhaghi/goforum/models"
	"gorm.io/gorm"
)

type PostServiceInterface interface {
	FindAll() (*[]models.Post, error)
	FindByID(id uint) (*models.Post, error)
	Create(p *models.Post) error
	Update(p *models.Post) error
	Delete(id uint) error
}

type postService struct {
	storage *gorm.DB
}

func NewPostService(storage *gorm.DB) PostServiceInterface {
	return &postService{
		storage: storage,
	}
}

func (ps *postService) FindAll() (*[]models.Post, error) {
	var posts *[]models.Post
	r := ps.storage.Find(&posts)
	return posts, r.Error
}

func (ps *postService) FindByID(id uint) (*models.Post, error) {
	var post *models.Post
	r := ps.storage.First(&post, id)
	return post, r.Error
}

func (ps *postService) Create(p *models.Post) error {
	r := ps.storage.Create(&p)
	return r.Error
}

func (ps *postService) Update(p *models.Post) error {
	r := ps.storage.Save(&p)
	return r.Error
}

func (ps *postService) Delete(id uint) error {
	r := ps.storage.Delete(&models.Post{}, id)
	return r.Error
}
