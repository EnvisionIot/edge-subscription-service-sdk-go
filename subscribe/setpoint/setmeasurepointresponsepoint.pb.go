// Code generated by protoc-gen-go. DO NOT EDIT.
// source: setmeasurepointresponsepoint.proto

package setpoint

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

type SetMeasurepointResponsePoint struct {
	RequestId                 *string  `protobuf:"bytes,1,req,name=requestId" json:"requestId,omitempty"`
	OrgId                     *string  `protobuf:"bytes,2,req,name=orgId" json:"orgId,omitempty"`
	CallType                  *string  `protobuf:"bytes,3,req,name=callType" json:"callType,omitempty"`
	SetMeasurepointChannelId  *string  `protobuf:"bytes,4,req,name=setMeasurepointChannelId" json:"setMeasurepointChannelId,omitempty"`
	ProductKey                *string  `protobuf:"bytes,5,req,name=productKey" json:"productKey,omitempty"`
	DeviceKey                 *string  `protobuf:"bytes,6,req,name=deviceKey" json:"deviceKey,omitempty"`
	AssetId                   *string  `protobuf:"bytes,7,req,name=assetId" json:"assetId,omitempty"`
	MeasurepointId            *string  `protobuf:"bytes,8,req,name=measurepointId" json:"measurepointId,omitempty"`
	CallbackUrl               *string  `protobuf:"bytes,9,req,name=callbackUrl" json:"callbackUrl,omitempty"`
	InputData                 *string  `protobuf:"bytes,10,req,name=inputData" json:"inputData,omitempty"`
	Status                    *int64   `protobuf:"varint,11,req,name=status" json:"status,omitempty"`
	Msg                       *string  `protobuf:"bytes,12,req,name=msg" json:"msg,omitempty"`
	Submsg                    *string  `protobuf:"bytes,13,req,name=submsg" json:"submsg,omitempty"`
	Timeout                   *int64   `protobuf:"varint,14,req,name=timeout" json:"timeout,omitempty"`
	GmtSetMeasurepointRequest *int64   `protobuf:"varint,15,req,name=gmtSetMeasurepointRequest" json:"gmtSetMeasurepointRequest,omitempty"`
	GmtSetMeasurepointReply   *int64   `protobuf:"varint,16,req,name=gmtSetMeasurepointReply" json:"gmtSetMeasurepointReply,omitempty"`
	Attr                      *string  `protobuf:"bytes,17,opt,name=attr" json:"attr,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *SetMeasurepointResponsePoint) Reset()         { *m = SetMeasurepointResponsePoint{} }
func (m *SetMeasurepointResponsePoint) String() string { return proto.CompactTextString(m) }
func (*SetMeasurepointResponsePoint) ProtoMessage()    {}
func (*SetMeasurepointResponsePoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4b414afba73e59, []int{0}
}

func (m *SetMeasurepointResponsePoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMeasurepointResponsePoint.Unmarshal(m, b)
}
func (m *SetMeasurepointResponsePoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMeasurepointResponsePoint.Marshal(b, m, deterministic)
}
func (m *SetMeasurepointResponsePoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMeasurepointResponsePoint.Merge(m, src)
}
func (m *SetMeasurepointResponsePoint) XXX_Size() int {
	return xxx_messageInfo_SetMeasurepointResponsePoint.Size(m)
}
func (m *SetMeasurepointResponsePoint) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMeasurepointResponsePoint.DiscardUnknown(m)
}

var xxx_messageInfo_SetMeasurepointResponsePoint proto.InternalMessageInfo

func (m *SetMeasurepointResponsePoint) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetOrgId() string {
	if m != nil && m.OrgId != nil {
		return *m.OrgId
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetCallType() string {
	if m != nil && m.CallType != nil {
		return *m.CallType
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetSetMeasurepointChannelId() string {
	if m != nil && m.SetMeasurepointChannelId != nil {
		return *m.SetMeasurepointChannelId
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetProductKey() string {
	if m != nil && m.ProductKey != nil {
		return *m.ProductKey
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetDeviceKey() string {
	if m != nil && m.DeviceKey != nil {
		return *m.DeviceKey
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetAssetId() string {
	if m != nil && m.AssetId != nil {
		return *m.AssetId
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetMeasurepointId() string {
	if m != nil && m.MeasurepointId != nil {
		return *m.MeasurepointId
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetCallbackUrl() string {
	if m != nil && m.CallbackUrl != nil {
		return *m.CallbackUrl
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetInputData() string {
	if m != nil && m.InputData != nil {
		return *m.InputData
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetStatus() int64 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *SetMeasurepointResponsePoint) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetSubmsg() string {
	if m != nil && m.Submsg != nil {
		return *m.Submsg
	}
	return ""
}

func (m *SetMeasurepointResponsePoint) GetTimeout() int64 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *SetMeasurepointResponsePoint) GetGmtSetMeasurepointRequest() int64 {
	if m != nil && m.GmtSetMeasurepointRequest != nil {
		return *m.GmtSetMeasurepointRequest
	}
	return 0
}

func (m *SetMeasurepointResponsePoint) GetGmtSetMeasurepointReply() int64 {
	if m != nil && m.GmtSetMeasurepointReply != nil {
		return *m.GmtSetMeasurepointReply
	}
	return 0
}

func (m *SetMeasurepointResponsePoint) GetAttr() string {
	if m != nil && m.Attr != nil {
		return *m.Attr
	}
	return ""
}

type SetMeasurepointResponsePoints struct {
	Points               []*SetMeasurepointResponsePoint `protobuf:"bytes,1,rep,name=points" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SetMeasurepointResponsePoints) Reset()         { *m = SetMeasurepointResponsePoints{} }
