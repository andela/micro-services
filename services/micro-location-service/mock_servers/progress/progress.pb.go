// Code generated by protoc-gen-go.
// source: progress/progress.proto
// DO NOT EDIT!

/*
Package progress is a generated protocol buffer package.

It is generated from these files:
	progress/progress.proto

It has these top-level messages:
	Empty
	ProgressData
	Query
	Progress
	ProgressList
	PeriodicProgress
*/
package progress

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProgressData struct {
	UserId   string  `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	LevelId  string  `protobuf:"bytes,2,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	CohortId string  `protobuf:"bytes,3,opt,name=cohort_id,json=cohortId" json:"cohort_id,omitempty"`
	Percent  float64 `protobuf:"fixed64,4,opt,name=percent" json:"percent,omitempty"`
}

func (m *ProgressData) Reset()                    { *m = ProgressData{} }
func (m *ProgressData) String() string            { return proto.CompactTextString(m) }
func (*ProgressData) ProtoMessage()               {}
func (*ProgressData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Query struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LevelId  string `protobuf:"bytes,2,opt,name=levelId" json:"levelId,omitempty"`
	CohortId string `protobuf:"bytes,3,opt,name=cohortId" json:"cohortId,omitempty"`
	Start    string `protobuf:"bytes,4,opt,name=start" json:"start,omitempty"`
	End      string `protobuf:"bytes,5,opt,name=end" json:"end,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Progress struct {
	Progress float64 `protobuf:"fixed64,1,opt,name=progress" json:"progress,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ProgressList struct {
	Values []float64 `protobuf:"fixed64,1,rep,name=values" json:"values,omitempty"`
}

