// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: task/task.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	"server/pkg/status"
	"server/pkg/user"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body     string                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreateAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UserId   int64                  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StatusId int64                  `protobuf:"varint,5,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateTaskRequest) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *CreateTaskRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTaskRequest) GetStatusId() int64 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskResponse) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   int64                  `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Title    string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body     string                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	CreateAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	User     *user.UserData         `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Status   *status.StatusData     `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskResponse) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *GetTaskResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetTaskResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *GetTaskResponse) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *GetTaskResponse) GetUser() *user.UserData {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetTaskResponse) GetStatus() *status.StatusData {
	if x != nil {
		return x.Status
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

var File_task_task_proto protoreflect.FileDescriptor

var file_task_task_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x32, 0xbe, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3f, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_task_task_proto_rawDescOnce sync.Once
	file_task_task_proto_rawDescData = file_task_task_proto_rawDesc
)

func file_task_task_proto_rawDescGZIP() []byte {
	file_task_task_proto_rawDescOnce.Do(func() {
		file_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_task_proto_rawDescData)
	})
	return file_task_task_proto_rawDescData
}

var file_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_task_task_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),     // 0: task.CreateTaskRequest
	(*CreateTaskResponse)(nil),    // 1: task.CreateTaskResponse
	(*GetTaskRequest)(nil),        // 2: task.GetTaskRequest
	(*GetTaskResponse)(nil),       // 3: task.GetTaskResponse
	(*DeleteTaskRequest)(nil),     // 4: task.DeleteTaskRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*user.UserData)(nil),         // 6: user.UserData
	(*status.StatusData)(nil),     // 7: status.StatusData
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_task_task_proto_depIdxs = []int32{
	5, // 0: task.CreateTaskRequest.create_at:type_name -> google.protobuf.Timestamp
	5, // 1: task.GetTaskResponse.create_at:type_name -> google.protobuf.Timestamp
	6, // 2: task.GetTaskResponse.user:type_name -> user.UserData
	7, // 3: task.GetTaskResponse.status:type_name -> status.StatusData
	0, // 4: task.Task.CreateTask:input_type -> task.CreateTaskRequest
	2, // 5: task.Task.GetTask:input_type -> task.GetTaskRequest
	4, // 6: task.Task.DeleteTask:input_type -> task.DeleteTaskRequest
	1, // 7: task.Task.CreateTask:output_type -> task.CreateTaskResponse
	3, // 8: task.Task.GetTask:output_type -> task.GetTaskResponse
	8, // 9: task.Task.DeleteTask:output_type -> google.protobuf.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_task_task_proto_init() }
func file_task_task_proto_init() {
	if File_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTaskRequest); i {
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
		file_task_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTaskResponse); i {
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
		file_task_task_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskRequest); i {
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
		file_task_task_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskResponse); i {
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
		file_task_task_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskRequest); i {
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
			RawDescriptor: file_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_task_proto_goTypes,
		DependencyIndexes: file_task_task_proto_depIdxs,
		MessageInfos:      file_task_task_proto_msgTypes,
	}.Build()
	File_task_task_proto = out.File
	file_task_task_proto_rawDesc = nil
	file_task_task_proto_goTypes = nil
	file_task_task_proto_depIdxs = nil
}
