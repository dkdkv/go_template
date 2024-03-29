// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// ImagesGet implements GET /images operation.
	//
	// GET /images
	ImagesGet(ctx context.Context) (ImagesGetRes, error)
	// ImagesIDDelete implements DELETE /images/{id} operation.
	//
	// DELETE /images/{id}
	ImagesIDDelete(ctx context.Context, params ImagesIDDeleteParams) (ImagesIDDeleteRes, error)
	// ImagesIDGet implements GET /images/{id} operation.
	//
	// GET /images/{id}
	ImagesIDGet(ctx context.Context, params ImagesIDGetParams) (ImagesIDGetRes, error)
	// ImagesIDPut implements PUT /images/{id} operation.
	//
	// PUT /images/{id}
	ImagesIDPut(ctx context.Context, req *CreateImage, params ImagesIDPutParams) (ImagesIDPutRes, error)
	// ImagesPost implements POST /images operation.
	//
	// POST /images
	ImagesPost(ctx context.Context, req *CreateImage) (ImagesPostRes, error)
	// PingGet implements GET /ping operation.
	//
	// GET /ping
	PingGet(ctx context.Context) (PingGetRes, error)
	// PostsGet implements GET /posts operation.
	//
	// GET /posts
	PostsGet(ctx context.Context) (PostsGetRes, error)
	// PostsIDDelete implements DELETE /posts/{id} operation.
	//
	// DELETE /posts/{id}
	PostsIDDelete(ctx context.Context, params PostsIDDeleteParams) (PostsIDDeleteRes, error)
	// PostsIDGet implements GET /posts/{id} operation.
	//
	// GET /posts/{id}
	PostsIDGet(ctx context.Context, params PostsIDGetParams) (PostsIDGetRes, error)
	// PostsIDPut implements PUT /posts/{id} operation.
	//
	// PUT /posts/{id}
	PostsIDPut(ctx context.Context, req *CreatePost, params PostsIDPutParams) (PostsIDPutRes, error)
	// PostsPost implements POST /posts operation.
	//
	// POST /posts
	PostsPost(ctx context.Context, req *CreatePost) (PostsPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
