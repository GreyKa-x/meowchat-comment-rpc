// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package commentrpc

import (
	"context"
	pb2 "github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                              = pb2.Comment
	CreateCommentRequest                 = pb2.CreateCommentRequest
	CreateCommentResponse                = pb2.CreateCommentResponse
	DeleteCommentByIdRequest             = pb2.DeleteCommentByIdRequest
	DeleteCommentByIdResponse            = pb2.DeleteCommentByIdResponse
	ListCommentByAuthorIdAndTypeRequest  = pb2.ListCommentByAuthorIdAndTypeRequest
	ListCommentByAuthorIdAndTypeResponse = pb2.ListCommentByAuthorIdAndTypeResponse
	ListCommentByParentRequest           = pb2.ListCommentByParentRequest
	ListCommentByParentResponse          = pb2.ListCommentByParentResponse
	ListCommentByReplyToAndTypeRequest   = pb2.ListCommentByReplyToAndTypeRequest
	ListCommentByReplyToAndTypeResponse  = pb2.ListCommentByReplyToAndTypeResponse
	RetrieveCommentByIdRequest           = pb2.RetrieveCommentByIdRequest
	RetrieveCommentByIdResponse          = pb2.RetrieveCommentByIdResponse
	UpdateCommentRequest                 = pb2.UpdateCommentRequest
	UpdateCommentResponse                = pb2.UpdateCommentResponse

	CommentRpc interface {
		// 创建
		CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
		// 修改
		UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
		// 删除
		DeleteComment(ctx context.Context, in *DeleteCommentByIdRequest, opts ...grpc.CallOption) (*DeleteCommentByIdResponse, error)
		// 根据 parentId 查找
		ListCommentByParent(ctx context.Context, in *ListCommentByParentRequest, opts ...grpc.CallOption) (*ListCommentByParentResponse, error)
		// 根据 id 查找
		RetrieveCommentById(ctx context.Context, in *RetrieveCommentByIdRequest, opts ...grpc.CallOption) (*RetrieveCommentByIdResponse, error)
		// 根据 authorId & type 查找
		ListCommentByAuthorIdAndType(ctx context.Context, in *ListCommentByAuthorIdAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByAuthorIdAndTypeResponse, error)
		// 根据 replyTo & type 查找
		ListCommentByReplyToAndType(ctx context.Context, in *ListCommentByReplyToAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByReplyToAndTypeResponse, error)
	}

	defaultCommentRpc struct {
		cli zrpc.Client
	}
)

func NewCommentRpc(cli zrpc.Client) CommentRpc {
	return &defaultCommentRpc{
		cli: cli,
	}
}

// 创建
func (m *defaultCommentRpc) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

// 修改
func (m *defaultCommentRpc) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.UpdateComment(ctx, in, opts...)
}

// 删除
func (m *defaultCommentRpc) DeleteComment(ctx context.Context, in *DeleteCommentByIdRequest, opts ...grpc.CallOption) (*DeleteCommentByIdResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

// 根据 parentId 查找
func (m *defaultCommentRpc) ListCommentByParent(ctx context.Context, in *ListCommentByParentRequest, opts ...grpc.CallOption) (*ListCommentByParentResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.ListCommentByParent(ctx, in, opts...)
}

// 根据 id 查找
func (m *defaultCommentRpc) RetrieveCommentById(ctx context.Context, in *RetrieveCommentByIdRequest, opts ...grpc.CallOption) (*RetrieveCommentByIdResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.RetrieveCommentById(ctx, in, opts...)
}

// 根据 authorId & type 查找
func (m *defaultCommentRpc) ListCommentByAuthorIdAndType(ctx context.Context, in *ListCommentByAuthorIdAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByAuthorIdAndTypeResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.ListCommentByAuthorIdAndType(ctx, in, opts...)
}

// 根据 replyTo & type 查找
func (m *defaultCommentRpc) ListCommentByReplyToAndType(ctx context.Context, in *ListCommentByReplyToAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByReplyToAndTypeResponse, error) {
	client := pb2.NewCommentRpcClient(m.cli.Conn())
	return client.ListCommentByReplyToAndType(ctx, in, opts...)
}