func (m *SetMeasurepointResponsePoints) String() string { return proto.CompactTextString(m) }
func (*SetMeasurepointResponsePoints) ProtoMessage()    {}
func (*SetMeasurepointResponsePoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4b414afba73e59, []int{1}
}

func (m *SetMeasurepointResponsePoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMeasurepointResponsePoints.Unmarshal(m, b)
}
func (m *SetMeasurepointResponsePoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMeasurepointResponsePoints.Marshal(b, m, deterministic)
}
func (m *SetMeasurepointResponsePoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMeasurepointResponsePoints.Merge(m, src)
}
func (m *SetMeasurepointResponsePoints) XXX_Size() int {
	return xxx_messageInfo_SetMeasurepointResponsePoints.Size(m)
}
func (m *SetMeasurepointResponsePoints) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMeasurepointResponsePoints.DiscardUnknown(m)
}

var xxx_messageInfo_SetMeasurepointResponsePoints proto.InternalMessageInfo

func (m *SetMeasurepointResponsePoints) GetPoints() []*SetMeasurepointResponsePoint {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*SetMeasurepointResponsePoint)(nil), "setmeasurepointresponsedata.SetMeasurepointResponsePoint")
	proto.RegisterType((*SetMeasurepointResponsePoints)(nil), "setmeasurepointresponsedata.SetMeasurepointResponsePoints")
}

func init() { proto.RegisterFile("setmeasurepointresponsepoint.proto", fileDescriptor_4b4b414afba73e59) }

