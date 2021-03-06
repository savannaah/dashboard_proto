// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package dashboardproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestHeader struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DashboardRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DashboardRequest) Reset()         { *m = DashboardRequest{} }
func (m *DashboardRequest) String() string { return proto.CompactTextString(m) }
func (*DashboardRequest) ProtoMessage()    {}
func (*DashboardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *DashboardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardRequest.Unmarshal(m, b)
}
func (m *DashboardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardRequest.Marshal(b, m, deterministic)
}
func (m *DashboardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardRequest.Merge(m, src)
}
func (m *DashboardRequest) XXX_Size() int {
	return xxx_messageInfo_DashboardRequest.Size(m)
}
func (m *DashboardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardRequest proto.InternalMessageInfo

func (m *DashboardRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DashboardRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ComponentRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ComponentRequest) Reset()         { *m = ComponentRequest{} }
func (m *ComponentRequest) String() string { return proto.CompactTextString(m) }
func (*ComponentRequest) ProtoMessage()    {}
func (*ComponentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ComponentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentRequest.Unmarshal(m, b)
}
func (m *ComponentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentRequest.Marshal(b, m, deterministic)
}
func (m *ComponentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentRequest.Merge(m, src)
}
func (m *ComponentRequest) XXX_Size() int {
	return xxx_messageInfo_ComponentRequest.Size(m)
}
func (m *ComponentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentRequest proto.InternalMessageInfo

func (m *ComponentRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ComponentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DashboardResponse struct {
	StatusCode           int32                   `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DashboardResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DashboardResponse) Reset()         { *m = DashboardResponse{} }
func (m *DashboardResponse) String() string { return proto.CompactTextString(m) }
func (*DashboardResponse) ProtoMessage()    {}
func (*DashboardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *DashboardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardResponse.Unmarshal(m, b)
}
func (m *DashboardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardResponse.Marshal(b, m, deterministic)
}
func (m *DashboardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardResponse.Merge(m, src)
}
func (m *DashboardResponse) XXX_Size() int {
	return xxx_messageInfo_DashboardResponse.Size(m)
}
func (m *DashboardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardResponse proto.InternalMessageInfo

func (m *DashboardResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DashboardResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DashboardResponse) GetData() *DashboardResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DashboardResponse_Component struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Html                 string   `protobuf:"bytes,3,opt,name=html,proto3" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DashboardResponse_Component) Reset()         { *m = DashboardResponse_Component{} }
func (m *DashboardResponse_Component) String() string { return proto.CompactTextString(m) }
func (*DashboardResponse_Component) ProtoMessage()    {}
func (*DashboardResponse_Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3, 0}
}

func (m *DashboardResponse_Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardResponse_Component.Unmarshal(m, b)
}
func (m *DashboardResponse_Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardResponse_Component.Marshal(b, m, deterministic)
}
func (m *DashboardResponse_Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardResponse_Component.Merge(m, src)
}
func (m *DashboardResponse_Component) XXX_Size() int {
	return xxx_messageInfo_DashboardResponse_Component.Size(m)
}
func (m *DashboardResponse_Component) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardResponse_Component.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardResponse_Component proto.InternalMessageInfo

func (m *DashboardResponse_Component) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DashboardResponse_Component) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DashboardResponse_Component) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

type DashboardResponse_Row struct {
	Id                   int32                          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Components           []*DashboardResponse_Component `protobuf:"bytes,2,rep,name=components,proto3" json:"components,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DashboardResponse_Row) Reset()         { *m = DashboardResponse_Row{} }
func (m *DashboardResponse_Row) String() string { return proto.CompactTextString(m) }
func (*DashboardResponse_Row) ProtoMessage()    {}
func (*DashboardResponse_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3, 1}
}

func (m *DashboardResponse_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardResponse_Row.Unmarshal(m, b)
}
func (m *DashboardResponse_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardResponse_Row.Marshal(b, m, deterministic)
}
func (m *DashboardResponse_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardResponse_Row.Merge(m, src)
}
func (m *DashboardResponse_Row) XXX_Size() int {
	return xxx_messageInfo_DashboardResponse_Row.Size(m)
}
func (m *DashboardResponse_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardResponse_Row.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardResponse_Row proto.InternalMessageInfo

func (m *DashboardResponse_Row) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DashboardResponse_Row) GetComponents() []*DashboardResponse_Component {
	if m != nil {
		return m.Components
	}
	return nil
}

type DashboardResponse_Data struct {
	Id                   int32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rows                 []*DashboardResponse_Row `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DashboardResponse_Data) Reset()         { *m = DashboardResponse_Data{} }
