// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/accountplus/genesis.proto

package types

import (
	fmt "fmt"
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

// AuthenticatorData represents a genesis exported account with Authenticators.
// The address is used as the key, and the account authenticators are stored in
// the authenticators field.
type AuthenticatorData struct {
	// address is an account address, one address can have many authenticators
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// authenticators are the account's authenticators, these can be multiple
	// types including SignatureVerification, AllOfs, CosmWasmAuthenticators, etc
	Authenticators []AccountAuthenticator `protobuf:"bytes,2,rep,name=authenticators,proto3" json:"authenticators"`
}

func (m *AuthenticatorData) Reset()         { *m = AuthenticatorData{} }
func (m *AuthenticatorData) String() string { return proto.CompactTextString(m) }
func (*AuthenticatorData) ProtoMessage()    {}
func (*AuthenticatorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d588eda993a7b7, []int{0}
}
func (m *AuthenticatorData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticatorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticatorData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticatorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticatorData.Merge(m, src)
}
func (m *AuthenticatorData) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticatorData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticatorData.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticatorData proto.InternalMessageInfo

func (m *AuthenticatorData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AuthenticatorData) GetAuthenticators() []AccountAuthenticator {
	if m != nil {
		return m.Authenticators
	}
	return nil
}

// Module genesis state
type GenesisState struct {
	Accounts []AccountState `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts"`
	// params define the parameters for the authenticator module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
	// next_authenticator_id is the next available authenticator ID.
	NextAuthenticatorId uint64 `protobuf:"varint,3,opt,name=next_authenticator_id,json=nextAuthenticatorId,proto3" json:"next_authenticator_id,omitempty"`
	// authenticator_data contains the data for multiple accounts, each with their
	// authenticators.
	AuthenticatorData []AuthenticatorData `protobuf:"bytes,4,rep,name=authenticator_data,json=authenticatorData,proto3" json:"authenticator_data"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d588eda993a7b7, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAccounts() []AccountState {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetNextAuthenticatorId() uint64 {
	if m != nil {
		return m.NextAuthenticatorId
	}
	return 0
}

func (m *GenesisState) GetAuthenticatorData() []AuthenticatorData {
	if m != nil {
		return m.AuthenticatorData
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticatorData)(nil), "bitoroprotocol.accountplus.AuthenticatorData")
	proto.RegisterType((*GenesisState)(nil), "bitoroprotocol.accountplus.GenesisState")
}

func init() {
	proto.RegisterFile("bitoroprotocol/accountplus/genesis.proto", fileDescriptor_c1d588eda993a7b7)
}

var fileDescriptor_c1d588eda993a7b7 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4a, 0xc3, 0x50,
	0x14, 0x86, 0x73, 0xdb, 0x52, 0xf5, 0x56, 0x84, 0x5e, 0x15, 0x42, 0x87, 0x18, 0xba, 0x98, 0xc1,
	0x26, 0x52, 0x57, 0x07, 0x2b, 0x82, 0xe8, 0x24, 0x15, 0x1c, 0x1c, 0x2c, 0x27, 0xc9, 0x25, 0x0d,
	0xb6, 0xb9, 0x21, 0x39, 0xc1, 0xfa, 0x10, 0x82, 0x0f, 0xe3, 0x43, 0x74, 0xec, 0xe8, 0x24, 0xd2,
	0xbe, 0x88, 0x24, 0xb9, 0x96, 0xa6, 0x62, 0xe8, 0x76, 0xca, 0xff, 0x9d, 0xff, 0x3f, 0x7f, 0x73,
	0xa9, 0x61, 0xfb, 0x28, 0x22, 0x11, 0x46, 0x02, 0x85, 0x23, 0x46, 0x16, 0x38, 0x8e, 0x48, 0x02,
	0x0c, 0x47, 0x49, 0x6c, 0x79, 0x3c, 0xe0, 0xb1, 0x1f, 0x9b, 0x99, 0xc8, 0x5a, 0x45, 0xd2, 0x5c,
	0x21, 0x5b, 0x07, 0x9e, 0xf0, 0x72, 0xc5, 0x4a, 0xa7, 0x7c, 0xa3, 0x75, 0x52, 0xe2, 0xbd, 0x32,
	0x4b, 0xfa, 0xb8, 0x84, 0x1e, 0x0b, 0x97, 0x8f, 0x36, 0x01, 0x43, 0x88, 0x60, 0x2c, 0xc1, 0xf6,
	0x1b, 0xa1, 0xcd, 0x5e, 0x82, 0x43, 0x1e, 0xa0, 0xef, 0x00, 0x8a, 0xe8, 0x0a, 0x10, 0x98, 0x4a,
	0xb7, 0xc0, 0x75, 0x23, 0x1e, 0xc7, 0x2a, 0xd1, 0x89, 0xb1, 0xd3, 0xff, 0xfd, 0xc9, 0x9e, 0xe8,
	0x1e, 0xac, 0xe2, 0xb1, 0x5a, 0xd1, 0xab, 0x46, 0xa3, 0x7b, 0x6a, 0xfe, 0x5f, 0xdd, 0xec, 0xe5,
	0x73, 0x21, 0xe7, 0xb2, 0x36, 0xfd, 0x3a, 0x52, 0xfa, 0x6b, 0x6e, 0xed, 0x8f, 0x0a, 0xdd, 0xbd,
	0xce, 0xff, 0xd3, 0x7b, 0x04, 0xe4, 0xec, 0x96, 0x6e, 0x4b, 0xab, 0xf4, 0x96, 0x34, 0xca, 0xd8,
	0x20, 0x2a, 0xdb, 0x95, 0x11, 0xcb, 0x7d, 0x76, 0x41, 0xeb, 0x79, 0x79, 0xb5, 0xa2, 0x13, 0xa3,
	0xd1, 0x6d, 0x97, 0x39, 0xdd, 0x65, 0xa4, 0xf4, 0x90, 0x7b, 0xac, 0x4b, 0x0f, 0x03, 0x3e, 0xc1,
	0x41, 0xe1, 0xea, 0x81, 0xef, 0xaa, 0x55, 0x9d, 0x18, 0xb5, 0xfe, 0x7e, 0x2a, 0x16, 0x6a, 0xde,
	0xb8, 0xcc, 0xa6, 0xac, 0x88, 0xbb, 0x80, 0xa0, 0xd6, 0xb2, 0x2e, 0x9d, 0xd2, 0x2e, 0xeb, 0xdf,
	0x45, 0x1e, 0xd3, 0x84, 0x3f, 0xc2, 0xc3, 0x74, 0xae, 0x91, 0xd9, 0x5c, 0x23, 0xdf, 0x73, 0x8d,
	0xbc, 0x2f, 0x34, 0x65, 0xb6, 0xd0, 0x94, 0xcf, 0x85, 0xa6, 0x3c, 0x9e, 0x7b, 0x3e, 0x0e, 0x13,
	0xdb, 0x74, 0xc4, 0xd8, 0xca, 0xb3, 0x3a, 0x01, 0xc7, 0x17, 0x11, 0x3d, 0x5b, 0xce, 0x10, 0xfc,
	0xc0, 0x5a, 0xbe, 0x91, 0x49, 0xe1, 0x95, 0xe0, 0x6b, 0xc8, 0x63, 0xbb, 0x9e, 0xa9, 0x67, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x53, 0xb3, 0x1b, 0x03, 0x03, 0x00, 0x00,
}

func (m *AuthenticatorData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticatorData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticatorData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authenticators) > 0 {
		for iNdEx := len(m.Authenticators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authenticators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticatorData) > 0 {
		for iNdEx := len(m.AuthenticatorData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuthenticatorData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.NextAuthenticatorId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextAuthenticatorId))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthenticatorData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Authenticators) > 0 {
		for _, e := range m.Authenticators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, e := range m.Accounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.NextAuthenticatorId != 0 {
		n += 1 + sovGenesis(uint64(m.NextAuthenticatorId))
	}
	if len(m.AuthenticatorData) > 0 {
		for _, e := range m.AuthenticatorData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthenticatorData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: AuthenticatorData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticatorData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authenticators = append(m.Authenticators, AccountAuthenticator{})
			if err := m.Authenticators[len(m.Authenticators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, AccountState{})
			if err := m.Accounts[len(m.Accounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAuthenticatorId", wireType)
			}
			m.NextAuthenticatorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextAuthenticatorId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticatorData = append(m.AuthenticatorData, AuthenticatorData{})
			if err := m.AuthenticatorData[len(m.AuthenticatorData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
