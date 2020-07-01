// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: data/imss.proto

package data

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UploadStatus int32

const (
	UploadStatus_UNKNOWN     UploadStatus = 0
	UploadStatus_IN_PROGRESS UploadStatus = 1
	UploadStatus_DONE        UploadStatus = 2
	UploadStatus_ERROR       UploadStatus = 3
)

// Enum value maps for UploadStatus.
var (
	UploadStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "IN_PROGRESS",
		2: "DONE",
		3: "ERROR",
	}
	UploadStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"IN_PROGRESS": 1,
		"DONE":        2,
		"ERROR":       3,
	}
)

func (x UploadStatus) Enum() *UploadStatus {
	p := new(UploadStatus)
	*p = x
	return p
}

func (x UploadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_data_imss_proto_enumTypes[0].Descriptor()
}

func (UploadStatus) Type() protoreflect.EnumType {
	return &file_data_imss_proto_enumTypes[0]
}

func (x UploadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatus.Descriptor instead.
func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return file_data_imss_proto_rawDescGZIP(), []int{0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                 // Session ID
	IsActive   bool                 `protobuf:"varint,2,opt,name=isActive,proto3" json:"isActive,omitempty"`    // Completion marker, only one active session at a time is allowed
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`   // Creation time
	FinishedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"` // Completion time
	Name       string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`             // Session name/title
	Images     []*Image             `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`         // List of related images (output only)
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_imss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_data_imss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_data_imss_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Session) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Session) GetFinishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Session) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Session) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // Image ID
	SessionId string               `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"` // Corresponding session ID
	LocalPath string               `protobuf:"bytes,3,opt,name=localPath,proto3" json:"localPath,omitempty"` // Path to image file on the workstation
	CloudId   string               `protobuf:"bytes,4,opt,name=cloudId,proto3" json:"cloudId,omitempty"`     // Remote path/ID of uploaded image
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // Creation time
	Size      uint64               `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`          // File size in bytes
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_imss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_data_imss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_data_imss_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Image) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

func (x *Image) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *Image) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Image) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ImageUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                 // Image upload ID
	JobId    string       `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId,omitempty"`                           // Related upload job ID
	ImageId  string       `protobuf:"bytes,3,opt,name=imageId,proto3" json:"imageId,omitempty"`                       // ID of the image being uploaded
	Progress uint64       `protobuf:"varint,4,opt,name=progress,proto3" json:"progress,omitempty"`                    // Uploaded bytes
	CloudId  string       `protobuf:"bytes,5,opt,name=cloudId,proto3" json:"cloudId,omitempty"`                       // Image ID in the cloud
	Status   UploadStatus `protobuf:"varint,6,opt,name=status,proto3,enum=imss.UploadStatus" json:"status,omitempty"` // Upload status
	Error    *Error       `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`                           // Error message, if any
}

func (x *ImageUpload) Reset() {
	*x = ImageUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_imss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUpload) ProtoMessage() {}

func (x *ImageUpload) ProtoReflect() protoreflect.Message {
	mi := &file_data_imss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUpload.ProtoReflect.Descriptor instead.
func (*ImageUpload) Descriptor() ([]byte, []int) {
	return file_data_imss_proto_rawDescGZIP(), []int{2}
}

func (x *ImageUpload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImageUpload) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *ImageUpload) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ImageUpload) GetProgress() uint64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *ImageUpload) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *ImageUpload) GetStatus() UploadStatus {
	if x != nil {
		return x.Status
	}
	return UploadStatus_UNKNOWN
}

func (x *ImageUpload) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UploadJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                 // Job ID
	Status     UploadStatus         `protobuf:"varint,2,opt,name=status,proto3,enum=imss.UploadStatus" json:"status,omitempty"` // Upload status
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`                   // Start time
	FinishedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`                 // End time
	Progress   uint64               `protobuf:"varint,5,opt,name=progress,proto3" json:"progress,omitempty"`                    // Uploaded bytes
	Size       uint64               `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`                            // Total amount of bytes to upload
	Recipient  string               `protobuf:"bytes,7,opt,name=recipient,proto3" json:"recipient,omitempty"`                   // User's email or other contact info
	Images     []*ImageUpload       `protobuf:"bytes,8,rep,name=images,proto3" json:"images,omitempty"`                         // Images upload data
	CloudLink  string               `protobuf:"bytes,9,opt,name=CloudLink,proto3" json:"CloudLink,omitempty"`                   // Link to share with a recipient
	Error      *Error               `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`                          // Error message, if any
	Name       string               `protobuf:"bytes,11,opt,name=Name,proto3" json:"Name,omitempty"`                            // Album name (optional)
}

