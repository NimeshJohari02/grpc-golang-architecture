// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0-devel
// 	protoc        v4.25.1
// source: train.proto

package train

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

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	EmailAddress string `protobuf:"bytes,4,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	HasTicket    bool   `protobuf:"varint,5,opt,name=hasTicket,proto3" json:"hasTicket,omitempty"`
	// Optional field to store the ticket
	Ticket *Ticket `protobuf:"bytes,6,opt,name=ticket,proto3,oneof" json:"ticket,omitempty"`
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

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
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

func (x *User) GetHasTicket() bool {
	if x != nil {
		return x.HasTicket
	}
	return false
}

func (x *User) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From          string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	UserId        string  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PriceOfTicket float32 `protobuf:"fixed32,4,opt,name=price_of_ticket,json=priceOfTicket,proto3" json:"price_of_ticket,omitempty"`
	Section       Section `protobuf:"varint,5,opt,name=section,proto3,enum=Section" json:"section,omitempty"`
	TrainId       string  `protobuf:"bytes,7,opt,name=train_id,json=trainId,proto3" json:"train_id,omitempty"`
	IsConfirmed   string  `protobuf:"bytes,8,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	Seat          string  `protobuf:"bytes,9,opt,name=seat,proto3" json:"seat,omitempty"`
	TicketId      string  `protobuf:"bytes,10,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
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

func (x *Ticket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket) GetPriceOfTicket() float32 {
	if x != nil {
		return x.PriceOfTicket
	}
	return 0
}

func (x *Ticket) GetSection() Section {
	if x != nil {
		return x.Section
	}
	return Section_SECTION_A
}

func (x *Ticket) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

func (x *Ticket) GetIsConfirmed() string {
	if x != nil {
		return x.IsConfirmed
	}
	return ""
}

func (x *Ticket) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

func (x *Ticket) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type PurchaseTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From          string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	UserId        string  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PriceOfTicket float32 `protobuf:"fixed32,4,opt,name=price_of_ticket,json=priceOfTicket,proto3" json:"price_of_ticket,omitempty"`
	Section       Section `protobuf:"varint,5,opt,name=section,proto3,enum=Section" json:"section,omitempty"`
	TrainId       string  `protobuf:"bytes,7,opt,name=train_id,json=trainId,proto3" json:"train_id,omitempty"`
}

func (x *PurchaseTicketRequest) Reset() {
	*x = PurchaseTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseTicketRequest) ProtoMessage() {}

func (x *PurchaseTicketRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PurchaseTicketRequest.ProtoReflect.Descriptor instead.
func (*PurchaseTicketRequest) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseTicketRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseTicketRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseTicketRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PurchaseTicketRequest) GetPriceOfTicket() float32 {
	if x != nil {
		return x.PriceOfTicket
	}
	return 0
}

func (x *PurchaseTicketRequest) GetSection() Section {
	if x != nil {
		return x.Section
	}
	return Section_SECTION_A
}

func (x *PurchaseTicketRequest) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

type ViewReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ViewReceiptRequest) Reset() {
	*x = ViewReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewReceiptRequest) ProtoMessage() {}

func (x *ViewReceiptRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ViewReceiptRequest.ProtoReflect.Descriptor instead.
func (*ViewReceiptRequest) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{3}
}

func (x *ViewReceiptRequest) GetUserId() string {
	if x != nil {
		return x.UserId
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

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
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

func (x *ViewUsersBySectionResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type RemoveUserFromTrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TrainId string `protobuf:"bytes,2,opt,name=train_id,json=trainId,proto3" json:"train_id,omitempty"`
}

func (x *RemoveUserFromTrainRequest) Reset() {
	*x = RemoveUserFromTrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromTrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromTrainRequest) ProtoMessage() {}

func (x *RemoveUserFromTrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromTrainRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserFromTrainRequest) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveUserFromTrainRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveUserFromTrainRequest) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[7]
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
	return file_train_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TrainId    string  `protobuf:"bytes,2,opt,name=train_id,json=trainId,proto3" json:"train_id,omitempty"`
	Section    Section `protobuf:"varint,3,opt,name=section,proto3,enum=Section" json:"section,omitempty"`
	NewSection Section `protobuf:"varint,4,opt,name=new_section,json=newSection,proto3,enum=Section" json:"new_section,omitempty"`
	NewSeat    string  `protobuf:"bytes,5,opt,name=new_seat,json=newSeat,proto3" json:"new_seat,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[8]
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
	return file_train_proto_rawDescGZIP(), []int{8}
}

