// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"github.com/xh-polaris/meowchat-comment-rpc/internal/logic"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"
)

type CommentRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCommentRpcServer
}

func NewCommentRpcServer(svcCtx *svc.ServiceContext) *CommentRpcServer {
	return &CommentRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建
func (s *CommentRpcServer) CreateComment(ctx context.Context, in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

// 修改
func (s *CommentRpcServer) UpdateComment(ctx context.Context, in *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

// 删除
func (s *CommentRpcServer) DeleteComment(ctx context.Context, in *pb.DeleteCommentByIdRequest) (*pb.DeleteCommentByIdResponse, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

// 根据 parentId 查找
func (s *CommentRpcServer) ListCommentByParent(ctx context.Context, in *pb.ListCommentByParentRequest) (*pb.ListCommentByParentResponse, error) {
	l := logic.NewListCommentByParentLogic(ctx, s.svcCtx)
	return l.ListCommentByParent(in)
}

// 根据 id 查找
func (s *CommentRpcServer) RetrieveCommentById(ctx context.Context, in *pb.RetrieveCommentByIdRequest) (*pb.RetrieveCommentByIdResponse, error) {
	l := logic.NewRetrieveCommentByIdLogic(ctx, s.svcCtx)
	return l.RetrieveCommentById(in)
}

// 根据 authorId & type 查找
func (s *CommentRpcServer) ListCommentByAuthorIdAndType(ctx context.Context, in *pb.ListCommentByAuthorIdAndTypeRequest) (*pb.ListCommentByAuthorIdAndTypeResponse, error) {
	l := logic.NewListCommentByAuthorIdAndTypeLogic(ctx, s.svcCtx)
	return l.ListCommentByAuthorIdAndType(in)
}

// 根据 replyTo & type 查找
func (s *CommentRpcServer) ListCommentByReplyToAndType(ctx context.Context, in *pb.ListCommentByReplyToAndTypeRequest) (*pb.ListCommentByReplyToAndTypeResponse, error) {
	l := logic.NewListCommentByReplyToAndTypeLogic(ctx, s.svcCtx)
	return l.ListCommentByReplyToAndType(in)
}
