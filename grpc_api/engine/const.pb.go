// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

package engine

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Lookup struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lookup) Reset()         { *m = Lookup{} }
func (m *Lookup) String() string { return proto.CompactTextString(m) }
func (*Lookup) ProtoMessage()    {}
func (*Lookup) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

func (m *Lookup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lookup.Unmarshal(m, b)
}
func (m *Lookup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lookup.Marshal(b, m, deterministic)
}
func (m *Lookup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lookup.Merge(m, src)
}
func (m *Lookup) XXX_Size() int {
	return xxx_messageInfo_Lookup.Size(m)
}
func (m *Lookup) XXX_DiscardUnknown() {
	xxx_messageInfo_Lookup.DiscardUnknown(m)
}

var xxx_messageInfo_Lookup proto.InternalMessageInfo

func (m *Lookup) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Lookup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FilterBetween struct {
	From                 int64    `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterBetween) Reset()         { *m = FilterBetween{} }
func (m *FilterBetween) String() string { return proto.CompactTextString(m) }
func (*FilterBetween) ProtoMessage()    {}
func (*FilterBetween) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

func (m *FilterBetween) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterBetween.Unmarshal(m, b)
}
func (m *FilterBetween) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterBetween.Marshal(b, m, deterministic)
}
func (m *FilterBetween) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterBetween.Merge(m, src)
}
func (m *FilterBetween) XXX_Size() int {
	return xxx_messageInfo_FilterBetween.Size(m)
}
func (m *FilterBetween) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterBetween.DiscardUnknown(m)
}

var xxx_messageInfo_FilterBetween proto.InternalMessageInfo

func (m *FilterBetween) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *FilterBetween) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

type ListRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ListRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type DomainRecord struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            int64    `protobuf:"varint,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            int64    `protobuf:"varint,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DomainRecord) Reset()         { *m = DomainRecord{} }
func (m *DomainRecord) String() string { return proto.CompactTextString(m) }
func (*DomainRecord) ProtoMessage()    {}
func (*DomainRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{3}
}

func (m *DomainRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainRecord.Unmarshal(m, b)
}
func (m *DomainRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainRecord.Marshal(b, m, deterministic)
}
func (m *DomainRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainRecord.Merge(m, src)
}
func (m *DomainRecord) XXX_Size() int {
	return xxx_messageInfo_DomainRecord.Size(m)
}
func (m *DomainRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DomainRecord proto.InternalMessageInfo

func (m *DomainRecord) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *DomainRecord) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *DomainRecord) GetCreatedBy() int64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *DomainRecord) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *DomainRecord) GetUpdatedBy() int64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

type ListForItemRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int32    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListForItemRequest) Reset()         { *m = ListForItemRequest{} }
func (m *ListForItemRequest) String() string { return proto.CompactTextString(m) }
func (*ListForItemRequest) ProtoMessage()    {}
func (*ListForItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{4}
}

