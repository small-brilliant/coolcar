// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: ai.proto

package aipb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Gender int32

const (
	Gender_G_NOT_SPECIFIED Gender = 0
	Gender_FEMALE          Gender = 1
	Gender_MALE            Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "G_NOT_SPECIFIED",
		1: "FEMALE",
		2: "MALE",
	}
	Gender_value = map[string]int32{
		"G_NOT_SPECIFIED": 0,
		"FEMALE":          1,
		"MALE":            2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_ai_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_ai_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{0}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicNumber       string `protobuf:"bytes,1,opt,name=lic_number,json=licNumber,proto3" json:"lic_number,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender          Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=ai.v1.Gender" json:"gender,omitempty"`
	BirthDateMillis int64  `protobuf:"varint,4,opt,name=birth_date_millis,json=birthDateMillis,proto3" json:"birth_date_millis,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{1}
}

func (x *Identity) GetLicNumber() string {
	if x != nil {
		return x.LicNumber
	}
	return ""
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Identity) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_G_NOT_SPECIFIED
}

func (x *Identity) GetBirthDateMillis() int64 {
	if x != nil {
		return x.BirthDateMillis
	}
	return 0
}

type LicIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photo  []byte `protobuf:"bytes,1,opt,name=photo,proto3" json:"photo,omitempty"`
	RealAi bool   `protobuf:"varint,2,opt,name=real_ai,json=realAi,proto3" json:"real_ai,omitempty"`
}

func (x *LicIdentityRequest) Reset() {
	*x = LicIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicIdentityRequest) ProtoMessage() {}

func (x *LicIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicIdentityRequest.ProtoReflect.Descriptor instead.
func (*LicIdentityRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{2}
}

func (x *LicIdentityRequest) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *LicIdentityRequest) GetRealAi() bool {
	if x != nil {
		return x.RealAi
	}
	return false
}

type MeasureDistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *Location `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *Location `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MeasureDistanceRequest) Reset() {
	*x = MeasureDistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasureDistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasureDistanceRequest) ProtoMessage() {}

func (x *MeasureDistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasureDistanceRequest.ProtoReflect.Descriptor instead.
func (*MeasureDistanceRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{3}
}

func (x *MeasureDistanceRequest) GetFrom() *Location {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MeasureDistanceRequest) GetTo() *Location {
	if x != nil {
		return x.To
	}
	return nil
}

type MeasureDistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistanceKm float32 `protobuf:"fixed32,1,opt,name=distance_km,json=distanceKm,proto3" json:"distance_km,omitempty"`
}

func (x *MeasureDistanceResponse) Reset() {
	*x = MeasureDistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasureDistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasureDistanceResponse) ProtoMessage() {}

func (x *MeasureDistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasureDistanceResponse.ProtoReflect.Descriptor instead.
func (*MeasureDistanceResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{4}
}

func (x *MeasureDistanceResponse) GetDistanceKm() float32 {
	if x != nil {
		return x.DistanceKm
	}
	return 0
}

type SimulateCarPosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string    `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Pos   *Location `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *SimulateCarPosRequest) Reset() {
	*x = SimulateCarPosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateCarPosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateCarPosRequest) ProtoMessage() {}

func (x *SimulateCarPosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateCarPosRequest.ProtoReflect.Descriptor instead.
func (*SimulateCarPosRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{5}
}

func (x *SimulateCarPosRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *SimulateCarPosRequest) GetPos() *Location {
	if x != nil {
		return x.Pos
	}
	return nil
}

type SimulateCarPosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SimulateCarPosResponse) Reset() {
	*x = SimulateCarPosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateCarPosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateCarPosResponse) ProtoMessage() {}

func (x *SimulateCarPosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateCarPosResponse.ProtoReflect.Descriptor instead.
func (*SimulateCarPosResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{6}
}

type EndSimulateCarPosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *EndSimulateCarPosRequest) Reset() {
	*x = EndSimulateCarPosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSimulateCarPosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSimulateCarPosRequest) ProtoMessage() {}

func (x *EndSimulateCarPosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSimulateCarPosRequest.ProtoReflect.Descriptor instead.
func (*EndSimulateCarPosRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{7}
}

func (x *EndSimulateCarPosRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type CarPosUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string    `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Pos   *Location `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *CarPosUpdate) Reset() {
	*x = CarPosUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarPosUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarPosUpdate) ProtoMessage() {}

func (x *CarPosUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarPosUpdate.ProtoReflect.Descriptor instead.
func (*CarPosUpdate) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{8}
}

func (x *CarPosUpdate) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *CarPosUpdate) GetPos() *Location {
	if x != nil {
		return x.Pos
	}
	return nil
}

type EndSimulateCarPosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndSimulateCarPosResponse) Reset() {
	*x = EndSimulateCarPosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSimulateCarPosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSimulateCarPosResponse) ProtoMessage() {}

func (x *EndSimulateCarPosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSimulateCarPosResponse.ProtoReflect.Descriptor instead.
func (*EndSimulateCarPosResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{9}
}

var File_ai_proto protoreflect.FileDescriptor

var file_ai_proto_rawDesc = []byte{
	0x0a, 0x08, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x69, 0x2e, 0x76,
	0x31, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x63, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x11, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x43, 0x0a, 0x12, 0x4c, 0x69,
	0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x61,
	0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x41, 0x69, 0x22,
	0x5e, 0x0a, 0x16, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1f,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x17, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x6d, 0x22, 0x51, 0x0a, 0x15, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x70,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x18,
	0x0a, 0x16, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x18, 0x45, 0x6e, 0x64, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x43,
	0x61, 0x72, 0x50, 0x6f, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x45, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x0f,
	0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x32, 0xbf, 0x02, 0x0a, 0x09, 0x41, 0x69, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x4c, 0x69, 0x63, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x63,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x50, 0x0a, 0x0f, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x50, 0x6f, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x6f,
	0x6c, 0x63, 0x61, 0x72, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ai_proto_rawDescOnce sync.Once
	file_ai_proto_rawDescData = file_ai_proto_rawDesc
)

func file_ai_proto_rawDescGZIP() []byte {
	file_ai_proto_rawDescOnce.Do(func() {
		file_ai_proto_rawDescData = protoimpl.X.CompressGZIP(file_ai_proto_rawDescData)
	})
	return file_ai_proto_rawDescData
}

var file_ai_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ai_proto_goTypes = []interface{}{
	(Gender)(0),                       // 0: ai.v1.Gender
	(*Location)(nil),                  // 1: ai.v1.Location
	(*Identity)(nil),                  // 2: ai.v1.Identity
	(*LicIdentityRequest)(nil),        // 3: ai.v1.LicIdentityRequest
	(*MeasureDistanceRequest)(nil),    // 4: ai.v1.MeasureDistanceRequest
	(*MeasureDistanceResponse)(nil),   // 5: ai.v1.MeasureDistanceResponse
	(*SimulateCarPosRequest)(nil),     // 6: ai.v1.SimulateCarPosRequest
	(*SimulateCarPosResponse)(nil),    // 7: ai.v1.SimulateCarPosResponse
	(*EndSimulateCarPosRequest)(nil),  // 8: ai.v1.EndSimulateCarPosRequest
	(*CarPosUpdate)(nil),              // 9: ai.v1.CarPosUpdate
	(*EndSimulateCarPosResponse)(nil), // 10: ai.v1.EndSimulateCarPosResponse
}
var file_ai_proto_depIdxs = []int32{
	0,  // 0: ai.v1.Identity.gender:type_name -> ai.v1.Gender
	1,  // 1: ai.v1.MeasureDistanceRequest.from:type_name -> ai.v1.Location
	1,  // 2: ai.v1.MeasureDistanceRequest.to:type_name -> ai.v1.Location
	1,  // 3: ai.v1.SimulateCarPosRequest.pos:type_name -> ai.v1.Location
	1,  // 4: ai.v1.CarPosUpdate.pos:type_name -> ai.v1.Location
	3,  // 5: ai.v1.AiService.LicIdentity:input_type -> ai.v1.LicIdentityRequest
	4,  // 6: ai.v1.AiService.MeasureDistance:input_type -> ai.v1.MeasureDistanceRequest
	6,  // 7: ai.v1.AiService.SimulateCarPos:input_type -> ai.v1.SimulateCarPosRequest
	8,  // 8: ai.v1.AiService.EndSimulateCarPos:input_type -> ai.v1.EndSimulateCarPosRequest
	2,  // 9: ai.v1.AiService.LicIdentity:output_type -> ai.v1.Identity
	5,  // 10: ai.v1.AiService.MeasureDistance:output_type -> ai.v1.MeasureDistanceResponse
	7,  // 11: ai.v1.AiService.SimulateCarPos:output_type -> ai.v1.SimulateCarPosResponse
	10, // 12: ai.v1.AiService.EndSimulateCarPos:output_type -> ai.v1.EndSimulateCarPosResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ai_proto_init() }
func file_ai_proto_init() {
	if File_ai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_ai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_ai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicIdentityRequest); i {
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
		file_ai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasureDistanceRequest); i {
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
		file_ai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasureDistanceResponse); i {
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
		file_ai_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateCarPosRequest); i {
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
		file_ai_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateCarPosResponse); i {
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
		file_ai_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSimulateCarPosRequest); i {
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
		file_ai_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarPosUpdate); i {
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
		file_ai_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSimulateCarPosResponse); i {
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
			RawDescriptor: file_ai_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ai_proto_goTypes,
		DependencyIndexes: file_ai_proto_depIdxs,
		EnumInfos:         file_ai_proto_enumTypes,
		MessageInfos:      file_ai_proto_msgTypes,
	}.Build()
	File_ai_proto = out.File
	file_ai_proto_rawDesc = nil
	file_ai_proto_goTypes = nil
	file_ai_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AiServiceClient is the client API for AiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AiServiceClient interface {
	LicIdentity(ctx context.Context, in *LicIdentityRequest, opts ...grpc.CallOption) (*Identity, error)
	MeasureDistance(ctx context.Context, in *MeasureDistanceRequest, opts ...grpc.CallOption) (*MeasureDistanceResponse, error)
	SimulateCarPos(ctx context.Context, in *SimulateCarPosRequest, opts ...grpc.CallOption) (*SimulateCarPosResponse, error)
	EndSimulateCarPos(ctx context.Context, in *EndSimulateCarPosRequest, opts ...grpc.CallOption) (*EndSimulateCarPosResponse, error)
}

type aiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAiServiceClient(cc grpc.ClientConnInterface) AiServiceClient {
	return &aiServiceClient{cc}
}

func (c *aiServiceClient) LicIdentity(ctx context.Context, in *LicIdentityRequest, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, "/ai.v1.AiService/LicIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiServiceClient) MeasureDistance(ctx context.Context, in *MeasureDistanceRequest, opts ...grpc.CallOption) (*MeasureDistanceResponse, error) {
	out := new(MeasureDistanceResponse)
	err := c.cc.Invoke(ctx, "/ai.v1.AiService/MeasureDistance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiServiceClient) SimulateCarPos(ctx context.Context, in *SimulateCarPosRequest, opts ...grpc.CallOption) (*SimulateCarPosResponse, error) {
	out := new(SimulateCarPosResponse)
	err := c.cc.Invoke(ctx, "/ai.v1.AiService/SimulateCarPos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiServiceClient) EndSimulateCarPos(ctx context.Context, in *EndSimulateCarPosRequest, opts ...grpc.CallOption) (*EndSimulateCarPosResponse, error) {
	out := new(EndSimulateCarPosResponse)
	err := c.cc.Invoke(ctx, "/ai.v1.AiService/EndSimulateCarPos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiServiceServer is the server API for AiService service.
type AiServiceServer interface {
	LicIdentity(context.Context, *LicIdentityRequest) (*Identity, error)
	MeasureDistance(context.Context, *MeasureDistanceRequest) (*MeasureDistanceResponse, error)
	SimulateCarPos(context.Context, *SimulateCarPosRequest) (*SimulateCarPosResponse, error)
	EndSimulateCarPos(context.Context, *EndSimulateCarPosRequest) (*EndSimulateCarPosResponse, error)
}

// UnimplementedAiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAiServiceServer struct {
}

func (*UnimplementedAiServiceServer) LicIdentity(context.Context, *LicIdentityRequest) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LicIdentity not implemented")
}
func (*UnimplementedAiServiceServer) MeasureDistance(context.Context, *MeasureDistanceRequest) (*MeasureDistanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeasureDistance not implemented")
}
func (*UnimplementedAiServiceServer) SimulateCarPos(context.Context, *SimulateCarPosRequest) (*SimulateCarPosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimulateCarPos not implemented")
}
func (*UnimplementedAiServiceServer) EndSimulateCarPos(context.Context, *EndSimulateCarPosRequest) (*EndSimulateCarPosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndSimulateCarPos not implemented")
}

func RegisterAiServiceServer(s *grpc.Server, srv AiServiceServer) {
	s.RegisterService(&_AiService_serviceDesc, srv)
}

func _AiService_LicIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LicIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServiceServer).LicIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.v1.AiService/LicIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServiceServer).LicIdentity(ctx, req.(*LicIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AiService_MeasureDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureDistanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServiceServer).MeasureDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.v1.AiService/MeasureDistance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServiceServer).MeasureDistance(ctx, req.(*MeasureDistanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AiService_SimulateCarPos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateCarPosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServiceServer).SimulateCarPos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.v1.AiService/SimulateCarPos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServiceServer).SimulateCarPos(ctx, req.(*SimulateCarPosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AiService_EndSimulateCarPos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndSimulateCarPosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServiceServer).EndSimulateCarPos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.v1.AiService/EndSimulateCarPos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServiceServer).EndSimulateCarPos(ctx, req.(*EndSimulateCarPosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.v1.AiService",
	HandlerType: (*AiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LicIdentity",
			Handler:    _AiService_LicIdentity_Handler,
		},
		{
			MethodName: "MeasureDistance",
			Handler:    _AiService_MeasureDistance_Handler,
		},
		{
			MethodName: "SimulateCarPos",
			Handler:    _AiService_SimulateCarPos_Handler,
		},
		{
			MethodName: "EndSimulateCarPos",
			Handler:    _AiService_EndSimulateCarPos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai.proto",
}
