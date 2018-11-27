// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/stu3/profile_config.proto

package google_fhir_proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PackageInfo struct {
	// Base url to use for all profiles defined here.
	// e.g., g.co/fhir/profiles
	BaseUrl              string   `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	ProtoPackage         string   `protobuf:"bytes,2,opt,name=proto_package,json=protoPackage,proto3" json:"proto_package,omitempty"`
	JavaProtoPackage     string   `protobuf:"bytes,3,opt,name=java_proto_package,json=javaProtoPackage,proto3" json:"java_proto_package,omitempty"`
	GoProtoPackage       string   `protobuf:"bytes,4,opt,name=go_proto_package,json=goProtoPackage,proto3" json:"go_proto_package,omitempty"`
	Publisher            string   `protobuf:"bytes,5,opt,name=publisher,proto3" json:"publisher,omitempty"`
	FhirVersion          string   `protobuf:"bytes,6,opt,name=fhir_version,json=fhirVersion,proto3" json:"fhir_version,omitempty"`
	TelcomUrl            string   `protobuf:"bytes,7,opt,name=telcom_url,json=telcomUrl,proto3" json:"telcom_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageInfo) Reset()         { *m = PackageInfo{} }
func (m *PackageInfo) String() string { return proto.CompactTextString(m) }
func (*PackageInfo) ProtoMessage()    {}
func (*PackageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f6d9e43a8e9ed3, []int{0}
}

func (m *PackageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageInfo.Unmarshal(m, b)
}
func (m *PackageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageInfo.Marshal(b, m, deterministic)
}
func (m *PackageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageInfo.Merge(m, src)
}
func (m *PackageInfo) XXX_Size() int {
	return xxx_messageInfo_PackageInfo.Size(m)
}
func (m *PackageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PackageInfo proto.InternalMessageInfo

func (m *PackageInfo) GetBaseUrl() string {
	if m != nil {
		return m.BaseUrl
	}
	return ""
}

func (m *PackageInfo) GetProtoPackage() string {
	if m != nil {
		return m.ProtoPackage
	}
	return ""
}

func (m *PackageInfo) GetJavaProtoPackage() string {
	if m != nil {
		return m.JavaProtoPackage
	}
	return ""
}

func (m *PackageInfo) GetGoProtoPackage() string {
	if m != nil {
		return m.GoProtoPackage
	}
	return ""
}

func (m *PackageInfo) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *PackageInfo) GetFhirVersion() string {
	if m != nil {
		return m.FhirVersion
	}
	return ""
}

func (m *PackageInfo) GetTelcomUrl() string {
	if m != nil {
		return m.TelcomUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*PackageInfo)(nil), "google.fhir.proto.PackageInfo")
}

func init() { proto.RegisterFile("proto/stu3/profile_config.proto", fileDescriptor_e7f6d9e43a8e9ed3) }

var fileDescriptor_e7f6d9e43a8e9ed3 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0xaa, 0xad, 0xbd, 0xad, 0x52, 0x03, 0xc2, 0x14, 0x14, 0xff, 0x36, 0x5d, 0xc8,
	0x74, 0xd1, 0x37, 0x70, 0xe7, 0x6e, 0x10, 0xea, 0x36, 0x64, 0x86, 0x4c, 0x1a, 0x4d, 0xe7, 0x86,
	0x9b, 0x99, 0xbe, 0x81, 0xef, 0x2d, 0xb9, 0x11, 0x74, 0xec, 0x2e, 0xf9, 0xce, 0x97, 0x03, 0x27,
	0x70, 0xe7, 0x09, 0x3b, 0x5c, 0x87, 0xae, 0xdf, 0xac, 0x3d, 0x61, 0x63, 0x9d, 0x96, 0x35, 0xb6,
	0x8d, 0x35, 0x05, 0x27, 0xe2, 0xca, 0x20, 0x1a, 0xa7, 0x8b, 0x66, 0x67, 0x29, 0xa1, 0xc7, 0xaf,
	0x11, 0xcc, 0x4a, 0x55, 0x7f, 0x2a, 0xa3, 0x5f, 0xdb, 0x06, 0xc5, 0x12, 0xce, 0x2b, 0x15, 0xb4,
	0xec, 0xc9, 0xe5, 0xd9, 0x7d, 0xb6, 0x9a, 0xbe, 0x4d, 0xe2, 0x7d, 0x4b, 0x4e, 0x3c, 0xc1, 0x05,
	0xbf, 0x91, 0x3e, 0xf9, 0xf9, 0x88, 0xf3, 0x39, 0xc3, 0x9f, 0x0e, 0xf1, 0x0c, 0xe2, 0x43, 0x1d,
	0x94, 0x1c, 0x9a, 0x27, 0x6c, 0x2e, 0x62, 0x52, 0xfe, 0xb5, 0x57, 0xb0, 0x30, 0xf8, 0xcf, 0x3d,
	0x65, 0xf7, 0xd2, 0xe0, 0xc0, 0xbc, 0x81, 0xa9, 0xef, 0x2b, 0x67, 0xc3, 0x4e, 0x53, 0x7e, 0xc6,
	0xca, 0x2f, 0x10, 0x0f, 0x30, 0x8f, 0x9b, 0xe4, 0x41, 0x53, 0xb0, 0xd8, 0xe6, 0x63, 0x16, 0x66,
	0x91, 0xbd, 0x27, 0x24, 0x6e, 0x01, 0x3a, 0xed, 0x6a, 0xdc, 0xf3, 0xb4, 0x49, 0x6a, 0x48, 0x64,
	0x4b, 0xee, 0x65, 0x09, 0xd7, 0x35, 0xee, 0x8b, 0xa3, 0x0f, 0x2a, 0xb3, 0x6a, 0xcc, 0x87, 0xcd,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xc8, 0x31, 0xf3, 0x5f, 0x01, 0x00, 0x00,
}
