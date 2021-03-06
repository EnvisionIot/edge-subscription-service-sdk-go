// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarmrecord.proto

package alarm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type AlarmRecord struct {
	EventId              *string  `protobuf:"bytes,1,req,name=eventId" json:"eventId,omitempty"`
	OuId                 *string  `protobuf:"bytes,2,req,name=ouId" json:"ouId,omitempty"`
	ModelId              *string  `protobuf:"bytes,3,req,name=modelId" json:"modelId,omitempty"`
	ModelPath            *string  `protobuf:"bytes,4,req,name=modelPath" json:"modelPath,omitempty"`
	AssetId              *string  `protobuf:"bytes,5,req,name=assetId" json:"assetId,omitempty"`
	PointId              *string  `protobuf:"bytes,6,req,name=pointId" json:"pointId,omitempty"`
	RuleId               *string  `protobuf:"bytes,7,req,name=ruleId" json:"ruleId,omitempty"`
	Value                *string  `protobuf:"bytes,8,req,name=value" json:"value,omitempty"`
	ServerityId          *string  `protobuf:"bytes,9,req,name=serverityId" json:"serverityId,omitempty"`
	ServerityDesc        *string  `protobuf:"bytes,10,req,name=serverityDesc" json:"serverityDesc,omitempty"`
	TypeId               *string  `protobuf:"bytes,11,req,name=typeId" json:"typeId,omitempty"`
	TypeDesc             *string  `protobuf:"bytes,12,req,name=typeDesc" json:"typeDesc,omitempty"`
	SubTypeId            *string  `protobuf:"bytes,13,req,name=subTypeId" json:"subTypeId,omitempty"`
	SubTypeDesc          *string  `protobuf:"bytes,14,req,name=subTypeDesc" json:"subTypeDesc,omitempty"`
	ContentId            *string  `protobuf:"bytes,15,req,name=contentId" json:"contentId,omitempty"`
	ContentDesc          *string  `protobuf:"bytes,16,req,name=contentDesc" json:"contentDesc,omitempty"`
	RuleDesc             *string  `protobuf:"bytes,17,req,name=ruleDesc" json:"ruleDesc,omitempty"`
	Tag                  *string  `protobuf:"bytes,18,req,name=tag" json:"tag,omitempty"`
	EventType            *int32   `protobuf:"varint,19,req,name=eventType" json:"eventType,omitempty"`
	OccurTime            *int64   `protobuf:"varint,20,req,name=occurTime" json:"occurTime,omitempty"`
	CreatTime            *int64   `protobuf:"varint,21,req,name=creatTime" json:"creatTime,omitempty"`
	LocalOccurTime       *string  `protobuf:"bytes,22,req,name=localOccurTime" json:"localOccurTime,omitempty"`
	UpdateTime           *int64   `protobuf:"varint,23,req,name=updateTime" json:"updateTime,omitempty"`
	RecoverTime          *int64   `protobuf:"varint,24,opt,name=recoverTime" json:"recoverTime,omitempty"`
	RecoverLocalTime     *string  `protobuf:"bytes,25,opt,name=recoverLocalTime" json:"recoverLocalTime,omitempty"`
	RecoverReason        *string  `protobuf:"bytes,26,opt,name=recoverReason" json:"recoverReason,omitempty"`
	AssetPath            []string `protobuf:"bytes,27,rep,name=assetPath" json:"assetPath,omitempty"`
	MaskedBy             []string `protobuf:"bytes,28,rep,name=maskedBy" json:"maskedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmRecord) Reset()         { *m = AlarmRecord{} }
func (m *AlarmRecord) String() string { return proto.CompactTextString(m) }
func (*AlarmRecord) ProtoMessage()    {}
func (*AlarmRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e11b19c64aad4b39, []int{0}
}

func (m *AlarmRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmRecord.Unmarshal(m, b)
}
func (m *AlarmRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmRecord.Marshal(b, m, deterministic)
}
func (m *AlarmRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmRecord.Merge(m, src)
}
func (m *AlarmRecord) XXX_Size() int {
	return xxx_messageInfo_AlarmRecord.Size(m)
}
func (m *AlarmRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmRecord proto.InternalMessageInfo

func (m *AlarmRecord) GetEventId() string {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return ""
}

func (m *AlarmRecord) GetOuId() string {
	if m != nil && m.OuId != nil {
		return *m.OuId
	}
	return ""
}

func (m *AlarmRecord) GetModelId() string {
	if m != nil && m.ModelId != nil {
		return *m.ModelId
	}
	return ""
}

func (m *AlarmRecord) GetModelPath() string {
	if m != nil && m.ModelPath != nil {
		return *m.ModelPath
	}
	return ""
}

func (m *AlarmRecord) GetAssetId() string {
	if m != nil && m.AssetId != nil {
		return *m.AssetId
	}
	return ""
}

func (m *AlarmRecord) GetPointId() string {
	if m != nil && m.PointId != nil {
		return *m.PointId
	}
	return ""
}

func (m *AlarmRecord) GetRuleId() string {
	if m != nil && m.RuleId != nil {
		return *m.RuleId
	}
	return ""
}

func (m *AlarmRecord) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *AlarmRecord) GetServerityId() string {
	if m != nil && m.ServerityId != nil {
		return *m.ServerityId
	}
	return ""
}

func (m *AlarmRecord) GetServerityDesc() string {
	if m != nil && m.ServerityDesc != nil {
		return *m.ServerityDesc
	}
	return ""
}

func (m *AlarmRecord) GetTypeId() string {
	if m != nil && m.TypeId != nil {
		return *m.TypeId
	}
	return ""
}

func (m *AlarmRecord) GetTypeDesc() string {
	if m != nil && m.TypeDesc != nil {
		return *m.TypeDesc
	}
	return ""
}

func (m *AlarmRecord) GetSubTypeId() string {
	if m != nil && m.SubTypeId != nil {
		return *m.SubTypeId
	}
	return ""
}

func (m *AlarmRecord) GetSubTypeDesc() string {
	if m != nil && m.SubTypeDesc != nil {
		return *m.SubTypeDesc
	}
	return ""
}

func (m *AlarmRecord) GetContentId() string {
	if m != nil && m.ContentId != nil {
		return *m.ContentId
	}
	return ""
}

func (m *AlarmRecord) GetContentDesc() string {
	if m != nil && m.ContentDesc != nil {
		return *m.ContentDesc
	}
	return ""
}

func (m *AlarmRecord) GetRuleDesc() string {
	if m != nil && m.RuleDesc != nil {
		return *m.RuleDesc
	}
	return ""
}

func (m *AlarmRecord) GetTag() string {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return ""
}

func (m *AlarmRecord) GetEventType() int32 {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return 0
}

func (m *AlarmRecord) GetOccurTime() int64 {
	if m != nil && m.OccurTime != nil {
		return *m.OccurTime
	}
	return 0
}

func (m *AlarmRecord) GetCreatTime() int64 {
	if m != nil && m.CreatTime != nil {
		return *m.CreatTime
	}
	return 0
}

func (m *AlarmRecord) GetLocalOccurTime() string {
	if m != nil && m.LocalOccurTime != nil {
		return *m.LocalOccurTime
	}
	return ""
}

func (m *AlarmRecord) GetUpdateTime() int64 {
	if m != nil && m.UpdateTime != nil {
		return *m.UpdateTime
	}
	return 0
}

func (m *AlarmRecord) GetRecoverTime() int64 {
	if m != nil && m.RecoverTime != nil {
		return *m.RecoverTime
	}
	return 0
}

func (m *AlarmRecord) GetRecoverLocalTime() string {
	if m != nil && m.RecoverLocalTime != nil {
		return *m.RecoverLocalTime
	}
	return ""
}

func (m *AlarmRecord) GetRecoverReason() string {
	if m != nil && m.RecoverReason != nil {
		return *m.RecoverReason
	}
	return ""
}

func (m *AlarmRecord) GetAssetPath() []string {
	if m != nil {
		return m.AssetPath
	}
	return nil
}

func (m *AlarmRecord) GetMaskedBy() []string {
	if m != nil {
		return m.MaskedBy
	}
	return nil
}

type AlarmRecords struct {
	Points               []*AlarmRecord `protobuf:"bytes,1,rep,name=points" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AlarmRecords) Reset()         { *m = AlarmRecords{} }
