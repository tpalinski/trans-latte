// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: crud.proto

package messages

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

// Message notifying the crud backend about new order
type NewOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date        string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *NewOrder) Reset() {
	*x = NewOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrder) ProtoMessage() {}

func (x *NewOrder) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrder.ProtoReflect.Descriptor instead.
func (*NewOrder) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{0}
}

func (x *NewOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewOrder) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewOrder) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewOrder) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

// Struct representing the state of the order cached in redis
type OrderStatusInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email             string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	ClientDescription string  `protobuf:"bytes,3,opt,name=clientDescription,proto3" json:"clientDescription,omitempty"`
	DateOrdered       string  `protobuf:"bytes,4,opt,name=dateOrdered,proto3" json:"dateOrdered,omitempty"`
	LastUpdated       string  `protobuf:"bytes,5,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	StatusDescription string  `protobuf:"bytes,6,opt,name=statusDescription,proto3" json:"statusDescription,omitempty"`
	Price             *string `protobuf:"bytes,7,opt,name=price,proto3,oneof" json:"price,omitempty"`
}

func (x *OrderStatusInfo) Reset() {
	*x = OrderStatusInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatusInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusInfo) ProtoMessage() {}

func (x *OrderStatusInfo) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatusInfo.ProtoReflect.Descriptor instead.
func (*OrderStatusInfo) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{1}
}

func (x *OrderStatusInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderStatusInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OrderStatusInfo) GetClientDescription() string {
	if x != nil {
		return x.ClientDescription
	}
	return ""
}

func (x *OrderStatusInfo) GetDateOrdered() string {
	if x != nil {
		return x.DateOrdered
	}
	return ""
}

func (x *OrderStatusInfo) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *OrderStatusInfo) GetStatusDescription() string {
	if x != nil {
		return x.StatusDescription
	}
	return ""
}

func (x *OrderStatusInfo) GetPrice() string {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return ""
}

var File_crud_proto protoreflect.FileDescriptor

var file_crud_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xfc,
	0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x5a,
	0x09, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_crud_proto_rawDescOnce sync.Once
	file_crud_proto_rawDescData = file_crud_proto_rawDesc
)

func file_crud_proto_rawDescGZIP() []byte {
	file_crud_proto_rawDescOnce.Do(func() {
		file_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_proto_rawDescData)
	})
	return file_crud_proto_rawDescData
}

var file_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crud_proto_goTypes = []interface{}{
	(*NewOrder)(nil),        // 0: messages.NewOrder
	(*OrderStatusInfo)(nil), // 1: messages.OrderStatusInfo
}
var file_crud_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crud_proto_init() }
func file_crud_proto_init() {
	if File_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOrder); i {
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
		file_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatusInfo); i {
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
	file_crud_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crud_proto_goTypes,
		DependencyIndexes: file_crud_proto_depIdxs,
		MessageInfos:      file_crud_proto_msgTypes,
	}.Build()
	File_crud_proto = out.File
	file_crud_proto_rawDesc = nil
	file_crud_proto_goTypes = nil
	file_crud_proto_depIdxs = nil
}