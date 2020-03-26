// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/legacy_app_install_ad_app_store.proto

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

// App store type in a legacy app install ad.
type LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore int32

const (
	// Not specified.
	LegacyAppInstallAdAppStoreEnum_UNSPECIFIED LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 0
	// Used for return value only. Represents value unknown in this version.
	LegacyAppInstallAdAppStoreEnum_UNKNOWN LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 1
	// Apple iTunes.
	LegacyAppInstallAdAppStoreEnum_APPLE_APP_STORE LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 2
	// Google Play.
	LegacyAppInstallAdAppStoreEnum_GOOGLE_PLAY LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 3
	// Windows Store.
	LegacyAppInstallAdAppStoreEnum_WINDOWS_STORE LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 4
	// Windows Phone Store.
	LegacyAppInstallAdAppStoreEnum_WINDOWS_PHONE_STORE LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 5
	// The app is hosted in a Chinese app store.
	LegacyAppInstallAdAppStoreEnum_CN_APP_STORE LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore = 6
)

var LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "APPLE_APP_STORE",
	3: "GOOGLE_PLAY",
	4: "WINDOWS_STORE",
	5: "WINDOWS_PHONE_STORE",
	6: "CN_APP_STORE",
}

var LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"APPLE_APP_STORE":     2,
	"GOOGLE_PLAY":         3,
	"WINDOWS_STORE":       4,
	"WINDOWS_PHONE_STORE": 5,
	"CN_APP_STORE":        6,
}

func (x LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore) String() string {
	return proto.EnumName(LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore_name, int32(x))
}

func (LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7fbe0fcc71f1681d, []int{0, 0}
}

// Container for enum describing app store type in a legacy app install ad.
type LegacyAppInstallAdAppStoreEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LegacyAppInstallAdAppStoreEnum) Reset()         { *m = LegacyAppInstallAdAppStoreEnum{} }
func (m *LegacyAppInstallAdAppStoreEnum) String() string { return proto.CompactTextString(m) }
func (*LegacyAppInstallAdAppStoreEnum) ProtoMessage()    {}
func (*LegacyAppInstallAdAppStoreEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbe0fcc71f1681d, []int{0}
}

func (m *LegacyAppInstallAdAppStoreEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LegacyAppInstallAdAppStoreEnum.Unmarshal(m, b)
}
func (m *LegacyAppInstallAdAppStoreEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LegacyAppInstallAdAppStoreEnum.Marshal(b, m, deterministic)
}
func (m *LegacyAppInstallAdAppStoreEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyAppInstallAdAppStoreEnum.Merge(m, src)
}
func (m *LegacyAppInstallAdAppStoreEnum) XXX_Size() int {
	return xxx_messageInfo_LegacyAppInstallAdAppStoreEnum.Size(m)
}
func (m *LegacyAppInstallAdAppStoreEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyAppInstallAdAppStoreEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyAppInstallAdAppStoreEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore", LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore_name, LegacyAppInstallAdAppStoreEnum_LegacyAppInstallAdAppStore_value)
	proto.RegisterType((*LegacyAppInstallAdAppStoreEnum)(nil), "google.ads.googleads.v3.enums.LegacyAppInstallAdAppStoreEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/legacy_app_install_ad_app_store.proto", fileDescriptor_7fbe0fcc71f1681d)
}

