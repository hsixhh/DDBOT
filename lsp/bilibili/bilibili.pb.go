// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bilibili.proto

package bilibili

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

type LiveStatus int32

const (
	LiveStatus_NoLiving LiveStatus = 0
	LiveStatus_Living   LiveStatus = 1
)

var LiveStatus_name = map[int32]string{
	0: "NoLiving",
	1: "Living",
}

var LiveStatus_value = map[string]int32{
	"NoLiving": 0,
	"Living":   1,
}

func (x LiveStatus) String() string {
	return proto.EnumName(LiveStatus_name, int32(x))
}

func (LiveStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{0}
}

type RoundStatus int32

const (
	RoundStatus_Off RoundStatus = 0
	RoundStatus_On  RoundStatus = 1
)

var RoundStatus_name = map[int32]string{
	0: "Off",
	1: "On",
}

var RoundStatus_value = map[string]int32{
	"Off": 0,
	"On":  1,
}

func (x RoundStatus) String() string {
	return proto.EnumName(RoundStatus_name, int32(x))
}

func (RoundStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{1}
}

type RoomStatus int32

const (
	RoomStatus_NonExist RoomStatus = 0
	RoomStatus_Exist    RoomStatus = 1
)

var RoomStatus_name = map[int32]string{
	0: "NonExist",
	1: "Exist",
}

var RoomStatus_value = map[string]int32{
	"NonExist": 0,
	"Exist":    1,
}

func (x RoomStatus) String() string {
	return proto.EnumName(RoomStatus_name, int32(x))
}

func (RoomStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{2}
}

type SpecialType int32

const (
	SpecialType_Normal       SpecialType = 0
	SpecialType_Pay          SpecialType = 1
	SpecialType_YearFestival SpecialType = 2
)

var SpecialType_name = map[int32]string{
	0: "Normal",
	1: "Pay",
	2: "YearFestival",
}

var SpecialType_value = map[string]int32{
	"Normal":       0,
	"Pay":          1,
	"YearFestival": 2,
}

func (x SpecialType) String() string {
	return proto.EnumName(SpecialType_name, int32(x))
}

func (SpecialType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{3}
}

type RoomInitResponse struct {
	Code                 int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *RoomInitResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RoomInitResponse) Reset()         { *m = RoomInitResponse{} }
func (m *RoomInitResponse) String() string { return proto.CompactTextString(m) }
func (*RoomInitResponse) ProtoMessage()    {}
func (*RoomInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{0}
}

func (m *RoomInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInitResponse.Unmarshal(m, b)
}
func (m *RoomInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInitResponse.Marshal(b, m, deterministic)
}
func (m *RoomInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInitResponse.Merge(m, src)
}
func (m *RoomInitResponse) XXX_Size() int {
	return xxx_messageInfo_RoomInitResponse.Size(m)
}
func (m *RoomInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInitResponse proto.InternalMessageInfo

func (m *RoomInitResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RoomInitResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RoomInitResponse) GetData() *RoomInitResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RoomInitResponse_Data struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	ShortId              int64    `protobuf:"varint,2,opt,name=short_id,json=shortId,proto3" json:"short_id,omitempty"`
	Uid                  int64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	IsHidden             bool     `protobuf:"varint,5,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	IsLocked             bool     `protobuf:"varint,6,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty"`
	LiveStatus           int32    `protobuf:"varint,8,opt,name=live_status,json=liveStatus,proto3" json:"live_status,omitempty"`
	IsSp                 int32    `protobuf:"varint,15,opt,name=is_sp,json=isSp,proto3" json:"is_sp,omitempty"`
	SpecialType          int32    `protobuf:"varint,16,opt,name=special_type,json=specialType,proto3" json:"special_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomInitResponse_Data) Reset()         { *m = RoomInitResponse_Data{} }
func (m *RoomInitResponse_Data) String() string { return proto.CompactTextString(m) }
func (*RoomInitResponse_Data) ProtoMessage()    {}
func (*RoomInitResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{0, 0}
}

func (m *RoomInitResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInitResponse_Data.Unmarshal(m, b)
}
func (m *RoomInitResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInitResponse_Data.Marshal(b, m, deterministic)
}
func (m *RoomInitResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInitResponse_Data.Merge(m, src)
}
func (m *RoomInitResponse_Data) XXX_Size() int {
	return xxx_messageInfo_RoomInitResponse_Data.Size(m)
}
func (m *RoomInitResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInitResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInitResponse_Data proto.InternalMessageInfo

func (m *RoomInitResponse_Data) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomInitResponse_Data) GetShortId() int64 {
	if m != nil {
		return m.ShortId
	}
	return 0
}

func (m *RoomInitResponse_Data) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RoomInitResponse_Data) GetIsHidden() bool {
	if m != nil {
		return m.IsHidden
	}
	return false
}

