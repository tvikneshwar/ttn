// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto
// DO NOT EDIT!

/*
	Package lorawan is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto

	It has these top-level messages:
		PrefixesRequest
		PrefixesResponse
		DevAddrRequest
		DevAddrResponse
*/
package lorawan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PrefixesRequest struct {
}

func (m *PrefixesRequest) Reset()                    { *m = PrefixesRequest{} }
func (m *PrefixesRequest) String() string            { return proto.CompactTextString(m) }
func (*PrefixesRequest) ProtoMessage()               {}
func (*PrefixesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{0} }

type PrefixesResponse struct {
	// The prefixes that are in use or available
	Prefixes []*PrefixesResponse_PrefixMapping `protobuf:"bytes,1,rep,name=prefixes" json:"prefixes,omitempty"`
}

func (m *PrefixesResponse) Reset()                    { *m = PrefixesResponse{} }
func (m *PrefixesResponse) String() string            { return proto.CompactTextString(m) }
func (*PrefixesResponse) ProtoMessage()               {}
func (*PrefixesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{1} }

func (m *PrefixesResponse) GetPrefixes() []*PrefixesResponse_PrefixMapping {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type PrefixesResponse_PrefixMapping struct {
	// The prefix that can be used
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Usage constraints of this prefix (see activation_constraints in device.proto)
	Usage []string `protobuf:"bytes,2,rep,name=usage" json:"usage,omitempty"`
}

func (m *PrefixesResponse_PrefixMapping) Reset()         { *m = PrefixesResponse_PrefixMapping{} }
func (m *PrefixesResponse_PrefixMapping) String() string { return proto.CompactTextString(m) }
func (*PrefixesResponse_PrefixMapping) ProtoMessage()    {}
func (*PrefixesResponse_PrefixMapping) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceAddress, []int{1, 0}
}

func (m *PrefixesResponse_PrefixMapping) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *PrefixesResponse_PrefixMapping) GetUsage() []string {
	if m != nil {
		return m.Usage
	}
	return nil
}

type DevAddrRequest struct {
	// The usage constraints (see activation_constraints in device.proto)
	Usage []string `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty"`
}

func (m *DevAddrRequest) Reset()                    { *m = DevAddrRequest{} }
func (m *DevAddrRequest) String() string            { return proto.CompactTextString(m) }
func (*DevAddrRequest) ProtoMessage()               {}
func (*DevAddrRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{2} }

func (m *DevAddrRequest) GetUsage() []string {
	if m != nil {
		return m.Usage
	}
	return nil
}

type DevAddrResponse struct {
	DevAddr *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
}

func (m *DevAddrResponse) Reset()                    { *m = DevAddrResponse{} }
func (m *DevAddrResponse) String() string            { return proto.CompactTextString(m) }
func (*DevAddrResponse) ProtoMessage()               {}
func (*DevAddrResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{3} }