func (m *DashboardResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DashboardResponse_Data) ProtoMessage()    {}
func (*DashboardResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3, 2}
}

func (m *DashboardResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardResponse_Data.Unmarshal(m, b)
}
func (m *DashboardResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardResponse_Data.Marshal(b, m, deterministic)
}
func (m *DashboardResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardResponse_Data.Merge(m, src)
}
func (m *DashboardResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DashboardResponse_Data.Size(m)
}
func (m *DashboardResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardResponse_Data proto.InternalMessageInfo

func (m *DashboardResponse_Data) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DashboardResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DashboardResponse_Data) GetRows() []*DashboardResponse_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ComponentResponse struct {
	StatusCode           int32                   `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *ComponentResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ComponentResponse) Reset()         { *m = ComponentResponse{} }
func (m *ComponentResponse) String() string { return proto.CompactTextString(m) }
func (*ComponentResponse) ProtoMessage()    {}
func (*ComponentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ComponentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentResponse.Unmarshal(m, b)
}
func (m *ComponentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentResponse.Marshal(b, m, deterministic)
}
func (m *ComponentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentResponse.Merge(m, src)
}
func (m *ComponentResponse) XXX_Size() int {
	return xxx_messageInfo_ComponentResponse.Size(m)
}
func (m *ComponentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentResponse proto.InternalMessageInfo

func (m *ComponentResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *ComponentResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ComponentResponse) GetData() *ComponentResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ComponentResponse_Value struct {
	Amount               string   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentResponse_Value) Reset()         { *m = ComponentResponse_Value{} }
func (m *ComponentResponse_Value) String() string { return proto.CompactTextString(m) }
func (*ComponentResponse_Value) ProtoMessage()    {}
func (*ComponentResponse_Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4, 0}
}

func (m *ComponentResponse_Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentResponse_Value.Unmarshal(m, b)
}
func (m *ComponentResponse_Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentResponse_Value.Marshal(b, m, deterministic)
}
func (m *ComponentResponse_Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentResponse_Value.Merge(m, src)
}
func (m *ComponentResponse_Value) XXX_Size() int {
	return xxx_messageInfo_ComponentResponse_Value.Size(m)
}
func (m *ComponentResponse_Value) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentResponse_Value.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentResponse_Value proto.InternalMessageInfo

func (m *ComponentResponse_Value) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type ComponentResponse_Chart struct {
	Summary              string                            `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Trend                string                            `protobuf:"bytes,2,opt,name=trend,proto3" json:"trend,omitempty"`
	Labels               []string                          `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	Series               []*ComponentResponse_Chart_Series `protobuf:"bytes,4,rep,name=series,proto3" json:"series,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ComponentResponse_Chart) Reset()         { *m = ComponentResponse_Chart{} }
func (m *ComponentResponse_Chart) String() string { return proto.CompactTextString(m) }
func (*ComponentResponse_Chart) ProtoMessage()    {}
func (*ComponentResponse_Chart) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4, 1}
}

func (m *ComponentResponse_Chart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentResponse_Chart.Unmarshal(m, b)
}
func (m *ComponentResponse_Chart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentResponse_Chart.Marshal(b, m, deterministic)
}
func (m *ComponentResponse_Chart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentResponse_Chart.Merge(m, src)
}
func (m *ComponentResponse_Chart) XXX_Size() int {
	return xxx_messageInfo_ComponentResponse_Chart.Size(m)
}
func (m *ComponentResponse_Chart) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentResponse_Chart.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentResponse_Chart proto.InternalMessageInfo

func (m *ComponentResponse_Chart) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ComponentResponse_Chart) GetTrend() string {
	if m != nil {
		return m.Trend
	}
	return ""
}

func (m *ComponentResponse_Chart) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ComponentResponse_Chart) GetSeries() []*ComponentResponse_Chart_Series {
	if m != nil {
		return m.Series
	}
	return nil
}

