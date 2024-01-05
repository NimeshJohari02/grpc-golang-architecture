// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0-devel
// 	protoc        v4.25.1
// source: train.proto

package train;

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

type Section int32

const (
	Section_SECTION_A Section = 0
	Section_SECTION_B Section = 1
)

// Enum value maps for Section.
var (
	Section_name = map[int32]string{
		0: "SECTION_A",
		1: "SECTION_B",
	}
	Section_value = map[string]int32{
		"SECTION_A": 0,
		"SECTION_B": 1,
	}
)

func (x Section) Enum() *Section {
	p := new(Section)
	*p = x
	return p
}

func (x Section) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Section) Descriptor() protoreflect.EnumDescriptor {
	return file_train_proto_enumTypes[0].Descriptor()
}

func (Section) Type() protoreflect.EnumType {
	return &file_train_proto_enumTypes[0]
}

func (x Section) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Section.Descriptor instead.
func (Section) EnumDescriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	EmailAddress string `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User      *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid float32 `protobuf:"fixed32,4,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	Section   Section `protobuf:"varint,5,opt,name=section,proto3,enum=Section" json:"section,omitempty"`
	Seat      string  `protobuf:"bytes,6,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{1}
}

func (x *Ticket) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Ticket) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Ticket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Ticket) GetPricePaid() float32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *Ticket) GetSection() Section {
	if x != nil {
		return x.Section
	}
	return Section_SECTION_A
}

func (x *Ticket) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Seat string `protobuf:"bytes,2,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{3}
}

func (x *ModifySeatRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ModifySeatRequest) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

type ViewUsersBySectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section Section `protobuf:"varint,1,opt,name=section,proto3,enum=Section" json:"section,omitempty"`
}

func (x *ViewUsersBySectionRequest) Reset() {
	*x = ViewUsersBySectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewUsersBySectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewUsersBySectionRequest) ProtoMessage() {}

func (x *ViewUsersBySectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewUsersBySectionRequest.ProtoReflect.Descriptor instead.
func (*ViewUsersBySectionRequest) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{4}
}

func (x *ViewUsersBySectionRequest) GetSection() Section {
	if x != nil {
		return x.Section
	}
	return Section_SECTION_A
}

type ViewUsersBySectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
}

func (x *ViewUsersBySectionResponse) Reset() {
	*x = ViewUsersBySectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewUsersBySectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewUsersBySectionResponse) ProtoMessage() {}

func (x *ViewUsersBySectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewUsersBySectionResponse.ProtoReflect.Descriptor instead.
func (*ViewUsersBySectionResponse) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{5}
}

func (x *ViewUsersBySectionResponse) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

var File_train_proto protoreflect.FileDescriptor

var file_train_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x42, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x3f, 0x0a, 0x19, 0x56, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x1a, 0x56, 0x69, 0x65, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2a, 0x27, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x10, 0x01,
	0x32, 0xfb, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x07, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x1a, 0x07, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0b, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x07, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x12, 0x56, 0x69,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x13,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61,
	0x74, 0x12, 0x12, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_train_proto_rawDescOnce sync.Once
	file_train_proto_rawDescData = file_train_proto_rawDesc
)

func file_train_proto_rawDescGZIP() []byte {
	file_train_proto_rawDescOnce.Do(func() {
		file_train_proto_rawDescData = protoimpl.X.CompressGZIP(file_train_proto_rawDescData)
	})
	return file_train_proto_rawDescData
}

var file_train_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_train_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_train_proto_goTypes = []interface{}{
	(Section)(0),                       // 0: Section
	(*User)(nil),                       // 1: User
	(*Ticket)(nil),                     // 2: Ticket
	(*RemoveUserResponse)(nil),         // 3: RemoveUserResponse
	(*ModifySeatRequest)(nil),          // 4: ModifySeatRequest
	(*ViewUsersBySectionRequest)(nil),  // 5: ViewUsersBySectionRequest
	(*ViewUsersBySectionResponse)(nil), // 6: ViewUsersBySectionResponse
}
var file_train_proto_depIdxs = []int32{
	1,  // 0: Ticket.user:type_name -> User
	0,  // 1: Ticket.section:type_name -> Section
	1,  // 2: ModifySeatRequest.user:type_name -> User
	0,  // 3: ViewUsersBySectionRequest.section:type_name -> Section
	2,  // 4: ViewUsersBySectionResponse.tickets:type_name -> Ticket
	2,  // 5: TrainTicketService.SubmitPurchase:input_type -> Ticket
	1,  // 6: TrainTicketService.ViewReceipt:input_type -> User
	5,  // 7: TrainTicketService.ViewUsersBySection:input_type -> ViewUsersBySectionRequest
	1,  // 8: TrainTicketService.RemoveUser:input_type -> User
	4,  // 9: TrainTicketService.ModifySeat:input_type -> ModifySeatRequest
	2,  // 10: TrainTicketService.SubmitPurchase:output_type -> Ticket
	2,  // 11: TrainTicketService.ViewReceipt:output_type -> Ticket
	6,  // 12: TrainTicketService.ViewUsersBySection:output_type -> ViewUsersBySectionResponse
	3,  // 13: TrainTicketService.RemoveUser:output_type -> RemoveUserResponse
	2,  // 14: TrainTicketService.ModifySeat:output_type -> Ticket
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_train_proto_init() }
func file_train_proto_init() {
	if File_train_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_train_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_train_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_train_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserResponse); i {
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
		file_train_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySeatRequest); i {
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
		file_train_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewUsersBySectionRequest); i {
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
		file_train_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewUsersBySectionResponse); i {
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
			RawDescriptor: file_train_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_train_proto_goTypes,
		DependencyIndexes: file_train_proto_depIdxs,
		EnumInfos:         file_train_proto_enumTypes,
		MessageInfos:      file_train_proto_msgTypes,
	}.Build()
	File_train_proto = out.File
	file_train_proto_rawDesc = nil
	file_train_proto_goTypes = nil
	file_train_proto_depIdxs = nil
}
