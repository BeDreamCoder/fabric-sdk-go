// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chaincode.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type InstallCCRequest struct {
	CcId                 string   `protobuf:"bytes,1,opt,name=cc_id,json=ccId,proto3" json:"cc_id,omitempty"`
	CcVersion            string   `protobuf:"bytes,2,opt,name=cc_version,json=ccVersion,proto3" json:"cc_version,omitempty"`
	CcPath               string   `protobuf:"bytes,3,opt,name=cc_path,json=ccPath,proto3" json:"cc_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallCCRequest) Reset()         { *m = InstallCCRequest{} }
func (m *InstallCCRequest) String() string { return proto.CompactTextString(m) }
func (*InstallCCRequest) ProtoMessage()    {}
func (*InstallCCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{0}
}

func (m *InstallCCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallCCRequest.Unmarshal(m, b)
}
func (m *InstallCCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallCCRequest.Marshal(b, m, deterministic)
}
func (m *InstallCCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallCCRequest.Merge(m, src)
}
func (m *InstallCCRequest) XXX_Size() int {
	return xxx_messageInfo_InstallCCRequest.Size(m)
}
func (m *InstallCCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallCCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstallCCRequest proto.InternalMessageInfo

func (m *InstallCCRequest) GetCcId() string {
	if m != nil {
		return m.CcId
	}
	return ""
}

func (m *InstallCCRequest) GetCcVersion() string {
	if m != nil {
		return m.CcVersion
	}
	return ""
}

func (m *InstallCCRequest) GetCcPath() string {
	if m != nil {
		return m.CcPath
	}
	return ""
}

type InstantiateCCRequest struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	CcId                 string   `protobuf:"bytes,2,opt,name=cc_id,json=ccId,proto3" json:"cc_id,omitempty"`
	CcVersion            string   `protobuf:"bytes,3,opt,name=cc_version,json=ccVersion,proto3" json:"cc_version,omitempty"`
	CcPath               string   `protobuf:"bytes,4,opt,name=cc_path,json=ccPath,proto3" json:"cc_path,omitempty"`
	Args                 [][]byte `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstantiateCCRequest) Reset()         { *m = InstantiateCCRequest{} }
func (m *InstantiateCCRequest) String() string { return proto.CompactTextString(m) }
func (*InstantiateCCRequest) ProtoMessage()    {}
func (*InstantiateCCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{1}
}

