// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/feed_attribute_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible data types for a feed attribute.
type FeedAttributeTypeEnum_FeedAttributeType int32

const (
	// Not specified.
	FeedAttributeTypeEnum_UNSPECIFIED FeedAttributeTypeEnum_FeedAttributeType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedAttributeTypeEnum_UNKNOWN FeedAttributeTypeEnum_FeedAttributeType = 1
	// Int64.
	FeedAttributeTypeEnum_INT64 FeedAttributeTypeEnum_FeedAttributeType = 2
	// Double.
	FeedAttributeTypeEnum_DOUBLE FeedAttributeTypeEnum_FeedAttributeType = 3
	// String.
	FeedAttributeTypeEnum_STRING FeedAttributeTypeEnum_FeedAttributeType = 4
	// Boolean.
	FeedAttributeTypeEnum_BOOLEAN FeedAttributeTypeEnum_FeedAttributeType = 5
	// Url.
	FeedAttributeTypeEnum_URL FeedAttributeTypeEnum_FeedAttributeType = 6
	// Datetime.
	FeedAttributeTypeEnum_DATE_TIME FeedAttributeTypeEnum_FeedAttributeType = 7
	// Int64 list.
	FeedAttributeTypeEnum_INT64_LIST FeedAttributeTypeEnum_FeedAttributeType = 8
	// Double (8 bytes) list.
	FeedAttributeTypeEnum_DOUBLE_LIST FeedAttributeTypeEnum_FeedAttributeType = 9
	// String list.
	FeedAttributeTypeEnum_STRING_LIST FeedAttributeTypeEnum_FeedAttributeType = 10
	// Boolean list.
	FeedAttributeTypeEnum_BOOLEAN_LIST FeedAttributeTypeEnum_FeedAttributeType = 11
	// Url list.
	FeedAttributeTypeEnum_URL_LIST FeedAttributeTypeEnum_FeedAttributeType = 12
	// Datetime list.
	FeedAttributeTypeEnum_DATE_TIME_LIST FeedAttributeTypeEnum_FeedAttributeType = 13
	// Price.
	FeedAttributeTypeEnum_PRICE FeedAttributeTypeEnum_FeedAttributeType = 14
)

var FeedAttributeTypeEnum_FeedAttributeType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INT64",
	3:  "DOUBLE",
	4:  "STRING",
	5:  "BOOLEAN",
	6:  "URL",
	7:  "DATE_TIME",
	8:  "INT64_LIST",
	9:  "DOUBLE_LIST",
	10: "STRING_LIST",
	11: "BOOLEAN_LIST",
	12: "URL_LIST",
	13: "DATE_TIME_LIST",
	14: "PRICE",
}

var FeedAttributeTypeEnum_FeedAttributeType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"INT64":          2,
	"DOUBLE":         3,
	"STRING":         4,
	"BOOLEAN":        5,
	"URL":            6,
	"DATE_TIME":      7,
	"INT64_LIST":     8,
	"DOUBLE_LIST":    9,
	"STRING_LIST":    10,
	"BOOLEAN_LIST":   11,
	"URL_LIST":       12,
	"DATE_TIME_LIST": 13,
	"PRICE":          14,
}

func (x FeedAttributeTypeEnum_FeedAttributeType) String() string {
	return proto.EnumName(FeedAttributeTypeEnum_FeedAttributeType_name, int32(x))
}

func (FeedAttributeTypeEnum_FeedAttributeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ad66d0f530d8c45, []int{0, 0}
}

// Container for enum describing possible data types for a feed attribute.
type FeedAttributeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeTypeEnum) Reset()         { *m = FeedAttributeTypeEnum{} }
func (m *FeedAttributeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeTypeEnum) ProtoMessage()    {}
func (*FeedAttributeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ad66d0f530d8c45, []int{0}
}

func (m *FeedAttributeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeTypeEnum.Unmarshal(m, b)
}
func (m *FeedAttributeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeTypeEnum.Marshal(b, m, deterministic)
}
func (m *FeedAttributeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeTypeEnum.Merge(m, src)
}
func (m *FeedAttributeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeTypeEnum.Size(m)
}
func (m *FeedAttributeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FeedAttributeTypeEnum_FeedAttributeType", FeedAttributeTypeEnum_FeedAttributeType_name, FeedAttributeTypeEnum_FeedAttributeType_value)
	proto.RegisterType((*FeedAttributeTypeEnum)(nil), "google.ads.googleads.v3.enums.FeedAttributeTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/feed_attribute_type.proto", fileDescriptor_7ad66d0f530d8c45)
}

