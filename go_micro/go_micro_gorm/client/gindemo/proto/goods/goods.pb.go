// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: proto/goods.proto

package goods

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

type GoodsModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle    string  `protobuf:"bytes,2,opt,name=subTitle,proto3" json:"subTitle,omitempty"`
	GoodsSn     string  `protobuf:"bytes,3,opt,name=goodsSn,proto3" json:"goodsSn,omitempty"`
	CateId      int32   `protobuf:"varint,4,opt,name=cateId,proto3" json:"cateId,omitempty"`
	ClickCount  int64   `protobuf:"varint,5,opt,name=clickCount,proto3" json:"clickCount,omitempty"`
	GoodsNumber int32   `protobuf:"varint,6,opt,name=goodsNumber,proto3" json:"goodsNumber,omitempty"`
	Price       float64 `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	Sort        int64   `protobuf:"varint,8,opt,name=sort,proto3" json:"sort,omitempty"`
	Status      int32   `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	AddTime     int64   `protobuf:"varint,10,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *GoodsModel) Reset() {
	*x = GoodsModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsModel) ProtoMessage() {}

func (x *GoodsModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsModel.ProtoReflect.Descriptor instead.
func (*GoodsModel) Descriptor() ([]byte, []int) {
	return file_proto_goods_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GoodsModel) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *GoodsModel) GetGoodsSn() string {
	if x != nil {
		return x.GoodsSn
	}
	return ""
}

func (x *GoodsModel) GetCateId() int32 {
	if x != nil {
		return x.CateId
	}
	return 0
}

func (x *GoodsModel) GetClickCount() int64 {
	if x != nil {
		return x.ClickCount
	}
	return 0
}

func (x *GoodsModel) GetGoodsNumber() int32 {
	if x != nil {
		return x.GoodsNumber
	}
	return 0
}

func (x *GoodsModel) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodsModel) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *GoodsModel) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GoodsModel) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_proto_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetGoodsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsList []*GoodsModel `protobuf:"bytes,1,rep,name=GoodsList,proto3" json:"GoodsList,omitempty"`
}

func (x *GetGoodsResponse) Reset() {
	*x = GetGoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsResponse) ProtoMessage() {}

func (x *GetGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsResponse) Descriptor() ([]byte, []int) {
	return file_proto_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GetGoodsResponse) GetGoodsList() []*GoodsModel {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

var File_proto_goods_proto protoreflect.FileDescriptor

var file_proto_goods_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x53, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x53, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x43,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0x46, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_goods_proto_rawDescOnce sync.Once
	file_proto_goods_proto_rawDescData = file_proto_goods_proto_rawDesc
)

func file_proto_goods_proto_rawDescGZIP() []byte {
	file_proto_goods_proto_rawDescOnce.Do(func() {
		file_proto_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_goods_proto_rawDescData)
	})
	return file_proto_goods_proto_rawDescData
}

var file_proto_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_goods_proto_goTypes = []interface{}{
	(*GoodsModel)(nil),       // 0: goods.GoodsModel
	(*GetGoodsRequest)(nil),  // 1: goods.GetGoodsRequest
	(*GetGoodsResponse)(nil), // 2: goods.GetGoodsResponse
}
var file_proto_goods_proto_depIdxs = []int32{
	0, // 0: goods.GetGoodsResponse.GoodsList:type_name -> goods.GoodsModel
	1, // 1: goods.Goods.GetGoods:input_type -> goods.GetGoodsRequest
	2, // 2: goods.Goods.GetGoods:output_type -> goods.GetGoodsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_goods_proto_init() }
func file_proto_goods_proto_init() {
	if File_proto_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsModel); i {
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
		file_proto_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_proto_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsResponse); i {
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
			RawDescriptor: file_proto_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_goods_proto_goTypes,
		DependencyIndexes: file_proto_goods_proto_depIdxs,
		MessageInfos:      file_proto_goods_proto_msgTypes,
	}.Build()
	File_proto_goods_proto = out.File
	file_proto_goods_proto_rawDesc = nil
	file_proto_goods_proto_goTypes = nil
	file_proto_goods_proto_depIdxs = nil
}
