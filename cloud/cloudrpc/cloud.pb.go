// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloud.proto

package cloudrpc

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

// JobSpec is the input for the RunJob service.
type JobSpec struct {
	// Version is the required InMAP version.
	Version string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	// Name is a user-specified name for the job.
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// Cmd is the command to be run, e.g., [inmap run steady]
	Cmd []string `protobuf:"bytes,3,rep,name=Cmd,proto3" json:"Cmd,omitempty"`
	// Args are the command line arguments, e.g., [--Layers, 2, --steady, true]
	Args []string `protobuf:"bytes,4,rep,name=Args,proto3" json:"Args,omitempty"`
	// MemoryGB specifies the required gigabytes of RAM memory for the
	// simulation.
	MemoryGB int32 `protobuf:"varint,5,opt,name=MemoryGB,proto3" json:"MemoryGB,omitempty"`
	// FileData holds the contents of any local files referred to by Args
	FileData             map[string][]byte `protobuf:"bytes,7,rep,name=FileData,proto3" json:"FileData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JobSpec) Reset()         { *m = JobSpec{} }
func (m *JobSpec) String() string { return proto.CompactTextString(m) }
func (*JobSpec) ProtoMessage()    {}
func (*JobSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_2d86b3326abf2f76, []int{0}
}
func (m *JobSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobSpec.Unmarshal(m, b)
}
func (m *JobSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobSpec.Marshal(b, m, deterministic)
}
func (dst *JobSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobSpec.Merge(dst, src)
}
func (m *JobSpec) XXX_Size() int {
	return xxx_messageInfo_JobSpec.Size(m)
}
func (m *JobSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_JobSpec.DiscardUnknown(m)
}

var xxx_messageInfo_JobSpec proto.InternalMessageInfo

func (m *JobSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JobSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobSpec) GetCmd() []string {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *JobSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *JobSpec) GetMemoryGB() int32 {
	if m != nil {
		return m.MemoryGB
	}
	return 0
}

func (m *JobSpec) GetFileData() map[string][]byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

type JobStatus struct {
	// Status holds the current status of the job.
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobStatus) Reset()         { *m = JobStatus{} }
func (m *JobStatus) String() string { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()    {}
func (*JobStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_2d86b3326abf2f76, []int{1}
}
func (m *JobStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatus.Unmarshal(m, b)
}
func (m *JobStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatus.Marshal(b, m, deterministic)
}
func (dst *JobStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatus.Merge(dst, src)
}
func (m *JobStatus) XXX_Size() int {
	return xxx_messageInfo_JobStatus.Size(m)
}
func (m *JobStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatus.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatus proto.InternalMessageInfo

func (m *JobStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type JobOutput struct {
	// Files holds the contents of each output file.
	Files                map[string][]byte `protobuf:"bytes,1,rep,name=Files,proto3" json:"Files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JobOutput) Reset()         { *m = JobOutput{} }
func (m *JobOutput) String() string { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()    {}
func (*JobOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_2d86b3326abf2f76, []int{2}
}
func (m *JobOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobOutput.Unmarshal(m, b)
}
func (m *JobOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobOutput.Marshal(b, m, deterministic)
}
func (dst *JobOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobOutput.Merge(dst, src)
}
func (m *JobOutput) XXX_Size() int {
	return xxx_messageInfo_JobOutput.Size(m)
}
func (m *JobOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_JobOutput.DiscardUnknown(m)
}

var xxx_messageInfo_JobOutput proto.InternalMessageInfo

func (m *JobOutput) GetFiles() map[string][]byte {
	if m != nil {
		return m.Files
	}
	return nil
}

type JobName struct {
	// Version is the required InMAP version.
	Version string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	// Name is a user-specified name for the job.
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobName) Reset()         { *m = JobName{} }
func (m *JobName) String() string { return proto.CompactTextString(m) }
func (*JobName) ProtoMessage()    {}
func (*JobName) Descriptor() ([]byte, []int) {
	return fileDescriptor_cloud_2d86b3326abf2f76, []int{3}
}
func (m *JobName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobName.Unmarshal(m, b)
}
func (m *JobName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobName.Marshal(b, m, deterministic)
}
func (dst *JobName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobName.Merge(dst, src)
}
func (m *JobName) XXX_Size() int {
	return xxx_messageInfo_JobName.Size(m)
}
func (m *JobName) XXX_DiscardUnknown() {
	xxx_messageInfo_JobName.DiscardUnknown(m)
}

var xxx_messageInfo_JobName proto.InternalMessageInfo

func (m *JobName) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JobName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*JobSpec)(nil), "cloudrpc.JobSpec")
	proto.RegisterMapType((map[string][]byte)(nil), "cloudrpc.JobSpec.FileDataEntry")
	proto.RegisterType((*JobStatus)(nil), "cloudrpc.JobStatus")
	proto.RegisterType((*JobOutput)(nil), "cloudrpc.JobOutput")
	proto.RegisterMapType((map[string][]byte)(nil), "cloudrpc.JobOutput.FilesEntry")
	proto.RegisterType((*JobName)(nil), "cloudrpc.JobName")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CloudRPCClient is the client API for CloudRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudRPCClient interface {
	// RunJob performs an InMAP simulation and returns the paths to the
	// output file(s).
	RunJob(ctx context.Context, in *JobSpec, opts ...grpc.CallOption) (*JobStatus, error)
	// OutputAddresses returns status and the addresses the output file(s) of the
	// requested simulation name.
	Status(ctx context.Context, in *JobName, opts ...grpc.CallOption) (*JobStatus, error)
	// Output returns the output file(s) of the
	// requested simulation name.
	Output(ctx context.Context, in *JobName, opts ...grpc.CallOption) (*JobOutput, error)
	// Delete deletes the specified simulation.
	Delete(ctx context.Context, in *JobName, opts ...grpc.CallOption) (*JobName, error)
}