func init() {
	proto.RegisterType((*PrefixesRequest)(nil), "lorawan.PrefixesRequest")
	proto.RegisterType((*PrefixesResponse)(nil), "lorawan.PrefixesResponse")
	proto.RegisterType((*PrefixesResponse_PrefixMapping)(nil), "lorawan.PrefixesResponse.PrefixMapping")
	proto.RegisterType((*DevAddrRequest)(nil), "lorawan.DevAddrRequest")
	proto.RegisterType((*DevAddrResponse)(nil), "lorawan.DevAddrResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DevAddrManager service

type DevAddrManagerClient interface {
	// Get all prefixes that are in use or available
	GetPrefixes(ctx context.Context, in *PrefixesRequest, opts ...grpc.CallOption) (*PrefixesResponse, error)
	// Request a device address
	GetDevAddr(ctx context.Context, in *DevAddrRequest, opts ...grpc.CallOption) (*DevAddrResponse, error)
}

type devAddrManagerClient struct {
	cc *grpc.ClientConn
}

func NewDevAddrManagerClient(cc *grpc.ClientConn) DevAddrManagerClient {
	return &devAddrManagerClient{cc}
}

func (c *devAddrManagerClient) GetPrefixes(ctx context.Context, in *PrefixesRequest, opts ...grpc.CallOption) (*PrefixesResponse, error) {
	out := new(PrefixesResponse)
	err := grpc.Invoke(ctx, "/lorawan.DevAddrManager/GetPrefixes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devAddrManagerClient) GetDevAddr(ctx context.Context, in *DevAddrRequest, opts ...grpc.CallOption) (*DevAddrResponse, error) {
	out := new(DevAddrResponse)
	err := grpc.Invoke(ctx, "/lorawan.DevAddrManager/GetDevAddr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DevAddrManager service

type DevAddrManagerServer interface {
	// Get all prefixes that are in use or available
	GetPrefixes(context.Context, *PrefixesRequest) (*PrefixesResponse, error)
	// Request a device address
	GetDevAddr(context.Context, *DevAddrRequest) (*DevAddrResponse, error)
}

func RegisterDevAddrManagerServer(s *grpc.Server, srv DevAddrManagerServer) {
	s.RegisterService(&_DevAddrManager_serviceDesc, srv)
}

func _DevAddrManager_GetPrefixes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevAddrManagerServer).GetPrefixes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DevAddrManager/GetPrefixes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevAddrManagerServer).GetPrefixes(ctx, req.(*PrefixesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevAddrManager_GetDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevAddrManagerServer).GetDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DevAddrManager/GetDevAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevAddrManagerServer).GetDevAddr(ctx, req.(*DevAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevAddrManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorawan.DevAddrManager",
	HandlerType: (*DevAddrManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrefixes",
			Handler:    _DevAddrManager_GetPrefixes_Handler,
		},
		{
			MethodName: "GetDevAddr",
			Handler:    _DevAddrManager_GetDevAddr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto",
}

func (m *PrefixesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *PrefixesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Prefixes) > 0 {
		for _, msg := range m.Prefixes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeviceAddress(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PrefixesResponse_PrefixMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixesResponse_PrefixMapping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Prefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceAddress(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *DevAddrRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevAddrRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *DevAddrResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevAddrResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DevAddr != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceAddress(dAtA, i, uint64(m.DevAddr.Size()))
		n1, err := m.DevAddr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64DeviceAddress(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32DeviceAddress(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeviceAddress(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PrefixesRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *PrefixesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Prefixes) > 0 {
		for _, e := range m.Prefixes {
			l = e.Size()
			n += 1 + l + sovDeviceAddress(uint64(l))
		}
	}
	return n
}

func (m *PrefixesResponse_PrefixMapping) Size() (n int) {
	var l int
	_ = l
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovDeviceAddress(uint64(l))
	}
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
			l = len(s)
			n += 1 + l + sovDeviceAddress(uint64(l))
		}
	}
	return n
}

func (m *DevAddrRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
			l = len(s)
			n += 1 + l + sovDeviceAddress(uint64(l))
		}
	}
	return n
}

func (m *DevAddrResponse) Size() (n int) {
	var l int
	_ = l
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovDeviceAddress(uint64(l))
	}
	return n
}

func sovDeviceAddress(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeviceAddress(x uint64) (n int) {
	return sovDeviceAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PrefixesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrefixesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *PrefixesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrefixesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefixes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefixes = append(m.Prefixes, &PrefixesResponse_PrefixMapping{})
			if err := m.Prefixes[len(m.Prefixes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *PrefixesResponse_PrefixMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrefixMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Usage = append(m.Usage, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *DevAddrRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevAddrRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevAddrRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Usage = append(m.Usage, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *DevAddrResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevAddrResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevAddrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func skipDeviceAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceAddress
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
					return 0, ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceAddress
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDeviceAddress
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceAddress
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDeviceAddress(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDeviceAddress = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceAddress   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto", fileDescriptorDeviceAddress)
}

var fileDescriptorDeviceAddress = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4f, 0xdb, 0x40,
	0x14, 0xec, 0x36, 0x6a, 0x3e, 0x36, 0x6d, 0xd3, 0xae, 0xaa, 0xd6, 0xf5, 0xc1, 0x8d, 0x72, 0x68,
	0x73, 0xa9, 0x2d, 0x05, 0xc4, 0x0d, 0x21, 0x02, 0x22, 0xe2, 0x10, 0x04, 0x56, 0x4e, 0x5c, 0xd0,
	0xc6, 0xfb, 0xe2, 0x58, 0x04, 0xaf, 0xd9, 0x5d, 0x27, 0xf0, 0x43, 0x40, 0xfc, 0x24, 0x8e, 0x9c,
	0x39, 0x20, 0x14, 0xfe, 0x08, 0xd2, 0x7a, 0xe3, 0x24, 0x7c, 0x08, 0x89, 0xdb, 0xbe, 0x99, 0x79,
	0xf3, 0xc6, 0x23, 0xe3, 0xdd, 0x30, 0x52, 0xc3, 0xb4, 0xef, 0x06, 0xfc, 0xc4, 0xeb, 0x0d, 0xa1,
	0x37, 0x8c, 0xe2, 0x50, 0xee, 0x81, 0x9a, 0x70, 0x71, 0xec, 0x29, 0x15, 0x7b, 0x34, 0x89, 0xbc,
	0x44, 0x70, 0xc5, 0x03, 0x3e, 0xf2, 0x46, 0x5c, 0xd0, 0x09, 0x8d, 0x3d, 0x06, 0xe3, 0x28, 0x80,
	0x23, 0xca, 0x98, 0x00, 0x29, 0x5d, 0xcd, 0x93, 0x92, 0x61, 0xed, 0xff, 0x0b, 0x9e, 0x21, 0x0f,
	0x79, 0xb6, 0xdf, 0x4f, 0x07, 0x7a, 0xd2, 0x83, 0x7e, 0x65, 0x7b, 0x8d, 0xef, 0xb8, 0xb6, 0x2f,
	0x60, 0x10, 0x9d, 0x81, 0xf4, 0xe1, 0x34, 0x05, 0xa9, 0x1a, 0x97, 0x08, 0x7f, 0x9b, 0x63, 0x32,
	0xe1, 0xb1, 0x04, 0xb2, 0x85, 0xcb, 0x89, 0xc1, 0x2c, 0x54, 0x2f, 0x34, 0xab, 0xad, 0x7f, 0xae,
	0x39, 0xe9, 0x3e, 0x15, 0x1b, 0xa0, 0x4b, 0x93, 0x24, 0x8a, 0x43, 0x3f, 0x5f, 0xb4, 0xd7, 0xf1,
	0x97, 0x25, 0x8a, 0xfc, 0xc4, 0xc5, 0x8c, 0xb4, 0x50, 0x1d, 0x35, 0x2b, 0xbe, 0x99, 0xc8, 0x0f,
	0xfc, 0x29, 0x95, 0x34, 0x04, 0xeb, 0x63, 0xbd, 0xd0, 0xac, 0xf8, 0xd9, 0xd0, 0xf8, 0x8b, 0xbf,
	0x6e, 0xc3, 0x78, 0x93, 0x31, 0x61, 0xa2, 0xce, 0x75, 0x68, 0x51, 0xc7, 0x70, 0x2d, 0xd7, 0x99,
	0xf8, 0x07, 0xb8, 0xcc, 0x60, 0xac, 0x3b, 0xd3, 0xa7, 0x3e, 0xb7, 0xd7, 0x6e, 0xef, 0xfe, 0xb4,
	0xde, 0xea, 0x3f, 0xe0, 0x02, 0x3c, 0x75, 0x9e, 0x80, 0x74, 0x67, 0x8e, 0x25, 0x96, 0x3d, 0x5a,
	0x17, 0x28, 0x8f, 0xd3, 0xa5, 0x31, 0x0d, 0x41, 0x90, 0x36, 0xae, 0x76, 0x40, 0xcd, 0xea, 0x20,
	0xd6, 0x0b, 0x0d, 0xe9, 0xdc, 0xf6, 0xef, 0x57, 0xbb, 0x23, 0x1b, 0x18, 0x77, 0x40, 0x19, 0x63,
	0xf2, 0x2b, 0x17, 0x2e, 0x7f, 0xb9, 0x6d, 0x3d, 0x27, 0x32, 0x83, 0xf6, 0xce, 0xf5, 0xd4, 0x41,
	0x37, 0x53, 0x07, 0xdd, 0x4f, 0x1d, 0x74, 0xf5, 0xe0, 0x7c, 0x38, 0x5c, 0x7d, 0xcf, 0x6f, 0xd6,
	0x2f, 0x6a, 0x64, 0xe5, 0x31, 0x00, 0x00, 0xff, 0xff, 0x04, 0x47, 0x60, 0xc6, 0xa5, 0x02, 0x00,
	0x00,
}
