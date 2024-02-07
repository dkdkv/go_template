package handler

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go_template/internal/main/service"
	"go_template/internal/oas"
)

type Handler struct {
	Params
	db *pgxpool.Pool
}

type Params struct {
	Logger *zap.Logger
}

func New(params Params, db *pgxpool.Pool) *Handler {
	return &Handler{
		Params: params,
		db:     db,
	}
}

func (h *Handler) PingGet(ctx context.Context) (oas.PingGetRes, error) {
	return &oas.PingGetOK{}, nil
}

func (h *Handler) ImagesGet(ctx context.Context) (oas.ImagesGetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) ImagesIDDelete(ctx context.Context, params oas.ImagesIDDeleteParams) (oas.ImagesIDDeleteRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) ImagesIDGet(ctx context.Context, params oas.ImagesIDGetParams) (oas.ImagesIDGetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) ImagesIDPut(ctx context.Context, req *oas.CreateImage, params oas.ImagesIDPutParams) (oas.ImagesIDPutRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) ImagesPost(ctx context.Context, req *oas.CreateImage) (oas.ImagesPostRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostsGet(ctx context.Context) (oas.PostsGetRes, error) {
	s := service.NewPostsService(h.db)
	posts, err := s.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	pagination := oas.Pagination{
		TotalItems:   10,
		ItemsPerPage: 10,
		CurrentPage:  10,
		TotalPages:   10,
	}
	return &oas.PostsGetOK{Data: posts, Pagination: pagination}, nil
}

func (h *Handler) PostsIDDelete(ctx context.Context, params oas.PostsIDDeleteParams) (oas.PostsIDDeleteRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostsIDGet(ctx context.Context, params oas.PostsIDGetParams) (oas.PostsIDGetRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostsIDPut(ctx context.Context, req *oas.CreatePost, params oas.PostsIDPutParams) (oas.PostsIDPutRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostsPost(ctx context.Context, req *oas.CreatePost) (oas.PostsPostRes, error) {
	//TODO implement me
	panic("implement me")
}
