// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Taks we have to do
type ToDo struct {
	// Unique integer identifier of the todo task
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the task
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Detail description of the todo task
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Date and time to remind the todo task
	Reminder             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ToDo) Reset()         { *m = ToDo{} }
func (m *ToDo) String() string { return proto.CompactTextString(m) }
func (*ToDo) ProtoMessage()    {}
func (*ToDo) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{0}
}

func (m *ToDo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDo.Unmarshal(m, b)
}
func (m *ToDo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDo.Marshal(b, m, deterministic)
}
func (m *ToDo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDo.Merge(m, src)
}
func (m *ToDo) XXX_Size() int {
	return xxx_messageInfo_ToDo.Size(m)
}
func (m *ToDo) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDo.DiscardUnknown(m)
}

var xxx_messageInfo_ToDo proto.InternalMessageInfo

func (m *ToDo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDo) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// Request data to create new todo task
type CreateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to add
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains data of created todo task
type CreateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created task
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request data to read todo task
type ReadRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains todo task data specified in by ID request
type ReadResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity read by ID
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Request data to update todo task
type UpdateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to update
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains status of update operation
type UpdateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed updated
	// Equals 1 in case of succesfull update
	Updated              int64    `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// Request data to delete todo task
type DeleteRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task to delete
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains status of delete operation
type DeleteResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have beed deleted
	// Equals 1 in case of succesfull delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// Request data to read all todo task
type ReadAllRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{9}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all todo tasks
type ReadAllResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// List of all todo tasks
	ToDos                []*ToDo  `protobuf:"bytes,2,rep,name=toDos,proto3" json:"toDos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b701c7b1c502fe, []int{10}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetToDos() []*ToDo {
	if m != nil {
		return m.ToDos
	}
	return nil
}

func init() {
	proto.RegisterType((*ToDo)(nil), "v1.ToDo")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v1.DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "v1.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "v1.ReadAllResponse")
}

func init() {
	proto.RegisterFile("todo-service.proto", fileDescriptor_80b701c7b1c502fe)
}