func (m *RoomInitResponse_Data) GetIsLocked() bool {
	if m != nil {
		return m.IsLocked
	}
	return false
}

func (m *RoomInitResponse_Data) GetLiveStatus() int32 {
	if m != nil {
		return m.LiveStatus
	}
	return 0
}

func (m *RoomInitResponse_Data) GetIsSp() int32 {
	if m != nil {
		return m.IsSp
	}
	return 0
}

func (m *RoomInitResponse_Data) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

type XSpaceAccInfoResponse struct {
	Code                 int32                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *XSpaceAccInfoResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *XSpaceAccInfoResponse) Reset()         { *m = XSpaceAccInfoResponse{} }
func (m *XSpaceAccInfoResponse) String() string { return proto.CompactTextString(m) }
func (*XSpaceAccInfoResponse) ProtoMessage()    {}
func (*XSpaceAccInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{1}
}

func (m *XSpaceAccInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSpaceAccInfoResponse.Unmarshal(m, b)
}
func (m *XSpaceAccInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSpaceAccInfoResponse.Marshal(b, m, deterministic)
}
func (m *XSpaceAccInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSpaceAccInfoResponse.Merge(m, src)
}
func (m *XSpaceAccInfoResponse) XXX_Size() int {
	return xxx_messageInfo_XSpaceAccInfoResponse.Size(m)
}
func (m *XSpaceAccInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_XSpaceAccInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_XSpaceAccInfoResponse proto.InternalMessageInfo

func (m *XSpaceAccInfoResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *XSpaceAccInfoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *XSpaceAccInfoResponse) GetData() *XSpaceAccInfoResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type XSpaceAccInfoResponse_Data struct {
	Mid                  int64                                `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	Name                 string                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level                int32                                `protobuf:"varint,7,opt,name=level,proto3" json:"level,omitempty"`
	LiveRoom             *XSpaceAccInfoResponse_Data_LiveRoom `protobuf:"bytes,17,opt,name=live_room,json=liveRoom,proto3" json:"live_room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *XSpaceAccInfoResponse_Data) Reset()         { *m = XSpaceAccInfoResponse_Data{} }
func (m *XSpaceAccInfoResponse_Data) String() string { return proto.CompactTextString(m) }
func (*XSpaceAccInfoResponse_Data) ProtoMessage()    {}
func (*XSpaceAccInfoResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{1, 0}
}

func (m *XSpaceAccInfoResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSpaceAccInfoResponse_Data.Unmarshal(m, b)
}
func (m *XSpaceAccInfoResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSpaceAccInfoResponse_Data.Marshal(b, m, deterministic)
}
func (m *XSpaceAccInfoResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSpaceAccInfoResponse_Data.Merge(m, src)
}
func (m *XSpaceAccInfoResponse_Data) XXX_Size() int {
	return xxx_messageInfo_XSpaceAccInfoResponse_Data.Size(m)
}
func (m *XSpaceAccInfoResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_XSpaceAccInfoResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_XSpaceAccInfoResponse_Data proto.InternalMessageInfo

func (m *XSpaceAccInfoResponse_Data) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *XSpaceAccInfoResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *XSpaceAccInfoResponse_Data) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *XSpaceAccInfoResponse_Data) GetLiveRoom() *XSpaceAccInfoResponse_Data_LiveRoom {
	if m != nil {
		return m.LiveRoom
	}
	return nil
}

type XSpaceAccInfoResponse_Data_LiveRoom struct {
	RoomStatus           int32    `protobuf:"varint,1,opt,name=roomStatus,proto3" json:"roomStatus,omitempty"`
	LiveStatus           int32    `protobuf:"varint,2,opt,name=liveStatus,proto3" json:"liveStatus,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Cover                string   `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	Roomid               int64    `protobuf:"varint,6,opt,name=roomid,proto3" json:"roomid,omitempty"`
	RoundStatus          int32    `protobuf:"varint,7,opt,name=roundStatus,proto3" json:"roundStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) Reset()         { *m = XSpaceAccInfoResponse_Data_LiveRoom{} }
func (m *XSpaceAccInfoResponse_Data_LiveRoom) String() string { return proto.CompactTextString(m) }
func (*XSpaceAccInfoResponse_Data_LiveRoom) ProtoMessage()    {}
func (*XSpaceAccInfoResponse_Data_LiveRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{1, 0, 0}
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XSpaceAccInfoResponse_Data_LiveRoom.Unmarshal(m, b)
}
func (m *XSpaceAccInfoResponse_Data_LiveRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XSpaceAccInfoResponse_Data_LiveRoom.Marshal(b, m, deterministic)
}
func (m *XSpaceAccInfoResponse_Data_LiveRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XSpaceAccInfoResponse_Data_LiveRoom.Merge(m, src)
}
func (m *XSpaceAccInfoResponse_Data_LiveRoom) XXX_Size() int {
	return xxx_messageInfo_XSpaceAccInfoResponse_Data_LiveRoom.Size(m)
}
func (m *XSpaceAccInfoResponse_Data_LiveRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_XSpaceAccInfoResponse_Data_LiveRoom.DiscardUnknown(m)
}

var xxx_messageInfo_XSpaceAccInfoResponse_Data_LiveRoom proto.InternalMessageInfo

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetRoomStatus() int32 {
	if m != nil {
		return m.RoomStatus
	}
	return 0
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetLiveStatus() int32 {
	if m != nil {
		return m.LiveStatus
	}
	return 0
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *XSpaceAccInfoResponse_Data_LiveRoom) GetRoundStatus() int32 {
	if m != nil {
		return m.RoundStatus
	}
	return 0
}

type DynamicSvrSpaceHistoryResponse struct {
	Code                 int32                                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DynamicSvrSpaceHistoryResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DynamicSvrSpaceHistoryResponse) Reset()         { *m = DynamicSvrSpaceHistoryResponse{} }
func (m *DynamicSvrSpaceHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*DynamicSvrSpaceHistoryResponse) ProtoMessage()    {}
func (*DynamicSvrSpaceHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{2}
}

func (m *DynamicSvrSpaceHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse.Unmarshal(m, b)
}
func (m *DynamicSvrSpaceHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse.Marshal(b, m, deterministic)
}
func (m *DynamicSvrSpaceHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse.Merge(m, src)
}
func (m *DynamicSvrSpaceHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse.Size(m)
}
func (m *DynamicSvrSpaceHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSvrSpaceHistoryResponse proto.InternalMessageInfo

func (m *DynamicSvrSpaceHistoryResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DynamicSvrSpaceHistoryResponse) GetData() *DynamicSvrSpaceHistoryResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DynamicSvrSpaceHistoryResponse_Data struct {
	Cards                []*DynamicSvrSpaceHistoryResponse_Data_Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	HasMore              int32                                       `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	NextOffset           int64                                       `protobuf:"varint,3,opt,name=next_offset,json=nextOffset,proto3" json:"next_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *DynamicSvrSpaceHistoryResponse_Data) Reset()         { *m = DynamicSvrSpaceHistoryResponse_Data{} }
func (m *DynamicSvrSpaceHistoryResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DynamicSvrSpaceHistoryResponse_Data) ProtoMessage()    {}
func (*DynamicSvrSpaceHistoryResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{2, 0}
}

func (m *DynamicSvrSpaceHistoryResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data.Unmarshal(m, b)
}
func (m *DynamicSvrSpaceHistoryResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data.Marshal(b, m, deterministic)
}
func (m *DynamicSvrSpaceHistoryResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data.Merge(m, src)
}
func (m *DynamicSvrSpaceHistoryResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data.Size(m)
}
func (m *DynamicSvrSpaceHistoryResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data proto.InternalMessageInfo

func (m *DynamicSvrSpaceHistoryResponse_Data) GetCards() []*DynamicSvrSpaceHistoryResponse_Data_Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DynamicSvrSpaceHistoryResponse_Data) GetHasMore() int32 {
	if m != nil {
		return m.HasMore
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data) GetNextOffset() int64 {
	if m != nil {
		return m.NextOffset
	}
	return 0
}

type DynamicSvrSpaceHistoryResponse_Data_Card struct {
	Card                 string                                          `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Desc                 *DynamicSvrSpaceHistoryResponse_Data_Card_Desc  `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Extra                *DynamicSvrSpaceHistoryResponse_Data_Card_Extra `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card) Reset() {
	*m = DynamicSvrSpaceHistoryResponse_Data_Card{}
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card) String() string { return proto.CompactTextString(m) }
func (*DynamicSvrSpaceHistoryResponse_Data_Card) ProtoMessage()    {}
func (*DynamicSvrSpaceHistoryResponse_Data_Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{2, 0, 0}
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card.Unmarshal(m, b)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card.Marshal(b, m, deterministic)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card.Merge(m, src)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card) XXX_Size() int {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card.Size(m)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card proto.InternalMessageInfo

func (m *DynamicSvrSpaceHistoryResponse_Data_Card) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card) GetDesc() *DynamicSvrSpaceHistoryResponse_Data_Card_Desc {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card) GetExtra() *DynamicSvrSpaceHistoryResponse_Data_Card_Extra {
	if m != nil {
		return m.Extra
	}
	return nil
}

type DynamicSvrSpaceHistoryResponse_Data_Card_Desc struct {
	DynamicId            int64    `protobuf:"varint,1,opt,name=dynamic_id,json=dynamicId,proto3" json:"dynamic_id,omitempty"`
	OrigDyId             int64    `protobuf:"varint,2,opt,name=orig_dy_id,json=origDyId,proto3" json:"orig_dy_id,omitempty"`
	OrigType             int32    `protobuf:"varint,3,opt,name=orig_type,json=origType,proto3" json:"orig_type,omitempty"`
	PreDyId              int32    `protobuf:"varint,4,opt,name=pre_dy_id,json=preDyId,proto3" json:"pre_dy_id,omitempty"`
	Comment              int32    `protobuf:"varint,5,opt,name=comment,proto3" json:"comment,omitempty"`
	Like                 int32    `protobuf:"varint,6,opt,name=like,proto3" json:"like,omitempty"`
	View                 int32    `protobuf:"varint,7,opt,name=view,proto3" json:"view,omitempty"`
	Repost               int32    `protobuf:"varint,8,opt,name=repost,proto3" json:"repost,omitempty"`
	Timestamp            int32    `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) Reset() {
	*m = DynamicSvrSpaceHistoryResponse_Data_Card_Desc{}
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) String() string {
	return proto.CompactTextString(m)
}
func (*DynamicSvrSpaceHistoryResponse_Data_Card_Desc) ProtoMessage() {}
func (*DynamicSvrSpaceHistoryResponse_Data_Card_Desc) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{2, 0, 0, 0}
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Desc.Unmarshal(m, b)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Desc.Marshal(b, m, deterministic)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Desc.Merge(m, src)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) XXX_Size() int {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Desc.Size(m)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Desc.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Desc proto.InternalMessageInfo

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetDynamicId() int64 {
	if m != nil {
		return m.DynamicId
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetOrigDyId() int64 {
	if m != nil {
		return m.OrigDyId
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetOrigType() int32 {
	if m != nil {
		return m.OrigType
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetPreDyId() int32 {
	if m != nil {
		return m.PreDyId
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetComment() int32 {
	if m != nil {
		return m.Comment
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetLike() int32 {
	if m != nil {
		return m.Like
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetView() int32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetRepost() int32 {
	if m != nil {
		return m.Repost
	}
	return 0
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Desc) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type DynamicSvrSpaceHistoryResponse_Data_Card_Extra struct {
	IsSpaceTop           int32    `protobuf:"varint,1,opt,name=is_space_top,json=isSpaceTop,proto3" json:"is_space_top,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) Reset() {
	*m = DynamicSvrSpaceHistoryResponse_Data_Card_Extra{}
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) String() string {
	return proto.CompactTextString(m)
}
func (*DynamicSvrSpaceHistoryResponse_Data_Card_Extra) ProtoMessage() {}
func (*DynamicSvrSpaceHistoryResponse_Data_Card_Extra) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0021bc90478bd43, []int{2, 0, 0, 1}
}

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Extra.Unmarshal(m, b)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Extra.Marshal(b, m, deterministic)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Extra.Merge(m, src)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) XXX_Size() int {
	return xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Extra.Size(m)
}
func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Extra.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSvrSpaceHistoryResponse_Data_Card_Extra proto.InternalMessageInfo

