// Code generated by protoc-gen-gogo.
// source: raft_internal.proto
// DO NOT EDIT!

package etcdserverpb

import (
	"fmt"

	proto "github.com/coreos/etcd/Godeps/_workspace/src/github.com/gogo/protobuf/proto"
)

import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An InternalRaftRequest is the union of all requests which can be
// sent via raft.
type InternalRaftRequest struct {
	ID          uint64              `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	V2          *Request            `protobuf:"bytes,2,opt,name=v2" json:"v2,omitempty"`
	Range       *RangeRequest       `protobuf:"bytes,3,opt,name=range" json:"range,omitempty"`
	Put         *PutRequest         `protobuf:"bytes,4,opt,name=put" json:"put,omitempty"`
	DeleteRange *DeleteRangeRequest `protobuf:"bytes,5,opt,name=delete_range" json:"delete_range,omitempty"`
	Txn         *TxnRequest         `protobuf:"bytes,6,opt,name=txn" json:"txn,omitempty"`
	Compaction  *CompactionRequest  `protobuf:"bytes,7,opt,name=compaction" json:"compaction,omitempty"`
	LeaseCreate *LeaseCreateRequest `protobuf:"bytes,8,opt,name=lease_create" json:"lease_create,omitempty"`
	LeaseRevoke *LeaseRevokeRequest `protobuf:"bytes,9,opt,name=lease_revoke" json:"lease_revoke,omitempty"`
	AuthEnable  *AuthEnableRequest  `protobuf:"bytes,10,opt,name=auth_enable" json:"auth_enable,omitempty"`
}

func (m *InternalRaftRequest) Reset()         { *m = InternalRaftRequest{} }
func (m *InternalRaftRequest) String() string { return proto.CompactTextString(m) }
func (*InternalRaftRequest) ProtoMessage()    {}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*InternalRaftRequest)(nil), "etcdserverpb.InternalRaftRequest")
	proto.RegisterType((*EmptyResponse)(nil), "etcdserverpb.EmptyResponse")
}
func (m *InternalRaftRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *InternalRaftRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.ID))
	}
	if m.V2 != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.V2.Size()))
		n1, err := m.V2.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Range != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Range.Size()))
		n2, err := m.Range.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Put != nil {
		data[i] = 0x22
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Put.Size()))
		n3, err := m.Put.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.DeleteRange != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.DeleteRange.Size()))
		n4, err := m.DeleteRange.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Txn != nil {
		data[i] = 0x32
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Txn.Size()))
		n5, err := m.Txn.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Compaction != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.Compaction.Size()))
		n6, err := m.Compaction.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.LeaseCreate != nil {
		data[i] = 0x42
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.LeaseCreate.Size()))
		n7, err := m.LeaseCreate.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.LeaseRevoke != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.LeaseRevoke.Size()))
		n8, err := m.LeaseRevoke.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.AuthEnable != nil {
		data[i] = 0x52
		i++
		i = encodeVarintRaftInternal(data, i, uint64(m.AuthEnable.Size()))
		n9, err := m.AuthEnable.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}

func (m *EmptyResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EmptyResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64RaftInternal(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32RaftInternal(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaftInternal(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *InternalRaftRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaftInternal(uint64(m.ID))
	}
	if m.V2 != nil {
		l = m.V2.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Put != nil {
		l = m.Put.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.DeleteRange != nil {
		l = m.DeleteRange.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Txn != nil {
		l = m.Txn.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Compaction != nil {
		l = m.Compaction.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseCreate != nil {
		l = m.LeaseCreate.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseRevoke != nil {
		l = m.LeaseRevoke.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthEnable != nil {
		l = m.AuthEnable.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	return n
}

func (m *EmptyResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovRaftInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaftInternal(x uint64) (n int) {
	return sovRaftInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalRaftRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InternalRaftRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalRaftRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.V2 == nil {
				m.V2 = &Request{}
			}
			if err := m.V2.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &RangeRequest{}
			}
			if err := m.Range.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Put", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Put == nil {
				m.Put = &PutRequest{}
			}
			if err := m.Put.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeleteRange == nil {
				m.DeleteRange = &DeleteRangeRequest{}
			}
			if err := m.DeleteRange.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Txn == nil {
				m.Txn = &TxnRequest{}
			}
			if err := m.Txn.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compaction == nil {
				m.Compaction = &CompactionRequest{}
			}
			if err := m.Compaction.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseCreate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseCreate == nil {
				m.LeaseCreate = &LeaseCreateRequest{}
			}
			if err := m.LeaseCreate.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseRevoke", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseRevoke == nil {
				m.LeaseRevoke = &LeaseRevokeRequest{}
			}
			if err := m.LeaseRevoke.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthEnable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthEnable == nil {
				m.AuthEnable = &AuthEnableRequest{}
			}
			if err := m.AuthEnable.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
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
func (m *EmptyResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
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
func skipRaftInternal(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaftInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaftInternal
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRaftInternal(data[start:])
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
	ErrInvalidLengthRaftInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaftInternal   = fmt.Errorf("proto: integer overflow")
)
