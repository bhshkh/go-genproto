// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/display_keyword_view.proto

package resources

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

// A display keyword view.
type DisplayKeywordView struct {
	// Output only. The resource name of the display keyword view.
	// Display Keyword view resource names have the form:
	//
	// `customers/{customer_id}/displayKeywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayKeywordView) Reset()         { *m = DisplayKeywordView{} }
func (m *DisplayKeywordView) String() string { return proto.CompactTextString(m) }
func (*DisplayKeywordView) ProtoMessage()    {}
func (*DisplayKeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b0a6220f800cacf, []int{0}
}

func (m *DisplayKeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayKeywordView.Unmarshal(m, b)
}
func (m *DisplayKeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayKeywordView.Marshal(b, m, deterministic)
}
func (m *DisplayKeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayKeywordView.Merge(m, src)
}
func (m *DisplayKeywordView) XXX_Size() int {
	return xxx_messageInfo_DisplayKeywordView.Size(m)
}
func (m *DisplayKeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayKeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayKeywordView proto.InternalMessageInfo

func (m *DisplayKeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*DisplayKeywordView)(nil), "google.ads.googleads.v3.resources.DisplayKeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/display_keyword_view.proto", fileDescriptor_5b0a6220f800cacf)
}

var fileDescriptor_5b0a6220f800cacf = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x1c, 0xc4, 0x49, 0x0a, 0x1f, 0x7c, 0x41, 0x2f, 0xb9, 0xa8, 0x45, 0xd0, 0x0a, 0x05, 0x41, 0xdc,
	0x3d, 0xe4, 0xb6, 0x0a, 0xb2, 0x45, 0x28, 0x28, 0x48, 0xe9, 0x21, 0x88, 0x04, 0xc2, 0x36, 0xbb,
	0xc6, 0xc5, 0x24, 0xff, 0x98, 0x4d, 0x53, 0x4a, 0xe9, 0xcd, 0x27, 0xf1, 0xe8, 0xa3, 0xf8, 0x14,
	0x7a, 0xed, 0x23, 0x78, 0x92, 0x76, 0xbb, 0x69, 0xa1, 0xa2, 0x78, 0x1b, 0xf8, 0xff, 0x66, 0x76,
	0x32, 0xc4, 0x39, 0x8f, 0x01, 0xe2, 0x44, 0x60, 0xc6, 0x15, 0xd6, 0x72, 0xae, 0x2a, 0x0f, 0x17,
	0x42, 0xc1, 0xb0, 0x88, 0x84, 0xc2, 0x5c, 0xaa, 0x3c, 0x61, 0xe3, 0xf0, 0x51, 0x8c, 0x47, 0x50,
	0xf0, 0xb0, 0x92, 0x62, 0x84, 0xf2, 0x02, 0x4a, 0x70, 0x5b, 0xda, 0x82, 0x18, 0x57, 0xa8, 0x76,
	0xa3, 0xca, 0x43, 0xb5, 0xbb, 0x79, 0x60, 0x1e, 0xc8, 0x25, 0xbe, 0x97, 0x22, 0xe1, 0xe1, 0x40,
	0x3c, 0xb0, 0x4a, 0x42, 0xa1, 0x33, 0x9a, 0x7b, 0x6b, 0x80, 0xb1, 0x2d, 0x4f, 0xfb, 0x6b, 0x27,
	0x96, 0x65, 0x50, 0xb2, 0x52, 0x42, 0xa6, 0xf4, 0xf5, 0xe8, 0xc3, 0x72, 0xdc, 0x4b, 0xdd, 0xed,
	0x5a, 0x57, 0xf3, 0xa5, 0x18, 0xb9, 0xb7, 0xce, 0xb6, 0x89, 0x09, 0x33, 0x96, 0x8a, 0x5d, 0xeb,
	0xd0, 0x3a, 0xfe, 0xdf, 0xf1, 0xde, 0x69, 0xe3, 0x93, 0x9e, 0x3a, 0x27, 0xab, 0x9e, 0x4b, 0x95,
	0x4b, 0x85, 0x22, 0x48, 0xf1, 0x66, 0x56, 0x7f, 0xcb, 0x24, 0xdd, 0xb0, 0x54, 0x90, 0xa7, 0x19,
	0xcd, 0xfe, 0xe4, 0x77, 0x2f, 0xa2, 0xa1, 0x2a, 0x21, 0x15, 0x85, 0xc2, 0x13, 0x23, 0xa7, 0x66,
	0xd0, 0x35, 0x50, 0xe1, 0xc9, 0x77, 0x2b, 0x4f, 0x3b, 0xcf, 0xb6, 0xd3, 0x8e, 0x20, 0x45, 0xbf,
	0xee, 0xdc, 0xd9, 0xd9, 0x7c, 0xbe, 0x37, 0x9f, 0xa9, 0x67, 0xdd, 0x5d, 0x2d, 0xdd, 0x31, 0x24,
	0x2c, 0x8b, 0x11, 0x14, 0x31, 0x8e, 0x45, 0xb6, 0x18, 0x11, 0xaf, 0xbe, 0xe1, 0x87, 0x5f, 0xe0,
	0xac, 0x56, 0x2f, 0x76, 0xa3, 0x4b, 0xe9, 0xab, 0xdd, 0xea, 0xea, 0x48, 0xca, 0x15, 0xd2, 0x72,
	0xae, 0x7c, 0x0f, 0xf5, 0x0d, 0xf9, 0x66, 0x98, 0x80, 0x72, 0x15, 0xd4, 0x4c, 0xe0, 0x7b, 0x41,
	0xcd, 0xcc, 0xec, 0xb6, 0x3e, 0x10, 0x42, 0xb9, 0x22, 0xa4, 0xa6, 0x08, 0xf1, 0x3d, 0x42, 0x6a,
	0x6e, 0xf0, 0x6f, 0x51, 0xd6, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xa4, 0x89, 0xd6, 0xae,
	0x02, 0x00, 0x00,
}