func (m *DynamicSvrSpaceHistoryResponse_Data_Card_Extra) GetIsSpaceTop() int32 {
	if m != nil {
		return m.IsSpaceTop
	}
	return 0
}

func init() {
	proto.RegisterEnum("bilibili.LiveStatus", LiveStatus_name, LiveStatus_value)
	proto.RegisterEnum("bilibili.RoundStatus", RoundStatus_name, RoundStatus_value)
	proto.RegisterEnum("bilibili.RoomStatus", RoomStatus_name, RoomStatus_value)
	proto.RegisterEnum("bilibili.SpecialType", SpecialType_name, SpecialType_value)
	proto.RegisterType((*RoomInitResponse)(nil), "bilibili.RoomInitResponse")
	proto.RegisterType((*RoomInitResponse_Data)(nil), "bilibili.RoomInitResponse.Data")
	proto.RegisterType((*XSpaceAccInfoResponse)(nil), "bilibili.XSpaceAccInfoResponse")
	proto.RegisterType((*XSpaceAccInfoResponse_Data)(nil), "bilibili.XSpaceAccInfoResponse.Data")
	proto.RegisterType((*XSpaceAccInfoResponse_Data_LiveRoom)(nil), "bilibili.XSpaceAccInfoResponse.Data.LiveRoom")
	proto.RegisterType((*DynamicSvrSpaceHistoryResponse)(nil), "bilibili.DynamicSvrSpaceHistoryResponse")
	proto.RegisterType((*DynamicSvrSpaceHistoryResponse_Data)(nil), "bilibili.DynamicSvrSpaceHistoryResponse.Data")
	proto.RegisterType((*DynamicSvrSpaceHistoryResponse_Data_Card)(nil), "bilibili.DynamicSvrSpaceHistoryResponse.Data.Card")
	proto.RegisterType((*DynamicSvrSpaceHistoryResponse_Data_Card_Desc)(nil), "bilibili.DynamicSvrSpaceHistoryResponse.Data.Card.Desc")
	proto.RegisterType((*DynamicSvrSpaceHistoryResponse_Data_Card_Extra)(nil), "bilibili.DynamicSvrSpaceHistoryResponse.Data.Card.Extra")
}

