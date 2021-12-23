package pb

import (
	"context"
)

type UserGrpc struct {
}

func NewUserGrpc() *UserGrpc {
	return &UserGrpc{}
}

func (ug *UserGrpc) Create(c context.Context, r *RequestCreate) (*ResponseCreate, error) {
	return &ResponseCreate{
		Id:        2,
		Username:  "max",
		Password:  "qwerty12345",
		State:     1,
		CreatedAt: "2021-12-06T18:23:08.328741Z",
	}, nil
}

func (ug *UserGrpc) Get(c context.Context, r *Request) (*Response, error) {
	return &Response{
		ID:        2,
		Username:  "max",
		State:     1,
		CreatedAt: "2021-12-06T18:23:08.328741Z",
	}, nil
}

func (ug *UserGrpc) Edit(c context.Context, r *RequestEdit) (*ResponseCreate, error) {
	return &ResponseCreate{
		Id:        0,
		Username:  "",
		Password:  "",
		State:     0,
		CreatedAt: "",
	}, nil
}

func (ug *UserGrpc) mustEmbedUnimplementedUserServer() {}
