// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: counter.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetListOrderFulfillmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListOrderFulfillmentRequest) Reset() {
	*x = GetListOrderFulfillmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOrderFulfillmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOrderFulfillmentRequest) ProtoMessage() {}

func (x *GetListOrderFulfillmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOrderFulfillmentRequest.ProtoReflect.Descriptor instead.
func (*GetListOrderFulfillmentRequest) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{0}
}

type GetListOrderFulfillmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderDto `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetListOrderFulfillmentResponse) Reset() {
	*x = GetListOrderFulfillmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOrderFulfillmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOrderFulfillmentResponse) ProtoMessage() {}

func (x *GetListOrderFulfillmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOrderFulfillmentResponse.ProtoReflect.Descriptor instead.
func (*GetListOrderFulfillmentResponse) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{1}
}

func (x *GetListOrderFulfillmentResponse) GetOrders() []*OrderDto {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderSource     int32          `protobuf:"varint,2,opt,name=order_source,json=orderSource,proto3" json:"order_source,omitempty"`
	LoyaltyMemberId string         `protobuf:"bytes,3,opt,name=loyalty_member_id,json=loyaltyMemberId,proto3" json:"loyalty_member_id,omitempty"`
	OrderStatus     int32          `protobuf:"varint,4,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
	Location        int32          `protobuf:"varint,5,opt,name=location,proto3" json:"location,omitempty"`
	LineItems       []*LineItemDto `protobuf:"bytes,6,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
}

func (x *OrderDto) Reset() {
	*x = OrderDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDto) ProtoMessage() {}

func (x *OrderDto) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDto.ProtoReflect.Descriptor instead.
func (*OrderDto) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{2}
}

func (x *OrderDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderDto) GetOrderSource() int32 {
	if x != nil {
		return x.OrderSource
	}
	return 0
}

func (x *OrderDto) GetLoyaltyMemberId() string {
	if x != nil {
		return x.LoyaltyMemberId
	}
	return ""
}

func (x *OrderDto) GetOrderStatus() int32 {
	if x != nil {
		return x.OrderStatus
	}
	return 0
}

func (x *OrderDto) GetLocation() int32 {
	if x != nil {
		return x.Location
	}
	return 0
}

func (x *OrderDto) GetLineItems() []*LineItemDto {
	if x != nil {
		return x.LineItems
	}
	return nil
}

type LineItemDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemType       int32   `protobuf:"varint,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	Name           string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price          float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	ItemStatus     int32   `protobuf:"varint,5,opt,name=item_status,json=itemStatus,proto3" json:"item_status,omitempty"`
	IsBaristaOrder bool    `protobuf:"varint,6,opt,name=is_barista_order,json=isBaristaOrder,proto3" json:"is_barista_order,omitempty"`
}

func (x *LineItemDto) Reset() {
	*x = LineItemDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItemDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemDto) ProtoMessage() {}

func (x *LineItemDto) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemDto.ProtoReflect.Descriptor instead.
func (*LineItemDto) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{3}
}

func (x *LineItemDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LineItemDto) GetItemType() int32 {
	if x != nil {
		return x.ItemType
	}
	return 0
}

func (x *LineItemDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LineItemDto) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *LineItemDto) GetItemStatus() int32 {
	if x != nil {
		return x.ItemStatus
	}
	return 0
}

func (x *LineItemDto) GetIsBaristaOrder() bool {
	if x != nil {
		return x.IsBaristaOrder
	}
	return false
}

type PlaceOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandType     int32                  `protobuf:"varint,1,opt,name=command_type,json=commandType,proto3" json:"command_type,omitempty"`
	OrderSource     int32                  `protobuf:"varint,2,opt,name=order_source,json=orderSource,proto3" json:"order_source,omitempty"`
	Location        int32                  `protobuf:"varint,3,opt,name=location,proto3" json:"location,omitempty"`
	LoyaltyMemberId string                 `protobuf:"bytes,4,opt,name=loyalty_member_id,json=loyaltyMemberId,proto3" json:"loyalty_member_id,omitempty"`
	BaristaItems    []*CommandItem         `protobuf:"bytes,5,rep,name=barista_items,json=baristaItems,proto3" json:"barista_items,omitempty"`
	KitchenItems    []*CommandItem         `protobuf:"bytes,6,rep,name=kitchen_items,json=kitchenItems,proto3" json:"kitchen_items,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PlaceOrderRequest) Reset() {
	*x = PlaceOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderRequest) ProtoMessage() {}

func (x *PlaceOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderRequest.ProtoReflect.Descriptor instead.
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceOrderRequest) GetCommandType() int32 {
	if x != nil {
		return x.CommandType
	}
	return 0
}

func (x *PlaceOrderRequest) GetOrderSource() int32 {
	if x != nil {
		return x.OrderSource
	}
	return 0
}

func (x *PlaceOrderRequest) GetLocation() int32 {
	if x != nil {
		return x.Location
	}
	return 0
}

func (x *PlaceOrderRequest) GetLoyaltyMemberId() string {
	if x != nil {
		return x.LoyaltyMemberId
	}
	return ""
}

func (x *PlaceOrderRequest) GetBaristaItems() []*CommandItem {
	if x != nil {
		return x.BaristaItems
	}
	return nil
}

func (x *PlaceOrderRequest) GetKitchenItems() []*CommandItem {
	if x != nil {
		return x.KitchenItems
	}
	return nil
}

func (x *PlaceOrderRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type PlaceOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlaceOrderResponse) Reset() {
	*x = PlaceOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResponse) ProtoMessage() {}

func (x *PlaceOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResponse.ProtoReflect.Descriptor instead.
func (*PlaceOrderResponse) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{5}
}

type CommandItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemType int32 `protobuf:"varint,1,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
}

