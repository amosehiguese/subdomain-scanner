// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: subdomain.proto

package subdomain

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *ApiQueryRequest) Reset() {
	*x = ApiQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiQueryRequest) ProtoMessage() {}

func (x *ApiQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiQueryRequest.ProtoReflect.Descriptor instead.
func (*ApiQueryRequest) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{0}
}

func (x *ApiQueryRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type ApiQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subdomains []string `protobuf:"bytes,1,rep,name=subdomains,proto3" json:"subdomains,omitempty"`
}

func (x *ApiQueryResponse) Reset() {
	*x = ApiQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiQueryResponse) ProtoMessage() {}

func (x *ApiQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiQueryResponse.ProtoReflect.Descriptor instead.
func (*ApiQueryResponse) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{1}
}

func (x *ApiQueryResponse) GetSubdomains() []string {
	if x != nil {
		return x.Subdomains
	}
	return nil
}

type BruteForceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *BruteForceRequest) Reset() {
	*x = BruteForceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BruteForceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BruteForceRequest) ProtoMessage() {}

func (x *BruteForceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BruteForceRequest.ProtoReflect.Descriptor instead.
func (*BruteForceRequest) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{2}
}

func (x *BruteForceRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type BruteForceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subdomains []string `protobuf:"bytes,1,rep,name=subdomains,proto3" json:"subdomains,omitempty"`
}

func (x *BruteForceResponse) Reset() {
	*x = BruteForceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BruteForceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BruteForceResponse) ProtoMessage() {}

func (x *BruteForceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BruteForceResponse.ProtoReflect.Descriptor instead.
func (*BruteForceResponse) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{3}
}

func (x *BruteForceResponse) GetSubdomains() []string {
	if x != nil {
		return x.Subdomains
	}
	return nil
}

type ResolveDnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *ResolveDnsRequest) Reset() {
	*x = ResolveDnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveDnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveDnsRequest) ProtoMessage() {}

func (x *ResolveDnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveDnsRequest.ProtoReflect.Descriptor instead.
func (*ResolveDnsRequest) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{4}
}

func (x *ResolveDnsRequest) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type ResolveDnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subdomain []string `protobuf:"bytes,1,rep,name=subdomain,proto3" json:"subdomain,omitempty"`
}

func (x *ResolveDnsResponse) Reset() {
	*x = ResolveDnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveDnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveDnsResponse) ProtoMessage() {}

func (x *ResolveDnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveDnsResponse.ProtoReflect.Descriptor instead.
func (*ResolveDnsResponse) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{5}
}

func (x *ResolveDnsResponse) GetSubdomain() []string {
	if x != nil {
		return x.Subdomain
	}
	return nil
}

type PortScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *PortScanRequest) Reset() {
	*x = PortScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortScanRequest) ProtoMessage() {}

func (x *PortScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortScanRequest.ProtoReflect.Descriptor instead.
func (*PortScanRequest) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{6}
}

func (x *PortScanRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Subdomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string  `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ports  []*Port `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *Subdomain) Reset() {
	*x = Subdomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subdomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subdomain) ProtoMessage() {}

func (x *Subdomain) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subdomain.ProtoReflect.Descriptor instead.
func (*Subdomain) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{7}
}

func (x *Subdomain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Subdomain) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnOpen bool   `protobuf:"varint,1,opt,name=connOpen,proto3" json:"connOpen,omitempty"`
	Port     uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subdomain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_subdomain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_subdomain_proto_rawDescGZIP(), []int{8}
}

func (x *Port) GetConnOpen() bool {
	if x != nil {
		return x.ConnOpen
	}
	return false
}

func (x *Port) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_subdomain_proto protoreflect.FileDescriptor

var file_subdomain_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x29, 0x0a, 0x0f,
	0x41, 0x70, 0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x32, 0x0a, 0x10, 0x41, 0x70, 0x69, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x42,
	0x72, 0x75, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x34, 0x0a, 0x12, 0x42, 0x72, 0x75, 0x74,
	0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x29,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x25, 0x0a,
	0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x22, 0x36, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x6e,
	0x4f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x6e,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x67, 0x0a, 0x0f, 0x41, 0x70, 0x69, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x70,
	0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41,
	0x70, 0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x6a, 0x0a, 0x0c, 0x42, 0x72, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x42, 0x79, 0x42, 0x72, 0x75, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x1c,
	0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x72, 0x75, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x72, 0x75, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x60, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x6e, 0x73,
	0x12, 0x1c, 0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32,
	0x5d, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x53, 0x63, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53,
	0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subdomain_proto_rawDescOnce sync.Once
	file_subdomain_proto_rawDescData = file_subdomain_proto_rawDesc
)

func file_subdomain_proto_rawDescGZIP() []byte {
	file_subdomain_proto_rawDescOnce.Do(func() {
		file_subdomain_proto_rawDescData = protoimpl.X.CompressGZIP(file_subdomain_proto_rawDescData)
	})
	return file_subdomain_proto_rawDescData
}

var file_subdomain_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_subdomain_proto_goTypes = []interface{}{
	(*ApiQueryRequest)(nil),    // 0: subdomain.ApiQueryRequest
	(*ApiQueryResponse)(nil),   // 1: subdomain.ApiQueryResponse
	(*BruteForceRequest)(nil),  // 2: subdomain.BruteForceRequest
	(*BruteForceResponse)(nil), // 3: subdomain.BruteForceResponse
	(*ResolveDnsRequest)(nil),  // 4: subdomain.ResolveDnsRequest
	(*ResolveDnsResponse)(nil), // 5: subdomain.ResolveDnsResponse
	(*PortScanRequest)(nil),    // 6: subdomain.PortScanRequest
	(*Subdomain)(nil),          // 7: subdomain.Subdomain
	(*Port)(nil),               // 8: subdomain.Port
}
var file_subdomain_proto_depIdxs = []int32{
	8, // 0: subdomain.Subdomain.ports:type_name -> subdomain.Port
	0, // 1: subdomain.ApiQueryService.GetSubdomainsByApiQuery:input_type -> subdomain.ApiQueryRequest
	2, // 2: subdomain.BruteService.GetSubdomainsByBruteForce:input_type -> subdomain.BruteForceRequest
	4, // 3: subdomain.ResolveDnsService.ResolveDns:input_type -> subdomain.ResolveDnsRequest
	6, // 4: subdomain.PortScanService.ScanForOpenPorts:input_type -> subdomain.PortScanRequest
	1, // 5: subdomain.ApiQueryService.GetSubdomainsByApiQuery:output_type -> subdomain.ApiQueryResponse
	3, // 6: subdomain.BruteService.GetSubdomainsByBruteForce:output_type -> subdomain.BruteForceResponse
	5, // 7: subdomain.ResolveDnsService.ResolveDns:output_type -> subdomain.ResolveDnsResponse
	7, // 8: subdomain.PortScanService.ScanForOpenPorts:output_type -> subdomain.Subdomain
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_subdomain_proto_init() }
func file_subdomain_proto_init() {
	if File_subdomain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subdomain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiQueryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BruteForceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BruteForceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveDnsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveDnsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortScanRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subdomain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subdomain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subdomain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_subdomain_proto_goTypes,
		DependencyIndexes: file_subdomain_proto_depIdxs,
		MessageInfos:      file_subdomain_proto_msgTypes,
	}.Build()
	File_subdomain_proto = out.File
	file_subdomain_proto_rawDesc = nil
	file_subdomain_proto_goTypes = nil
	file_subdomain_proto_depIdxs = nil
}
