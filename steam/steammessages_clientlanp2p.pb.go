// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.0
// source: steammessages_clientlanp2p.proto

package steam

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CMsgClientLANP2PRequestChunks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkKeys []*CMsgClientLANP2PRequestChunks_ChunkKey `protobuf:"bytes,1,rep,name=chunk_keys,json=chunkKeys" json:"chunk_keys,omitempty"`
}

func (x *CMsgClientLANP2PRequestChunks) Reset() {
	*x = CMsgClientLANP2PRequestChunks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientlanp2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLANP2PRequestChunks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLANP2PRequestChunks) ProtoMessage() {}

func (x *CMsgClientLANP2PRequestChunks) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientlanp2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLANP2PRequestChunks.ProtoReflect.Descriptor instead.
func (*CMsgClientLANP2PRequestChunks) Descriptor() ([]byte, []int) {
	return file_steammessages_clientlanp2p_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgClientLANP2PRequestChunks) GetChunkKeys() []*CMsgClientLANP2PRequestChunks_ChunkKey {
	if x != nil {
		return x.ChunkKeys
	}
	return nil
}

type CMsgClientLANP2PRequestChunksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkResponses []*CMsgClientLANP2PRequestChunksResponse_ChunkData `protobuf:"bytes,1,rep,name=chunk_responses,json=chunkResponses" json:"chunk_responses,omitempty"`
}

func (x *CMsgClientLANP2PRequestChunksResponse) Reset() {
	*x = CMsgClientLANP2PRequestChunksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientlanp2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLANP2PRequestChunksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLANP2PRequestChunksResponse) ProtoMessage() {}

func (x *CMsgClientLANP2PRequestChunksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientlanp2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLANP2PRequestChunksResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientLANP2PRequestChunksResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_clientlanp2p_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgClientLANP2PRequestChunksResponse) GetChunkResponses() []*CMsgClientLANP2PRequestChunksResponse_ChunkData {
	if x != nil {
		return x.ChunkResponses
	}
	return nil
}

type CMsgClientLANP2PRequestChunks_ChunkKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepotId *uint32 `protobuf:"varint,1,opt,name=depot_id,json=depotId" json:"depot_id,omitempty"`
	Sha     []byte  `protobuf:"bytes,2,opt,name=sha" json:"sha,omitempty"`
}

func (x *CMsgClientLANP2PRequestChunks_ChunkKey) Reset() {
	*x = CMsgClientLANP2PRequestChunks_ChunkKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientlanp2p_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLANP2PRequestChunks_ChunkKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLANP2PRequestChunks_ChunkKey) ProtoMessage() {}

func (x *CMsgClientLANP2PRequestChunks_ChunkKey) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientlanp2p_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLANP2PRequestChunks_ChunkKey.ProtoReflect.Descriptor instead.
func (*CMsgClientLANP2PRequestChunks_ChunkKey) Descriptor() ([]byte, []int) {
	return file_steammessages_clientlanp2p_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CMsgClientLANP2PRequestChunks_ChunkKey) GetDepotId() uint32 {
	if x != nil && x.DepotId != nil {
		return *x.DepotId
	}
	return 0
}

func (x *CMsgClientLANP2PRequestChunks_ChunkKey) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

