package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
		GetByID(ctx context.Context, id int64) (*Post, error)
		List(ctx context.Context) ([]*Post, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, post *Post) error
		GetUserFeed(ctx context.Context, id int64) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(ctx context.Context, user *User) error
		GetByID(ctx context.Context, id int64) (*User, error)
	}
	Comments interface {
		Create(ctx context.Context, comment *Comment) error
		GetPostID(ctx context.Context, postId int64) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerId, userId int64) error
		UnFollow(ctx context.Context, followerId, userId int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db: db},
		Users:     &UserStore{db: db},
		Comments:  &CommentStore{db: db},
		Followers: &FollowerStore{db: db},
	}
}
