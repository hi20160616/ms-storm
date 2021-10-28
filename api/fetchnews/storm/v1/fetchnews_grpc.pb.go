// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	v1 "github.com/hi20160616/fetchnews-api/proto/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FetchClient is the client API for Fetch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchClient interface {
	ListArticles(ctx context.Context, in *v1.ListArticlesRequest, opts ...grpc.CallOption) (*v1.ListArticlesResponse, error)
	GetArticle(ctx context.Context, in *v1.GetArticleRequest, opts ...grpc.CallOption) (*v1.Article, error)
	SearchArticles(ctx context.Context, in *v1.SearchArticlesRequest, opts ...grpc.CallOption) (*v1.SearchArticlesResponse, error)
}

type fetchClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchClient(cc grpc.ClientConnInterface) FetchClient {
	return &fetchClient{cc}
}

func (c *fetchClient) ListArticles(ctx context.Context, in *v1.ListArticlesRequest, opts ...grpc.CallOption) (*v1.ListArticlesResponse, error) {
	out := new(v1.ListArticlesResponse)
	err := c.cc.Invoke(ctx, "/fetchnews.storm.v1.Fetch/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) GetArticle(ctx context.Context, in *v1.GetArticleRequest, opts ...grpc.CallOption) (*v1.Article, error) {
	out := new(v1.Article)
	err := c.cc.Invoke(ctx, "/fetchnews.storm.v1.Fetch/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) SearchArticles(ctx context.Context, in *v1.SearchArticlesRequest, opts ...grpc.CallOption) (*v1.SearchArticlesResponse, error) {
	out := new(v1.SearchArticlesResponse)
	err := c.cc.Invoke(ctx, "/fetchnews.storm.v1.Fetch/SearchArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchServer is the server API for Fetch service.
// All implementations must embed UnimplementedFetchServer
// for forward compatibility
type FetchServer interface {
	ListArticles(context.Context, *v1.ListArticlesRequest) (*v1.ListArticlesResponse, error)
	GetArticle(context.Context, *v1.GetArticleRequest) (*v1.Article, error)
	SearchArticles(context.Context, *v1.SearchArticlesRequest) (*v1.SearchArticlesResponse, error)
	mustEmbedUnimplementedFetchServer()
}

// UnimplementedFetchServer must be embedded to have forward compatible implementations.
type UnimplementedFetchServer struct {
}

func (UnimplementedFetchServer) ListArticles(context.Context, *v1.ListArticlesRequest) (*v1.ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedFetchServer) GetArticle(context.Context, *v1.GetArticleRequest) (*v1.Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedFetchServer) SearchArticles(context.Context, *v1.SearchArticlesRequest) (*v1.SearchArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArticles not implemented")
}
func (UnimplementedFetchServer) mustEmbedUnimplementedFetchServer() {}

// UnsafeFetchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchServer will
// result in compilation errors.
type UnsafeFetchServer interface {
	mustEmbedUnimplementedFetchServer()
}

func RegisterFetchServer(s grpc.ServiceRegistrar, srv FetchServer) {
	s.RegisterService(&Fetch_ServiceDesc, srv)
}

func _Fetch_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetchnews.storm.v1.Fetch/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).ListArticles(ctx, req.(*v1.ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetchnews.storm.v1.Fetch/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).GetArticle(ctx, req.(*v1.GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_SearchArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SearchArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).SearchArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetchnews.storm.v1.Fetch/SearchArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).SearchArticles(ctx, req.(*v1.SearchArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fetch_ServiceDesc is the grpc.ServiceDesc for Fetch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fetch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fetchnews.storm.v1.Fetch",
	HandlerType: (*FetchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArticles",
			Handler:    _Fetch_ListArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Fetch_GetArticle_Handler,
		},
		{
			MethodName: "SearchArticles",
			Handler:    _Fetch_SearchArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/fetchnews/storm/v1/fetchnews.proto",
}