func (m *AlarmRecords) String() string { return proto.CompactTextString(m) }
func (*AlarmRecords) ProtoMessage()    {}
func (*AlarmRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_e11b19c64aad4b39, []int{1}
}

func (m *AlarmRecords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmRecords.Unmarshal(m, b)
}
func (m *AlarmRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmRecords.Marshal(b, m, deterministic)
}
func (m *AlarmRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmRecords.Merge(m, src)
}
func (m *AlarmRecords) XXX_Size() int {
	return xxx_messageInfo_AlarmRecords.Size(m)
}
func (m *AlarmRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmRecords.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmRecords proto.InternalMessageInfo

func (m *AlarmRecords) GetPoints() []*AlarmRecord {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*AlarmRecord)(nil), "alarmdata.AlarmRecord")
	proto.RegisterType((*AlarmRecords)(nil), "alarmdata.AlarmRecords")
}

func init() { proto.RegisterFile("alarmrecord.proto", fileDescriptor_e11b19c64aad4b39) }

var fileDescriptor_e11b19c64aad4b39 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xc9, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x6b, 0xcb, 0x4b, 0x34, 0xf2, 0x2a, 0x37, 0xee, 0x74, 0x39, 0x08, 0x01, 0x5a, 0xe8,
	0xe4, 0x43, 0x0f, 0xbd, 0xd7, 0xc8, 0x45, 0x40, 0x81, 0x06, 0x41, 0x5e, 0x60, 0x22, 0x4e, 0x53,
	0xa3, 0x92, 0x29, 0x90, 0x94, 0x00, 0xbf, 0x4c, 0x9f, 0x35, 0xe0, 0xc8, 0x14, 0x72, 0xfc, 0x3f,
	0x7e, 0xf8, 0xb9, 0x0c, 0x61, 0x4b, 0x15, 0x99, 0xda, 0x70, 0xa9, 0x8d, 0x3a, 0x34, 0x46, 0x3b,
	0x9d, 0xc6, 0x82, 0x14, 0x39, 0xba, 0xfb, 0x3f, 0x81, 0xe4, 0xa7, 0x4f, 0x8f, 0x22, 0xa4, 0x6b,
	0x98, 0x73, 0xc7, 0x67, 0x57, 0x28, 0x1c, 0x65, 0xe3, 0x3c, 0x4e, 0x17, 0x30, 0xd1, 0x6d, 0xa1,
	0x70, 0x2c, 0x69, 0x0d, 0xf3, 0x5a, 0x2b, 0xae, 0x0a, 0x85, 0x91, 0x80, 0x2d, 0xc4, 0x02, 0x1e,
	0xc8, 0xfd, 0xc5, 0x49, 0x70, 0xc8, 0x5a, 0xf6, 0x15, 0xd3, 0x00, 0x1a, 0x7d, 0x92, 0xce, 0x99,
	0x80, 0x15, 0xcc, 0x4c, 0x5b, 0x71, 0xa1, 0x70, 0x2e, 0x79, 0x09, 0xd3, 0x8e, 0xaa, 0x96, 0xf1,
	0x46, 0xe2, 0x0e, 0x12, 0xcb, 0xa6, 0x63, 0x73, 0x72, 0x97, 0x42, 0x61, 0x2c, 0xf0, 0x16, 0x96,
	0x03, 0xbc, 0x67, 0x5b, 0x22, 0x84, 0x2a, 0x77, 0x69, 0x7c, 0x55, 0x22, 0x79, 0x03, 0x37, 0x3e,
	0x8b, 0xb1, 0x08, 0x27, 0xb4, 0xed, 0xf3, 0x53, 0x2f, 0x2d, 0x87, 0x0d, 0x7a, 0x24, 0xde, 0x2a,
	0x78, 0xa5, 0x3e, 0xbb, 0xfe, 0xee, 0xeb, 0xe0, 0x5d, 0x91, 0x78, 0x9b, 0xb0, 0x83, 0x3f, 0xbc,
	0x90, 0xad, 0x90, 0x04, 0x22, 0x47, 0x2f, 0x98, 0x86, 0x1a, 0x79, 0x40, 0xdf, 0x8e, 0xbb, 0x6c,
	0x9c, 0x4f, 0x3d, 0xd2, 0x65, 0xd9, 0x9a, 0xa7, 0x53, 0xcd, 0xf8, 0x3e, 0x1b, 0xe7, 0x91, 0x6c,
	0x66, 0x98, 0x9c, 0xa0, 0x5b, 0x41, 0x7b, 0x58, 0x55, 0xba, 0xa4, 0xea, 0xf7, 0xa0, 0xee, 0xa5,
	0x30, 0x05, 0x68, 0x1b, 0x45, 0x8e, 0x85, 0x7d, 0x10, 0x77, 0x07, 0x89, 0x1f, 0x68, 0xc7, 0xbd,
	0x88, 0xd9, 0x28, 0x8f, 0x52, 0x84, 0xcd, 0x15, 0xfe, 0xf2, 0x3d, 0xb2, 0xf2, 0x31, 0x1b, 0xf5,
	0x6f, 0x77, 0x5d, 0x79, 0x64, 0xb2, 0xfa, 0x8c, 0x9f, 0x04, 0x6f, 0x21, 0x96, 0x41, 0xc9, 0xec,
	0x3e, 0x67, 0x51, 0x7f, 0xb9, 0x9a, 0xec, 0x3f, 0x56, 0xc7, 0x0b, 0x7e, 0xf1, 0xe4, 0xee, 0x07,
	0x2c, 0xde, 0xfc, 0x0f, 0x9b, 0x7e, 0x83, 0x99, 0x0c, 0xd3, 0xe2, 0x28, 0x8b, 0xf2, 0xe4, 0xfb,
	0xfe, 0x30, 0x7c, 0xa6, 0xc3, 0x1b, 0xf1, 0xf8, 0x15, 0xf6, 0xa5, 0xae, 0xfb, 0x0f, 0xf7, 0xdc,
	0xfe, 0x39, 0xbc, 0xf0, 0x99, 0x0d, 0x39, 0x56, 0xc7, 0x58, 0xb4, 0x7b, 0x72, 0xf4, 0xf0, 0xee,
	0x35, 0x00, 0x00, 0xff, 0xff, 0x72, 0x5a, 0x71, 0xc0, 0x9e, 0x02, 0x00, 0x00,
}
