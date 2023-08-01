// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: submission.proto

package pb

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

// status
type GetMySubmissionStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignmentID string `protobuf:"bytes,1,opt,name=AssignmentID,proto3" json:"AssignmentID,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetMySubmissionStatusReq) Reset() {
	*x = GetMySubmissionStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMySubmissionStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMySubmissionStatusReq) ProtoMessage() {}

func (x *GetMySubmissionStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMySubmissionStatusReq.ProtoReflect.Descriptor instead.
func (*GetMySubmissionStatusReq) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{0}
}

func (x *GetMySubmissionStatusReq) GetAssignmentID() string {
	if x != nil {
		return x.AssignmentID
	}
	return ""
}

func (x *GetMySubmissionStatusReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetMySubmissionStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetMySubmissionStatusResp) Reset() {
	*x = GetMySubmissionStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMySubmissionStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMySubmissionStatusResp) ProtoMessage() {}

func (x *GetMySubmissionStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMySubmissionStatusResp.ProtoReflect.Descriptor instead.
func (*GetMySubmissionStatusResp) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{1}
}

func (x *GetMySubmissionStatusResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// set
type SetSubmissionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignmentID string   `protobuf:"bytes,1,opt,name=AssignmentID,proto3" json:"AssignmentID,omitempty"`
	UserId       string   `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Urls         []string `protobuf:"bytes,3,rep,name=Urls,proto3" json:"Urls,omitempty"`
}

func (x *SetSubmissionReq) Reset() {
	*x = SetSubmissionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSubmissionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSubmissionReq) ProtoMessage() {}

func (x *SetSubmissionReq) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSubmissionReq.ProtoReflect.Descriptor instead.
func (*SetSubmissionReq) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{2}
}

func (x *SetSubmissionReq) GetAssignmentID() string {
	if x != nil {
		return x.AssignmentID
	}
	return ""
}

func (x *SetSubmissionReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetSubmissionReq) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

type SetSubmissionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
}

func (x *SetSubmissionResp) Reset() {
	*x = SetSubmissionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSubmissionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSubmissionResp) ProtoMessage() {}

func (x *SetSubmissionResp) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSubmissionResp.ProtoReflect.Descriptor instead.
func (*SetSubmissionResp) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{3}
}

func (x *SetSubmissionResp) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

