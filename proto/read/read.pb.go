// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/provisioning-service/proto/read/read.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_service_provisioning_read is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/provisioning-service/proto/read/read.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_service_provisioning_read

import proto "github.com/HailoOSS/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Request struct {
	ServiceName      *string `protobuf:"bytes,1,req,name=serviceName" json:"serviceName,omitempty"`
	ServiceVersion   *uint64 `protobuf:"varint,2,req,name=serviceVersion" json:"serviceVersion,omitempty"`
	MachineClass     *string `protobuf:"bytes,3,req,name=machineClass" json:"machineClass,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetServiceVersion() uint64 {
	if m != nil && m.ServiceVersion != nil {
		return *m.ServiceVersion
	}
	return 0
}

func (m *Request) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

type Response struct {
	ServiceName      *string `protobuf:"bytes,1,req,name=serviceName" json:"serviceName,omitempty"`
	ServiceVersion   *uint64 `protobuf:"varint,2,req,name=serviceVersion" json:"serviceVersion,omitempty"`
	MachineClass     *string `protobuf:"bytes,3,req,name=machineClass" json:"machineClass,omitempty"`
	NoFileSoftLimit  *uint64 `protobuf:"varint,4,opt,name=noFileSoftLimit" json:"noFileSoftLimit,omitempty"`
	NoFileHardLimit  *uint64 `protobuf:"varint,5,opt,name=noFileHardLimit" json:"noFileHardLimit,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Response) GetServiceVersion() uint64 {
	if m != nil && m.ServiceVersion != nil {
		return *m.ServiceVersion
	}
	return 0
}

func (m *Response) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

func (m *Response) GetNoFileSoftLimit() uint64 {
	if m != nil && m.NoFileSoftLimit != nil {
		return *m.NoFileSoftLimit
	}
	return 0
}

func (m *Response) GetNoFileHardLimit() uint64 {
	if m != nil && m.NoFileHardLimit != nil {
		return *m.NoFileHardLimit
	}
	return 0
}

func init() {
}
