package responses

import (
	"time"

	"github.com/hamedhaghi/goforum/models"
)

type CommentResponse struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	Content   string    `json:"content"`
	Votes     int       `json:"votes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CommentListResponse(comments []models.Comment) *[]CommentResponse {
	list := []CommentResponse{}
	for _, c := range comments {
		r := CommentSingleResponse(c)
		list = append(list, *r)
	}
	return &list
}

func CommentSingleResponse(c models.Comment) *CommentResponse {
	return &CommentResponse{
		ID:        c.ID,
		PostID:    c.PostID,
		Content:   c.Content,
		Votes:     c.Votes,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