var fileDescriptor_80b701c7b1c502fe = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x7e, 0x36, 0xbb, 0x3b, 0xa1, 0xd9, 0xc5, 0x70, 0x88, 0x22, 0x04, 0x91, 0x4f, 0xe5,
	0x40, 0xa2, 0x06, 0x89, 0xd3, 0x8a, 0x15, 0xda, 0x3e, 0x41, 0x28, 0x0f, 0x90, 0xd6, 0x43, 0x65,
	0x29, 0xad, 0x43, 0xe2, 0xf6, 0x11, 0x78, 0x6d, 0x90, 0xed, 0xb8, 0x6d, 0x22, 0xd2, 0x0b, 0xb7,
	0xcc, 0x97, 0x6f, 0xe6, 0xfb, 0xe6, 0xc7, 0x40, 0xa4, 0x60, 0xe2, 0x53, 0x87, 0xed, 0x91, 0x6f,
	0x30, 0x6b, 0x5a, 0x21, 0x05, 0x71, 0x8f, 0x8b, 0xe4, 0xc3, 0x56, 0x88, 0x6d, 0x8d, 0xb9, 0x46,
	0xd6, 0x87, 0x9f, 0xb9, 0xe4, 0x3b, 0xec, 0x64, 0xb5, 0x6b, 0x0c, 0x89, 0xfe, 0x76, 0xc0, 0x5f,
	0x89, 0xa5, 0x20, 0x11, 0xb8, 0x9c, 0xc5, 0x4e, 0xea, 0xcc, 0xbd, 0xd2, 0xe5, 0x8c, 0xbc, 0x85,
	0x1b, 0xc9, 0x65, 0x8d, 0xb1, 0x9b, 0x3a, 0xf3, 0xfb, 0xd2, 0x04, 0x24, 0x85, 0x90, 0x61, 0xb7,
	0x69, 0x79, 0x23, 0xb9, 0xd8, 0xc7, 0x9e, 0xfe, 0x77, 0x09, 0x91, 0x2f, 0x70, 0xd7, 0xe2, 0x8e,
	0xef, 0x19, 0xb6, 0xb1, 0x9f, 0x3a, 0xf3, 0xb0, 0x48, 0x32, 0x63, 0x22, 0xb3, 0x26, 0xb2, 0x95,
	0x35, 0x51, 0x9e, 0xb8, 0xf4, 0x19, 0x66, 0x2f, 0x2d, 0x56, 0x12, 0x4b, 0xfc, 0x75, 0xc0, 0x4e,
	0x92, 0x47, 0xf0, 0xaa, 0x86, 0x6b, 0x47, 0xf7, 0xa5, 0xfa, 0x24, 0xef, 0xc0, 0x97, 0x62, 0x29,
	0xb4, 0xa3, 0xb0, 0xb8, 0xcb, 0x8e, 0x8b, 0x4c, 0x59, 0x2f, 0x35, 0x4a, 0x0b, 0x88, 0x6c, 0x81,
	0xae, 0x11, 0xfb, 0x0e, 0xff, 0x51, 0xc1, 0x34, 0xe9, 0xda, 0x26, 0x69, 0x0e, 0x61, 0x89, 0x15,
	0x9b, 0x96, 0x1c, 0x27, 0x7c, 0x85, 0x57, 0x26, 0x61, 0x52, 0xe2, 0xba, 0xc9, 0x67, 0x98, 0xfd,
	0x68, 0xd8, 0x7f, 0x74, 0xf9, 0x04, 0x91, 0x2d, 0x30, 0x69, 0x21, 0x86, 0xdb, 0x83, 0xe6, 0x58,
	0xe7, 0x36, 0xa4, 0x0b, 0x98, 0x2d, 0xb1, 0xc6, 0x6b, 0xf2, 0xe3, 0x8e, 0x9f, 0x20, 0xb2, 0x29,
	0xd7, 0x04, 0x99, 0xe6, 0x9c, 0x04, 0xfb, 0x90, 0x52, 0x88, 0xd4, 0xbc, 0xbe, 0xd5, 0xf5, 0xa4,
	0x22, 0x7d, 0x81, 0x87, 0x13, 0x67, 0x52, 0xe2, 0x3d, 0xdc, 0xa8, 0xfe, 0xbb, 0xd8, 0x4d, 0xbd,
	0xc1, 0x58, 0x0c, 0x5c, 0xfc, 0x71, 0x20, 0x54, 0xf1, 0x77, 0xf3, 0x04, 0x48, 0x0e, 0x81, 0xb9,
	0x06, 0xf2, 0x5a, 0x51, 0x07, 0xa7, 0x95, 0x90, 0x4b, 0xa8, 0x97, 0xfc, 0x08, 0xbe, 0x72, 0x41,
	0x1e, 0xd4, 0xbf, 0x8b, 0xa3, 0x48, 0x1e, 0xcf, 0x40, 0x4f, 0xcd, 0x21, 0x30, 0x3b, 0x30, 0xb5,
	0x07, 0x0b, 0x35, 0xb5, 0x47, 0x2b, 0xca, 0x21, 0x30, 0x33, 0x34, 0x09, 0x83, 0x15, 0x98, 0x84,
	0xd1, 0x88, 0x0b, 0xb8, 0xed, 0x47, 0x42, 0x88, 0x95, 0x3f, 0xcf, 0x30, 0x79, 0x33, 0xc0, 0x4c,
	0xce, 0x3a, 0xd0, 0xcf, 0xeb, 0xf3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x15, 0x17, 0x41,
	0x0b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToDoServiceClient interface {
	// Create new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update todo task
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete todo task
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Read all todo tasks
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConn
}

func NewToDoServiceClient(cc grpc.ClientConn) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
type ToDoServiceServer interface {
	// Create new todo task
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update todo task
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete todo task
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Read all todo tasks
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedToDoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (*UnimplementedToDoServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedToDoServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedToDoServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedToDoServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedToDoServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterToDoServiceServer(s *grpc.Server, srv ToDoServiceServer) {
	s.RegisterService(&_ToDoService_serviceDesc, srv)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ToDoService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ToDoService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo-service.proto",
}