var fileDescriptor_7ad66d0f530d8c45 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x24, 0x0e, 0x4d, 0x9a, 0x97, 0x34, 0x2c, 0x2b, 0xc1, 0x01, 0xd1, 0x43, 0xfb, 0x01, 0xeb,
	0x83, 0x11, 0x48, 0xcb, 0x69, 0xdd, 0x6c, 0x23, 0x0b, 0xe3, 0x58, 0x8e, 0x1d, 0x24, 0x64, 0xc9,
	0x72, 0xf1, 0x62, 0x59, 0x6a, 0xbc, 0x56, 0xd6, 0xa9, 0xd4, 0x3b, 0x5f, 0xc2, 0x91, 0x4f, 0xe1,
	0x53, 0xca, 0x4f, 0xa0, 0xf5, 0xd6, 0xbe, 0x54, 0x70, 0xb1, 0xc6, 0xf3, 0x66, 0x66, 0x9f, 0xe6,
	0xc1, 0x87, 0x52, 0xca, 0xf2, 0x56, 0xd8, 0x79, 0xa1, 0x6c, 0x03, 0x35, 0xba, 0x73, 0x6c, 0x51,
	0x1f, 0xf7, 0xca, 0xfe, 0x2e, 0x44, 0x91, 0xe5, 0x6d, 0x7b, 0xa8, 0x6e, 0x8e, 0xad, 0xc8, 0xda,
	0xfb, 0x46, 0x90, 0xe6, 0x20, 0x5b, 0x89, 0xcf, 0x8d, 0x9a, 0xe4, 0x85, 0x22, 0x83, 0x91, 0xdc,
	0x39, 0xa4, 0x33, 0xbe, 0x79, 0xdb, 0xe7, 0x36, 0x95, 0x9d, 0xd7, 0xb5, 0x6c, 0xf3, 0xb6, 0x92,
	0xb5, 0x32, 0xe6, 0xcb, 0x1f, 0x16, 0xbc, 0xba, 0x16, 0xa2, 0x60, 0x7d, 0x72, 0x7c, 0xdf, 0x08,
	0x5e, 0x1f, 0xf7, 0x97, 0x0f, 0x23, 0x78, 0xf9, 0x64, 0x82, 0x5f, 0xc0, 0x3c, 0x09, 0xb6, 0x21,
	0xbf, 0xf2, 0xae, 0x3d, 0xbe, 0x42, 0xcf, 0xf0, 0x1c, 0xa6, 0x49, 0xf0, 0x29, 0xd8, 0x7c, 0x09,
	0xd0, 0x08, 0xcf, 0xe0, 0xc4, 0x0b, 0xe2, 0xf7, 0xef, 0x90, 0x85, 0x01, 0x26, 0xab, 0x4d, 0xe2,
	0xfa, 0x1c, 0x8d, 0x35, 0xde, 0xc6, 0x91, 0x17, 0xac, 0xd1, 0x73, 0xad, 0x77, 0x37, 0x1b, 0x9f,
	0xb3, 0x00, 0x9d, 0xe0, 0x29, 0x8c, 0x93, 0xc8, 0x47, 0x13, 0x7c, 0x06, 0xb3, 0x15, 0x8b, 0x79,
	0x16, 0x7b, 0x9f, 0x39, 0x9a, 0xe2, 0x25, 0x40, 0x97, 0x93, 0xf9, 0xde, 0x36, 0x46, 0xa7, 0xfa,
	0x55, 0x13, 0x66, 0x88, 0x99, 0x26, 0x4c, 0xa2, 0x21, 0x00, 0x23, 0x58, 0x3c, 0xc6, 0x1a, 0x66,
	0x8e, 0x17, 0x70, 0x9a, 0x44, 0xbe, 0xf9, 0x5b, 0x60, 0x0c, 0xcb, 0xe1, 0x01, 0xc3, 0x9d, 0xe9,
	0x6d, 0xc3, 0xc8, 0xbb, 0xe2, 0x68, 0xe9, 0xfe, 0x19, 0xc1, 0xc5, 0x37, 0xb9, 0x27, 0xff, 0xad,
	0xd2, 0x7d, 0xfd, 0xa4, 0x8f, 0x50, 0x97, 0x18, 0x8e, 0xbe, 0xba, 0x8f, 0xc6, 0x52, 0xde, 0xe6,
	0x75, 0x49, 0xe4, 0xa1, 0xb4, 0x4b, 0x51, 0x77, 0x15, 0xf7, 0xc7, 0x6c, 0x2a, 0xf5, 0x8f, 0xdb,
	0x7e, 0xec, 0xbe, 0x3f, 0xad, 0xf1, 0x9a, 0xb1, 0x5f, 0xd6, 0xf9, 0xda, 0x44, 0xb1, 0x42, 0x11,
	0x03, 0x35, 0xda, 0x39, 0x44, 0x5f, 0x45, 0xfd, 0xee, 0xe7, 0x29, 0x2b, 0x54, 0x3a, 0xcc, 0xd3,
	0x9d, 0x93, 0x76, 0xf3, 0x07, 0xeb, 0xc2, 0x90, 0x94, 0xb2, 0x42, 0x51, 0x3a, 0x28, 0x28, 0xdd,
	0x39, 0x94, 0x76, 0x9a, 0x9b, 0x49, 0xb7, 0x98, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x52, 0x2c,
	0x81, 0xd4, 0x73, 0x02, 0x00, 0x00,
}