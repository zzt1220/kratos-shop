// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cart/v1/cart.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CartInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	GoodsId    int64  `protobuf:"varint,3,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	GoodsSn    string `protobuf:"bytes,4,opt,name=goodsSn,proto3" json:"goodsSn,omitempty"`
	GoodsName  string `protobuf:"bytes,5,opt,name=goodsName,proto3" json:"goodsName,omitempty"`
	SkuId      int64  `protobuf:"varint,6,opt,name=skuId,proto3" json:"skuId,omitempty"`
	GoodsPrice int64  `protobuf:"varint,7,opt,name=goodsPrice,proto3" json:"goodsPrice,omitempty"`
	GoodsNum   int32  `protobuf:"varint,8,opt,name=goodsNum,proto3" json:"goodsNum,omitempty"`
	IsSelect   bool   `protobuf:"varint,9,opt,name=isSelect,proto3" json:"isSelect,omitempty"`
}

func (x *CartInfoReply) Reset() {
	*x = CartInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartInfoReply) ProtoMessage() {}

func (x *CartInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartInfoReply.ProtoReflect.Descriptor instead.
func (*CartInfoReply) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartInfoReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartInfoReply) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CartInfoReply) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *CartInfoReply) GetGoodsSn() string {
	if x != nil {
		return x.GoodsSn
	}
	return ""
}

func (x *CartInfoReply) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *CartInfoReply) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CartInfoReply) GetGoodsPrice() int64 {
	if x != nil {
		return x.GoodsPrice
	}
	return 0
}

func (x *CartInfoReply) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *CartInfoReply) GetIsSelect() bool {
	if x != nil {
		return x.IsSelect
	}
	return false
}

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	GoodsId    int64  `protobuf:"varint,3,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	GoodsSn    string `protobuf:"bytes,4,opt,name=goodsSn,proto3" json:"goodsSn,omitempty"`
	GoodsName  string `protobuf:"bytes,5,opt,name=goodsName,proto3" json:"goodsName,omitempty"`
	SkuId      int64  `protobuf:"varint,6,opt,name=skuId,proto3" json:"skuId,omitempty"`
	GoodsPrice int64  `protobuf:"varint,7,opt,name=goodsPrice,proto3" json:"goodsPrice,omitempty"`
	GoodsNum   int32  `protobuf:"varint,8,opt,name=goodsNum,proto3" json:"goodsNum,omitempty"`
	IsSelect   bool   `protobuf:"varint,9,opt,name=isSelect,proto3" json:"isSelect,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCartRequest) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *CreateCartRequest) GetGoodsSn() string {
	if x != nil {
		return x.GoodsSn
	}
	return ""
}

func (x *CreateCartRequest) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *CreateCartRequest) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CreateCartRequest) GetGoodsPrice() int64 {
	if x != nil {
		return x.GoodsPrice
	}
	return 0
}

func (x *CreateCartRequest) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *CreateCartRequest) GetIsSelect() bool {
	if x != nil {
		return x.IsSelect
	}
	return false
}

type UpdateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GoodsNum int32 `protobuf:"varint,2,opt,name=goodsNum,proto3" json:"goodsNum,omitempty"`
}

func (x *UpdateCartRequest) Reset() {
	*x = UpdateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartRequest) ProtoMessage() {}

func (x *UpdateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartRequest.ProtoReflect.Descriptor instead.
func (*UpdateCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCartRequest) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

type UpdateCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCartReply) Reset() {
	*x = UpdateCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartReply) ProtoMessage() {}

func (x *UpdateCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartReply.ProtoReflect.Descriptor instead.
func (*UpdateCartReply) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{3}
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{4}
}

func (x *CheckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCartRequest) Reset() {
	*x = DeleteCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartRequest) ProtoMessage() {}

func (x *DeleteCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{5}
}

type DeleteCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCartReply) Reset() {
	*x = DeleteCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartReply) ProtoMessage() {}

func (x *DeleteCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartReply.ProtoReflect.Descriptor instead.
func (*DeleteCartReply) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{6}
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{7}
}

type GetCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCartReply) Reset() {
	*x = GetCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReply) ProtoMessage() {}

func (x *GetCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReply.ProtoReflect.Descriptor instead.
func (*GetCartReply) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{8}
}

type ListCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ListCartRequest) Reset() {
	*x = ListCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartRequest) ProtoMessage() {}