// info
type GetSubmissionInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignmentID string `protobuf:"bytes,1,opt,name=AssignmentID,proto3" json:"AssignmentID,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetSubmissionInfoReq) Reset() {
	*x = GetSubmissionInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmissionInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmissionInfoReq) ProtoMessage() {}

func (x *GetSubmissionInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmissionInfoReq.ProtoReflect.Descriptor instead.
func (*GetSubmissionInfoReq) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{4}
}

func (x *GetSubmissionInfoReq) GetAssignmentID() string {
	if x != nil {
		return x.AssignmentID
	}
	return ""
}

func (x *GetSubmissionInfoReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetSubmissionInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionID string   `protobuf:"bytes,1,opt,name=SubmissionID,proto3" json:"SubmissionID,omitempty"`
	Urls         []string `protobuf:"bytes,2,rep,name=Urls,proto3" json:"Urls,omitempty"`
}

func (x *GetSubmissionInfoResp) Reset() {
	*x = GetSubmissionInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmissionInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmissionInfoResp) ProtoMessage() {}

func (x *GetSubmissionInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmissionInfoResp.ProtoReflect.Descriptor instead.
func (*GetSubmissionInfoResp) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{5}
}

func (x *GetSubmissionInfoResp) GetSubmissionID() string {
	if x != nil {
		return x.SubmissionID
	}
	return ""
}

func (x *GetSubmissionInfoResp) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

// all status
type GetAllSubmissionStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignmentID string `protobuf:"bytes,1,opt,name=AssignmentID,proto3" json:"AssignmentID,omitempty"`
	Page         int64  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (x *GetAllSubmissionStatusReq) Reset() {
	*x = GetAllSubmissionStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSubmissionStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSubmissionStatusReq) ProtoMessage() {}

func (x *GetAllSubmissionStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSubmissionStatusReq.ProtoReflect.Descriptor instead.
func (*GetAllSubmissionStatusReq) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllSubmissionStatusReq) GetAssignmentID() string {
	if x != nil {
		return x.AssignmentID
	}
	return ""
}

func (x *GetAllSubmissionStatusReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Completion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Grade  string `protobuf:"bytes,3,opt,name=Grade,proto3" json:"Grade,omitempty"`
	School string `protobuf:"bytes,4,opt,name=School,proto3" json:"School,omitempty"`
	Status string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Completion) Reset() {
	*x = Completion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Completion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Completion) ProtoMessage() {}

func (x *Completion) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Completion.ProtoReflect.Descriptor instead.
func (*Completion) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{7}
}

func (x *Completion) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Completion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Completion) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *Completion) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *Completion) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetAllSubmissionStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completions []*Completion `protobuf:"bytes,1,rep,name=Completions,proto3" json:"Completions,omitempty"`
}

func (x *GetAllSubmissionStatusResp) Reset() {
	*x = GetAllSubmissionStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSubmissionStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSubmissionStatusResp) ProtoMessage() {}

func (x *GetAllSubmissionStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSubmissionStatusResp.ProtoReflect.Descriptor instead.
func (*GetAllSubmissionStatusResp) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllSubmissionStatusResp) GetCompletions() []*Completion {
	if x != nil {
		return x.Completions
	}
	return nil
}

var File_submission_proto protoreflect.FileDescriptor

var file_submission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x56,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x62, 0x0a, 0x10, 0x53,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x55, 0x72, 0x6c, 0x73, 0x22,
	0x27, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x22, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x72, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x53, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x7e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x56, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x38, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x89, 0x03, 0x0a, 0x10, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x64, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x25,
	0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x67, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x26,
	0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submission_proto_rawDescOnce sync.Once
	file_submission_proto_rawDescData = file_submission_proto_rawDesc
)

func file_submission_proto_rawDescGZIP() []byte {
	file_submission_proto_rawDescOnce.Do(func() {
		file_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_submission_proto_rawDescData)
	})
	return file_submission_proto_rawDescData
}

var file_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_submission_proto_goTypes = []interface{}{
	(*GetMySubmissionStatusReq)(nil),   // 0: submission.GetMySubmissionStatusReq
	(*GetMySubmissionStatusResp)(nil),  // 1: submission.GetMySubmissionStatusResp
	(*SetSubmissionReq)(nil),           // 2: submission.SetSubmissionReq
	(*SetSubmissionResp)(nil),          // 3: submission.SetSubmissionResp
	(*GetSubmissionInfoReq)(nil),       // 4: submission.GetSubmissionInfoReq
	(*GetSubmissionInfoResp)(nil),      // 5: submission.GetSubmissionInfoResp
	(*GetAllSubmissionStatusReq)(nil),  // 6: submission.GetAllSubmissionStatusReq
	(*Completion)(nil),                 // 7: submission.Completion
	(*GetAllSubmissionStatusResp)(nil), // 8: submission.GetAllSubmissionStatusResp
}
var file_submission_proto_depIdxs = []int32{
	7, // 0: submission.GetAllSubmissionStatusResp.Completions:type_name -> submission.Completion
	0, // 1: submission.submissionClient.GetMySubmissionStatus:input_type -> submission.GetMySubmissionStatusReq
	2, // 2: submission.submissionClient.SetSubmission:input_type -> submission.SetSubmissionReq
	4, // 3: submission.submissionClient.GetSubmissionInfo:input_type -> submission.GetSubmissionInfoReq
	6, // 4: submission.submissionClient.GetAllSubmissionStatus:input_type -> submission.GetAllSubmissionStatusReq
	1, // 5: submission.submissionClient.GetMySubmissionStatus:output_type -> submission.GetMySubmissionStatusResp
	3, // 6: submission.submissionClient.SetSubmission:output_type -> submission.SetSubmissionResp
	5, // 7: submission.submissionClient.GetSubmissionInfo:output_type -> submission.GetSubmissionInfoResp
	8, // 8: submission.submissionClient.GetAllSubmissionStatus:output_type -> submission.GetAllSubmissionStatusResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submission_proto_init() }
func file_submission_proto_init() {
	if File_submission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMySubmissionStatusReq); i {
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
		file_submission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMySubmissionStatusResp); i {
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
		file_submission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSubmissionReq); i {
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
		file_submission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSubmissionResp); i {
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
		file_submission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmissionInfoReq); i {
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
		file_submission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmissionInfoResp); i {
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
		file_submission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSubmissionStatusReq); i {
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
		file_submission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Completion); i {
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
		file_submission_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSubmissionStatusResp); i {
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
			RawDescriptor: file_submission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submission_proto_goTypes,
		DependencyIndexes: file_submission_proto_depIdxs,
		MessageInfos:      file_submission_proto_msgTypes,
	}.Build()
	File_submission_proto = out.File
	file_submission_proto_rawDesc = nil
	file_submission_proto_goTypes = nil
	file_submission_proto_depIdxs = nil
}