func (x *UploadJob) Reset() {
	*x = UploadJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_imss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadJob) ProtoMessage() {}

func (x *UploadJob) ProtoReflect() protoreflect.Message {
	mi := &file_data_imss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadJob.ProtoReflect.Descriptor instead.
func (*UploadJob) Descriptor() ([]byte, []int) {
	return file_data_imss_proto_rawDescGZIP(), []int{3}
}

func (x *UploadJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadJob) GetStatus() UploadStatus {
	if x != nil {
		return x.Status
	}
	return UploadStatus_UNKNOWN
}

func (x *UploadJob) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UploadJob) GetFinishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *UploadJob) GetProgress() uint64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *UploadJob) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadJob) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *UploadJob) GetImages() []*ImageUpload {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *UploadJob) GetCloudLink() string {
	if x != nil {
		return x.CloudLink
	}
	return ""
}

func (x *UploadJob) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UploadJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // Error code
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Error message
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_imss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_data_imss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_data_imss_proto_rawDescGZIP(), []int{4}
}

func (x *Error) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_data_imss_proto protoreflect.FileDescriptor

var file_data_imss_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x69, 0x6d, 0x73, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6d, 0x73,
	0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22,
	0xbb, 0x01, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xd2, 0x01,
	0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x69, 0x6d, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x69, 0x6d, 0x73, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x8b, 0x03, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4a, 0x6f, 0x62,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x69, 0x6d, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x69, 0x6d, 0x73, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6d, 0x73, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x41, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x6e, 0x69, 0x6e, 0x2f, 0x69,
	0x6d, 0x73, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_imss_proto_rawDescOnce sync.Once
	file_data_imss_proto_rawDescData = file_data_imss_proto_rawDesc
)

func file_data_imss_proto_rawDescGZIP() []byte {
	file_data_imss_proto_rawDescOnce.Do(func() {
		file_data_imss_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_imss_proto_rawDescData)
	})
	return file_data_imss_proto_rawDescData
}

var file_data_imss_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_imss_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_data_imss_proto_goTypes = []interface{}{
	(UploadStatus)(0),           // 0: imss.UploadStatus
	(*Session)(nil),             // 1: imss.Session
	(*Image)(nil),               // 2: imss.Image
	(*ImageUpload)(nil),         // 3: imss.ImageUpload
	(*UploadJob)(nil),           // 4: imss.UploadJob
	(*Error)(nil),               // 5: imss.Error
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_data_imss_proto_depIdxs = []int32{
	6,  // 0: imss.Session.createdAt:type_name -> google.protobuf.Timestamp
	6,  // 1: imss.Session.finishedAt:type_name -> google.protobuf.Timestamp
	2,  // 2: imss.Session.images:type_name -> imss.Image
	6,  // 3: imss.Image.createdAt:type_name -> google.protobuf.Timestamp
	0,  // 4: imss.ImageUpload.status:type_name -> imss.UploadStatus
	5,  // 5: imss.ImageUpload.error:type_name -> imss.Error
	0,  // 6: imss.UploadJob.status:type_name -> imss.UploadStatus
	6,  // 7: imss.UploadJob.createdAt:type_name -> google.protobuf.Timestamp
	6,  // 8: imss.UploadJob.finishedAt:type_name -> google.protobuf.Timestamp
	3,  // 9: imss.UploadJob.images:type_name -> imss.ImageUpload
	5,  // 10: imss.UploadJob.error:type_name -> imss.Error
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_data_imss_proto_init() }
func file_data_imss_proto_init() {
	if File_data_imss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_imss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_data_imss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_data_imss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUpload); i {
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
		file_data_imss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadJob); i {
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
		file_data_imss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_data_imss_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_imss_proto_goTypes,
		DependencyIndexes: file_data_imss_proto_depIdxs,
		EnumInfos:         file_data_imss_proto_enumTypes,
		MessageInfos:      file_data_imss_proto_msgTypes,
	}.Build()
	File_data_imss_proto = out.File
	file_data_imss_proto_rawDesc = nil
	file_data_imss_proto_goTypes = nil
	file_data_imss_proto_depIdxs = nil
}
