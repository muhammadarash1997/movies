// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movies.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Movie struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Rating               float64  `protobuf:"fixed64,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	CreatedAt            int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{0}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Movie) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Movie) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Movie) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Movie) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GetMovieListResponse struct {
	Movies               []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMovieListResponse) Reset()         { *m = GetMovieListResponse{} }
func (m *GetMovieListResponse) String() string { return proto.CompactTextString(m) }
func (*GetMovieListResponse) ProtoMessage()    {}
func (*GetMovieListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{1}
}

func (m *GetMovieListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMovieListResponse.Unmarshal(m, b)
}
func (m *GetMovieListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMovieListResponse.Marshal(b, m, deterministic)
}
func (m *GetMovieListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMovieListResponse.Merge(m, src)
}
func (m *GetMovieListResponse) XXX_Size() int {
	return xxx_messageInfo_GetMovieListResponse.Size(m)
}
func (m *GetMovieListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMovieListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMovieListResponse proto.InternalMessageInfo

func (m *GetMovieListResponse) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type Request struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_546d72ade507cae9, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Movie)(nil), "grpc.Movie")
	proto.RegisterType((*GetMovieListResponse)(nil), "grpc.GetMovieListResponse")
	proto.RegisterType((*Request)(nil), "grpc.Request")
}

func init() { proto.RegisterFile("movies.proto", fileDescriptor_546d72ade507cae9) }

var fileDescriptor_546d72ade507cae9 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x6d, 0xd2, 0x36, 0xfd, 0x7a, 0xd3, 0x76, 0x31, 0x94, 0x92, 0x2f, 0x22, 0x84, 0xb8, 0xc9,
	0x2a, 0xc5, 0x2a, 0x6e, 0x5c, 0x15, 0x2a, 0x6e, 0x74, 0x13, 0x71, 0x2d, 0x69, 0xe6, 0x1a, 0x06,
	0xd2, 0x4e, 0xcc, 0xdc, 0x16, 0x7c, 0x21, 0x5f, 0xc2, 0x97, 0x93, 0xcc, 0x24, 0x10, 0x8a, 0x5d,
	0xb8, 0xbc, 0xe7, 0x67, 0xce, 0xc9, 0x09, 0x4c, 0x76, 0xf2, 0x28, 0x50, 0xc5, 0x65, 0x25, 0x49,
	0xb2, 0x41, 0x5e, 0x95, 0x99, 0x7f, 0x91, 0x4b, 0x99, 0x17, 0xb8, 0xd4, 0xd8, 0xf6, 0xf0, 0xbe,
	0xc4, 0x5d, 0x49, 0x9f, 0x46, 0x12, 0x7e, 0x5b, 0x30, 0x7c, 0xae, 0x3d, 0x6c, 0x06, 0xb6, 0xe0,
	0x9e, 0x15, 0x58, 0x51, 0x3f, 0xb1, 0x05, 0x67, 0x73, 0x18, 0x92, 0xa0, 0x02, 0x3d, 0x3b, 0xb0,
	0xa2, 0x71, 0x62, 0x0e, 0x16, 0x80, 0xcb, 0x51, 0x65, 0x95, 0x28, 0x49, 0xc8, 0xbd, 0xd7, 0xd7,
	0x5c, 0x17, 0x62, 0x0b, 0x70, 0xaa, 0x94, 0xc4, 0x3e, 0xf7, 0x06, 0x81, 0x15, 0x59, 0x49, 0x73,
	0xd5, 0xef, 0x89, 0x5d, 0x9a, 0xa3, 0x37, 0x34, 0xef, 0xe9, 0x83, 0x5d, 0x02, 0x64, 0x15, 0xa6,
	0x84, 0xfc, 0x2d, 0x25, 0xcf, 0xd1, 0xe9, 0xe3, 0x06, 0x59, 0x53, 0x4d, 0x1f, 0x4a, 0xde, 0xd2,
	0x23, 0x43, 0x37, 0xc8, 0x9a, 0xc2, 0x7b, 0x98, 0x3f, 0x22, 0xe9, 0xfe, 0x4f, 0x42, 0x51, 0x82,
	0xaa, 0x94, 0x7b, 0x85, 0xec, 0x0a, 0x1c, 0x33, 0x84, 0x67, 0x05, 0xfd, 0xc8, 0x5d, 0xb9, 0x71,
	0xbd, 0x44, 0xac, 0x85, 0x49, 0x43, 0x85, 0xff, 0x61, 0x94, 0xe0, 0xc7, 0x01, 0x15, 0x9d, 0x7e,
	0xfb, 0xea, 0xcb, 0x86, 0xa9, 0x16, 0xab, 0x17, 0xac, 0x8e, 0x22, 0x43, 0xb6, 0x81, 0x49, 0x37,
	0x89, 0x2d, 0x62, 0xb3, 0x6a, 0xdc, 0xae, 0x1a, 0x3f, 0xd4, 0xab, 0xfa, 0xbe, 0x49, 0xfa, 0xad,
	0x55, 0xd8, 0x63, 0x31, 0xcc, 0x5a, 0x66, 0x83, 0x94, 0x8a, 0x82, 0x4d, 0x8d, 0xbe, 0x29, 0xe2,
	0x77, 0x8b, 0x86, 0x3d, 0x76, 0x0d, 0xff, 0xd6, 0x9c, 0x9b, 0xff, 0xd3, 0xa5, 0xfc, 0x33, 0xf1,
	0x61, 0x8f, 0xdd, 0x82, 0xfb, 0xaa, 0xf7, 0xf9, 0x93, 0xeb, 0x0e, 0xdc, 0x0d, 0x16, 0xd8, 0xba,
	0x4e, 0x5a, 0x9d, 0xf5, 0x6d, 0x1d, 0x8d, 0xdc, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x25,
	0x3c, 0x79, 0x78, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoviesServiceClient is the client API for MoviesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoviesServiceClient interface {
	GetMovieList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetMovieListResponse, error)
	GetMovieDetail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Movie, error)
	AddMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteMovie(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
}