type ComponentResponse_Chart_Series struct {
	Values               []int64  `protobuf:"varint,1,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentResponse_Chart_Series) Reset()         { *m = ComponentResponse_Chart_Series{} }
func (m *ComponentResponse_Chart_Series) String() string { return proto.CompactTextString(m) }
func (*ComponentResponse_Chart_Series) ProtoMessage()    {}
func (*ComponentResponse_Chart_Series) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4, 1, 0}
}

func (m *ComponentResponse_Chart_Series) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentResponse_Chart_Series.Unmarshal(m, b)
}
func (m *ComponentResponse_Chart_Series) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentResponse_Chart_Series.Marshal(b, m, deterministic)
}
func (m *ComponentResponse_Chart_Series) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentResponse_Chart_Series.Merge(m, src)
}
func (m *ComponentResponse_Chart_Series) XXX_Size() int {
	return xxx_messageInfo_ComponentResponse_Chart_Series.Size(m)
}
func (m *ComponentResponse_Chart_Series) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentResponse_Chart_Series.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentResponse_Chart_Series proto.InternalMessageInfo

func (m *ComponentResponse_Chart_Series) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type ComponentResponse_Data struct {
	Value                *ComponentResponse_Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Chart                *ComponentResponse_Chart `protobuf:"bytes,2,opt,name=chart,proto3" json:"chart,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ComponentResponse_Data) Reset()         { *m = ComponentResponse_Data{} }
func (m *ComponentResponse_Data) String() string { return proto.CompactTextString(m) }
func (*ComponentResponse_Data) ProtoMessage()    {}
func (*ComponentResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4, 2}
}

