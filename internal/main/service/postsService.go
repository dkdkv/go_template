package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go_template/internal/oas"
	"go_template/internal/pkg/repository"
)

type PostsService struct {
	db *pgxpool.Pool
}

func NewPostsService(db *pgxpool.Pool) *PostsService {
	return &PostsService{
		db: db,
	}
}

func (s *PostsService) GetPosts(ctx context.Context) ([]oas.Post, error) {
	r := repository.NewPostsRepository(s.db)
	posts, err := r.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostsService) GetPost(ctx context.Context, id int) (*oas.Post, error) {
	r := repository.NewPostsRepository(s.db)
	post, err := r.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostsService) CreatePost(ctx context.Context, post *oas.CreatePost) (*oas.Post, error) {
	r := repository.NewPostsRepository(s.db)
	createdPost, err := r.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func (s *PostsService) UpdatePost(ctx context.Context, post *oas.CreatePost, id int) (*oas.Post, error) {
	r := repository.NewPostsRepository(s.db)
	updatedPost, err := r.UpdatePost(ctx, post, id)
	if err != nil {
		return nil, err
	}
	return updatedPost, nil
}

func (s *PostsService) DeletePost(ctx context.Context, id int) (*int, error) {
	r := repository.NewPostsRepository(s.db)
	deletedPost, err := r.DeletePost(ctx, id)
	if err != nil {
		return nil, err
	}
	return deletedPost, nil
}
