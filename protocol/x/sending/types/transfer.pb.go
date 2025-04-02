// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/sending/transfer.proto

package types

import (
	fmt "fmt"
	types "github.com/bitoro-network/chain/protocol/x/subaccounts/types"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Transfer represents a single transfer between two subaccounts.
type Transfer struct {
	// The sender subaccount ID.
	Sender types.SubaccountId `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender"`
	// The recipient subaccount ID.
	Recipient types.SubaccountId `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient"`
	// Id of the asset to transfer.
	AssetId uint32 `protobuf:"varint,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The amount of asset to transfer
	Amount uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5df1f9e206fc6e2, []int{0}
}
func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return m.Size()
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetSender() types.SubaccountId {
	if m != nil {
		return m.Sender
	}
	return types.SubaccountId{}
}

func (m *Transfer) GetRecipient() types.SubaccountId {
	if m != nil {
		return m.Recipient
	}
	return types.SubaccountId{}
}

func (m *Transfer) GetAssetId() uint32 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *Transfer) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// MsgDepositToSubaccount represents a single transfer from an `x/bank`
// account to an `x/subaccounts` subaccount.
type MsgDepositToSubaccount struct {
	// The sender wallet address.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// The recipient subaccount ID.
	Recipient types.SubaccountId `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient"`
	// Id of the asset to transfer.
	AssetId uint32 `protobuf:"varint,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The number of quantums of asset to transfer.
	Quantums uint64 `protobuf:"varint,4,opt,name=quantums,proto3" json:"quantums,omitempty"`
}

func (m *MsgDepositToSubaccount) Reset()         { *m = MsgDepositToSubaccount{} }
func (m *MsgDepositToSubaccount) String() string { return proto.CompactTextString(m) }
func (*MsgDepositToSubaccount) ProtoMessage()    {}
func (*MsgDepositToSubaccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5df1f9e206fc6e2, []int{1}
}
func (m *MsgDepositToSubaccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositToSubaccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositToSubaccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositToSubaccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositToSubaccount.Merge(m, src)
}
func (m *MsgDepositToSubaccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositToSubaccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositToSubaccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositToSubaccount proto.InternalMessageInfo

func (m *MsgDepositToSubaccount) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgDepositToSubaccount) GetRecipient() types.SubaccountId {
	if m != nil {
		return m.Recipient
	}
	return types.SubaccountId{}
}

func (m *MsgDepositToSubaccount) GetAssetId() uint32 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *MsgDepositToSubaccount) GetQuantums() uint64 {
	if m != nil {
		return m.Quantums
	}
	return 0
}

// MsgWithdrawFromSubaccount represents a single transfer from an
// `x/subaccounts` subaccount to an `x/bank` account.
type MsgWithdrawFromSubaccount struct {
	// The sender subaccount ID.
	Sender types.SubaccountId `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender"`
	// The recipient wallet address.
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// Id of the asset to transfer.
	AssetId uint32 `protobuf:"varint,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The number of quantums of asset to transfer.
	Quantums uint64 `protobuf:"varint,4,opt,name=quantums,proto3" json:"quantums,omitempty"`
}

func (m *MsgWithdrawFromSubaccount) Reset()         { *m = MsgWithdrawFromSubaccount{} }
func (m *MsgWithdrawFromSubaccount) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFromSubaccount) ProtoMessage()    {}
func (*MsgWithdrawFromSubaccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5df1f9e206fc6e2, []int{2}
}
func (m *MsgWithdrawFromSubaccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawFromSubaccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawFromSubaccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawFromSubaccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFromSubaccount.Merge(m, src)
}
func (m *MsgWithdrawFromSubaccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawFromSubaccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFromSubaccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFromSubaccount proto.InternalMessageInfo

func (m *MsgWithdrawFromSubaccount) GetSender() types.SubaccountId {
	if m != nil {
		return m.Sender
	}
	return types.SubaccountId{}
}

func (m *MsgWithdrawFromSubaccount) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *MsgWithdrawFromSubaccount) GetAssetId() uint32 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *MsgWithdrawFromSubaccount) GetQuantums() uint64 {
	if m != nil {
		return m.Quantums
	}
	return 0
}

// MsgSendFromModuleToAccount represents a single transfer from a module
// to an `x/bank` account (can be either a module account address or a user
// account address).
// Should only be executed by governance.
type MsgSendFromModuleToAccount struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// The sender module name.
	SenderModuleName string `protobuf:"bytes,2,opt,name=sender_module_name,json=senderModuleName,proto3" json:"sender_module_name,omitempty"`
	// The recipient account address (can be either a module account address
	// or a user account address).
	Recipient string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// The coin to transfer, which specifies both denom and amount.
	Coin types1.Coin `protobuf:"bytes,4,opt,name=coin,proto3" json:"coin"`
}

func (m *MsgSendFromModuleToAccount) Reset()         { *m = MsgSendFromModuleToAccount{} }
func (m *MsgSendFromModuleToAccount) String() string { return proto.CompactTextString(m) }
func (*MsgSendFromModuleToAccount) ProtoMessage()    {}
func (*MsgSendFromModuleToAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5df1f9e206fc6e2, []int{3}
}
func (m *MsgSendFromModuleToAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendFromModuleToAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendFromModuleToAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendFromModuleToAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendFromModuleToAccount.Merge(m, src)
}
func (m *MsgSendFromModuleToAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendFromModuleToAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendFromModuleToAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendFromModuleToAccount proto.InternalMessageInfo

func (m *MsgSendFromModuleToAccount) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSendFromModuleToAccount) GetSenderModuleName() string {
	if m != nil {
		return m.SenderModuleName
	}
	return ""
}

func (m *MsgSendFromModuleToAccount) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *MsgSendFromModuleToAccount) GetCoin() types1.Coin {
	if m != nil {
		return m.Coin
	}
	return types1.Coin{}
}

func init() {
	proto.RegisterType((*Transfer)(nil), "bitoroprotocol.sending.Transfer")
	proto.RegisterType((*MsgDepositToSubaccount)(nil), "bitoroprotocol.sending.MsgDepositToSubaccount")
	proto.RegisterType((*MsgWithdrawFromSubaccount)(nil), "bitoroprotocol.sending.MsgWithdrawFromSubaccount")
	proto.RegisterType((*MsgSendFromModuleToAccount)(nil), "bitoroprotocol.sending.MsgSendFromModuleToAccount")
}

func init() {
	proto.RegisterFile("bitoroprotocol/sending/transfer.proto", fileDescriptor_e5df1f9e206fc6e2)
}

var fileDescriptor_e5df1f9e206fc6e2 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x6d, 0x88, 0xc9, 0x04, 0x45, 0x96, 0x12, 0x93, 0x1c, 0xd6, 0x10, 0x10, 0x82,
	0xda, 0x1d, 0xd3, 0x42, 0x91, 0xde, 0x1a, 0xa5, 0x50, 0x30, 0x1e, 0x36, 0x01, 0xc1, 0x4b, 0x98,
	0xdd, 0x1d, 0x37, 0x83, 0xdd, 0x79, 0x71, 0x66, 0xb6, 0xb5, 0x57, 0xff, 0x02, 0xff, 0x14, 0x0f,
	0xfe, 0x11, 0x3d, 0x16, 0xbd, 0x78, 0x10, 0x95, 0xe4, 0xe0, 0x9f, 0xa1, 0xec, 0xce, 0x90, 0x4d,
	0x7a, 0xa9, 0xa8, 0xf4, 0x94, 0x79, 0xf9, 0xbe, 0x5f, 0x9f, 0xf7, 0x78, 0x8b, 0xef, 0x05, 0x5c,
	0x83, 0x84, 0x99, 0x04, 0x0d, 0x21, 0x1c, 0x13, 0xc5, 0x44, 0xc4, 0x45, 0x4c, 0xb4, 0xa4, 0x42,
	0xbd, 0x62, 0xd2, 0xcb, 0x15, 0xa7, 0xb1, 0xee, 0xe6, 0x59, 0xb7, 0x76, 0x2b, 0x04, 0x95, 0x80,
	0x9a, 0xe4, 0x02, 0x31, 0x86, 0x09, 0x69, 0xbb, 0xc6, 0x22, 0x01, 0x55, 0x8c, 0x9c, 0xf4, 0x03,
	0xa6, 0x69, 0x9f, 0x84, 0xc0, 0x85, 0xd5, 0xef, 0x58, 0x3d, 0x51, 0x31, 0x39, 0xe9, 0x67, 0x3f,
	0x56, 0xd8, 0x8a, 0x21, 0x36, 0x95, 0x48, 0xf6, 0xb2, 0xff, 0x3e, 0xb8, 0xdc, 0x68, 0x1a, 0xd0,
	0x30, 0x84, 0x54, 0x68, 0xb5, 0xf2, 0x36, 0xce, 0xdd, 0xcf, 0x08, 0x57, 0xc7, 0x96, 0xc0, 0x39,
	0xc4, 0x95, 0xac, 0x5d, 0x26, 0x9b, 0xa8, 0x83, 0x7a, 0xf5, 0x9d, 0x9e, 0x77, 0x19, 0xa6, 0x48,
	0xe5, 0x8d, 0x96, 0xef, 0xa3, 0x68, 0x50, 0x3e, 0xff, 0x76, 0xb7, 0xe4, 0xdb, 0x68, 0xe7, 0x19,
	0xae, 0x49, 0x16, 0xf2, 0x19, 0x67, 0x42, 0x37, 0x37, 0xfe, 0x2a, 0x55, 0x91, 0xc0, 0x69, 0xe1,
	0x2a, 0x55, 0x8a, 0xe9, 0x09, 0x8f, 0x9a, 0x9b, 0x1d, 0xd4, 0xbb, 0xe9, 0xdf, 0xc8, 0xed, 0xa3,
	0xc8, 0x69, 0xe0, 0x0a, 0x4d, 0xb2, 0xb8, 0x66, 0xb9, 0x83, 0x7a, 0x65, 0xdf, 0x5a, 0xdd, 0xaf,
	0x08, 0x37, 0x86, 0x2a, 0x7e, 0xca, 0x66, 0xa0, 0xb8, 0x1e, 0x43, 0x51, 0xc0, 0x79, 0xb4, 0xc6,
	0x58, 0x1b, 0x34, 0x3f, 0x7d, 0xdc, 0xde, 0xb2, 0xeb, 0x38, 0x88, 0x22, 0xc9, 0x94, 0x1a, 0x69,
	0xc9, 0x45, 0x7c, 0xfd, 0x34, 0x6d, 0x5c, 0x7d, 0x93, 0x52, 0xa1, 0xd3, 0x44, 0x59, 0x9e, 0xa5,
	0xbd, 0x5f, 0x7f, 0xf7, 0xf3, 0xc3, 0x7d, 0xdb, 0x51, 0xf7, 0x3b, 0xc2, 0xad, 0xa1, 0x8a, 0x5f,
	0x70, 0x3d, 0x8d, 0x24, 0x3d, 0x3d, 0x94, 0x90, 0xac, 0x10, 0x16, 0x5b, 0xdc, 0xf8, 0xa7, 0x2d,
	0xee, 0xad, 0x72, 0x5f, 0x35, 0xac, 0xff, 0x4c, 0xf8, 0x0b, 0xe1, 0xf6, 0x50, 0xc5, 0x23, 0x26,
	0xa2, 0x8c, 0x6e, 0x08, 0x51, 0x7a, 0xcc, 0xc6, 0x70, 0x60, 0x11, 0xf7, 0x70, 0x8d, 0xa6, 0x7a,
	0x0a, 0x92, 0xeb, 0xb3, 0xab, 0x5b, 0x5b, 0xba, 0x3a, 0x0f, 0xb1, 0x63, 0x0a, 0x4c, 0x92, 0x3c,
	0xe3, 0x44, 0xd0, 0x84, 0xe5, 0x63, 0xaa, 0xf9, 0xb7, 0x8d, 0x62, 0x4a, 0x3d, 0xa7, 0x09, 0x5b,
	0x1f, 0xc0, 0xe6, 0x9f, 0x0f, 0x60, 0x17, 0x97, 0xb3, 0xeb, 0xcd, 0x09, 0xeb, 0x3b, 0x2d, 0xcf,
	0xfa, 0x67, 0xe7, 0xed, 0xd9, 0xf3, 0xf6, 0x9e, 0x00, 0x17, 0x76, 0xde, 0xb9, 0xf3, 0xfe, 0xad,
	0x0c, 0xbf, 0x68, 0x75, 0xe0, 0x9f, 0xcf, 0x5d, 0x74, 0x31, 0x77, 0xd1, 0x8f, 0xb9, 0x8b, 0xde,
	0x2f, 0xdc, 0xd2, 0xc5, 0xc2, 0x2d, 0x7d, 0x59, 0xb8, 0xa5, 0x97, 0x8f, 0x63, 0xae, 0xa7, 0x69,
	0xe0, 0x85, 0x90, 0x10, 0xb3, 0xd9, 0x6d, 0xc1, 0xf4, 0x29, 0xc8, 0xd7, 0x24, 0x9c, 0x52, 0x2e,
	0xc8, 0xf2, 0xf2, 0xdf, 0x16, 0x1f, 0xa9, 0xb3, 0x19, 0x53, 0x41, 0x25, 0x57, 0x76, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0xf7, 0x30, 0x0f, 0xcb, 0x04, 0x00, 0x00,
}

func (m *Transfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if m.AssetId != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Recipient.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTransfer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Sender.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTransfer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgDepositToSubaccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositToSubaccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositToSubaccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Quantums != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.Quantums))
		i--
		dAtA[i] = 0x20
	}
	if m.AssetId != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Recipient.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTransfer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawFromSubaccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawFromSubaccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawFromSubaccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Quantums != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.Quantums))
		i--
		dAtA[i] = 0x20
	}
	if m.AssetId != 0 {
		i = encodeVarintTransfer(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Sender.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTransfer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendFromModuleToAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendFromModuleToAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendFromModuleToAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTransfer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SenderModuleName) > 0 {
		i -= len(m.SenderModuleName)
		copy(dAtA[i:], m.SenderModuleName)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.SenderModuleName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTransfer(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Sender.Size()
	n += 1 + l + sovTransfer(uint64(l))
	l = m.Recipient.Size()
	n += 1 + l + sovTransfer(uint64(l))
	if m.AssetId != 0 {
		n += 1 + sovTransfer(uint64(m.AssetId))
	}
	if m.Amount != 0 {
		n += 1 + sovTransfer(uint64(m.Amount))
	}
	return n
}

func (m *MsgDepositToSubaccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = m.Recipient.Size()
	n += 1 + l + sovTransfer(uint64(l))
	if m.AssetId != 0 {
		n += 1 + sovTransfer(uint64(m.AssetId))
	}
	if m.Quantums != 0 {
		n += 1 + sovTransfer(uint64(m.Quantums))
	}
	return n
}

func (m *MsgWithdrawFromSubaccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = m.Sender.Size()
	n += 1 + l + sovTransfer(uint64(l))
	if m.AssetId != 0 {
		n += 1 + sovTransfer(uint64(m.AssetId))
	}
	if m.Quantums != 0 {
		n += 1 + sovTransfer(uint64(m.Quantums))
	}
	return n
}

func (m *MsgSendFromModuleToAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.SenderModuleName)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTransfer(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovTransfer(uint64(l))
	return n
}

func sovTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransfer(x uint64) (n int) {
	return sovTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Transfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Recipient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDepositToSubaccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDepositToSubaccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositToSubaccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Recipient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			m.Quantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgWithdrawFromSubaccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgWithdrawFromSubaccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawFromSubaccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			m.Quantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendFromModuleToAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransfer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendFromModuleToAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendFromModuleToAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransfer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransfer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransfer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTransfer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransfer = fmt.Errorf("proto: unexpected end of group")
)