func (m *ComponentResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentResponse_Data.Unmarshal(m, b)
}
func (m *ComponentResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentResponse_Data.Marshal(b, m, deterministic)
}
func (m *ComponentResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentResponse_Data.Merge(m, src)
}
func (m *ComponentResponse_Data) XXX_Size() int {
	return xxx_messageInfo_ComponentResponse_Data.Size(m)
}
func (m *ComponentResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentResponse_Data proto.InternalMessageInfo

func (m *ComponentResponse_Data) GetValue() *ComponentResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ComponentResponse_Data) GetChart() *ComponentResponse_Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "dashboardproto.RequestHeader")
	proto.RegisterType((*DashboardRequest)(nil), "dashboardproto.DashboardRequest")
	proto.RegisterType((*ComponentRequest)(nil), "dashboardproto.ComponentRequest")
	proto.RegisterType((*DashboardResponse)(nil), "dashboardproto.DashboardResponse")
	proto.RegisterType((*DashboardResponse_Component)(nil), "dashboardproto.DashboardResponse.Component")
	proto.RegisterType((*DashboardResponse_Row)(nil), "dashboardproto.DashboardResponse.Row")
	proto.RegisterType((*DashboardResponse_Data)(nil), "dashboardproto.DashboardResponse.Data")
	proto.RegisterType((*ComponentResponse)(nil), "dashboardproto.ComponentResponse")
	proto.RegisterType((*ComponentResponse_Value)(nil), "dashboardproto.ComponentResponse.Value")
	proto.RegisterType((*ComponentResponse_Chart)(nil), "dashboardproto.ComponentResponse.Chart")
	proto.RegisterType((*ComponentResponse_Chart_Series)(nil), "dashboardproto.ComponentResponse.Chart.Series")
	proto.RegisterType((*ComponentResponse_Data)(nil), "dashboardproto.ComponentResponse.Data")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0x26, 0x93, 0x1f, 0x49, 0x2d, 0xbb, 0xec, 0xb6, 0x22, 0x21, 0xa0, 0xc6, 0x81, 0xd5, 0x01,
	0x21, 0x87, 0x11, 0x0f, 0x0a, 0x9e, 0xb2, 0xb8, 0x82, 0xb7, 0x16, 0x14, 0x04, 0x0f, 0x3d, 0x93,
	0xc2, 0x19, 0x9c, 0xa4, 0xc7, 0xee, 0xce, 0x0e, 0xde, 0x7d, 0x09, 0x1f, 0xc5, 0x9b, 0x2f, 0xe4,
	0x3b, 0x48, 0x57, 0x3a, 0x63, 0xd6, 0xe8, 0x66, 0x2f, 0x7b, 0xcb, 0xd7, 0xa9, 0xef, 0xa7, 0xba,
	0xba, 0xe0, 0x50, 0xa3, 0xba, 0x58, 0x2f, 0x31, 0xdf, 0x2a, 0x69, 0x24, 0x3b, 0x2a, 0x85, 0x5e,
	0x2d, 0xa4, 0x50, 0x25, 0xe1, 0xe9, 0x29, 0x1c, 0x72, 0xfc, 0xd2, 0xa0, 0x36, 0xaf, 0x51, 0x94,
	0xa8, 0xd8, 0x1d, 0x08, 0x8d, 0xfc, 0x8c, 0x75, 0xe2, 0x65, 0xde, 0x2c, 0xe6, 0x2d, 0x98, 0x7e,
	0x84, 0xe3, 0xb3, 0x8e, 0xe8, 0xea, 0xd9, 0x33, 0x88, 0x56, 0xc4, 0xa1, 0xd2, 0x83, 0xf9, 0xbd,
	0xfc, 0xb2, 0x76, 0x7e, 0x49, 0x98, 0xbb, 0x62, 0xc6, 0x20, 0xa8, 0x45, 0x85, 0xc9, 0x84, 0xf4,
	0xe9, 0xdb, 0xca, 0x17, 0xb2, 0xda, 0xca, 0x1a, 0x6b, 0x73, 0x03, 0xf2, 0xdf, 0x7d, 0x38, 0xe9,
	0xc5, 0xd7, 0x5b, 0x59, 0x6b, 0x64, 0xf7, 0x01, 0xb4, 0x11, 0xa6, 0xd1, 0x85, 0x2c, 0x91, 0x4c,
	0x42, 0xde, 0x3b, 0x61, 0x09, 0xdc, 0xaa, 0x50, 0x6b, 0xf1, 0xa9, 0x13, 0xeb, 0x20, 0x7b, 0x01,
	0x41, 0x29, 0x8c, 0x48, 0x7c, 0x0a, 0xf6, 0xe8, 0xef, 0x60, 0x03, 0xab, 0xfc, 0x4c, 0x18, 0xc1,
	0x89, 0x93, 0x16, 0x10, 0xef, 0x5b, 0x65, 0x47, 0x30, 0x59, 0x97, 0xce, 0x7a, 0xb2, 0x2e, 0xff,
	0x15, 0xde, 0x9e, 0xad, 0x4c, 0xb5, 0x21, 0xb3, 0x98, 0xd3, 0x77, 0xba, 0x00, 0x9f, 0xcb, 0xdd,
	0x80, 0xfe, 0x06, 0x60, 0xd9, 0x69, 0xeb, 0x64, 0x92, 0xf9, 0xb3, 0x83, 0xf9, 0x93, 0xf1, 0x74,
	0x7f, 0xae, 0xbe, 0x47, 0x4f, 0x11, 0x02, 0x1b, 0xfb, 0x5a, 0x19, 0x9f, 0x43, 0xa0, 0xe4, 0x4e,
	0x27, 0x3e, 0x59, 0x9e, 0x8e, 0x5b, 0x72, 0xb9, 0xe3, 0x44, 0x99, 0xfe, 0xf2, 0xe1, 0xa4, 0x37,
	0xfb, 0x9b, 0x9e, 0xcd, 0xc0, 0xaa, 0x3f, 0x9b, 0x07, 0x10, 0xbe, 0x13, 0x9b, 0x06, 0xd9, 0x5d,
	0x88, 0x44, 0x25, 0x9b, 0xda, 0xb8, 0x2d, 0x70, 0x28, 0xfd, 0xe1, 0x41, 0x58, 0xac, 0x84, 0x32,
	0x36, 0x80, 0x6e, 0xaa, 0x4a, 0xa8, 0xaf, 0xae, 0xa4, 0x83, 0xb4, 0x40, 0x0a, 0xeb, 0xd2, 0x05,
	0x6b, 0x81, 0x55, 0xdc, 0x88, 0x05, 0x6e, 0xda, 0x3b, 0x8a, 0xb9, 0x43, 0xec, 0x15, 0x44, 0x1a,
	0xd5, 0x1a, 0x75, 0x12, 0xd0, 0xdd, 0xe5, 0xe3, 0x81, 0x29, 0x40, 0xfe, 0x96, 0x58, 0xdc, 0xb1,
	0xd3, 0x0c, 0xa2, 0xf6, 0xc4, 0x3a, 0x51, 0x13, 0x3a, 0xf1, 0x32, 0x7f, 0xe6, 0x73, 0x87, 0xd2,
	0x6f, 0x9e, 0x1b, 0xe8, 0x4b, 0x08, 0x2f, 0xec, 0x91, 0xdb, 0xab, 0xc7, 0xe3, 0x8e, 0xa4, 0xc0,
	0x5b, 0x96, 0xa5, 0x2f, 0x6d, 0x02, 0xea, 0xef, 0x5a, 0x74, 0x0a, 0xcc, 0x5b, 0xd6, 0xfc, 0xa7,
	0x07, 0xf1, 0xfe, 0x3d, 0xb0, 0xf7, 0x70, 0x7c, 0x8e, 0x66, 0x8f, 0x29, 0x5f, 0x76, 0xc5, 0xf3,
	0xa1, 0x8d, 0x4f, 0x1f, 0x8e, 0x3e, 0x30, 0xf6, 0x01, 0x6e, 0x9f, 0xa3, 0xd9, 0x67, 0x29, 0x64,
	0x6d, 0xec, 0xc2, 0x65, 0x57, 0xa4, 0xfd, 0x8f, 0xf6, 0xa0, 0x9f, 0x45, 0x44, 0x3f, 0x9e, 0xfe,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x52, 0x0b, 0x1d, 0xb3, 0x5b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DashboardClient is the client API for Dashboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DashboardClient interface {
	GetDashboardData(ctx context.Context, in *DashboardRequest, opts ...grpc.CallOption) (*DashboardResponse, error)
	GetComponentContent(ctx context.Context, in *ComponentRequest, opts ...grpc.CallOption) (*ComponentResponse, error)
}

