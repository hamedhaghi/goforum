package responses

import (
	"time"

	"github.com/hamedhaghi/goforum/models"
)

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	ThreadID  uint      `json:"thread_id"`
	Content   string    `json:"content"`
	Votes     int       `json:"votes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PostListResponse(posts []models.Post) *[]PostResponse {
	list := []PostResponse{}
	for _, p := range posts {
		r := PostSingleResponse(p)
		list = append(list, *r)
	}
	return &list
}

func PostSingleResponse(p models.Post) *PostResponse {
	return &PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		ThreadID:  p.ThreadID,
		Content:   p.Content,
		Votes:     p.Votes,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