type moviesServiceClient struct {
	cc *grpc.ClientConn
}

func NewMoviesServiceClient(cc *grpc.ClientConn) MoviesServiceClient {
	return &moviesServiceClient{cc}
}

func (c *moviesServiceClient) GetMovieList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetMovieListResponse, error) {
	out := new(GetMovieListResponse)
	err := c.cc.Invoke(ctx, "/grpc.MoviesService/GetMovieList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) GetMovieDetail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/grpc.MoviesService/GetMovieDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) AddMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.MoviesService/AddMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) UpdateMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.MoviesService/UpdateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) DeleteMovie(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.MoviesService/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServiceServer is the server API for MoviesService service.
type MoviesServiceServer interface {
	GetMovieList(context.Context, *empty.Empty) (*GetMovieListResponse, error)
	GetMovieDetail(context.Context, *Request) (*Movie, error)
	AddMovie(context.Context, *Movie) (*empty.Empty, error)
	UpdateMovie(context.Context, *Movie) (*empty.Empty, error)
	DeleteMovie(context.Context, *Request) (*empty.Empty, error)
}

// UnimplementedMoviesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMoviesServiceServer struct {
}

func (*UnimplementedMoviesServiceServer) GetMovieList(ctx context.Context, req *empty.Empty) (*GetMovieListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieList not implemented")
}
func (*UnimplementedMoviesServiceServer) GetMovieDetail(ctx context.Context, req *Request) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieDetail not implemented")
}
func (*UnimplementedMoviesServiceServer) AddMovie(ctx context.Context, req *Movie) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (*UnimplementedMoviesServiceServer) UpdateMovie(ctx context.Context, req *Movie) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (*UnimplementedMoviesServiceServer) DeleteMovie(ctx context.Context, req *Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}

func RegisterMoviesServiceServer(s *grpc.Server, srv MoviesServiceServer) {
	s.RegisterService(&_MoviesService_serviceDesc, srv)
}

func _MoviesService_GetMovieList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).GetMovieList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MoviesService/GetMovieList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).GetMovieList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_GetMovieDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).GetMovieDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MoviesService/GetMovieDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).GetMovieDetail(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MoviesService/AddMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).AddMovie(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MoviesService/UpdateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).UpdateMovie(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MoviesService/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).DeleteMovie(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoviesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MoviesService",
	HandlerType: (*MoviesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieList",
			Handler:    _MoviesService_GetMovieList_Handler,
		},
		{
			MethodName: "GetMovieDetail",
			Handler:    _MoviesService_GetMovieDetail_Handler,
		},
		{
			MethodName: "AddMovie",
			Handler:    _MoviesService_AddMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MoviesService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MoviesService_DeleteMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movies.proto",
}