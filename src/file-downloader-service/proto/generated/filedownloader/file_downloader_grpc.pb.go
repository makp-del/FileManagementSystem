// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: file_downloader.proto

package filedownloader

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
	FileDownloaderService_DownloadFile_FullMethodName = "/filedownloader.FileDownloaderService/DownloadFile"
)

// FileDownloaderServiceClient is the client API for FileDownloaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// File Downloader Service definition
type FileDownloaderServiceClient interface {
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
}

type fileDownloaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileDownloaderServiceClient(cc grpc.ClientConnInterface) FileDownloaderServiceClient {
	return &fileDownloaderServiceClient{cc}
}

func (c *fileDownloaderServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, FileDownloaderService_DownloadFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileDownloaderServiceServer is the server API for FileDownloaderService service.
// All implementations must embed UnimplementedFileDownloaderServiceServer
// for forward compatibility.
//
// File Downloader Service definition
type FileDownloaderServiceServer interface {
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
	mustEmbedUnimplementedFileDownloaderServiceServer()
}

// UnimplementedFileDownloaderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileDownloaderServiceServer struct{}

func (UnimplementedFileDownloaderServiceServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileDownloaderServiceServer) mustEmbedUnimplementedFileDownloaderServiceServer() {}
func (UnimplementedFileDownloaderServiceServer) testEmbeddedByValue()                               {}

// UnsafeFileDownloaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileDownloaderServiceServer will
// result in compilation errors.
type UnsafeFileDownloaderServiceServer interface {
	mustEmbedUnimplementedFileDownloaderServiceServer()
}

func RegisterFileDownloaderServiceServer(s grpc.ServiceRegistrar, srv FileDownloaderServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileDownloaderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileDownloaderService_ServiceDesc, srv)
}

func _FileDownloaderService_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileDownloaderServiceServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileDownloaderService_DownloadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileDownloaderServiceServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileDownloaderService_ServiceDesc is the grpc.ServiceDesc for FileDownloaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileDownloaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filedownloader.FileDownloaderService",
	HandlerType: (*FileDownloaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadFile",
			Handler:    _FileDownloaderService_DownloadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file_downloader.proto",
}