func (x *ListCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartRequest.ProtoReflect.Descriptor instead.
func (*ListCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{9}
}

func (x *ListCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CartListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*CartInfoReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *CartListReply) Reset() {
	*x = CartListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_v1_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartListReply) ProtoMessage() {}

func (x *CartListReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_v1_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartListReply.ProtoReflect.Descriptor instead.
func (*CartListReply) Descriptor() ([]byte, []int) {
	return file_cart_v1_cart_proto_rawDescGZIP(), []int{10}
}

func (x *CartListReply) GetResults() []*CartInfoReply {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_cart_v1_cart_proto protoreflect.FileDescriptor

var file_cart_v1_cart_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x53, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x53, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x4e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x22, 0xc3, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x53, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x6e, 0x12, 0x25, 0x0a,
	0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x08,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75,
	0x6d, 0x12, 0x23, 0x0a, 0x08, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69, 0x73,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x22, 0x3f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x0a, 0x0d, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x29, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x43, 0x61,
	0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x8a, 0x02,
	0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x63, 0x61,
	0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_v1_cart_proto_rawDescOnce sync.Once
	file_cart_v1_cart_proto_rawDescData = file_cart_v1_cart_proto_rawDesc
)

func file_cart_v1_cart_proto_rawDescGZIP() []byte {
	file_cart_v1_cart_proto_rawDescOnce.Do(func() {
		file_cart_v1_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_v1_cart_proto_rawDescData)
	})
	return file_cart_v1_cart_proto_rawDescData
}

var file_cart_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cart_v1_cart_proto_goTypes = []interface{}{
	(*CartInfoReply)(nil),     // 0: cart.v1.CartInfoReply
	(*CreateCartRequest)(nil), // 1: cart.v1.CreateCartRequest
	(*UpdateCartRequest)(nil), // 2: cart.v1.UpdateCartRequest
	(*UpdateCartReply)(nil),   // 3: cart.v1.UpdateCartReply
	(*CheckResponse)(nil),     // 4: cart.v1.CheckResponse
	(*DeleteCartRequest)(nil), // 5: cart.v1.DeleteCartRequest
	(*DeleteCartReply)(nil),   // 6: cart.v1.DeleteCartReply
	(*GetCartRequest)(nil),    // 7: cart.v1.GetCartRequest
	(*GetCartReply)(nil),      // 8: cart.v1.GetCartReply
	(*ListCartRequest)(nil),   // 9: cart.v1.ListCartRequest
	(*CartListReply)(nil),     // 10: cart.v1.CartListReply
}
var file_cart_v1_cart_proto_depIdxs = []int32{
	0,  // 0: cart.v1.CartListReply.results:type_name -> cart.v1.CartInfoReply
	1,  // 1: cart.v1.Cart.CreateCart:input_type -> cart.v1.CreateCartRequest
	2,  // 2: cart.v1.Cart.UpdateCart:input_type -> cart.v1.UpdateCartRequest
	5,  // 3: cart.v1.Cart.DeleteCart:input_type -> cart.v1.DeleteCartRequest
	9,  // 4: cart.v1.Cart.ListCart:input_type -> cart.v1.ListCartRequest
	0,  // 5: cart.v1.Cart.CreateCart:output_type -> cart.v1.CartInfoReply
	4,  // 6: cart.v1.Cart.UpdateCart:output_type -> cart.v1.CheckResponse
	4,  // 7: cart.v1.Cart.DeleteCart:output_type -> cart.v1.CheckResponse
	10, // 8: cart.v1.Cart.ListCart:output_type -> cart.v1.CartListReply
	5,  // [5:9] is the sub-list for method output_type
	1,  // [1:5] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_cart_v1_cart_proto_init() }
func file_cart_v1_cart_proto_init() {
	if File_cart_v1_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_v1_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartInfoReply); i {
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
		file_cart_v1_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartRequest); i {
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
		file_cart_v1_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartRequest); i {
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
		file_cart_v1_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartReply); i {
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
		file_cart_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_cart_v1_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartRequest); i {
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
		file_cart_v1_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartReply); i {
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
		file_cart_v1_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
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
		file_cart_v1_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReply); i {
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
		file_cart_v1_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartRequest); i {
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
		file_cart_v1_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartListReply); i {
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
			RawDescriptor: file_cart_v1_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_v1_cart_proto_goTypes,
		DependencyIndexes: file_cart_v1_cart_proto_depIdxs,
		MessageInfos:      file_cart_v1_cart_proto_msgTypes,
	}.Build()
	File_cart_v1_cart_proto = out.File
	file_cart_v1_cart_proto_rawDesc = nil
	file_cart_v1_cart_proto_goTypes = nil
	file_cart_v1_cart_proto_depIdxs = nil
}
