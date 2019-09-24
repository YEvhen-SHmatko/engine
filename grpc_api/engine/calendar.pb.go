// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

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

type ListCalendar struct {
	Items []*Calendar `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListCalendar) Reset()                    { *m = ListCalendar{} }
func (m *ListCalendar) String() string            { return proto.CompactTextString(m) }
func (*ListCalendar) ProtoMessage()               {}
func (*ListCalendar) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ListCalendar) GetItems() []*Calendar {
	if m != nil {
		return m.Items
	}
	return nil
}

type Calendar struct {
	Id          int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	DomainId    int64   `protobuf:"varint,3,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	Start       int64   `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	Finish      int64   `protobuf:"varint,5,opt,name=finish" json:"finish,omitempty"`
	Timezone    *Lookup `protobuf:"bytes,6,opt,name=timezone" json:"timezone,omitempty"`
	Description string  `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
}

func (m *Calendar) Reset()                    { *m = Calendar{} }
func (m *Calendar) String() string            { return proto.CompactTextString(m) }
func (*Calendar) ProtoMessage()               {}
func (*Calendar) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *Calendar) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Calendar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Calendar) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *Calendar) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Calendar) GetFinish() int64 {
	if m != nil {
		return m.Finish
	}
	return 0
}

func (m *Calendar) GetTimezone() *Lookup {
	if m != nil {
		return m.Timezone
	}
	return nil
}

func (m *Calendar) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ListTimezoneResponse struct {
	Items []*Timezone `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListTimezoneResponse) Reset()                    { *m = ListTimezoneResponse{} }
func (m *ListTimezoneResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTimezoneResponse) ProtoMessage()               {}
func (*ListTimezoneResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *ListTimezoneResponse) GetItems() []*Timezone {
	if m != nil {
		return m.Items
	}
	return nil
}

type Timezone struct {
	Id     int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Offset string `protobuf:"bytes,3,opt,name=offset" json:"offset,omitempty"`
}

func (m *Timezone) Reset()                    { *m = Timezone{} }
func (m *Timezone) String() string            { return proto.CompactTextString(m) }
func (*Timezone) ProtoMessage()               {}
func (*Timezone) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *Timezone) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Timezone) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Timezone) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

type AcceptOfDay struct {
	Id             int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	WeekDay        int32 `protobuf:"varint,3,opt,name=week_day,json=weekDay" json:"week_day,omitempty"`
	StartTimeOfDay int32 `protobuf:"varint,4,opt,name=start_time_of_day,json=startTimeOfDay" json:"start_time_of_day,omitempty"`
	EndTimeOfDay   int32 `protobuf:"varint,5,opt,name=end_time_of_day,json=endTimeOfDay" json:"end_time_of_day,omitempty"`
}

func (m *AcceptOfDay) Reset()                    { *m = AcceptOfDay{} }
func (m *AcceptOfDay) String() string            { return proto.CompactTextString(m) }
func (*AcceptOfDay) ProtoMessage()               {}
func (*AcceptOfDay) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *AcceptOfDay) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AcceptOfDay) GetWeekDay() int32 {
	if m != nil {
		return m.WeekDay
	}
	return 0
}

func (m *AcceptOfDay) GetStartTimeOfDay() int32 {
	if m != nil {
		return m.StartTimeOfDay
	}
	return 0
}

func (m *AcceptOfDay) GetEndTimeOfDay() int32 {
	if m != nil {
		return m.EndTimeOfDay
	}
	return 0
}

type AcceptOfDayReqeust struct {
	CalendarId int64 `protobuf:"varint,1,opt,name=calendar_id,json=calendarId" json:"calendar_id,omitempty"`
}

func (m *AcceptOfDayReqeust) Reset()                    { *m = AcceptOfDayReqeust{} }
func (m *AcceptOfDayReqeust) String() string            { return proto.CompactTextString(m) }
func (*AcceptOfDayReqeust) ProtoMessage()               {}
func (*AcceptOfDayReqeust) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *AcceptOfDayReqeust) GetCalendarId() int64 {
	if m != nil {
		return m.CalendarId
	}
	return 0
}

