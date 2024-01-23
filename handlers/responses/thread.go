package responses

import (
	"time"

	"github.com/hamedhaghi/goforum/models"
)

type ThreadResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ThreadListResponse(threads []models.Thread) *[]ThreadResponse {
	list := []ThreadResponse{}
	for _, t := range threads {
		r := ThreadSingleResponse(t)
		list = append(list, *r)
	}
	return &list
}

func ThreadSingleResponse(t models.Thread) *ThreadResponse {
	return &ThreadResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
