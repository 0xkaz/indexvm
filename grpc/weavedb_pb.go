package grpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WeaveDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Query   string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Nocache bool   `protobuf:"varint,3,opt,name=nocache,proto3" json:"nocache,omitempty"`
}

func (x *WeaveDBRequest) Reset() {
	*x = WeaveDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weavedb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeaveDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeaveDBRequest) ProtoMessage() {}

func (x *WeaveDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weavedb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeaveDBRequest.ProtoReflect.Descriptor instead.
func (*WeaveDBRequest) Descriptor() ([]byte, []int) {
	return file_weavedb_proto_rawDescGZIP(), []int{0}
}

func (x *WeaveDBRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *WeaveDBRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *WeaveDBRequest) GetNocache() bool {
	if x != nil {
		return x.Nocache
	}
	return false
}

type WeaveDBReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *WeaveDBReply) Reset() {
	*x = WeaveDBReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weavedb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeaveDBReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeaveDBReply) ProtoMessage() {}

func (x *WeaveDBReply) ProtoReflect() protoreflect.Message {
	mi := &file_weavedb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeaveDBReply.ProtoReflect.Descriptor instead.
func (*WeaveDBReply) Descriptor() ([]byte, []int) {
	return file_weavedb_proto_rawDescGZIP(), []int{1}
}

func (x *WeaveDBReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *WeaveDBReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //string  = 2;
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weavedb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weavedb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_weavedb_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weavedb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_weavedb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_weavedb_proto_rawDescGZIP(), []int{3}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weavedb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weavedb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_weavedb_proto_rawDescGZIP(), []int{4}
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weavedb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_weavedb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_weavedb_proto_rawDescGZIP(), []int{5}
}

func (x *PingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_weavedb_proto protoreflect.FileDescriptor

var file_weavedb_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x76, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x77, 0x65, 0x61, 0x76, 0x65, 0x64, 0x62, 0x22, 0x58, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x76,
	0x65, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x57, 0x65, 0x61, 0x76, 0x65, 0x44, 0x42, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x22, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xad,
	0x01, 0x0a, 0x02, 0x44, 0x42, 0x12, 0x39, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17,
	0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x64, 0x62, 0x2e, 0x57, 0x65, 0x61, 0x76, 0x65, 0x44, 0x42,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x64,
	0x62, 0x2e, 0x57, 0x65, 0x61, 0x76, 0x65, 0x44, 0x42, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x08, 0x73, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x2e, 0x77,
	0x65, 0x61, 0x76, 0x65, 0x64, 0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x64, 0x62, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x14, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x64, 0x62, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65,
	0x64, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weavedb_proto_rawDescOnce sync.Once
	file_weavedb_proto_rawDescData = file_weavedb_proto_rawDesc
)

func file_weavedb_proto_rawDescGZIP() []byte {
	file_weavedb_proto_rawDescOnce.Do(func() {
		file_weavedb_proto_rawDescData = protoimpl.X.CompressGZIP(file_weavedb_proto_rawDescData)
	})
	return file_weavedb_proto_rawDescData
}

var file_weavedb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_weavedb_proto_goTypes = []interface{}{
	(*WeaveDBRequest)(nil), // 0: weavedb.WeaveDBRequest
	(*WeaveDBReply)(nil),   // 1: weavedb.WeaveDBReply
	(*HelloRequest)(nil),   // 2: weavedb.HelloRequest
	(*HelloReply)(nil),     // 3: weavedb.HelloReply
	(*PingRequest)(nil),    // 4: weavedb.PingRequest
	(*PingReply)(nil),      // 5: weavedb.PingReply
}
var file_weavedb_proto_depIdxs = []int32{
	0, // 0: weavedb.DB.query:input_type -> weavedb.WeaveDBRequest
	2, // 1: weavedb.DB.sayHello:input_type -> weavedb.HelloRequest
	4, // 2: weavedb.DB.ping:input_type -> weavedb.PingRequest
	1, // 3: weavedb.DB.query:output_type -> weavedb.WeaveDBReply
	3, // 4: weavedb.DB.sayHello:output_type -> weavedb.HelloReply
	5, // 5: weavedb.DB.ping:output_type -> weavedb.PingReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weavedb_proto_init() }
func file_weavedb_proto_init() {
	if File_weavedb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weavedb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeaveDBRequest); i {
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
		file_weavedb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeaveDBReply); i {
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
		file_weavedb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_weavedb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_weavedb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_weavedb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
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
			RawDescriptor: file_weavedb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weavedb_proto_goTypes,
		DependencyIndexes: file_weavedb_proto_depIdxs,
		MessageInfos:      file_weavedb_proto_msgTypes,
	}.Build()
	File_weavedb_proto = out.File
	file_weavedb_proto_rawDesc = nil
	file_weavedb_proto_goTypes = nil
	file_weavedb_proto_depIdxs = nil
}
