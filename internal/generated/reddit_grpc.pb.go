// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: internal/generated/reddit.proto

package redditpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RedditService_RegisterUser_FullMethodName      = "/reddit.RedditService/RegisterUser"
	RedditService_CreateSubreddit_FullMethodName   = "/reddit.RedditService/CreateSubreddit"
	RedditService_JoinSubreddit_FullMethodName     = "/reddit.RedditService/JoinSubreddit"
	RedditService_CreatePost_FullMethodName        = "/reddit.RedditService/CreatePost"
	RedditService_Repost_FullMethodName            = "/reddit.RedditService/Repost"
	RedditService_AddComment_FullMethodName        = "/reddit.RedditService/AddComment"
	RedditService_UpvotePost_FullMethodName        = "/reddit.RedditService/UpvotePost"
	RedditService_DownvotePost_FullMethodName      = "/reddit.RedditService/DownvotePost"
	RedditService_FetchFeed_FullMethodName         = "/reddit.RedditService/FetchFeed"
	RedditService_SendDirectMessage_FullMethodName = "/reddit.RedditService/SendDirectMessage"
	RedditService_GetDirectMessages_FullMethodName = "/reddit.RedditService/GetDirectMessages"
	RedditService_GetUserProfile_FullMethodName    = "/reddit.RedditService/GetUserProfile"
)

// RedditServiceClient is the client API for RedditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the RedditService
type RedditServiceClient interface {
	// User-related RPCs
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Subreddit-related RPCs
	CreateSubreddit(ctx context.Context, in *CreateSubredditRequest, opts ...grpc.CallOption) (*CreateSubredditResponse, error)
	JoinSubreddit(ctx context.Context, in *JoinSubredditRequest, opts ...grpc.CallOption) (*JoinSubredditResponse, error)
	// Post-related RPCs
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	Repost(ctx context.Context, in *RepostRequest, opts ...grpc.CallOption) (*RepostResponse, error)
	// Comment-related RPCs
	AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	// Voting RPCs
	UpvotePost(ctx context.Context, in *UpvoteRequest, opts ...grpc.CallOption) (*UpvoteResponse, error)
	DownvotePost(ctx context.Context, in *DownvoteRequest, opts ...grpc.CallOption) (*DownvoteResponse, error)
	// Feed RPCs
	FetchFeed(ctx context.Context, in *FetchFeedRequest, opts ...grpc.CallOption) (*FetchFeedResponse, error)
	// Direct Messaging RPC
	SendDirectMessage(ctx context.Context, in *DirectMessageRequest, opts ...grpc.CallOption) (*DirectMessageResponse, error)
	GetDirectMessages(ctx context.Context, in *GetDirectMessagesRequest, opts ...grpc.CallOption) (*GetDirectMessagesResponse, error)
	// User profile-related RPCs
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error)
}

type redditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRedditServiceClient(cc grpc.ClientConnInterface) RedditServiceClient {
	return &redditServiceClient{cc}
}

func (c *redditServiceClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, RedditService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) CreateSubreddit(ctx context.Context, in *CreateSubredditRequest, opts ...grpc.CallOption) (*CreateSubredditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSubredditResponse)
	err := c.cc.Invoke(ctx, RedditService_CreateSubreddit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) JoinSubreddit(ctx context.Context, in *JoinSubredditRequest, opts ...grpc.CallOption) (*JoinSubredditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinSubredditResponse)
	err := c.cc.Invoke(ctx, RedditService_JoinSubreddit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, RedditService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) Repost(ctx context.Context, in *RepostRequest, opts ...grpc.CallOption) (*RepostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepostResponse)
	err := c.cc.Invoke(ctx, RedditService_Repost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, RedditService_AddComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) UpvotePost(ctx context.Context, in *UpvoteRequest, opts ...grpc.CallOption) (*UpvoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpvoteResponse)
	err := c.cc.Invoke(ctx, RedditService_UpvotePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) DownvotePost(ctx context.Context, in *DownvoteRequest, opts ...grpc.CallOption) (*DownvoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownvoteResponse)
	err := c.cc.Invoke(ctx, RedditService_DownvotePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) FetchFeed(ctx context.Context, in *FetchFeedRequest, opts ...grpc.CallOption) (*FetchFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchFeedResponse)
	err := c.cc.Invoke(ctx, RedditService_FetchFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) SendDirectMessage(ctx context.Context, in *DirectMessageRequest, opts ...grpc.CallOption) (*DirectMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectMessageResponse)
	err := c.cc.Invoke(ctx, RedditService_SendDirectMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) GetDirectMessages(ctx context.Context, in *GetDirectMessagesRequest, opts ...grpc.CallOption) (*GetDirectMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDirectMessagesResponse)
	err := c.cc.Invoke(ctx, RedditService_GetDirectMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserProfileResponse)
	err := c.cc.Invoke(ctx, RedditService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedditServiceServer is the server API for RedditService service.
// All implementations must embed UnimplementedRedditServiceServer
// for forward compatibility.
//
// Define the RedditService
type RedditServiceServer interface {
	// User-related RPCs
	RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Subreddit-related RPCs
	CreateSubreddit(context.Context, *CreateSubredditRequest) (*CreateSubredditResponse, error)
	JoinSubreddit(context.Context, *JoinSubredditRequest) (*JoinSubredditResponse, error)
	// Post-related RPCs
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	Repost(context.Context, *RepostRequest) (*RepostResponse, error)
	// Comment-related RPCs
	AddComment(context.Context, *CommentRequest) (*CommentResponse, error)
	// Voting RPCs
	UpvotePost(context.Context, *UpvoteRequest) (*UpvoteResponse, error)
	DownvotePost(context.Context, *DownvoteRequest) (*DownvoteResponse, error)
	// Feed RPCs
	FetchFeed(context.Context, *FetchFeedRequest) (*FetchFeedResponse, error)
	// Direct Messaging RPC
	SendDirectMessage(context.Context, *DirectMessageRequest) (*DirectMessageResponse, error)
	GetDirectMessages(context.Context, *GetDirectMessagesRequest) (*GetDirectMessagesResponse, error)
	// User profile-related RPCs
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error)
	mustEmbedUnimplementedRedditServiceServer()
}

// UnimplementedRedditServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRedditServiceServer struct{}

func (UnimplementedRedditServiceServer) RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedRedditServiceServer) CreateSubreddit(context.Context, *CreateSubredditRequest) (*CreateSubredditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubreddit not implemented")
}
func (UnimplementedRedditServiceServer) JoinSubreddit(context.Context, *JoinSubredditRequest) (*JoinSubredditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSubreddit not implemented")
}
func (UnimplementedRedditServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedRedditServiceServer) Repost(context.Context, *RepostRequest) (*RepostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Repost not implemented")
}
func (UnimplementedRedditServiceServer) AddComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedRedditServiceServer) UpvotePost(context.Context, *UpvoteRequest) (*UpvoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvotePost not implemented")
}
func (UnimplementedRedditServiceServer) DownvotePost(context.Context, *DownvoteRequest) (*DownvoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvotePost not implemented")
}
func (UnimplementedRedditServiceServer) FetchFeed(context.Context, *FetchFeedRequest) (*FetchFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchFeed not implemented")
}
func (UnimplementedRedditServiceServer) SendDirectMessage(context.Context, *DirectMessageRequest) (*DirectMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDirectMessage not implemented")
}
func (UnimplementedRedditServiceServer) GetDirectMessages(context.Context, *GetDirectMessagesRequest) (*GetDirectMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirectMessages not implemented")
}
func (UnimplementedRedditServiceServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedRedditServiceServer) mustEmbedUnimplementedRedditServiceServer() {}
func (UnimplementedRedditServiceServer) testEmbeddedByValue()                       {}

// UnsafeRedditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedditServiceServer will
// result in compilation errors.
type UnsafeRedditServiceServer interface {
	mustEmbedUnimplementedRedditServiceServer()
}

func RegisterRedditServiceServer(s grpc.ServiceRegistrar, srv RedditServiceServer) {
	// If the following call pancis, it indicates UnimplementedRedditServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RedditService_ServiceDesc, srv)
}

func _RedditService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_CreateSubreddit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubredditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).CreateSubreddit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_CreateSubreddit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).CreateSubreddit(ctx, req.(*CreateSubredditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_JoinSubreddit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinSubredditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).JoinSubreddit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_JoinSubreddit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).JoinSubreddit(ctx, req.(*JoinSubredditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_Repost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).Repost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_Repost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).Repost(ctx, req.(*RepostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).AddComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_UpvotePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpvoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).UpvotePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_UpvotePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).UpvotePost(ctx, req.(*UpvoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_DownvotePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownvoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).DownvotePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_DownvotePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).DownvotePost(ctx, req.(*DownvoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_FetchFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).FetchFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_FetchFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).FetchFeed(ctx, req.(*FetchFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_SendDirectMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).SendDirectMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_SendDirectMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).SendDirectMessage(ctx, req.(*DirectMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_GetDirectMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirectMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).GetDirectMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_GetDirectMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).GetDirectMessages(ctx, req.(*GetDirectMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RedditService_ServiceDesc is the grpc.ServiceDesc for RedditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reddit.RedditService",
	HandlerType: (*RedditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _RedditService_RegisterUser_Handler,
		},
		{
			MethodName: "CreateSubreddit",
			Handler:    _RedditService_CreateSubreddit_Handler,
		},
		{
			MethodName: "JoinSubreddit",
			Handler:    _RedditService_JoinSubreddit_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _RedditService_CreatePost_Handler,
		},
		{
			MethodName: "Repost",
			Handler:    _RedditService_Repost_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _RedditService_AddComment_Handler,
		},
		{
			MethodName: "UpvotePost",
			Handler:    _RedditService_UpvotePost_Handler,
		},
		{
			MethodName: "DownvotePost",
			Handler:    _RedditService_DownvotePost_Handler,
		},
		{
			MethodName: "FetchFeed",
			Handler:    _RedditService_FetchFeed_Handler,
		},
		{
			MethodName: "SendDirectMessage",
			Handler:    _RedditService_SendDirectMessage_Handler,
		},
		{
			MethodName: "GetDirectMessages",
			Handler:    _RedditService_GetDirectMessages_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _RedditService_GetUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/generated/reddit.proto",
}