func (x *ModifySeatRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ModifySeatRequest) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

func (x *ModifySeatRequest) GetSection() Section {
	if x != nil {
		return x.Section
	}
	return Section_SECTION_A
}

func (x *ModifySeatRequest) GetNewSection() Section {
	if x != nil {
		return x.NewSection
	}
	return Section_SECTION_A
}

func (x *ModifySeatRequest) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

type WasModifiedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *WasModifiedResponse) Reset() {
	*x = WasModifiedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasModifiedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasModifiedResponse) ProtoMessage() {}

func (x *WasModifiedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasModifiedResponse.ProtoReflect.Descriptor instead.
func (*WasModifiedResponse) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{9}
}

func (x *WasModifiedResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_train_proto protoreflect.FileDescriptor

var file_train_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24,
	0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x80, 0x02, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x22, 0x2d, 0x0a, 0x12, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3f, 0x0a, 0x19, 0x76, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x39, 0x0a, 0x1a, 0x76, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x1a, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x2e, 0x0a,
	0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xb1, 0x01,
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0b, 0x6e,
	0x65, 0x77, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x61,
	0x74, 0x22, 0x2f, 0x0a, 0x13, 0x77, 0x61, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2a, 0x27, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x10, 0x01, 0x32, 0xc4, 0x02, 0x0a, 0x12,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x12, 0x13, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x4d, 0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x6f, 0x6d, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x77,
	0x61, 0x73, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_train_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_train_proto_goTypes = []interface{}{
	(Section)(0),                       // 0: Section
	(*User)(nil),                       // 1: User
	(*Ticket)(nil),                     // 2: Ticket
	(*PurchaseTicketRequest)(nil),      // 3: purchaseTicketRequest
	(*ViewReceiptRequest)(nil),         // 4: viewReceiptRequest
	(*ViewUsersBySectionRequest)(nil),  // 5: viewUsersBySectionRequest
	(*ViewUsersBySectionResponse)(nil), // 6: viewUsersBySectionResponse
	(*RemoveUserFromTrainRequest)(nil), // 7: removeUserFromTrainRequest
	(*RemoveUserResponse)(nil),         // 8: removeUserResponse
	(*ModifySeatRequest)(nil),          // 9: modifySeatRequest
	(*WasModifiedResponse)(nil),        // 10: wasModifiedResponse
}
var file_train_proto_depIdxs = []int32{
	2,  // 0: User.ticket:type_name -> Ticket
	0,  // 1: Ticket.section:type_name -> Section
	0,  // 2: purchaseTicketRequest.section:type_name -> Section
	0,  // 3: viewUsersBySectionRequest.section:type_name -> Section
	1,  // 4: viewUsersBySectionResponse.users:type_name -> User
	0,  // 5: modifySeatRequest.section:type_name -> Section
	0,  // 6: modifySeatRequest.new_section:type_name -> Section
	3,  // 7: TrainTicketService.purchaseTicket:input_type -> purchaseTicketRequest
	4,  // 8: TrainTicketService.ViewReceipt:input_type -> viewReceiptRequest
	5,  // 9: TrainTicketService.ViewUsersBySection:input_type -> viewUsersBySectionRequest
	7,  // 10: TrainTicketService.RemoveUserFromTrain:input_type -> removeUserFromTrainRequest
	9,  // 11: TrainTicketService.ModifySeat:input_type -> modifySeatRequest
	2,  // 12: TrainTicketService.purchaseTicket:output_type -> Ticket
	2,  // 13: TrainTicketService.ViewReceipt:output_type -> Ticket
	6,  // 14: TrainTicketService.ViewUsersBySection:output_type -> viewUsersBySectionResponse
	8,  // 15: TrainTicketService.RemoveUserFromTrain:output_type -> removeUserResponse
	10, // 16: TrainTicketService.ModifySeat:output_type -> wasModifiedResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
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
			switch v := v.(*PurchaseTicketRequest); i {
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
			switch v := v.(*ViewReceiptRequest); i {
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
		file_train_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromTrainRequest); i {
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
		file_train_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_train_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_train_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasModifiedResponse); i {
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
	file_train_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_train_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
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