func (x *CommandItem) Reset() {
	*x = CommandItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandItem) ProtoMessage() {}

func (x *CommandItem) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandItem.ProtoReflect.Descriptor instead.
func (*CommandItem) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{6}
}

func (x *CommandItem) GetItemType() int32 {
	if x != nil {
		return x.ItemType
	}
	return 0
}

var File_counter_proto protoreflect.FileDescriptor

var file_counter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31,
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x44, 0x74, 0x6f, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69,
	0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f,
	0x62, 0x61, 0x72, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x42, 0x61, 0x72, 0x69, 0x73, 0x74, 0x61, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0xcd, 0x02, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x6c,
	0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0d, 0x62, 0x61, 0x72, 0x69, 0x73,
	0x74, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x69, 0x73, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x37, 0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x6b, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x32, 0xfe, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd2, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x68, 0x92, 0x41, 0x47, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x66, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x25, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x20, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x96, 0x01,
	0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x53, 0x92, 0x41, 0x37, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0e,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1d,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x6f,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x8c, 0x01, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x6f, 0x61, 0x6e, 0x67, 0x63, 0x6e, 0x39, 0x35, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x69, 0x6d, 0x69, 0x74, 0x61, 0x74, 0x65,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counter_proto_rawDescOnce sync.Once
	file_counter_proto_rawDescData = file_counter_proto_rawDesc
)

func file_counter_proto_rawDescGZIP() []byte {
	file_counter_proto_rawDescOnce.Do(func() {
		file_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_counter_proto_rawDescData)
	})
	return file_counter_proto_rawDescData
}

var file_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_counter_proto_goTypes = []interface{}{
	(*GetListOrderFulfillmentRequest)(nil),  // 0: proto.GetListOrderFulfillmentRequest
	(*GetListOrderFulfillmentResponse)(nil), // 1: proto.GetListOrderFulfillmentResponse
	(*OrderDto)(nil),                        // 2: proto.OrderDto
	(*LineItemDto)(nil),                     // 3: proto.LineItemDto
	(*PlaceOrderRequest)(nil),               // 4: proto.PlaceOrderRequest
	(*PlaceOrderResponse)(nil),              // 5: proto.PlaceOrderResponse
	(*CommandItem)(nil),                     // 6: proto.CommandItem
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
}
var file_counter_proto_depIdxs = []int32{
	2, // 0: proto.GetListOrderFulfillmentResponse.orders:type_name -> proto.OrderDto
	3, // 1: proto.OrderDto.line_items:type_name -> proto.LineItemDto
	6, // 2: proto.PlaceOrderRequest.barista_items:type_name -> proto.CommandItem
	6, // 3: proto.PlaceOrderRequest.kitchen_items:type_name -> proto.CommandItem
	7, // 4: proto.PlaceOrderRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 5: proto.CounterService.GetListOrderFulfillment:input_type -> proto.GetListOrderFulfillmentRequest
	4, // 6: proto.CounterService.PlaceOrder:input_type -> proto.PlaceOrderRequest
	1, // 7: proto.CounterService.GetListOrderFulfillment:output_type -> proto.GetListOrderFulfillmentResponse
	5, // 8: proto.CounterService.PlaceOrder:output_type -> proto.PlaceOrderResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_counter_proto_init() }
func file_counter_proto_init() {
	if File_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOrderFulfillmentRequest); i {
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
		file_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOrderFulfillmentResponse); i {
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
		file_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDto); i {
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
		file_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItemDto); i {
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
		file_counter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderRequest); i {
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
		file_counter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderResponse); i {
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
		file_counter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandItem); i {
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
			RawDescriptor: file_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counter_proto_goTypes,
		DependencyIndexes: file_counter_proto_depIdxs,
		MessageInfos:      file_counter_proto_msgTypes,
	}.Build()
	File_counter_proto = out.File
	file_counter_proto_rawDesc = nil
	file_counter_proto_goTypes = nil
	file_counter_proto_depIdxs = nil
}
