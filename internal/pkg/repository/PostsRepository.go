package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go_template/internal/main/model"
	"go_template/internal/oas"
)

type PostsRepository struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewPostsRepository(db *pgxpool.Pool) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}

func (r *PostsRepository) GetPosts(ctx context.Context) ([]oas.Post, error) {
	var posts []oas.Post
	rows, err := r.db.Query(ctx, "SELECT * FROM posts")
	if err != nil {
		r.log.Error("Error getting posts", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post oas.Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.CoverURL)
		if err != nil {
			r.log.Error("Error scanning posts", zap.Error(err))
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostsRepository) GetPost(ctx context.Context, id int) (*oas.Post, error) {
	var post oas.Post
	err := r.db.QueryRow(ctx, "SELECT * FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content, &post.CoverURL)
	if err != nil {
		r.log.Error("Error getting post", zap.Error(err))
		return nil, err
	}
	return &post, nil
}

func (r *PostsRepository) CreatePost(ctx context.Context, post *oas.CreatePost) (*oas.Post, error) {
	var posts model.PostsModel

	//	build struct to pass to the database
	posts.Title = post.Title
	if post.Content.IsSet() {
		posts.Content.String = post.Content.Value
	}
	if post.CoverURL.IsSet() {
		posts.CoverURL.String = post.CoverURL.Value
	}

	//	prepare the query
	query := `INSERT INTO posts (title, content, cover_url) VALUES ($1, $2, $3)`

	//	execute the query
	err := r.db.QueryRow(ctx, query, posts.Title, posts.Content, posts.CoverURL).Scan(&posts.ID)
	if err != nil {
		r.log.Error("Error creating post", zap.Error(err))
		return nil, err
	}

	return &oas.Post{
		ID:       posts.ID,
		Title:    posts.Title,
		Content:  oas.OptString{Value: posts.Content.String, Set: posts.Content.Valid},
		CoverURL: oas.OptString{Value: posts.CoverURL.String, Set: posts.CoverURL.Valid},
	}, nil
}

func (r *PostsRepository) UpdatePost(ctx context.Context, post *oas.CreatePost, id int) (*oas.Post, error) {
	var posts model.PostsModel

	//	build struct to pass to the database
	posts.ID = id
	posts.Title = post.Title
	if post.Content.IsSet() {
		posts.Content.String = post.Content.Value
	}
	if post.CoverURL.IsSet() {
		posts.CoverURL.String = post.CoverURL.Value
	}

	//	prepare the query
	query := `UPDATE posts SET title = $1, content = $2, cover_url = $3 WHERE id = $4`

	//	execute the query
	_, err := r.db.Exec(ctx, query, posts.Title, posts.Content, posts.CoverURL, posts.ID)
	if err != nil {
		r.log.Error("Error updating post", zap.Error(err))
		return nil, err
	}

	return &oas.Post{
		ID:       posts.ID,
		Title:    posts.Title,
		Content:  oas.OptString{Value: posts.Content.String, Set: posts.Content.Valid},
		CoverURL: oas.OptString{Value: posts.CoverURL.String, Set: posts.CoverURL.Valid},
	}, nil
}

func (r *PostsRepository) DeletePost(ctx context.Context, id int) (*int, error) {
	//	prepare the query
	query := `DELETE FROM posts WHERE id = $1`

	//	execute the query
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Error deleting post", zap.Error(err))
		return nil, err
	}
	return &id, nil
}
