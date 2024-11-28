package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	User      User   `json:"user"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetPostID(ctx context.Context, postId int64) ([]Comment, error) {
	query := `
	SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id FROM comments c
	JOIN users ON users.id = c.user_id
	WHERE c.post_id = $1
	ORDER BY c.created_at DESC;
	`

	rows, err := s.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.User.ID, &c.User.Username, &c.Content, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)

	}
	return comments, nil

}