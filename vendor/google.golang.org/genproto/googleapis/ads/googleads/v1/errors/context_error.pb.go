// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/context_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible context errors.
type ContextErrorEnum_ContextError int32

const (
	// Enum unspecified.
	ContextErrorEnum_UNSPECIFIED ContextErrorEnum_ContextError = 0
	// The received error code is not known in this version.
	ContextErrorEnum_UNKNOWN ContextErrorEnum_ContextError = 1
	// The operation is not allowed for the given context.
	ContextErrorEnum_OPERATION_NOT_PERMITTED_FOR_CONTEXT ContextErrorEnum_ContextError = 2
	// The operation is not allowed for removed resources.
	ContextErrorEnum_OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE ContextErrorEnum_ContextError = 3
)

var ContextErrorEnum_ContextError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPERATION_NOT_PERMITTED_FOR_CONTEXT",
	3: "OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE",
}
var ContextErrorEnum_ContextError_value = map[string]int32{
	"UNSPECIFIED":                         0,
	"UNKNOWN":                             1,
	"OPERATION_NOT_PERMITTED_FOR_CONTEXT": 2,
	"OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE": 3,
}

func (x ContextErrorEnum_ContextError) String() string {
	return proto.EnumName(ContextErrorEnum_ContextError_name, int32(x))
}
func (ContextErrorEnum_ContextError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_context_error_947690d5590518f2, []int{0, 0}
}

// Container for enum describing possible context errors.
type ContextErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContextErrorEnum) Reset()         { *m = ContextErrorEnum{} }
func (m *ContextErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ContextErrorEnum) ProtoMessage()    {}
func (*ContextErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_error_947690d5590518f2, []int{0}
}
func (m *ContextErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContextErrorEnum.Unmarshal(m, b)
}
func (m *ContextErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContextErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ContextErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContextErrorEnum.Merge(dst, src)
}
func (m *ContextErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ContextErrorEnum.Size(m)
}
func (m *ContextErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ContextErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ContextErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ContextErrorEnum)(nil), "google.ads.googleads.v1.errors.ContextErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.ContextErrorEnum_ContextError", ContextErrorEnum_ContextError_name, ContextErrorEnum_ContextError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/context_error.proto", fileDescriptor_context_error_947690d5590518f2)
}

var fileDescriptor_context_error_947690d5590518f2 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x4d, 0x0a, 0x0a, 0x53, 0xc1, 0x98, 0xa5, 0x48, 0x17, 0x71, 0xe1, 0x46, 0x26, 0x46,
	0x77, 0xe3, 0x2a, 0x4d, 0xa6, 0x25, 0x48, 0x67, 0x42, 0x9a, 0x46, 0x91, 0x40, 0x88, 0x4d, 0x18,
	0x0a, 0xed, 0x4c, 0xc9, 0xc4, 0xe2, 0x0d, 0xbc, 0x84, 0x2b, 0x97, 0x1e, 0xc5, 0xa3, 0x88, 0x87,
	0x90, 0x64, 0x6c, 0xe8, 0xc6, 0xae, 0xf2, 0xf2, 0xf3, 0xbd, 0x37, 0xff, 0xfb, 0xc1, 0x0d, 0x13,
	0x82, 0x2d, 0x4b, 0x3b, 0x2f, 0xa4, 0xad, 0x64, 0xa3, 0x36, 0x8e, 0x5d, 0x56, 0x95, 0xa8, 0xa4,
	0x3d, 0x17, 0xbc, 0x2e, 0x5f, 0xeb, 0xac, 0xfd, 0x85, 0xeb, 0x4a, 0xd4, 0xc2, 0x1c, 0x28, 0x10,
	0xe6, 0x85, 0x84, 0x9d, 0x07, 0x6e, 0x1c, 0xa8, 0x3c, 0x67, 0xe7, 0xdb, 0xcc, 0xf5, 0xc2, 0xce,
	0x39, 0x17, 0x75, 0x5e, 0x2f, 0x04, 0x97, 0xca, 0x6d, 0xbd, 0x6b, 0xc0, 0xf0, 0x54, 0x2a, 0x6e,
	0x78, 0xcc, 0x5f, 0x56, 0xd6, 0x9b, 0x06, 0x8e, 0x77, 0x87, 0xe6, 0x09, 0xe8, 0xcf, 0xc8, 0x34,
	0xc4, 0x5e, 0x30, 0x0a, 0xb0, 0x6f, 0x1c, 0x98, 0x7d, 0x70, 0x34, 0x23, 0xf7, 0x84, 0x3e, 0x10,
	0x43, 0x33, 0x2f, 0xc1, 0x05, 0x0d, 0x71, 0xe4, 0xc6, 0x01, 0x25, 0x19, 0xa1, 0x71, 0x16, 0xe2,
	0x68, 0x12, 0xc4, 0x31, 0xf6, 0xb3, 0x11, 0x8d, 0x32, 0x8f, 0x92, 0x18, 0x3f, 0xc6, 0x86, 0x6e,
	0x5e, 0x83, 0xab, 0x7d, 0x60, 0x84, 0x27, 0x34, 0xc1, 0x7e, 0x16, 0xe1, 0x29, 0x9d, 0x45, 0x1e,
	0x36, 0x7a, 0xc3, 0x1f, 0x0d, 0x58, 0x73, 0xb1, 0x82, 0xfb, 0x3b, 0x0e, 0x4f, 0x77, 0xb7, 0x0d,
	0x9b, 0x62, 0xa1, 0xf6, 0xe4, 0xff, 0x99, 0x98, 0x58, 0xe6, 0x9c, 0x41, 0x51, 0x31, 0x9b, 0x95,
	0xbc, 0xad, 0xbd, 0x3d, 0xee, 0x7a, 0x21, 0xff, 0xbb, 0xf5, 0x9d, 0xfa, 0x7c, 0xe8, 0xbd, 0xb1,
	0xeb, 0x7e, 0xea, 0x83, 0xb1, 0x0a, 0x73, 0x0b, 0x09, 0x95, 0x6c, 0x54, 0xe2, 0xc0, 0xf6, 0x49,
	0xf9, 0xb5, 0x05, 0x52, 0xb7, 0x90, 0x69, 0x07, 0xa4, 0x89, 0x93, 0x2a, 0xe0, 0x5b, 0xb7, 0xd4,
	0x14, 0x21, 0xb7, 0x90, 0x08, 0x75, 0x08, 0x42, 0x89, 0x83, 0x90, 0x82, 0x9e, 0x0f, 0xdb, 0xed,
	0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa0, 0xf3, 0x92, 0x08, 0x02, 0x00, 0x00,
}