func (m *ProgressList) Reset()                    { *m = ProgressList{} }
func (m *ProgressList) String() string            { return proto.CompactTextString(m) }
func (*ProgressList) ProtoMessage()               {}
func (*ProgressList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PeriodicProgress struct {
	Progress map[string]*ProgressList `protobuf:"bytes,1,rep,name=progress" json:"progress,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PeriodicProgress) Reset()                    { *m = PeriodicProgress{} }
func (m *PeriodicProgress) String() string            { return proto.CompactTextString(m) }
func (*PeriodicProgress) ProtoMessage()               {}
func (*PeriodicProgress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PeriodicProgress) GetProgress() map[string]*ProgressList {
	if m != nil {
		return m.Progress
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "progress.Empty")
	proto.RegisterType((*ProgressData)(nil), "progress.ProgressData")
	proto.RegisterType((*Query)(nil), "progress.Query")
	proto.RegisterType((*Progress)(nil), "progress.Progress")
	proto.RegisterType((*ProgressList)(nil), "progress.ProgressList")
	proto.RegisterType((*PeriodicProgress)(nil), "progress.PeriodicProgress")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Micro service

type MicroClient interface {
	ListAll(ctx context.Context, in *Query, opts ...grpc.CallOption) (*PeriodicProgress, error)
	Create(ctx context.Context, in *ProgressData, opts ...grpc.CallOption) (*Empty, error)
	GetUserProgress(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Progress, error)
	GetCohortAverage(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Progress, error)
	GetLevelAverage(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Progress, error)
}

type microClient struct {
	cc *grpc.ClientConn
}

func NewMicroClient(cc *grpc.ClientConn) MicroClient {
	return &microClient{cc}
}

func (c *microClient) ListAll(ctx context.Context, in *Query, opts ...grpc.CallOption) (*PeriodicProgress, error) {
	out := new(PeriodicProgress)
	err := grpc.Invoke(ctx, "/progress.micro/ListAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microClient) Create(ctx context.Context, in *ProgressData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/progress.micro/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microClient) GetUserProgress(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Progress, error) {
	out := new(Progress)
	err := grpc.Invoke(ctx, "/progress.micro/GetUserProgress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microClient) GetCohortAverage(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Progress, error) {
	out := new(Progress)
	err := grpc.Invoke(ctx, "/progress.micro/GetCohortAverage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microClient) GetLevelAverage(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Progress, error) {
	out := new(Progress)
	err := grpc.Invoke(ctx, "/progress.micro/GetLevelAverage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Micro service

type MicroServer interface {
	ListAll(context.Context, *Query) (*PeriodicProgress, error)
	Create(context.Context, *ProgressData) (*Empty, error)
	GetUserProgress(context.Context, *Query) (*Progress, error)
	GetCohortAverage(context.Context, *Query) (*Progress, error)
	GetLevelAverage(context.Context, *Query) (*Progress, error)
}

func RegisterMicroServer(s *grpc.Server, srv MicroServer) {
	s.RegisterService(&_Micro_serviceDesc, srv)
}

func _Micro_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/progress.micro/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServer).ListAll(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Micro_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgressData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/progress.micro/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServer).Create(ctx, req.(*ProgressData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Micro_GetUserProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServer).GetUserProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/progress.micro/GetUserProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServer).GetUserProgress(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Micro_GetCohortAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServer).GetCohortAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/progress.micro/GetCohortAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServer).GetCohortAverage(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Micro_GetLevelAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServer).GetLevelAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/progress.micro/GetLevelAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServer).GetLevelAverage(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _Micro_serviceDesc = grpc.ServiceDesc{
	ServiceName: "progress.micro",
	HandlerType: (*MicroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAll",
			Handler:    _Micro_ListAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Micro_Create_Handler,
		},
		{
			MethodName: "GetUserProgress",
			Handler:    _Micro_GetUserProgress_Handler,
		},
		{
			MethodName: "GetCohortAverage",
			Handler:    _Micro_GetCohortAverage_Handler,
		},
		{
			MethodName: "GetLevelAverage",
			Handler:    _Micro_GetLevelAverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("progress/progress.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x76, 0xfd, 0xc1, 0xd0, 0x16, 0xb4, 0xaa, 0xc0, 0x75, 0x2f, 0x95, 0x0f, 0x88, 0x43,
	0x45, 0x25, 0xb8, 0xd0, 0xde, 0x10, 0xa0, 0xaa, 0x12, 0x07, 0xea, 0x2a, 0xe7, 0xc8, 0xb1, 0x47,
	0xc4, 0x8a, 0xc1, 0x68, 0xbd, 0x46, 0xf2, 0x0f, 0xca, 0x31, 0xff, 0x31, 0xbb, 0xeb, 0x2f, 0x3e,
	0x12, 0x29, 0xb9, 0xcd, 0xec, 0x9b, 0x79, 0x6f, 0xfc, 0x66, 0x0c, 0x83, 0x03, 0x4d, 0xb6, 0x14,
	0xd3, 0xf4, 0x67, 0x15, 0x8c, 0x79, 0xc0, 0x12, 0x62, 0x55, 0xb9, 0x6b, 0x82, 0xbe, 0xda, 0x1d,
	0x58, 0xee, 0xe6, 0xf0, 0x71, 0x53, 0x3e, 0x2e, 0x7d, 0xe6, 0x93, 0x01, 0x98, 0x59, 0x8a, 0xf4,
	0x36, 0x0a, 0x6d, 0xe5, 0xbb, 0x32, 0x6a, 0x7b, 0x86, 0x48, 0xff, 0x86, 0xe4, 0x2b, 0x58, 0x31,
	0x1e, 0x31, 0x16, 0x88, 0x2a, 0x11, 0x53, 0xe6, 0x1c, 0xfa, 0x06, 0xed, 0x20, 0xb9, 0x4f, 0x28,
	0x13, 0x98, 0x26, 0x31, 0xab, 0x78, 0xe0, 0xa0, 0x0d, 0xe6, 0x01, 0x69, 0x80, 0x7b, 0x66, 0x7f,
	0xe0, 0x90, 0xe2, 0x55, 0xa9, 0x9b, 0x81, 0xfe, 0x2f, 0x43, 0x9a, 0x93, 0xcf, 0xa0, 0xd6, 0x72,
	0x3c, 0x12, 0x2d, 0x25, 0xf5, 0xa5, 0x92, 0x03, 0x35, 0xf1, 0x95, 0xd0, 0x17, 0xd0, 0x53, 0xe6,
	0xd3, 0x42, 0xa6, 0xed, 0x15, 0x09, 0xe9, 0x81, 0x86, 0xfb, 0xd0, 0xd6, 0xe5, 0x9b, 0x08, 0xdd,
	0x21, 0x58, 0xd5, 0x17, 0x0b, 0xbe, 0xca, 0x12, 0xa9, 0xaf, 0x78, 0x8d, 0x45, 0xc3, 0xc6, 0x99,
	0x75, 0x94, 0x32, 0xd2, 0x07, 0xe3, 0xe8, 0xc7, 0x19, 0x8a, 0x4a, 0x8d, 0x57, 0x96, 0x99, 0xfb,
	0xa4, 0x40, 0x6f, 0x83, 0x34, 0x4a, 0xc2, 0x28, 0xa8, 0x89, 0x97, 0x67, 0xc4, 0xda, 0xa8, 0x33,
	0x19, 0x8d, 0xeb, 0x65, 0x5c, 0x56, 0x8f, 0xab, 0x60, 0xb5, 0x67, 0x34, 0x6f, 0x46, 0x70, 0xfe,
	0xc3, 0xa7, 0x33, 0x48, 0x7c, 0xcd, 0x03, 0xe6, 0xa5, 0x55, 0x22, 0x24, 0x3f, 0x40, 0x97, 0x73,
	0x48, 0xa7, 0x3a, 0x93, 0xfe, 0x89, 0xca, 0xc9, 0xf0, 0x5e, 0x51, 0xf4, 0x5b, 0x9d, 0x29, 0x93,
	0x47, 0x15, 0xf4, 0x5d, 0x14, 0xd0, 0x84, 0xcc, 0xc0, 0x14, 0xe0, 0x3c, 0x8e, 0x49, 0xb7, 0xe9,
	0x93, 0x3b, 0x71, 0x9c, 0xd7, 0xc7, 0x75, 0x5b, 0x64, 0x0a, 0xc6, 0x82, 0xa2, 0xcf, 0x90, 0xbc,
	0x20, 0x28, 0xee, 0xc8, 0x39, 0x21, 0x2c, 0x0e, 0xad, 0xc5, 0xe5, 0xba, 0x7f, 0x90, 0xdd, 0xf0,
	0x73, 0xaa, 0x6d, 0xba, 0x92, 0x25, 0xd7, 0x74, 0xbc, 0xf3, 0x17, 0xf4, 0x78, 0xe7, 0x42, 0x6e,
	0x7a, 0x7e, 0x44, 0xea, 0x6f, 0xf1, 0xad, 0xad, 0x85, 0xe8, 0x5a, 0xdc, 0xcf, 0xfb, 0x3a, 0xef,
	0x0c, 0xf9, 0xcf, 0x4c, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x9a, 0x8c, 0x24, 0x4e, 0x03,
	0x00, 0x00,
}