var fileDescriptor_4b4b414afba73e59 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x69, 0xd2, 0xaf, 0x4c, 0xfa, 0x19, 0xa4, 0x32, 0xa8, 0x20, 0x42, 0x4f, 0x39, 0xe5,
	0xc0, 0x05, 0x71, 0x2d, 0x5c, 0x22, 0x84, 0x54, 0xf1, 0xf1, 0x07, 0xb8, 0xf1, 0x10, 0x02, 0x49,
	0x1c, 0xec, 0xf1, 0x4a, 0xfd, 0xdf, 0xf7, 0xb0, 0x8a, 0xbb, 0x5d, 0x75, 0xb5, 0xdb, 0x1e, 0xdf,
	0x6f, 0x5e, 0xde, 0xc4, 0xcf, 0x86, 0x8d, 0x21, 0xae, 0x49, 0x18, 0xab, 0xa9, 0x55, 0x65, 0xc3,
	0x9a, 0x4c, 0xab, 0x1a, 0x73, 0x14, 0x69, 0xab, 0x15, 0xab, 0x68, 0x7d, 0xc1, 0x23, 0x05, 0x8b,
	0xcd, 0xad, 0x07, 0x6f, 0x7e, 0x10, 0x7f, 0x3b, 0x9b, 0x7f, 0xbf, 0x9f, 0xef, 0x3a, 0x11, 0x2d,
	0x21, 0xd0, 0xf4, 0xdf, 0x92, 0xe1, 0x4c, 0x62, 0x2f, 0xf6, 0x92, 0x20, 0x9a, 0xc2, 0x40, 0xe9,
	0x22, 0x93, 0xe8, 0x39, 0xb9, 0x80, 0x71, 0x2e, 0xaa, 0xea, 0xe7, 0xa1, 0x25, 0xf4, 0x1d, 0x89,
	0x01, 0xcd, 0xe3, 0xcc, 0xcf, 0x7f, 0x44, 0xd3, 0x50, 0x95, 0x49, 0xec, 0x3b, 0x47, 0x04, 0xd0,
	0x6a, 0x25, 0x6d, 0xce, 0x5f, 0xe9, 0x80, 0x03, 0xc7, 0x96, 0x10, 0x48, 0xba, 0x29, 0x73, 0xea,
	0xd0, 0xd0, 0xa1, 0x39, 0x8c, 0x84, 0x31, 0xd4, 0xad, 0x1e, 0x39, 0xb0, 0x82, 0xd9, 0xf9, 0x51,
	0x32, 0x89, 0x63, 0xc7, 0x5f, 0x42, 0xd8, 0xfd, 0xc3, 0x5e, 0xe4, 0xff, 0x7e, 0xe9, 0x0a, 0x83,
	0x53, 0x60, 0xd9, 0xb4, 0x96, 0xbf, 0x08, 0x16, 0x08, 0x0e, 0xcd, 0x60, 0x68, 0x58, 0xb0, 0x35,
	0x18, 0xc6, 0x5e, 0xe2, 0x47, 0x21, 0xf8, 0xb5, 0x29, 0x70, 0xf2, 0x30, 0xb4, 0xfb, 0x4e, 0x4f,
	0x4f, 0xdb, 0xb9, 0xac, 0x49, 0x59, 0xc6, 0x99, 0x73, 0xbf, 0x87, 0xd7, 0x45, 0xcd, 0x4f, 0xea,
	0x72, 0xed, 0xe0, 0xdc, 0x59, 0xde, 0xc1, 0xab, 0xe7, 0x2c, 0x6d, 0x75, 0xc0, 0x85, 0x33, 0x4c,
	0xa0, 0x2f, 0x98, 0x35, 0x2e, 0xe3, 0x5e, 0x12, 0x6c, 0xfe, 0xc2, 0xdb, 0x6b, 0xed, 0x9b, 0x28,
	0x83, 0xa1, 0xc3, 0x06, 0x7b, 0xb1, 0x9f, 0x84, 0x1f, 0x3e, 0xa5, 0x57, 0x6e, 0x33, 0xbd, 0x96,
	0xb5, 0xfd, 0x08, 0xab, 0x5c, 0xd5, 0xc7, 0x47, 0xb1, 0xb7, 0xbf, 0xd3, 0x82, 0x1a, 0xd2, 0x82,
	0x49, 0x6e, 0xd7, 0x17, 0xbe, 0xeb, 0x8a, 0xdb, 0xbd, 0xb8, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x65,
	0x39, 0xf6, 0x9e, 0x65, 0x02, 0x00, 0x00,
}
