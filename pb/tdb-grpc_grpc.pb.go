// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotesServiceClient is the client API for NotesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotesServiceClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	ReadNote(ctx context.Context, in *ReadNoteRequest, opts ...grpc.CallOption) (*ReadNoteResponse, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error)
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error)
}

type notesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesServiceClient(cc grpc.ClientConnInterface) NotesServiceClient {
	return &notesServiceClient{cc}
}

func (c *notesServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, "/tdb.NotesService/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesServiceClient) ReadNote(ctx context.Context, in *ReadNoteRequest, opts ...grpc.CallOption) (*ReadNoteResponse, error) {
	out := new(ReadNoteResponse)
	err := c.cc.Invoke(ctx, "/tdb.NotesService/ReadNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesServiceClient) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error) {
	out := new(UpdateNoteResponse)
	err := c.cc.Invoke(ctx, "/tdb.NotesService/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesServiceClient) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error) {
	out := new(DeleteNoteResponse)
	err := c.cc.Invoke(ctx, "/tdb.NotesService/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServiceServer is the server API for NotesService service.
// All implementations must embed UnimplementedNotesServiceServer
// for forward compatibility
type NotesServiceServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	ReadNote(context.Context, *ReadNoteRequest) (*ReadNoteResponse, error)
	UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error)
	DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error)
	mustEmbedUnimplementedNotesServiceServer()
}

// UnimplementedNotesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServiceServer struct {
}

func (UnimplementedNotesServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNotesServiceServer) ReadNote(context.Context, *ReadNoteRequest) (*ReadNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNote not implemented")
}
func (UnimplementedNotesServiceServer) UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedNotesServiceServer) DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedNotesServiceServer) mustEmbedUnimplementedNotesServiceServer() {}

// UnsafeNotesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotesServiceServer will
// result in compilation errors.
type UnsafeNotesServiceServer interface {
	mustEmbedUnimplementedNotesServiceServer()
}

func RegisterNotesServiceServer(s grpc.ServiceRegistrar, srv NotesServiceServer) {
	s.RegisterService(&NotesService_ServiceDesc, srv)
}

func _NotesService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdb.NotesService/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotesService_ReadNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).ReadNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdb.NotesService/ReadNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).ReadNote(ctx, req.(*ReadNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotesService_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdb.NotesService/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotesService_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tdb.NotesService/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).DeleteNote(ctx, req.(*DeleteNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotesService_ServiceDesc is the grpc.ServiceDesc for NotesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tdb.NotesService",
	HandlerType: (*NotesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _NotesService_CreateNote_Handler,
		},
		{
			MethodName: "ReadNote",
			Handler:    _NotesService_ReadNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _NotesService_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _NotesService_DeleteNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/tdb-grpc.proto",
}