func init() {
	proto.RegisterFile("bilibili.proto", fileDescriptor_d0021bc90478bd43)
}

var fileDescriptor_d0021bc90478bd43 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xe4, 0x44,
	0x10, 0x8e, 0xc7, 0xe3, 0x8c, 0x5d, 0x8e, 0xc0, 0x34, 0x7f, 0x66, 0x58, 0xb2, 0x43, 0x04, 0x68,
	0x88, 0xc4, 0x1c, 0x66, 0x91, 0xd8, 0xeb, 0x42, 0x16, 0x65, 0x20, 0x24, 0xa8, 0x67, 0x0f, 0xc0,
	0xc5, 0xea, 0xb5, 0x6b, 0x92, 0xd6, 0xda, 0x6e, 0xab, 0xbb, 0x33, 0x64, 0x5e, 0x8f, 0x57, 0x40,
	0x08, 0x71, 0xe2, 0xca, 0x03, 0x70, 0xe4, 0x80, 0xba, 0xec, 0x89, 0x27, 0x08, 0x21, 0x36, 0x87,
	0x48, 0x55, 0x5f, 0xc5, 0x5f, 0x55, 0x7f, 0xf5, 0x33, 0xf0, 0xca, 0x73, 0x59, 0x4a, 0xf7, 0x37,
	0x6b, 0xb4, 0xb2, 0x8a, 0x85, 0x5b, 0xff, 0xe8, 0xd7, 0x01, 0x24, 0x5c, 0xa9, 0x6a, 0x51, 0x4b,
	0xcb, 0xd1, 0x34, 0xaa, 0x36, 0xc8, 0x18, 0x0c, 0x73, 0x55, 0x60, 0xea, 0x4d, 0xbc, 0x69, 0xc0,
	0xc9, 0x66, 0x29, 0x8c, 0x2a, 0x34, 0x46, 0x5c, 0x62, 0x3a, 0x98, 0x78, 0xd3, 0x88, 0x6f, 0x5d,
	0xf6, 0x08, 0x86, 0x85, 0xb0, 0x22, 0xf5, 0x27, 0xde, 0x34, 0x9e, 0x3f, 0x9c, 0xdd, 0xe6, 0xfa,
	0x27, 0xef, 0xec, 0x44, 0x58, 0xc1, 0xe9, 0x9f, 0xc7, 0xbf, 0x7b, 0x30, 0x74, 0x2e, 0x7b, 0x1b,
	0x46, 0x5a, 0xa9, 0x2a, 0x93, 0x05, 0xa5, 0xf3, 0xf9, 0xbe, 0x73, 0x17, 0x05, 0x7b, 0x07, 0x42,
	0x73, 0xa5, 0xb4, 0x75, 0x91, 0x01, 0x45, 0x46, 0xe4, 0x2f, 0x0a, 0x96, 0x80, 0x7f, 0x2d, 0x0b,
	0x4a, 0xe8, 0x73, 0x67, 0xb2, 0x77, 0x21, 0x92, 0x26, 0xbb, 0x92, 0x45, 0x81, 0x75, 0x1a, 0x4c,
	0xbc, 0x69, 0xc8, 0x43, 0x69, 0x4e, 0xc9, 0xef, 0x82, 0xa5, 0xca, 0x5f, 0x60, 0x91, 0xee, 0x6f,
	0x83, 0x67, 0xe4, 0xb3, 0x87, 0x10, 0x97, 0x72, 0x8d, 0x99, 0xb1, 0xc2, 0x5e, 0x9b, 0x34, 0xa4,
	0x27, 0x83, 0x83, 0x96, 0x84, 0xb0, 0xd7, 0x21, 0x90, 0x26, 0x33, 0x4d, 0xfa, 0x6a, 0xab, 0x86,
	0x34, 0xcb, 0x86, 0xbd, 0x0f, 0x07, 0xa6, 0xc1, 0x5c, 0x8a, 0x32, 0xb3, 0x9b, 0x06, 0xd3, 0x84,
	0x62, 0x71, 0x87, 0x3d, 0xdb, 0x34, 0x78, 0xf4, 0x8b, 0x0f, 0x6f, 0x7e, 0xb7, 0x6c, 0x44, 0x8e,
	0x4f, 0xf2, 0x7c, 0x51, 0xaf, 0xd4, 0x3d, 0xe5, 0x7d, 0x7c, 0x47, 0xde, 0x0f, 0x7a, 0x79, 0xff,
	0x95, 0x7c, 0x57, 0xe3, 0x9f, 0x07, 0x9d, 0xc6, 0x09, 0xf8, 0xd5, 0xad, 0xbe, 0xce, 0x74, 0x25,
	0xd4, 0xa2, 0xda, 0xe6, 0x22, 0x9b, 0xbd, 0x01, 0x41, 0x89, 0x6b, 0x2c, 0xd3, 0x11, 0xd5, 0xd5,
	0x3a, 0xec, 0x2b, 0x88, 0x48, 0x1f, 0xd7, 0x95, 0xf4, 0x35, 0xaa, 0xe1, 0x93, 0xff, 0x53, 0xc3,
	0xec, 0x4c, 0xae, 0xd1, 0x4d, 0x00, 0x0f, 0xcb, 0xce, 0x1a, 0xff, 0xe4, 0x41, 0xb8, 0x85, 0xd9,
	0x21, 0x80, 0xe3, 0x6c, 0x55, 0xee, 0xb4, 0xd8, 0x41, 0x5c, 0xbc, 0xef, 0x02, 0x15, 0x7a, 0xb7,
	0x2f, 0x6e, 0x08, 0x74, 0x49, 0xb2, 0x44, 0xdc, 0x99, 0xee, 0x01, 0x56, 0xda, 0x12, 0xd3, 0x21,
	0x61, 0xad, 0xe3, 0xd0, 0x5c, 0xad, 0x51, 0xd3, 0x58, 0x44, 0xbc, 0x75, 0xd8, 0x5b, 0x40, 0x73,
	0x26, 0xdb, 0x81, 0xe8, 0xa6, 0x4e, 0x16, 0x6c, 0x02, 0xb1, 0x56, 0xd7, 0x75, 0xd1, 0xa5, 0x6d,
	0xa5, 0xd8, 0x85, 0x8e, 0xfe, 0x0a, 0xe0, 0xf0, 0x64, 0x53, 0x8b, 0x4a, 0xe6, 0xcb, 0xb5, 0xa6,
	0xf7, 0x9f, 0x4a, 0x63, 0x95, 0xde, 0xdc, 0xb3, 0xc1, 0x4f, 0xee, 0x34, 0x78, 0x47, 0xdc, 0xff,
	0xce, 0xb2, 0xdb, 0xe9, 0x3f, 0x86, 0x5d, 0xa7, 0x4f, 0x21, 0xc8, 0x85, 0x2e, 0x9c, 0x9e, 0xfe,
	0x34, 0x9e, 0xcf, 0x5f, 0x8a, 0x6c, 0xf6, 0x85, 0xd0, 0x05, 0x6f, 0x09, 0xdc, 0xfa, 0x5d, 0x09,
	0x93, 0x55, 0x4a, 0x63, 0x27, 0xfe, 0xe8, 0x4a, 0x98, 0x6f, 0x94, 0x46, 0xb7, 0x32, 0x35, 0xde,
	0xd8, 0x4c, 0xad, 0x56, 0x06, 0x6d, 0xb7, 0x86, 0xe0, 0xa0, 0x0b, 0x42, 0xc6, 0xbf, 0xf9, 0x30,
	0x74, 0x5c, 0x24, 0x84, 0xd0, 0xed, 0xe4, 0x45, 0x9c, 0x6c, 0xf6, 0x35, 0x0c, 0x0b, 0x34, 0x39,
	0x91, 0xc6, 0xf3, 0xcf, 0x5e, 0xbe, 0xc2, 0xd9, 0x09, 0x9a, 0x9c, 0x13, 0x09, 0x3b, 0x87, 0x00,
	0x6f, 0xac, 0x16, 0xd4, 0xdc, 0x78, 0xfe, 0xf8, 0x1e, 0x6c, 0x4f, 0xdd, 0xf7, 0xbc, 0xa5, 0x19,
	0xff, 0xe9, 0xce, 0x92, 0x23, 0x7e, 0x0f, 0xa0, 0x68, 0x19, 0xfa, 0xcb, 0x14, 0x75, 0xc8, 0xa2,
	0x60, 0x0f, 0x00, 0x94, 0x96, 0x97, 0x59, 0xb1, 0xe9, 0xcf, 0x53, 0xe8, 0x90, 0x93, 0xcd, 0x82,
	0xae, 0x11, 0x45, 0xe9, 0x34, 0xf8, 0x24, 0x1e, 0x05, 0xdd, 0x5d, 0x60, 0x63, 0x88, 0x1a, 0x8d,
	0xdd, 0x97, 0xc3, 0x56, 0xd9, 0x46, 0x23, 0x7d, 0x98, 0xc2, 0x28, 0x57, 0x55, 0x85, 0xb5, 0xa5,
	0x07, 0x05, 0x7c, 0xeb, 0x3a, 0x25, 0x4b, 0xf9, 0x02, 0x69, 0x5a, 0x03, 0x4e, 0xb6, 0xc3, 0xd6,
	0x12, 0x7f, 0xec, 0x86, 0x94, 0x6c, 0x9a, 0x6b, 0x6c, 0x94, 0xb1, 0xdd, 0x25, 0xeb, 0x3c, 0xf6,
	0x00, 0x22, 0x2b, 0x2b, 0x34, 0x56, 0x54, 0x4d, 0x1a, 0x51, 0xa8, 0x07, 0xc6, 0x1f, 0x43, 0x40,
	0x32, 0xb0, 0x09, 0x1c, 0xd0, 0xb1, 0x13, 0x39, 0x66, 0x56, 0x35, 0xdb, 0xb5, 0x74, 0x37, 0x4f,
	0xe4, 0xf8, 0x4c, 0x35, 0xc7, 0x1f, 0x01, 0x9c, 0xf5, 0x4b, 0x78, 0x00, 0xe1, 0xb9, 0x3a, 0x93,
	0x6b, 0x59, 0x5f, 0x26, 0x7b, 0x0c, 0x60, 0xbf, 0xb3, 0xbd, 0xe3, 0x43, 0x88, 0x79, 0xbf, 0x35,
	0x6c, 0x04, 0xfe, 0xc5, 0x6a, 0x95, 0xec, 0xb1, 0x7d, 0x18, 0x5c, 0xd4, 0x89, 0x77, 0xfc, 0x21,
	0x00, 0xef, 0x97, 0x9d, 0x78, 0xea, 0xa7, 0x37, 0xd2, 0xd8, 0x64, 0x8f, 0x45, 0xae, 0x1c, 0x67,
	0x7a, 0xc7, 0x9f, 0x42, 0xbc, 0xec, 0x8f, 0xaa, 0xcb, 0x70, 0xae, 0x74, 0x25, 0xca, 0x64, 0xcf,
	0x51, 0x7e, 0x2b, 0x36, 0x89, 0xc7, 0x12, 0x38, 0xf8, 0x1e, 0x85, 0xfe, 0x12, 0x8d, 0x95, 0x6b,
	0x51, 0x26, 0x83, 0xcf, 0xe1, 0x87, 0xdb, 0x5f, 0xb8, 0xe7, 0xfb, 0xf4, 0x93, 0xf7, 0xe8, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x07, 0x78, 0x9b, 0x04, 0x07, 0x00, 0x00,
}