// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_team.proto

package engine

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

type AgentTeam struct {
	Id                int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	DomainId          int64  `protobuf:"varint,2,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	Name              string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Description       string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Strategy          string `protobuf:"bytes,5,opt,name=strategy" json:"strategy,omitempty"`
	MaxNoAnswer       int32  `protobuf:"varint,6,opt,name=max_no_answer,json=maxNoAnswer" json:"max_no_answer,omitempty"`
	WrapUpTime        int32  `protobuf:"varint,7,opt,name=wrap_up_time,json=wrapUpTime" json:"wrap_up_time,omitempty"`
	RejectDelayTime   int32  `protobuf:"varint,8,opt,name=reject_delay_time,json=rejectDelayTime" json:"reject_delay_time,omitempty"`
	BusyDelayTime     int32  `protobuf:"varint,9,opt,name=busy_delay_time,json=busyDelayTime" json:"busy_delay_time,omitempty"`
	NoAnswerDelayTime int32  `protobuf:"varint,10,opt,name=no_answer_delay_time,json=noAnswerDelayTime" json:"no_answer_delay_time,omitempty"`
	CallTimeout       int32  `protobuf:"varint,11,opt,name=call_timeout,json=callTimeout" json:"call_timeout,omitempty"`
	UpdatedAt         int64  `protobuf:"varint,12,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *AgentTeam) Reset()                    { *m = AgentTeam{} }
func (m *AgentTeam) String() string            { return proto.CompactTextString(m) }
func (*AgentTeam) ProtoMessage()               {}
func (*AgentTeam) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *AgentTeam) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgentTeam) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *AgentTeam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentTeam) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AgentTeam) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *AgentTeam) GetMaxNoAnswer() int32 {
	if m != nil {
		return m.MaxNoAnswer
	}
	return 0
}

func (m *AgentTeam) GetWrapUpTime() int32 {
	if m != nil {
		return m.WrapUpTime
	}
	return 0
}

func (m *AgentTeam) GetRejectDelayTime() int32 {
	if m != nil {
		return m.RejectDelayTime
	}
	return 0
}

func (m *AgentTeam) GetBusyDelayTime() int32 {
	if m != nil {
		return m.BusyDelayTime
	}
	return 0
}

func (m *AgentTeam) GetNoAnswerDelayTime() int32 {
	if m != nil {
		return m.NoAnswerDelayTime
	}
	return 0
}

func (m *AgentTeam) GetCallTimeout() int32 {
	if m != nil {
		return m.CallTimeout
	}
	return 0
}