type dashboardClient struct {
	cc *grpc.ClientConn
}

func NewDashboardClient(cc *grpc.ClientConn) DashboardClient {
	return &dashboardClient{cc}
}

func (c *dashboardClient) GetDashboardData(ctx context.Context, in *DashboardRequest, opts ...grpc.CallOption) (*DashboardResponse, error) {
	out := new(DashboardResponse)
	err := c.cc.Invoke(ctx, "/dashboardproto.Dashboard/GetDashboardData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetComponentContent(ctx context.Context, in *ComponentRequest, opts ...grpc.CallOption) (*ComponentResponse, error) {
	out := new(ComponentResponse)
	err := c.cc.Invoke(ctx, "/dashboardproto.Dashboard/GetComponentContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServer is the server API for Dashboard service.
type DashboardServer interface {
	GetDashboardData(context.Context, *DashboardRequest) (*DashboardResponse, error)
	GetComponentContent(context.Context, *ComponentRequest) (*ComponentResponse, error)
}

// UnimplementedDashboardServer can be embedded to have forward compatible implementations.
type UnimplementedDashboardServer struct {
}

func (*UnimplementedDashboardServer) GetDashboardData(ctx context.Context, req *DashboardRequest) (*DashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboardData not implemented")
}
func (*UnimplementedDashboardServer) GetComponentContent(ctx context.Context, req *ComponentRequest) (*ComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentContent not implemented")
}

func RegisterDashboardServer(s *grpc.Server, srv DashboardServer) {
	s.RegisterService(&_Dashboard_serviceDesc, srv)
}

func _Dashboard_GetDashboardData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).GetDashboardData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboardproto.Dashboard/GetDashboardData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).GetDashboardData(ctx, req.(*DashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_GetComponentContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).GetComponentContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboardproto.Dashboard/GetComponentContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).GetComponentContent(ctx, req.(*ComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dashboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dashboardproto.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDashboardData",
			Handler:    _Dashboard_GetDashboardData_Handler,
		},
		{
			MethodName: "GetComponentContent",
			Handler:    _Dashboard_GetComponentContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
