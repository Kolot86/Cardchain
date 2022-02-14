// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/collection.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type CStatus int32

const (
	CStatus_design    CStatus = 0
	CStatus_finalized CStatus = 1
	CStatus_active    CStatus = 2
	CStatus_archived  CStatus = 3
)

var CStatus_name = map[int32]string{
	0: "design",
	1: "finalized",
	2: "active",
	3: "archived",
}

var CStatus_value = map[string]int32{
	"design":    0,
	"finalized": 1,
	"active":    2,
	"archived":  3,
}

func (x CStatus) String() string {
	return proto.EnumName(CStatus_name, int32(x))
}

func (CStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fdd74a16aa19652d, []int{0}
}

type Collection struct {
	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cards        []uint64 `protobuf:"varint,2,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	Artist       string   `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	StoryWriter  string   `protobuf:"bytes,4,opt,name=storyWriter,proto3" json:"storyWriter,omitempty"`
	Contributors []string `protobuf:"bytes,5,rep,name=contributors,proto3" json:"contributors,omitempty"`
	Story        string   `protobuf:"bytes,6,opt,name=story,proto3" json:"story,omitempty"`
	Artwork      []byte   `protobuf:"bytes,7,opt,name=artwork,proto3" json:"artwork,omitempty"`
	Status       CStatus  `protobuf:"varint,8,opt,name=status,proto3,enum=DecentralCardGame.cardchain.cardchain.CStatus" json:"status,omitempty"`
	TimeStamp    int64    `protobuf:"varint,9,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdd74a16aa19652d, []int{0}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Collection) GetCards() []uint64 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *Collection) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *Collection) GetStoryWriter() string {
	if m != nil {
		return m.StoryWriter
	}
	return ""
}

func (m *Collection) GetContributors() []string {
	if m != nil {
		return m.Contributors
	}
	return nil
}

func (m *Collection) GetStory() string {
	if m != nil {
		return m.Story
	}
	return ""
}

func (m *Collection) GetArtwork() []byte {
	if m != nil {
		return m.Artwork
	}
	return nil
}

func (m *Collection) GetStatus() CStatus {
	if m != nil {
		return m.Status
	}
	return CStatus_design
}

func (m *Collection) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("DecentralCardGame.cardchain.cardchain.CStatus", CStatus_name, CStatus_value)
	proto.RegisterType((*Collection)(nil), "DecentralCardGame.cardchain.cardchain.Collection")
}

func init() { proto.RegisterFile("cardchain/collection.proto", fileDescriptor_fdd74a16aa19652d) }

var fileDescriptor_fdd74a16aa19652d = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x6a, 0xe3, 0x40,
	0x10, 0x86, 0xb5, 0x96, 0x2d, 0x5b, 0x73, 0xbe, 0x43, 0x2c, 0xc7, 0xb1, 0x1c, 0x87, 0x10, 0x86,
	0x03, 0x71, 0x85, 0x04, 0x97, 0x26, 0x55, 0x8a, 0x28, 0x24, 0xbd, 0x5c, 0x04, 0xd2, 0xad, 0x57,
	0x1b, 0x7b, 0x89, 0xa4, 0x35, 0xbb, 0x63, 0x27, 0xce, 0x53, 0xe4, 0x39, 0xf2, 0x24, 0x29, 0x5d,
	0xa6, 0x0c, 0xf6, 0x8b, 0x04, 0xcb, 0x8e, 0x9d, 0x90, 0x26, 0xdd, 0xfc, 0xff, 0xfc, 0x1f, 0xbb,
	0xc3, 0x0f, 0xbf, 0x05, 0x37, 0x85, 0x98, 0x70, 0x55, 0xa7, 0x42, 0x97, 0xa5, 0x14, 0xa8, 0x74,
	0x9d, 0x4c, 0x8d, 0x46, 0x4d, 0xff, 0x9e, 0x49, 0x21, 0x6b, 0x34, 0xbc, 0xcc, 0xb8, 0x29, 0x2e,
	0x78, 0x25, 0x93, 0x7d, 0xfa, 0x30, 0x0d, 0x1e, 0x5b, 0x00, 0xd9, 0x9e, 0xa5, 0x14, 0xda, 0x35,
	0xaf, 0x24, 0x23, 0x11, 0x89, 0xfd, 0xbc, 0x99, 0xe9, 0x4f, 0xe8, 0x6c, 0xf2, 0x96, 0xb5, 0x22,
	0x37, 0x6e, 0xe7, 0x5b, 0x41, 0x7f, 0x81, 0xc7, 0x0d, 0x2a, 0x8b, 0xcc, 0x6d, 0xb2, 0x3b, 0x45,
	0x23, 0xf8, 0x66, 0x51, 0x9b, 0xc5, 0xa5, 0x51, 0x28, 0x0d, 0x6b, 0x37, 0xcb, 0xf7, 0x16, 0x1d,
	0x40, 0x5f, 0xe8, 0x1a, 0x8d, 0x1a, 0xcd, 0x50, 0x1b, 0xcb, 0x3a, 0x91, 0x1b, 0xfb, 0xf9, 0x07,
	0x6f, 0xf3, 0x66, 0x83, 0x30, 0xaf, 0xe1, 0xb7, 0x82, 0x32, 0xe8, 0x72, 0x83, 0xb7, 0xda, 0xdc,
	0xb0, 0x6e, 0x44, 0xe2, 0x7e, 0xfe, 0x26, 0xe9, 0x39, 0x78, 0x16, 0x39, 0xce, 0x2c, 0xeb, 0x45,
	0x24, 0xfe, 0xf1, 0x3f, 0x49, 0xbe, 0x74, 0x7e, 0x92, 0x0d, 0x1b, 0x2a, 0xdf, 0xd1, 0xf4, 0x0f,
	0xf8, 0xa8, 0x2a, 0x39, 0x44, 0x5e, 0x4d, 0x99, 0x1f, 0x91, 0xd8, 0xcd, 0x0f, 0xc6, 0xbf, 0x13,
	0xe8, 0xee, 0x00, 0x0a, 0xe0, 0x15, 0xd2, 0xaa, 0x71, 0x1d, 0x38, 0xf4, 0x3b, 0xf8, 0xd7, 0xaa,
	0xe6, 0xa5, 0xba, 0x97, 0x45, 0x40, 0x36, 0x2b, 0x2e, 0x50, 0xcd, 0x65, 0xd0, 0xa2, 0x7d, 0xe8,
	0x71, 0x23, 0x26, 0x6a, 0x2e, 0x8b, 0xc0, 0x3d, 0xcd, 0x9f, 0x56, 0x21, 0x59, 0xae, 0x42, 0xf2,
	0xb2, 0x0a, 0xc9, 0xc3, 0x3a, 0x74, 0x96, 0xeb, 0xd0, 0x79, 0x5e, 0x87, 0xce, 0xd5, 0xf1, 0x58,
	0xe1, 0x64, 0x36, 0x4a, 0x84, 0xae, 0xd2, 0x4f, 0x3f, 0x4f, 0xb3, 0x7d, 0xcd, 0x77, 0xe9, 0xa1,
	0x72, 0x5c, 0x4c, 0xa5, 0x1d, 0x79, 0x4d, 0xdd, 0x47, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xc6, 0x6a, 0x10, 0x0c, 0x02, 0x00, 0x00,
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Artwork) > 0 {
		i -= len(m.Artwork)
		copy(dAtA[i:], m.Artwork)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Artwork)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Story) > 0 {
		i -= len(m.Story)
		copy(dAtA[i:], m.Story)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Story)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Contributors) > 0 {
		for iNdEx := len(m.Contributors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contributors[iNdEx])
			copy(dAtA[i:], m.Contributors[iNdEx])
			i = encodeVarintCollection(dAtA, i, uint64(len(m.Contributors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.StoryWriter) > 0 {
		i -= len(m.StoryWriter)
		copy(dAtA[i:], m.StoryWriter)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.StoryWriter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Artist) > 0 {
		i -= len(m.Artist)
		copy(dAtA[i:], m.Artist)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Artist)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cards) > 0 {
		dAtA2 := make([]byte, len(m.Cards)*10)
		var j1 int
		for _, num := range m.Cards {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintCollection(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.Cards) > 0 {
		l = 0
		for _, e := range m.Cards {
			l += sovCollection(uint64(e))
		}
		n += 1 + sovCollection(uint64(l)) + l
	}
	l = len(m.Artist)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.StoryWriter)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.Contributors) > 0 {
		for _, s := range m.Contributors {
			l = len(s)
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	l = len(m.Story)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Artwork)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovCollection(uint64(m.Status))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovCollection(uint64(m.TimeStamp))
	}
	return n
}

func sovCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollection(x uint64) (n int) {
	return sovCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
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
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCollection
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Cards = append(m.Cards, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCollection
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCollection
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCollection
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Cards) == 0 {
					m.Cards = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCollection
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Cards = append(m.Cards, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Cards", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Artist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Artist = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoryWriter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoryWriter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contributors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contributors = append(m.Contributors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Story", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Story = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Artwork", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Artwork = append(m.Artwork[:0], dAtA[iNdEx:postIndex]...)
			if m.Artwork == nil {
				m.Artwork = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollection
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
func skipCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
				return 0, ErrInvalidLengthCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollection = fmt.Errorf("proto: unexpected end of group")
)