func (m *ListForItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListForItemRequest.Unmarshal(m, b)
}
func (m *ListForItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListForItemRequest.Marshal(b, m, deterministic)
}
func (m *ListForItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListForItemRequest.Merge(m, src)
}
func (m *ListForItemRequest) XXX_Size() int {
	return xxx_messageInfo_ListForItemRequest.Size(m)
}
func (m *ListForItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListForItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListForItemRequest proto.InternalMessageInfo

func (m *ListForItemRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ListForItemRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ListForItemRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListForItemRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ItemRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequest) Reset()         { *m = ItemRequest{} }
func (m *ItemRequest) String() string { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()    {}
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{5}
}

func (m *ItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequest.Unmarshal(m, b)
}
func (m *ItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequest.Marshal(b, m, deterministic)
}
func (m *ItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequest.Merge(m, src)
}
func (m *ItemRequest) XXX_Size() int {
	return xxx_messageInfo_ItemRequest.Size(m)
}
func (m *ItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequest proto.InternalMessageInfo

func (m *ItemRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Lookup)(nil), "engine.Lookup")
	proto.RegisterType((*FilterBetween)(nil), "engine.FilterBetween")
	proto.RegisterType((*ListRequest)(nil), "engine.ListRequest")
	proto.RegisterType((*DomainRecord)(nil), "engine.DomainRecord")
	proto.RegisterType((*ListForItemRequest)(nil), "engine.ListForItemRequest")
	proto.RegisterType((*ItemRequest)(nil), "engine.ItemRequest")
	proto.RegisterType((*Response)(nil), "engine.Response")
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xa5, 0xed, 0x56, 0xd7, 0x6f, 0xea, 0x21, 0x07, 0x2d, 0x88, 0x30, 0x72, 0xda, 0x41, 0xbc,
	0xec, 0xe6, 0xcd, 0x21, 0x83, 0xc1, 0x4e, 0xf9, 0x03, 0xa3, 0x5b, 0x3e, 0x47, 0xd0, 0x26, 0x35,
	0xf9, 0x8a, 0xd4, 0x3f, 0xe4, 0xdf, 0x94, 0x24, 0x9d, 0x56, 0x10, 0xd9, 0xed, 0xe5, 0xbd, 0xbe,
	0xbc, 0xaf, 0xef, 0x0b, 0x4c, 0xf7, 0x46, 0x3b, 0xba, 0x6f, 0xac, 0x21, 0xc3, 0x72, 0xd4, 0x07,
	0xa5, 0x91, 0xdf, 0x41, 0xbe, 0x31, 0xe6, 0xa5, 0x6d, 0xd8, 0x25, 0xa4, 0x4a, 0x96, 0xc9, 0x2c,
	0x99, 0x67, 0x22, 0x55, 0x92, 0x31, 0x18, 0xe9, 0xaa, 0xc6, 0x32, 0x9d, 0x25, 0xf3, 0x42, 0x04,
	0xcc, 0x17, 0x70, 0xb1, 0x52, 0xaf, 0x84, 0x76, 0x89, 0xf4, 0x8e, 0xa8, 0xfd, 0x47, 0xcf, 0xd6,
	0xd4, 0xbd, 0x2d, 0x60, 0x7f, 0x11, 0x99, 0x60, 0xcb, 0x44, 0x4a, 0x86, 0x0b, 0x98, 0x6e, 0x94,
	0x23, 0x81, 0x6f, 0x2d, 0x3a, 0x62, 0x37, 0x50, 0x48, 0x53, 0x57, 0x4a, 0x6f, 0xbf, 0xe3, 0x26,
	0x91, 0x58, 0x87, 0x50, 0xa7, 0x3e, 0x62, 0xe8, 0x58, 0x04, 0xec, 0xb9, 0xa6, 0x3a, 0x60, 0x99,
	0x45, 0xce, 0x63, 0xfe, 0x99, 0xc0, 0xf9, 0x53, 0x30, 0x09, 0xdc, 0x1b, 0x2b, 0xff, 0xbf, 0xf5,
	0x16, 0x60, 0x6f, 0xb1, 0x22, 0x94, 0xdb, 0x8a, 0xfa, 0xc9, 0x8a, 0x9e, 0x79, 0xa4, 0xa1, 0xbc,
	0xeb, 0x42, 0xcc, 0x8f, 0xbc, 0xec, 0xbc, 0xdc, 0x36, 0xf2, 0xe8, 0x1e, 0x45, 0xb9, 0x67, 0xa2,
	0xfb, 0x28, 0xef, 0xba, 0x72, 0xfc, 0x4b, 0x5e, 0x76, 0xdc, 0x02, 0xf3, 0x7f, 0xbf, 0x32, 0x76,
	0x4d, 0x58, 0x9f, 0x54, 0xc2, 0x35, 0x9c, 0x29, 0xc2, 0xda, 0x4b, 0x71, 0xd6, 0xdc, 0x1f, 0x07,
	0xed, 0x64, 0x7f, 0xb4, 0x33, 0x1a, 0xb4, 0xf3, 0x00, 0xd3, 0x93, 0xc3, 0xe2, 0xda, 0xd3, 0xe3,
	0xda, 0x39, 0x87, 0x89, 0x40, 0xd7, 0x18, 0xed, 0x90, 0x5d, 0x41, 0xee, 0xa8, 0xa2, 0xd6, 0x05,
	0x57, 0x21, 0xfa, 0xd3, 0x2e, 0x0f, 0x6f, 0x68, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x62,
	0x08, 0xf0, 0x52, 0x02, 0x00, 0x00,
}