type cloudRPCClient struct {
	cc *grpc.ClientConn
}

func NewCloudRPCClient(cc *grpc.ClientConn) CloudRPCClient {
	return &cloudRPCClient{cc}
}

func (c *cloudRPCClient) RunJob(ctx context.Context, in *JobSpec, opts ...grpc.CallOption) (*JobStatus, error) {
	out := new(JobStatus)
	err := c.cc.Invoke(ctx, "/cloudrpc.CloudRPC/RunJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRPCClient) Status(ctx context.Context, in *JobName, opts ...grpc.CallOption) (*JobStatus, error) {
	out := new(JobStatus)
	err := c.cc.Invoke(ctx, "/cloudrpc.CloudRPC/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRPCClient) Output(ctx context.Context, in *JobName, opts ...grpc.CallOption) (*JobOutput, error) {
	out := new(JobOutput)
	err := c.cc.Invoke(ctx, "/cloudrpc.CloudRPC/Output", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRPCClient) Delete(ctx context.Context, in *JobName, opts ...grpc.CallOption) (*JobName, error) {
	out := new(JobName)
	err := c.cc.Invoke(ctx, "/cloudrpc.CloudRPC/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudRPCServer is the server API for CloudRPC service.
type CloudRPCServer interface {
	// RunJob performs an InMAP simulation and returns the paths to the
	// output file(s).
	RunJob(context.Context, *JobSpec) (*JobStatus, error)
	// OutputAddresses returns status and the addresses the output file(s) of the
	// requested simulation name.
	Status(context.Context, *JobName) (*JobStatus, error)
	// Output returns the output file(s) of the
	// requested simulation name.
	Output(context.Context, *JobName) (*JobOutput, error)
	// Delete deletes the specified simulation.
	Delete(context.Context, *JobName) (*JobName, error)
}

func RegisterCloudRPCServer(s *grpc.Server, srv CloudRPCServer) {
	s.RegisterService(&_CloudRPC_serviceDesc, srv)
}

func _CloudRPC_RunJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRPCServer).RunJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudrpc.CloudRPC/RunJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRPCServer).RunJob(ctx, req.(*JobSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRPC_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRPCServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudrpc.CloudRPC/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRPCServer).Status(ctx, req.(*JobName))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRPC_Output_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRPCServer).Output(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudrpc.CloudRPC/Output",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRPCServer).Output(ctx, req.(*JobName))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRPC_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRPCServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudrpc.CloudRPC/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRPCServer).Delete(ctx, req.(*JobName))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudrpc.CloudRPC",
	HandlerType: (*CloudRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunJob",
			Handler:    _CloudRPC_RunJob_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _CloudRPC_Status_Handler,
		},
		{
			MethodName: "Output",
			Handler:    _CloudRPC_Output_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudRPC_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud.proto",
}

func init() { proto.RegisterFile("cloud.proto", fileDescriptor_cloud_2d86b3326abf2f76) }

var fileDescriptor_cloud_2d86b3326abf2f76 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x9b, 0xa6, 0x49, 0xd3, 0xdb, 0xff, 0x1f, 0xf4, 0x2a, 0x32, 0x64, 0xa1, 0x21, 0x6e,
	0xb2, 0x0a, 0x52, 0x05, 0x8b, 0x5d, 0x69, 0xab, 0x42, 0xc1, 0x0f, 0x46, 0x70, 0x9f, 0xb6, 0x83,
	0x14, 0xd3, 0x4e, 0x48, 0x26, 0x42, 0xf1, 0x45, 0x7d, 0x0f, 0x5f, 0x40, 0xe6, 0x23, 0x95, 0x50,
	0x51, 0xba, 0xbb, 0xe7, 0xe4, 0x9c, 0xc9, 0x2f, 0x77, 0x02, 0xdd, 0x69, 0xca, 0xcb, 0x59, 0x9c,
	0xe5, 0x5c, 0x70, 0xf4, 0x94, 0xc8, 0xb3, 0x69, 0xf8, 0x69, 0x41, 0x7b, 0xcc, 0x27, 0x4f, 0x19,
	0x9b, 0x22, 0x81, 0xf6, 0x33, 0xcb, 0x8b, 0x39, 0x5f, 0x12, 0x2b, 0xb0, 0xa2, 0x0e, 0xad, 0x24,
	0x22, 0xb4, 0xee, 0x93, 0x05, 0x23, 0x4d, 0x65, 0xab, 0x19, 0x77, 0xc0, 0x1e, 0x2e, 0x66, 0xc4,
	0x0e, 0xec, 0xa8, 0x43, 0xe5, 0x28, 0x53, 0x97, 0xf9, 0x4b, 0x41, 0x5a, 0xca, 0x52, 0x33, 0xfa,
	0xe0, 0xdd, 0xb1, 0x05, 0xcf, 0x57, 0xb7, 0x57, 0xc4, 0x09, 0xac, 0xc8, 0xa1, 0x6b, 0x8d, 0x03,
	0xf0, 0x6e, 0xe6, 0x29, 0x1b, 0x25, 0x22, 0x21, 0xed, 0xc0, 0x8e, 0xba, 0xbd, 0xa3, 0xb8, 0x02,
	0x8b, 0x0d, 0x54, 0x5c, 0x25, 0xae, 0x97, 0x22, 0x5f, 0xd1, 0x75, 0xc1, 0x1f, 0xc0, 0xff, 0xda,
	0x23, 0xc9, 0xf3, 0xca, 0x56, 0x86, 0x5c, 0x8e, 0xb8, 0x0f, 0xce, 0x5b, 0x92, 0x96, 0x1a, 0xfb,
	0x1f, 0xd5, 0xe2, 0xa2, 0xd9, 0xb7, 0xc2, 0x63, 0xe8, 0xc8, 0xf3, 0x45, 0x22, 0xca, 0x02, 0x0f,
	0xc0, 0xd5, 0x93, 0xe9, 0x1a, 0x15, 0xbe, 0xab, 0xd0, 0x43, 0x29, 0xb2, 0x52, 0xe0, 0x19, 0x38,
	0xf2, 0x75, 0x32, 0x23, 0x41, 0x0f, 0x6b, 0xa0, 0x3a, 0xa3, 0x50, 0x0b, 0xcd, 0xa9, 0xc3, 0x7e,
	0x1f, 0xe0, 0xdb, 0xdc, 0x8a, 0xf0, 0x5c, 0x5d, 0x8b, 0x5a, 0xf4, 0x56, 0xd7, 0xd2, 0xfb, 0xb0,
	0xc0, 0x1b, 0x4a, 0x36, 0xfa, 0x38, 0xc4, 0x1e, 0xb8, 0xb4, 0x5c, 0x8e, 0xf9, 0x04, 0x77, 0x37,
	0x36, 0xeb, 0xef, 0xd5, 0x2d, 0xfd, 0xd1, 0x0d, 0xd9, 0x31, 0x8b, 0xa9, 0x77, 0xe4, 0xe9, 0xbf,
	0x74, 0xcc, 0x9e, 0xfe, 0xec, 0xe8, 0x5c, 0xd8, 0xc0, 0x13, 0x70, 0x47, 0x2c, 0x65, 0x82, 0xfd,
	0xd4, 0xd9, 0xb4, 0xc2, 0xc6, 0xc4, 0x55, 0x3f, 0xef, 0xe9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6a, 0xbc, 0x1d, 0xa9, 0xcb, 0x02, 0x00, 0x00,
}
