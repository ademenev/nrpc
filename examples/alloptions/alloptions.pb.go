// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alloptions.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	alloptions.proto

It has these top-level messages:
	StringArg
	SimpleStringReply
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/rapidloop/nrpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StringArg struct {
	Arg1 string `protobuf:"bytes,1,opt,name=arg1" json:"arg1,omitempty"`
}

func (m *StringArg) Reset()                    { *m = StringArg{} }
func (m *StringArg) String() string            { return proto.CompactTextString(m) }
func (*StringArg) ProtoMessage()               {}
func (*StringArg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringArg) GetArg1() string {
	if m != nil {
		return m.Arg1
	}
	return ""
}

type SimpleStringReply struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *SimpleStringReply) Reset()                    { *m = SimpleStringReply{} }
func (m *SimpleStringReply) String() string            { return proto.CompactTextString(m) }
func (*SimpleStringReply) ProtoMessage()               {}
func (*SimpleStringReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SimpleStringReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*StringArg)(nil), "main.StringArg")
	proto.RegisterType((*SimpleStringReply)(nil), "main.SimpleStringReply")
}

func init() { proto.RegisterFile("alloptions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x1c, 0xc6, 0x37, 0x1a, 0x75, 0x3b, 0xba, 0x9b, 0x3a, 0x2b, 0x9a, 0x06, 0x41, 0x09, 0x1e, 0x56,
	0x90, 0x2c, 0x1b, 0x4f, 0x1e, 0xd7, 0x3d, 0x08, 0x42, 0x16, 0x69, 0xc0, 0x3d, 0x96, 0xe9, 0x64,
	0xc8, 0x8e, 0x64, 0x32, 0xb3, 0x93, 0x7f, 0x16, 0xbc, 0xf6, 0xd8, 0x53, 0x8f, 0x3e, 0x47, 0x21,
	0x2f, 0xe0, 0xd1, 0x97, 0xf0, 0xe8, 0x03, 0xf4, 0x05, 0x24, 0x33, 0xb1, 0xb6, 0x85, 0x62, 0x65,
	0x2f, 0xc9, 0xe4, 0x3f, 0xdf, 0x7c, 0xdf, 0x97, 0x5f, 0x08, 0xea, 0x93, 0xa2, 0x90, 0x0a, 0xb8,
	0x2c, 0xab, 0x48, 0x69, 0x09, 0x12, 0xbb, 0x82, 0xf0, 0x32, 0x78, 0x95, 0x73, 0xb8, 0xaa, 0xc7,
	0x11, 0x95, 0xe2, 0x44, 0x13, 0xc5, 0xb3, 0x42, 0x4a, 0x75, 0x52, 0x6a, 0x45, 0xcd, 0xc5, 0x6a,
	0xc3, 0x17, 0xa8, 0x97, 0x82, 0xe6, 0x65, 0x7e, 0xa6, 0x73, 0x8c, 0x91, 0x4b, 0x74, 0x7e, 0xea,
	0x3b, 0x2f, 0x9d, 0xe3, 0xde, 0xd0, 0xac, 0xc3, 0xd7, 0xe8, 0x71, 0xca, 0x85, 0x2a, 0x98, 0x95,
	0x0d, 0x99, 0x2a, 0xbe, 0xe2, 0x27, 0xe8, 0x9e, 0x6e, 0x17, 0x9d, 0xd2, 0x3e, 0xc4, 0xbf, 0xee,
	0xa0, 0x7e, 0x7a, 0x43, 0xcf, 0xeb, 0x0a, 0xa4, 0x48, 0xeb, 0xf1, 0x17, 0x46, 0x01, 0x5f, 0xa0,
	0x83, 0x04, 0xac, 0x83, 0x3d, 0xeb, 0x45, 0x6d, 0xbd, 0x68, 0x99, 0x1a, 0x3c, 0xeb, 0x06, 0x9b,
	0x29, 0xe1, 0xd1, 0x64, 0x3e, 0xf0, 0x04, 0x8c, 0x2a, 0xb3, 0x33, 0x32, 0x21, 0xf8, 0x0d, 0x7a,
	0x98, 0xc0, 0x67, 0xc9, 0xb3, 0x2d, 0x6e, 0x28, 0x32, 0x6f, 0xd7, 0x2a, 0xc2, 0x3d, 0xfc, 0xae,
	0x55, 0x5f, 0xc8, 0x21, 0xbb, 0xae, 0x59, 0x05, 0xd8, 0xb3, 0x9b, 0xcb, 0xc1, 0xf6, 0xec, 0x3d,
	0x7c, 0x86, 0xbc, 0x04, 0x52, 0xd0, 0x8c, 0x08, 0x96, 0xfd, 0x6f, 0x75, 0x77, 0x36, 0x1f, 0x38,
	0xf8, 0x1c, 0x3d, 0xfd, 0xd3, 0xf5, 0x7a, 0xdd, 0x69, 0xa5, 0xe5, 0x3f, 0x4c, 0x02, 0xfc, 0x7d,
	0xe1, 0x1f, 0x52, 0xc3, 0x74, 0x54, 0x59, 0xa8, 0xf1, 0x4f, 0xc7, 0x90, 0xee, 0x18, 0x7f, 0x22,
	0x9a, 0x88, 0x0a, 0x7f, 0x44, 0x47, 0x09, 0x5c, 0x72, 0xb8, 0x5a, 0x1f, 0xef, 0x14, 0x75, 0x38,
	0x9d, 0x0f, 0xee, 0x0a, 0x75, 0x6a, 0x6f, 0x31, 0x3e, 0x46, 0x3d, 0xcb, 0x6d, 0xb3, 0xec, 0xc1,
	0x5f, 0x82, 0x16, 0xd3, 0x07, 0x84, 0x57, 0x08, 0x5f, 0x76, 0xa1, 0xbb, 0x83, 0x7e, 0xd0, 0x25,
	0x07, 0x8f, 0x7e, 0x2c, 0xfc, 0x7d, 0x5a, 0x70, 0x56, 0x02, 0xcf, 0xe2, 0x04, 0xf5, 0x97, 0x87,
	0x53, 0xa6, 0x6f, 0x38, 0x65, 0xb7, 0xf8, 0x98, 0xef, 0x9f, 0x4f, 0x1a, 0xdf, 0xd5, 0x52, 0xc2,
	0xb4, 0xf1, 0xf7, 0x79, 0x59, 0x01, 0x29, 0x29, 0x9b, 0x35, 0xbe, 0xf3, 0xad, 0xf1, 0x9d, 0xf1,
	0x7d, 0xf3, 0x2f, 0xbc, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x6f, 0xc0, 0xe5, 0x4b, 0x03,
	0x00, 0x00,
}