var fileDescriptor_7fbe0fcc71f1681d = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x6b, 0xa3, 0x40,
	0x18, 0x5d, 0xcd, 0x6e, 0x16, 0x26, 0xbb, 0xc4, 0x35, 0x87, 0x85, 0xb0, 0xd9, 0x92, 0xfc, 0x80,
	0xf1, 0xe0, 0x6d, 0x7a, 0x9a, 0x24, 0xd6, 0x4a, 0x45, 0x87, 0xda, 0x24, 0xb4, 0x08, 0x32, 0x8d,
	0x22, 0x82, 0x99, 0x19, 0x32, 0x26, 0xd0, 0x3f, 0xd3, 0x43, 0x8f, 0xfd, 0x17, 0xbd, 0xf6, 0xa7,
	0xf4, 0xdc, 0x1f, 0x50, 0x1c, 0x63, 0xe8, 0x25, 0xbd, 0xc8, 0xf3, 0x7b, 0xef, 0x7b, 0x4f, 0xdf,
	0x07, 0x66, 0x39, 0xe7, 0x79, 0x99, 0x59, 0x34, 0x95, 0x56, 0x03, 0x6b, 0xb4, 0xb7, 0xad, 0x8c,
	0xed, 0x36, 0xd2, 0x2a, 0xb3, 0x9c, 0xae, 0x1f, 0x12, 0x2a, 0x44, 0x52, 0x30, 0x59, 0xd1, 0xb2,
	0x4c, 0x68, 0xaa, 0x5e, 0x65, 0xc5, 0xb7, 0x19, 0x14, 0x5b, 0x5e, 0x71, 0x73, 0xd4, 0x6c, 0x42,
	0x9a, 0x4a, 0x78, 0x34, 0x81, 0x7b, 0x1b, 0x2a, 0x93, 0xe1, 0xbf, 0x36, 0x43, 0x14, 0x16, 0x65,
	0x8c, 0x57, 0xb4, 0x2a, 0x38, 0x93, 0xcd, 0xf2, 0xe4, 0x45, 0x03, 0xff, 0x7d, 0x15, 0x83, 0x85,
	0xf0, 0x9a, 0x10, 0x9c, 0x62, 0x21, 0xa2, 0x3a, 0xc1, 0x61, 0xbb, 0xcd, 0xe4, 0x51, 0x03, 0xc3,
	0xd3, 0x12, 0xb3, 0x0f, 0x7a, 0x8b, 0x20, 0x22, 0xce, 0xcc, 0xbb, 0xf0, 0x9c, 0xb9, 0xf1, 0xcd,
	0xec, 0x81, 0x9f, 0x8b, 0xe0, 0x2a, 0x08, 0x57, 0x81, 0xa1, 0x99, 0x03, 0xd0, 0xc7, 0x84, 0xf8,
	0x4e, 0x82, 0x09, 0x49, 0xa2, 0x9b, 0xf0, 0xda, 0x31, 0xf4, 0x7a, 0xc5, 0x0d, 0x43, 0xd7, 0x77,
	0x12, 0xe2, 0xe3, 0x5b, 0xa3, 0x63, 0xfe, 0x01, 0xbf, 0x57, 0x5e, 0x30, 0x0f, 0x57, 0xd1, 0x41,
	0xf3, 0xdd, 0xfc, 0x0b, 0x06, 0xed, 0x88, 0x5c, 0x86, 0x81, 0x73, 0x20, 0x7e, 0x98, 0x06, 0xf8,
	0x35, 0x0b, 0x3e, 0xd9, 0x75, 0xa7, 0xef, 0x1a, 0x18, 0xaf, 0xf9, 0x06, 0x7e, 0xd9, 0xc3, 0xf4,
	0xec, 0xf4, 0x3f, 0x90, 0xba, 0x0a, 0xa2, 0xdd, 0x4d, 0x0f, 0x0e, 0x39, 0x2f, 0x29, 0xcb, 0x21,
	0xdf, 0xe6, 0x56, 0x9e, 0x31, 0x55, 0x54, 0x7b, 0x1e, 0x51, 0xc8, 0x13, 0xd7, 0x3a, 0x57, 0xcf,
	0x27, 0xbd, 0xe3, 0x62, 0xfc, 0xac, 0x8f, 0xdc, 0xc6, 0x0a, 0xa7, 0x12, 0x36, 0xb0, 0x46, 0x4b,
	0x1b, 0xd6, 0x95, 0xca, 0xd7, 0x96, 0x8f, 0x71, 0x2a, 0xe3, 0x23, 0x1f, 0x2f, 0xed, 0x58, 0xf1,
	0x6f, 0xfa, 0xb8, 0x19, 0x22, 0x84, 0x53, 0x89, 0xd0, 0x51, 0x81, 0xd0, 0xd2, 0x46, 0x48, 0x69,
	0xee, 0xbb, 0xea, 0xc3, 0xec, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xbf, 0x9c, 0xfa, 0x45,
	0x02, 0x00, 0x00,
}