type CMsgClientLANP2PRequestChunksResponse_ChunkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     *uint32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	DepotId    *uint32 `protobuf:"varint,2,opt,name=depot_id,json=depotId" json:"depot_id,omitempty"`
	Sha        []byte  `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	ChunkData  []byte  `protobuf:"bytes,4,opt,name=chunk_data,json=chunkData" json:"chunk_data,omitempty"`
	Encrypted  *bool   `protobuf:"varint,5,opt,name=encrypted" json:"encrypted,omitempty"`
	Compressed *bool   `protobuf:"varint,6,opt,name=compressed" json:"compressed,omitempty"`
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) Reset() {
	*x = CMsgClientLANP2PRequestChunksResponse_ChunkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientlanp2p_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLANP2PRequestChunksResponse_ChunkData) ProtoMessage() {}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientlanp2p_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLANP2PRequestChunksResponse_ChunkData.ProtoReflect.Descriptor instead.
func (*CMsgClientLANP2PRequestChunksResponse_ChunkData) Descriptor() ([]byte, []int) {
	return file_steammessages_clientlanp2p_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) GetDepotId() uint32 {
	if x != nil && x.DepotId != nil {
		return *x.DepotId
	}
	return 0
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) GetEncrypted() bool {
	if x != nil && x.Encrypted != nil {
		return *x.Encrypted
	}
	return false
}

func (x *CMsgClientLANP2PRequestChunksResponse_ChunkData) GetCompressed() bool {
	if x != nil && x.Compressed != nil {
		return *x.Compressed
	}
	return false
}

var File_steammessages_clientlanp2p_proto protoreflect.FileDescriptor

var file_steammessages_clientlanp2p_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x6c, 0x61, 0x6e, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a,
	0x1d, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x41, 0x4e, 0x50, 0x32,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x46,
	0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x41, 0x4e, 0x50, 0x32, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x37, 0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x68, 0x61, 0x22,
	0xb2, 0x02, 0x0a, 0x25, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x41,
	0x4e, 0x50, 0x32, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x41, 0x4e, 0x50, 0x32, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x1a, 0xad, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65,
	0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65,
	0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x42, 0x35, 0x48, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x62, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x3b, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x80, 0x01, 0x00,
}

var (
	file_steammessages_clientlanp2p_proto_rawDescOnce sync.Once
	file_steammessages_clientlanp2p_proto_rawDescData = file_steammessages_clientlanp2p_proto_rawDesc
)

func file_steammessages_clientlanp2p_proto_rawDescGZIP() []byte {
	file_steammessages_clientlanp2p_proto_rawDescOnce.Do(func() {
		file_steammessages_clientlanp2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_clientlanp2p_proto_rawDescData)
	})
	return file_steammessages_clientlanp2p_proto_rawDescData
}

var file_steammessages_clientlanp2p_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_steammessages_clientlanp2p_proto_goTypes = []interface{}{
	(*CMsgClientLANP2PRequestChunks)(nil),                   // 0: CMsgClientLANP2PRequestChunks
	(*CMsgClientLANP2PRequestChunksResponse)(nil),           // 1: CMsgClientLANP2PRequestChunksResponse
	(*CMsgClientLANP2PRequestChunks_ChunkKey)(nil),          // 2: CMsgClientLANP2PRequestChunks.ChunkKey
	(*CMsgClientLANP2PRequestChunksResponse_ChunkData)(nil), // 3: CMsgClientLANP2PRequestChunksResponse.ChunkData
}
var file_steammessages_clientlanp2p_proto_depIdxs = []int32{
	2, // 0: CMsgClientLANP2PRequestChunks.chunk_keys:type_name -> CMsgClientLANP2PRequestChunks.ChunkKey
	3, // 1: CMsgClientLANP2PRequestChunksResponse.chunk_responses:type_name -> CMsgClientLANP2PRequestChunksResponse.ChunkData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_steammessages_clientlanp2p_proto_init() }
func file_steammessages_clientlanp2p_proto_init() {
	if File_steammessages_clientlanp2p_proto != nil {
		return
	}
	file_steammessages_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_clientlanp2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLANP2PRequestChunks); i {
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
		file_steammessages_clientlanp2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLANP2PRequestChunksResponse); i {
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
		file_steammessages_clientlanp2p_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLANP2PRequestChunks_ChunkKey); i {
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
		file_steammessages_clientlanp2p_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLANP2PRequestChunksResponse_ChunkData); i {
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
			RawDescriptor: file_steammessages_clientlanp2p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steammessages_clientlanp2p_proto_goTypes,
		DependencyIndexes: file_steammessages_clientlanp2p_proto_depIdxs,
		MessageInfos:      file_steammessages_clientlanp2p_proto_msgTypes,
	}.Build()
	File_steammessages_clientlanp2p_proto = out.File
	file_steammessages_clientlanp2p_proto_rawDesc = nil
	file_steammessages_clientlanp2p_proto_goTypes = nil
	file_steammessages_clientlanp2p_proto_depIdxs = nil
}