type ListAcceptOfDay struct {
	Items []*AcceptOfDay `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListAcceptOfDay) Reset()                    { *m = ListAcceptOfDay{} }
func (m *ListAcceptOfDay) String() string            { return proto.CompactTextString(m) }
func (*ListAcceptOfDay) ProtoMessage()               {}
func (*ListAcceptOfDay) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *ListAcceptOfDay) GetItems() []*AcceptOfDay {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ListCalendar)(nil), "engine.ListCalendar")
	proto.RegisterType((*Calendar)(nil), "engine.Calendar")
	proto.RegisterType((*ListTimezoneResponse)(nil), "engine.ListTimezoneResponse")
	proto.RegisterType((*Timezone)(nil), "engine.Timezone")
	proto.RegisterType((*AcceptOfDay)(nil), "engine.AcceptOfDay")
	proto.RegisterType((*AcceptOfDayReqeust)(nil), "engine.AcceptOfDayReqeust")
	proto.RegisterType((*ListAcceptOfDay)(nil), "engine.ListAcceptOfDay")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CalendarApi service

type CalendarApiClient interface {
	// - POST /calendar
	Create(ctx context.Context, in *Calendar, opts ...grpc.CallOption) (*Calendar, error)
	// -GET /calendar
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCalendar, error)
	// -GET /calendar/:id
	Get(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*Calendar, error)
	// -PUT /calendar/:id
	Update(ctx context.Context, in *Calendar, opts ...grpc.CallOption) (*Calendar, error)
	// -DELETE /calendar/:id
	Remove(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*Calendar, error)
	// -GET /calendar/timezones
	GetTimezones(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListTimezoneResponse, error)
	// -GET /calendar/:id/accept
	GetAcceptOfDay(ctx context.Context, in *AcceptOfDayReqeust, opts ...grpc.CallOption) (*ListAcceptOfDay, error)
}

type calendarApiClient struct {
	cc *grpc.ClientConn
}

func NewCalendarApiClient(cc *grpc.ClientConn) CalendarApiClient {
	return &calendarApiClient{cc}
}

func (c *calendarApiClient) Create(ctx context.Context, in *Calendar, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCalendar, error) {
	out := new(ListCalendar)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) Get(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) Update(ctx context.Context, in *Calendar, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) Remove(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) GetTimezones(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListTimezoneResponse, error) {
	out := new(ListTimezoneResponse)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/GetTimezones", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarApiClient) GetAcceptOfDay(ctx context.Context, in *AcceptOfDayReqeust, opts ...grpc.CallOption) (*ListAcceptOfDay, error) {
	out := new(ListAcceptOfDay)
	err := grpc.Invoke(ctx, "/engine.CalendarApi/GetAcceptOfDay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalendarApi service

type CalendarApiServer interface {
	// - POST /calendar
	Create(context.Context, *Calendar) (*Calendar, error)
	// -GET /calendar
	List(context.Context, *ListRequest) (*ListCalendar, error)
	// -GET /calendar/:id
	Get(context.Context, *ItemRequest) (*Calendar, error)
	// -PUT /calendar/:id
	Update(context.Context, *Calendar) (*Calendar, error)
	// -DELETE /calendar/:id
	Remove(context.Context, *ItemRequest) (*Calendar, error)
	// -GET /calendar/timezones
	GetTimezones(context.Context, *ListRequest) (*ListTimezoneResponse, error)
	// -GET /calendar/:id/accept
	GetAcceptOfDay(context.Context, *AcceptOfDayReqeust) (*ListAcceptOfDay, error)
}

func RegisterCalendarApiServer(s *grpc.Server, srv CalendarApiServer) {
	s.RegisterService(&_CalendarApi_serviceDesc, srv)
}

func _CalendarApi_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Calendar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).Create(ctx, req.(*Calendar))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).Get(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Calendar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).Update(ctx, req.(*Calendar))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).Remove(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_GetTimezones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).GetTimezones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/GetTimezones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).GetTimezones(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarApi_GetAcceptOfDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptOfDayReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarApiServer).GetAcceptOfDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarApi/GetAcceptOfDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarApiServer).GetAcceptOfDay(ctx, req.(*AcceptOfDayReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.CalendarApi",
	HandlerType: (*CalendarApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CalendarApi_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CalendarApi_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CalendarApi_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CalendarApi_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _CalendarApi_Remove_Handler,
		},
		{
			MethodName: "GetTimezones",
			Handler:    _CalendarApi_GetTimezones_Handler,
		},
		{
			MethodName: "GetAcceptOfDay",
			Handler:    _CalendarApi_GetAcceptOfDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0xf4, 0x23, 0xdb, 0xde, 0x94, 0xae, 0x8e, 0x65, 0x8d, 0x55, 0x30, 0x04, 0x94, 0xae,
	0x0f, 0x01, 0xbb, 0xe8, 0x93, 0x08, 0x4b, 0xc5, 0x52, 0x58, 0x10, 0x06, 0x7d, 0x0e, 0x31, 0x73,
	0xa3, 0xc3, 0x9a, 0x99, 0x6c, 0x66, 0xaa, 0xac, 0x3f, 0xc1, 0x57, 0x7f, 0x94, 0x7f, 0x4b, 0x32,
	0xf9, 0x60, 0x6a, 0x2b, 0xf4, 0x2d, 0xf7, 0xdc, 0x73, 0xee, 0xc7, 0x99, 0x4b, 0x60, 0x9a, 0x26,
	0xdf, 0x50, 0xb0, 0xa4, 0x8c, 0x8a, 0x52, 0x6a, 0x49, 0x5c, 0x14, 0x5f, 0xb8, 0xc0, 0xb9, 0x97,
	0x4a, 0xa1, 0x74, 0x0d, 0x86, 0xaf, 0x61, 0x72, 0xcd, 0x95, 0x5e, 0x35, 0x54, 0xf2, 0x1c, 0x86,
	0x5c, 0x63, 0xae, 0x7c, 0x27, 0xe8, 0x2f, 0xbc, 0xe5, 0xbd, 0xa8, 0x16, 0x45, 0x2d, 0x81, 0xd6,
	0xe9, 0xf0, 0x8f, 0x03, 0xa3, 0x4e, 0x34, 0x85, 0x1e, 0x67, 0xbe, 0x13, 0x38, 0x8b, 0x3e, 0xed,
	0x71, 0x46, 0x08, 0x0c, 0x44, 0x92, 0xa3, 0xdf, 0x0b, 0x9c, 0xc5, 0x98, 0x9a, 0x6f, 0xf2, 0x18,
	0xc6, 0x4c, 0xe6, 0x09, 0x17, 0x31, 0x67, 0x7e, 0xdf, 0x50, 0x47, 0x35, 0xb0, 0x61, 0x64, 0x06,
	0x43, 0xa5, 0x93, 0x52, 0xfb, 0x03, 0x93, 0xa8, 0x03, 0x72, 0x0e, 0x6e, 0xc6, 0x05, 0x57, 0x5f,
	0xfd, 0xa1, 0x81, 0x9b, 0x88, 0xbc, 0x80, 0x91, 0xe6, 0x39, 0xfe, 0x94, 0x02, 0x7d, 0x37, 0x70,
	0x16, 0xde, 0x72, 0xda, 0x8e, 0x79, 0x2d, 0xe5, 0xcd, 0xb6, 0xa0, 0x5d, 0x9e, 0x04, 0xe0, 0x31,
	0x54, 0x69, 0xc9, 0x0b, 0xcd, 0xa5, 0xf0, 0x4f, 0xcd, 0x44, 0x36, 0x14, 0xbe, 0x85, 0x59, 0xe5,
	0xc0, 0xc7, 0x46, 0x41, 0x51, 0x15, 0x52, 0x28, 0xfc, 0xaf, 0x13, 0x1d, 0xb1, 0x71, 0xe2, 0x3d,
	0x8c, 0x5a, 0xe8, 0x28, 0x23, 0xce, 0xc1, 0x95, 0x59, 0xa6, 0x50, 0x1b, 0x17, 0xc6, 0xb4, 0x89,
	0xc2, 0x5f, 0x0e, 0x78, 0x57, 0x69, 0x8a, 0x85, 0xfe, 0x90, 0xbd, 0x4b, 0xee, 0xf6, 0x6a, 0x3d,
	0x82, 0xd1, 0x0f, 0xc4, 0x9b, 0x98, 0x25, 0x77, 0x46, 0x39, 0xa4, 0xa7, 0x55, 0x5c, 0x51, 0x2f,
	0xe0, 0xbe, 0x71, 0x2c, 0xae, 0xd6, 0x8e, 0x65, 0x66, 0x38, 0x03, 0xc3, 0x99, 0x9a, 0x44, 0x35,
	0x60, 0x5d, 0xf5, 0x19, 0x9c, 0xa1, 0x60, 0x3b, 0xc4, 0xa1, 0x21, 0x4e, 0x50, 0xb0, 0x8e, 0x16,
	0xbe, 0x02, 0x62, 0xcd, 0x42, 0xf1, 0x16, 0xb7, 0x4a, 0x93, 0xa7, 0xe0, 0xb5, 0x37, 0x15, 0x77,
	0xb3, 0x41, 0x0b, 0x6d, 0x58, 0xf8, 0x06, 0xce, 0x2a, 0x2f, 0xed, 0x35, 0x2e, 0x76, 0x6d, 0x7c,
	0xd0, 0xda, 0x68, 0x97, 0xaf, 0x19, 0xcb, 0xdf, 0x7d, 0xf0, 0xda, 0x9b, 0xba, 0x2a, 0x38, 0x89,
	0xc0, 0x5d, 0x95, 0x98, 0x68, 0x24, 0x7b, 0x67, 0x38, 0xdf, 0x43, 0xc2, 0x13, 0x72, 0x09, 0x83,
	0xaa, 0x3b, 0xe9, 0x7a, 0x54, 0x11, 0xc5, 0xdb, 0x2d, 0x2a, 0x3d, 0x9f, 0xd9, 0xa0, 0x25, 0x8a,
	0xa0, 0xbf, 0x46, 0x4b, 0xb3, 0xd1, 0x98, 0xb7, 0x9a, 0x43, 0x4d, 0x22, 0x70, 0x3f, 0x15, 0xec,
	0xf8, 0xa1, 0x5e, 0x82, 0x4b, 0x31, 0x97, 0xdf, 0xf1, 0xf8, 0x16, 0x2b, 0x98, 0xac, 0xb1, 0x3b,
	0x48, 0x75, 0x78, 0x9f, 0x27, 0x36, 0xf8, 0xef, 0xf1, 0x86, 0x27, 0x64, 0x0d, 0xd3, 0x35, 0xee,
	0xbc, 0xc4, 0xfc, 0x90, 0xf5, 0xf5, 0xcb, 0xce, 0x1f, 0xda, 0xd5, 0xac, 0x7c, 0x78, 0xf2, 0xd9,
	0x35, 0x3f, 0x8a, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x17, 0x87, 0x37, 0x4f, 0x04,
	0x00, 0x00,
}