func (m *AgentTeam) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ListAgentTeam struct {
	Items []*AgentTeam `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListAgentTeam) Reset()                    { *m = ListAgentTeam{} }
func (m *ListAgentTeam) String() string            { return proto.CompactTextString(m) }
func (*ListAgentTeam) ProtoMessage()               {}
func (*ListAgentTeam) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *ListAgentTeam) GetItems() []*AgentTeam {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentTeam)(nil), "engine.AgentTeam")
	proto.RegisterType((*ListAgentTeam)(nil), "engine.ListAgentTeam")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AgentTeamApi service

type AgentTeamApiClient interface {
	// POST /call_center/agent_team
	Create(ctx context.Context, in *AgentTeam, opts ...grpc.CallOption) (*AgentTeam, error)
	// GET /call_center/agent_team
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentTeam, error)
	// GET /call_center/agent_team/:ID
	Get(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*AgentTeam, error)
	// PUT /call_center/agent_team/:ID
	Update(ctx context.Context, in *AgentTeam, opts ...grpc.CallOption) (*AgentTeam, error)
	// DELETE /call_center/agent_team/:ID
	Remove(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*AgentTeam, error)
}

type agentTeamApiClient struct {
	cc *grpc.ClientConn
}

func NewAgentTeamApiClient(cc *grpc.ClientConn) AgentTeamApiClient {
	return &agentTeamApiClient{cc}
}

func (c *agentTeamApiClient) Create(ctx context.Context, in *AgentTeam, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := grpc.Invoke(ctx, "/engine.AgentTeamApi/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamApiClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentTeam, error) {
	out := new(ListAgentTeam)
	err := grpc.Invoke(ctx, "/engine.AgentTeamApi/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamApiClient) Get(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := grpc.Invoke(ctx, "/engine.AgentTeamApi/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamApiClient) Update(ctx context.Context, in *AgentTeam, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := grpc.Invoke(ctx, "/engine.AgentTeamApi/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamApiClient) Remove(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := grpc.Invoke(ctx, "/engine.AgentTeamApi/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AgentTeamApi service

type AgentTeamApiServer interface {
	// POST /call_center/agent_team
	Create(context.Context, *AgentTeam) (*AgentTeam, error)
	// GET /call_center/agent_team
	List(context.Context, *ListRequest) (*ListAgentTeam, error)
	// GET /call_center/agent_team/:ID
	Get(context.Context, *ItemRequest) (*AgentTeam, error)
	// PUT /call_center/agent_team/:ID
	Update(context.Context, *AgentTeam) (*AgentTeam, error)
	// DELETE /call_center/agent_team/:ID
	Remove(context.Context, *ItemRequest) (*AgentTeam, error)
}

func RegisterAgentTeamApiServer(s *grpc.Server, srv AgentTeamApiServer) {
	s.RegisterService(&_AgentTeamApi_serviceDesc, srv)
}

func _AgentTeamApi_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentTeam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamApiServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamApi/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamApiServer).Create(ctx, req.(*AgentTeam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamApi_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamApiServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamApi/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamApiServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamApi_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamApiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamApi/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamApiServer).Get(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentTeam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamApiServer).Update(ctx, req.(*AgentTeam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamApi_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamApiServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamApi/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamApiServer).Remove(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentTeamApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.AgentTeamApi",
	HandlerType: (*AgentTeamApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AgentTeamApi_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AgentTeamApi_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AgentTeamApi_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AgentTeamApi_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _AgentTeamApi_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_team.proto",
}

func init() { proto.RegisterFile("agent_team.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xd3, 0x86, 0xe6, 0xa5, 0x65, 0xd4, 0x80, 0x64, 0x15, 0x21, 0x85, 0x1e, 0xa0,
	0xe2, 0xd0, 0x4d, 0x85, 0x03, 0xd7, 0x0a, 0x24, 0x34, 0x09, 0x71, 0x88, 0xba, 0x73, 0xe4, 0xd5,
	0x4f, 0x95, 0x51, 0x6d, 0x87, 0xf8, 0x85, 0xad, 0x1f, 0x82, 0x33, 0x5f, 0x17, 0xc5, 0xde, 0xd2,
	0x48, 0xe3, 0xb0, 0x5b, 0xfd, 0xfb, 0xff, 0xfe, 0x7a, 0xcf, 0x75, 0xe0, 0xb9, 0xd8, 0xa3, 0xa1,
	0x92, 0x50, 0xe8, 0x55, 0x55, 0x5b, 0xb2, 0x2c, 0x41, 0xb3, 0x57, 0x06, 0xe7, 0xd9, 0xce, 0x1a,
	0x47, 0x01, 0x2e, 0xfe, 0xc6, 0x90, 0x6e, 0x5a, 0x73, 0x8b, 0x42, 0xb3, 0x67, 0x30, 0x50, 0x92,
	0x47, 0x79, 0xb4, 0x8c, 0x8b, 0x81, 0x92, 0xec, 0x35, 0xa4, 0xd2, 0x6a, 0xa1, 0x4c, 0xa9, 0x24,
	0x1f, 0x78, 0x3c, 0x0e, 0xe0, 0x52, 0x32, 0x06, 0x43, 0x23, 0x34, 0xf2, 0x38, 0x8f, 0x96, 0x69,
	0xe1, 0x7f, 0xb3, 0x1c, 0x32, 0x89, 0x6e, 0x57, 0xab, 0x8a, 0x94, 0x35, 0x7c, 0xe8, 0xa3, 0x3e,
	0x62, 0x73, 0x18, 0x3b, 0xaa, 0x05, 0xe1, 0xfe, 0xc8, 0x47, 0x3e, 0xee, 0xce, 0x6c, 0x01, 0x53,
	0x2d, 0x6e, 0x4b, 0x63, 0x4b, 0x61, 0xdc, 0x0d, 0xd6, 0x3c, 0xc9, 0xa3, 0xe5, 0xa8, 0xc8, 0xb4,
	0xb8, 0xfd, 0x61, 0x37, 0x1e, 0xb1, 0x1c, 0x26, 0x37, 0xb5, 0xa8, 0xca, 0xa6, 0x2a, 0x49, 0x69,
	0xe4, 0x4f, 0xbd, 0x02, 0x2d, 0xbb, 0xaa, 0xb6, 0x4a, 0x23, 0xfb, 0x00, 0xb3, 0x1a, 0x7f, 0xe2,
	0x8e, 0x4a, 0x89, 0x07, 0x71, 0x0c, 0xda, 0xd8, 0x6b, 0x67, 0x21, 0xf8, 0xda, 0x72, 0xef, 0xbe,
	0x83, 0xb3, 0xeb, 0xc6, 0x1d, 0xfb, 0x66, 0xea, 0xcd, 0x69, 0x8b, 0x4f, 0xde, 0x39, 0xbc, 0xec,
	0xb6, 0xea, 0xcb, 0xe0, 0xe5, 0x99, 0xb9, 0xdb, 0xee, 0x54, 0x78, 0x0b, 0x93, 0x9d, 0x38, 0x1c,
	0xbc, 0x65, 0x1b, 0xe2, 0x59, 0xb8, 0x49, 0xcb, 0xb6, 0x01, 0xb1, 0x37, 0x00, 0x4d, 0x25, 0x05,
	0xa1, 0x2c, 0x05, 0xf1, 0x89, 0xff, 0x77, 0xd3, 0x3b, 0xb2, 0xa1, 0xc5, 0x67, 0x98, 0x7e, 0x57,
	0x8e, 0x4e, 0x8f, 0xf3, 0x1e, 0x46, 0x8a, 0x50, 0x3b, 0x1e, 0xe5, 0xf1, 0x32, 0x5b, 0xcf, 0x56,
	0xe1, 0x3d, 0x57, 0x9d, 0x51, 0x84, 0x7c, 0xfd, 0x67, 0x00, 0x93, 0x0e, 0x6e, 0x2a, 0xc5, 0x2e,
	0x20, 0xf9, 0x52, 0xa3, 0x20, 0x64, 0x0f, 0x4b, 0xf3, 0x87, 0x68, 0xf1, 0x84, 0x7d, 0x82, 0x61,
	0x3b, 0x9c, 0xbd, 0xb8, 0x0f, 0xdb, 0x53, 0x81, 0xbf, 0x1a, 0x74, 0x34, 0x7f, 0xd5, 0x87, 0xfd,
	0xd6, 0x39, 0xc4, 0xdf, 0xb0, 0x57, 0xba, 0x24, 0xd4, 0xf7, 0xa5, 0xff, 0x8e, 0xb9, 0x80, 0xe4,
	0xca, 0x5f, 0xf8, 0xd1, 0x8b, 0xad, 0x21, 0x29, 0x50, 0xdb, 0xdf, 0xf8, 0xf8, 0x29, 0xd7, 0x89,
	0xff, 0xd4, 0x3f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xda, 0x5e, 0x61, 0xdf, 0x13, 0x03, 0x00,
	0x00,
}