func (m *InstantiateCCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstantiateCCRequest.Unmarshal(m, b)
}
func (m *InstantiateCCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstantiateCCRequest.Marshal(b, m, deterministic)
}
func (m *InstantiateCCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstantiateCCRequest.Merge(m, src)
}
func (m *InstantiateCCRequest) XXX_Size() int {
	return xxx_messageInfo_InstantiateCCRequest.Size(m)
}
func (m *InstantiateCCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstantiateCCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstantiateCCRequest proto.InternalMessageInfo

func (m *InstantiateCCRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *InstantiateCCRequest) GetCcId() string {
	if m != nil {
		return m.CcId
	}
	return ""
}

func (m *InstantiateCCRequest) GetCcVersion() string {
	if m != nil {
		return m.CcVersion
	}
	return ""
}

func (m *InstantiateCCRequest) GetCcPath() string {
	if m != nil {
		return m.CcPath
	}
	return ""
}

func (m *InstantiateCCRequest) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

type InstantiateCCResponse struct {
	Status               StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=protos.StatusCode" json:"status,omitempty"`
	TransactionId        string     `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InstantiateCCResponse) Reset()         { *m = InstantiateCCResponse{} }
func (m *InstantiateCCResponse) String() string { return proto.CompactTextString(m) }
func (*InstantiateCCResponse) ProtoMessage()    {}
func (*InstantiateCCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{2}
}

func (m *InstantiateCCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstantiateCCResponse.Unmarshal(m, b)
}
func (m *InstantiateCCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstantiateCCResponse.Marshal(b, m, deterministic)
}
func (m *InstantiateCCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstantiateCCResponse.Merge(m, src)
}
func (m *InstantiateCCResponse) XXX_Size() int {
	return xxx_messageInfo_InstantiateCCResponse.Size(m)
}
func (m *InstantiateCCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstantiateCCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstantiateCCResponse proto.InternalMessageInfo

func (m *InstantiateCCResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_SUCCESS
}

func (m *InstantiateCCResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type UpgradeCCRequest struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	CcId                 string   `protobuf:"bytes,2,opt,name=cc_id,json=ccId,proto3" json:"cc_id,omitempty"`
	CcVersion            string   `protobuf:"bytes,3,opt,name=cc_version,json=ccVersion,proto3" json:"cc_version,omitempty"`
	CcPath               string   `protobuf:"bytes,4,opt,name=cc_path,json=ccPath,proto3" json:"cc_path,omitempty"`
	Args                 [][]byte `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradeCCRequest) Reset()         { *m = UpgradeCCRequest{} }
func (m *UpgradeCCRequest) String() string { return proto.CompactTextString(m) }
func (*UpgradeCCRequest) ProtoMessage()    {}
func (*UpgradeCCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{3}
}

func (m *UpgradeCCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeCCRequest.Unmarshal(m, b)
}
func (m *UpgradeCCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeCCRequest.Marshal(b, m, deterministic)
}
func (m *UpgradeCCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeCCRequest.Merge(m, src)
}
func (m *UpgradeCCRequest) XXX_Size() int {
	return xxx_messageInfo_UpgradeCCRequest.Size(m)
}
func (m *UpgradeCCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeCCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeCCRequest proto.InternalMessageInfo

func (m *UpgradeCCRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *UpgradeCCRequest) GetCcId() string {
	if m != nil {
		return m.CcId
	}
	return ""
}

func (m *UpgradeCCRequest) GetCcVersion() string {
	if m != nil {
		return m.CcVersion
	}
	return ""
}

func (m *UpgradeCCRequest) GetCcPath() string {
	if m != nil {
		return m.CcPath
	}
	return ""
}

func (m *UpgradeCCRequest) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

type UpgradeCCResponse struct {
	Status               StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=protos.StatusCode" json:"status,omitempty"`
	TransactionId        string     `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpgradeCCResponse) Reset()         { *m = UpgradeCCResponse{} }
func (m *UpgradeCCResponse) String() string { return proto.CompactTextString(m) }
func (*UpgradeCCResponse) ProtoMessage()    {}
func (*UpgradeCCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{4}
}

func (m *UpgradeCCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeCCResponse.Unmarshal(m, b)
}
func (m *UpgradeCCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeCCResponse.Marshal(b, m, deterministic)
}
func (m *UpgradeCCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeCCResponse.Merge(m, src)
}
func (m *UpgradeCCResponse) XXX_Size() int {
	return xxx_messageInfo_UpgradeCCResponse.Size(m)
}
func (m *UpgradeCCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeCCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeCCResponse proto.InternalMessageInfo

func (m *UpgradeCCResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_SUCCESS
}

func (m *UpgradeCCResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type InvokeCCRequest struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	CcId                 string   `protobuf:"bytes,2,opt,name=cc_id,json=ccId,proto3" json:"cc_id,omitempty"`
	Func                 string   `protobuf:"bytes,3,opt,name=func,proto3" json:"func,omitempty"`
	Args                 [][]byte `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeCCRequest) Reset()         { *m = InvokeCCRequest{} }
func (m *InvokeCCRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeCCRequest) ProtoMessage()    {}
func (*InvokeCCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{5}
}

func (m *InvokeCCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeCCRequest.Unmarshal(m, b)
}
func (m *InvokeCCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeCCRequest.Marshal(b, m, deterministic)
}
func (m *InvokeCCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeCCRequest.Merge(m, src)
}
func (m *InvokeCCRequest) XXX_Size() int {
	return xxx_messageInfo_InvokeCCRequest.Size(m)
}
func (m *InvokeCCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeCCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeCCRequest proto.InternalMessageInfo

func (m *InvokeCCRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *InvokeCCRequest) GetCcId() string {
	if m != nil {
		return m.CcId
	}
	return ""
}

func (m *InvokeCCRequest) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *InvokeCCRequest) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

type InvokeCCResponse struct {
	Status               StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=protos.StatusCode" json:"status,omitempty"`
	TransactionId        string     `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Payload              []byte     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InvokeCCResponse) Reset()         { *m = InvokeCCResponse{} }
func (m *InvokeCCResponse) String() string { return proto.CompactTextString(m) }
func (*InvokeCCResponse) ProtoMessage()    {}
func (*InvokeCCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{6}
}

func (m *InvokeCCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeCCResponse.Unmarshal(m, b)
}
func (m *InvokeCCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeCCResponse.Marshal(b, m, deterministic)
}
func (m *InvokeCCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeCCResponse.Merge(m, src)
}
func (m *InvokeCCResponse) XXX_Size() int {
	return xxx_messageInfo_InvokeCCResponse.Size(m)
}
func (m *InvokeCCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeCCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeCCResponse proto.InternalMessageInfo

func (m *InvokeCCResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_SUCCESS
}

func (m *InvokeCCResponse) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *InvokeCCResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type QueryCCRequest struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	CcId                 string   `protobuf:"bytes,2,opt,name=cc_id,json=ccId,proto3" json:"cc_id,omitempty"`
	Func                 string   `protobuf:"bytes,3,opt,name=func,proto3" json:"func,omitempty"`
	Args                 [][]byte `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryCCRequest) Reset()         { *m = QueryCCRequest{} }
func (m *QueryCCRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCCRequest) ProtoMessage()    {}
func (*QueryCCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{7}
}

func (m *QueryCCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCCRequest.Unmarshal(m, b)
}
func (m *QueryCCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCCRequest.Marshal(b, m, deterministic)
}
func (m *QueryCCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCCRequest.Merge(m, src)
}
func (m *QueryCCRequest) XXX_Size() int {
	return xxx_messageInfo_QueryCCRequest.Size(m)
}
func (m *QueryCCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCCRequest proto.InternalMessageInfo

func (m *QueryCCRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *QueryCCRequest) GetCcId() string {
	if m != nil {
		return m.CcId
	}
	return ""
}

func (m *QueryCCRequest) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *QueryCCRequest) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

type QueryCCResponse struct {
	Status               StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=protos.StatusCode" json:"status,omitempty"`
	Payload              []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryCCResponse) Reset()         { *m = QueryCCResponse{} }
func (m *QueryCCResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCCResponse) ProtoMessage()    {}
func (*QueryCCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97136ef4b384cc22, []int{8}
}

func (m *QueryCCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCCResponse.Unmarshal(m, b)
}
func (m *QueryCCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCCResponse.Marshal(b, m, deterministic)
}
func (m *QueryCCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCCResponse.Merge(m, src)
}
func (m *QueryCCResponse) XXX_Size() int {
	return xxx_messageInfo_QueryCCResponse.Size(m)
}
func (m *QueryCCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCCResponse proto.InternalMessageInfo

func (m *QueryCCResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_SUCCESS
}

func (m *QueryCCResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallCCRequest)(nil), "protos.InstallCCRequest")
	proto.RegisterType((*InstantiateCCRequest)(nil), "protos.InstantiateCCRequest")
	proto.RegisterType((*InstantiateCCResponse)(nil), "protos.InstantiateCCResponse")
	proto.RegisterType((*UpgradeCCRequest)(nil), "protos.UpgradeCCRequest")
	proto.RegisterType((*UpgradeCCResponse)(nil), "protos.UpgradeCCResponse")
	proto.RegisterType((*InvokeCCRequest)(nil), "protos.InvokeCCRequest")
	proto.RegisterType((*InvokeCCResponse)(nil), "protos.InvokeCCResponse")
	proto.RegisterType((*QueryCCRequest)(nil), "protos.QueryCCRequest")
	proto.RegisterType((*QueryCCResponse)(nil), "protos.QueryCCResponse")
}

func init() { proto.RegisterFile("chaincode.proto", fileDescriptor_97136ef4b384cc22) }

var fileDescriptor_97136ef4b384cc22 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0x12, 0x37, 0xc1, 0xa3, 0x34, 0x29, 0x43, 0x20, 0xc6, 0xa2, 0x52, 0x65, 0x09, 0xa9,
	0x42, 0x6a, 0x40, 0xe5, 0x0a, 0x42, 0x6a, 0x4e, 0xbe, 0x20, 0x30, 0x02, 0x24, 0x2e, 0xd1, 0x76,
	0xbc, 0x49, 0x4c, 0x9d, 0x5d, 0xb3, 0xbb, 0x89, 0xd4, 0x13, 0xff, 0x81, 0x03, 0xff, 0x8e, 0xff,
	0x82, 0xea, 0x6f, 0xa7, 0x94, 0x03, 0x28, 0x88, 0x93, 0xd7, 0x6f, 0xc6, 0xf3, 0xde, 0x1b, 0xcf,
	0x0e, 0x0c, 0x69, 0xc9, 0x22, 0x41, 0x32, 0xe4, 0x93, 0x44, 0x49, 0x23, 0xb1, 0x9b, 0x3e, 0xb4,
	0xdb, 0x27, 0xb9, 0x5a, 0x49, 0x91, 0xa1, 0xde, 0x0c, 0x0e, 0x7d, 0xa1, 0x0d, 0x8b, 0xe3, 0xe9,
	0x34, 0xe0, 0x5f, 0xd6, 0x5c, 0x1b, 0xbc, 0x07, 0xfb, 0x44, 0xb3, 0x28, 0x74, 0x5a, 0xc7, 0xad,
	0x13, 0x3b, 0xb0, 0x88, 0xfc, 0x10, 0x8f, 0x00, 0x88, 0x66, 0x1b, 0xae, 0x74, 0x24, 0x85, 0xd3,
	0x4e, 0x23, 0x36, 0xd1, 0x87, 0x0c, 0xc0, 0x31, 0xf4, 0x88, 0x66, 0x09, 0x33, 0x4b, 0xa7, 0x93,
	0xc6, 0xba, 0x44, 0x6f, 0x98, 0x59, 0x7a, 0xdf, 0x5b, 0x30, 0x4a, 0x19, 0x84, 0x89, 0x98, 0xe1,
	0x15, 0xcb, 0x75, 0xc1, 0x25, 0x13, 0x82, 0xc7, 0x15, 0x95, 0x9d, 0x23, 0x7e, 0x58, 0x89, 0x68,
	0xdf, 0x2a, 0xa2, 0xf3, 0x1b, 0x11, 0x56, 0x5d, 0x04, 0x22, 0x58, 0x4c, 0x2d, 0xb4, 0xb3, 0x7f,
	0xdc, 0x39, 0xe9, 0x07, 0xe9, 0xd9, 0xfb, 0x0c, 0xf7, 0xb7, 0x74, 0xe9, 0x44, 0x0a, 0xcd, 0xf1,
	0x09, 0x74, 0xb5, 0x61, 0x66, 0xad, 0x53, 0x51, 0x83, 0x33, 0xcc, 0x5a, 0xa5, 0x27, 0xef, 0x52,
	0x74, 0x2a, 0x43, 0x1e, 0xe4, 0x19, 0xf8, 0x18, 0x06, 0x46, 0x31, 0xa1, 0x19, 0x99, 0x48, 0x8a,
	0x4a, 0xee, 0x41, 0x0d, 0xf5, 0x43, 0xef, 0x5b, 0x0b, 0x0e, 0xdf, 0x27, 0x0b, 0xc5, 0xc2, 0xff,
	0xa8, 0x01, 0x73, 0xb8, 0x5b, 0xd3, 0xb4, 0x3b, 0xf3, 0x2b, 0x18, 0xfa, 0x62, 0x23, 0x2f, 0xff,
	0xd2, 0x3a, 0x82, 0x35, 0x5f, 0x0b, 0xca, 0x4d, 0xa7, 0xe7, 0xd2, 0x96, 0x55, 0xb3, 0xf5, 0xf5,
	0x7a, 0xa2, 0x0b, 0xba, 0x9d, 0xb9, 0x42, 0x07, 0x7a, 0x09, 0xbb, 0x8a, 0x25, 0x0b, 0x53, 0x45,
	0xfd, 0xa0, 0x78, 0xf5, 0x62, 0x18, 0xbc, 0x5d, 0x73, 0x75, 0xf5, 0x6f, 0xec, 0x7e, 0x84, 0x61,
	0xc9, 0xf6, 0x07, 0x6e, 0x6b, 0x36, 0xda, 0x0d, 0x1b, 0x67, 0x3f, 0xda, 0x60, 0x4f, 0x8b, 0x1d,
	0x82, 0x2f, 0xc1, 0x2e, 0xf7, 0x04, 0x3a, 0x45, 0xc1, 0xed, 0xd5, 0xe1, 0x8e, 0x4a, 0x2a, 0xae,
	0x36, 0x5c, 0x65, 0x84, 0xde, 0x1e, 0xbe, 0x86, 0x83, 0xc6, 0x65, 0xc3, 0x47, 0x8d, 0x12, 0x5b,
	0xbb, 0xc1, 0x3d, 0xba, 0x25, 0x9a, 0x19, 0xf4, 0xf6, 0xf0, 0x1c, 0xec, 0x72, 0x76, 0x2b, 0x39,
	0xdb, 0x57, 0xcc, 0x7d, 0xf8, 0x8b, 0x48, 0x59, 0xe3, 0x15, 0xdc, 0x29, 0x06, 0x05, 0xc7, 0x15,
	0x61, 0x63, 0x52, 0x5d, 0xe7, 0x66, 0xa0, 0x2c, 0xf0, 0x02, 0x7a, 0x79, 0xeb, 0xf1, 0x41, 0x91,
	0xd6, 0xfc, 0xf3, 0xee, 0xf8, 0x06, 0x5e, 0x7c, 0x7d, 0xfe, 0x0c, 0x1c, 0xa9, 0x16, 0x93, 0x39,
	0xbb, 0x50, 0x11, 0x9d, 0xea, 0xf0, 0xf2, 0x74, 0x21, 0xf3, 0xec, 0x4f, 0xa3, 0x06, 0xfa, 0x34,
	0x43, 0x2f, 0xb2, 0x0d, 0xfe, 0xfc, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x67, 0x69, 0x35,
	0xdb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChaincodeClient is the client API for Chaincode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChaincodeClient interface {
	InstallCC(ctx context.Context, in *InstallCCRequest, opts ...grpc.CallOption) (*ServerStatus, error)
	InstantiateCC(ctx context.Context, in *InstantiateCCRequest, opts ...grpc.CallOption) (*InstantiateCCResponse, error)
	UpgradeCC(ctx context.Context, in *UpgradeCCRequest, opts ...grpc.CallOption) (*UpgradeCCResponse, error)
	InvokeCC(ctx context.Context, in *InvokeCCRequest, opts ...grpc.CallOption) (*InvokeCCResponse, error)
	QueryCC(ctx context.Context, in *QueryCCRequest, opts ...grpc.CallOption) (*QueryCCResponse, error)
}

type chaincodeClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeClient(cc *grpc.ClientConn) ChaincodeClient {
	return &chaincodeClient{cc}
}

func (c *chaincodeClient) InstallCC(ctx context.Context, in *InstallCCRequest, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := c.cc.Invoke(ctx, "/protos.Chaincode/InstallCC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) InstantiateCC(ctx context.Context, in *InstantiateCCRequest, opts ...grpc.CallOption) (*InstantiateCCResponse, error) {
	out := new(InstantiateCCResponse)
	err := c.cc.Invoke(ctx, "/protos.Chaincode/InstantiateCC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) UpgradeCC(ctx context.Context, in *UpgradeCCRequest, opts ...grpc.CallOption) (*UpgradeCCResponse, error) {
	out := new(UpgradeCCResponse)
	err := c.cc.Invoke(ctx, "/protos.Chaincode/UpgradeCC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) InvokeCC(ctx context.Context, in *InvokeCCRequest, opts ...grpc.CallOption) (*InvokeCCResponse, error) {
	out := new(InvokeCCResponse)
	err := c.cc.Invoke(ctx, "/protos.Chaincode/InvokeCC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeClient) QueryCC(ctx context.Context, in *QueryCCRequest, opts ...grpc.CallOption) (*QueryCCResponse, error) {
	out := new(QueryCCResponse)
	err := c.cc.Invoke(ctx, "/protos.Chaincode/QueryCC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaincodeServer is the server API for Chaincode service.
type ChaincodeServer interface {
	InstallCC(context.Context, *InstallCCRequest) (*ServerStatus, error)
	InstantiateCC(context.Context, *InstantiateCCRequest) (*InstantiateCCResponse, error)
	UpgradeCC(context.Context, *UpgradeCCRequest) (*UpgradeCCResponse, error)
	InvokeCC(context.Context, *InvokeCCRequest) (*InvokeCCResponse, error)
	QueryCC(context.Context, *QueryCCRequest) (*QueryCCResponse, error)
}

// UnimplementedChaincodeServer can be embedded to have forward compatible implementations.
type UnimplementedChaincodeServer struct {
}

func (*UnimplementedChaincodeServer) InstallCC(ctx context.Context, req *InstallCCRequest) (*ServerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallCC not implemented")
}
func (*UnimplementedChaincodeServer) InstantiateCC(ctx context.Context, req *InstantiateCCRequest) (*InstantiateCCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstantiateCC not implemented")
}
func (*UnimplementedChaincodeServer) UpgradeCC(ctx context.Context, req *UpgradeCCRequest) (*UpgradeCCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeCC not implemented")
}
func (*UnimplementedChaincodeServer) InvokeCC(ctx context.Context, req *InvokeCCRequest) (*InvokeCCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeCC not implemented")
}
func (*UnimplementedChaincodeServer) QueryCC(ctx context.Context, req *QueryCCRequest) (*QueryCCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCC not implemented")
}

func RegisterChaincodeServer(s *grpc.Server, srv ChaincodeServer) {
	s.RegisterService(&_Chaincode_serviceDesc, srv)
}

func _Chaincode_InstallCC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).InstallCC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chaincode/InstallCC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).InstallCC(ctx, req.(*InstallCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_InstantiateCC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstantiateCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).InstantiateCC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chaincode/InstantiateCC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).InstantiateCC(ctx, req.(*InstantiateCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_UpgradeCC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).UpgradeCC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chaincode/UpgradeCC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).UpgradeCC(ctx, req.(*UpgradeCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_InvokeCC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).InvokeCC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chaincode/InvokeCC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).InvokeCC(ctx, req.(*InvokeCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chaincode_QueryCC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeServer).QueryCC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chaincode/QueryCC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeServer).QueryCC(ctx, req.(*QueryCCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chaincode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Chaincode",
	HandlerType: (*ChaincodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstallCC",
			Handler:    _Chaincode_InstallCC_Handler,
		},
		{
			MethodName: "InstantiateCC",
			Handler:    _Chaincode_InstantiateCC_Handler,
		},
		{
			MethodName: "UpgradeCC",
			Handler:    _Chaincode_UpgradeCC_Handler,
		},
		{
			MethodName: "InvokeCC",
			Handler:    _Chaincode_InvokeCC_Handler,
		},
		{
			MethodName: "QueryCC",
			Handler:    _Chaincode_QueryCC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chaincode.proto",
}
