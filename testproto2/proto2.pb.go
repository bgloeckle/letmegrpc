// Code generated by protoc-gen-gogo.
// source: proto2.proto
// DO NOT EDIT!

/*
Package proto2 is a generated protocol buffer package.

It is generated from these files:
	proto2.proto

It has these top-level messages:
	Artist
	Song
	Album
*/
package proto2

import proto "github.com/gogo/protobuf/proto"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Instrument int32

const (
	Instrument_Voice  Instrument = 0
	Instrument_Guitar Instrument = 1
	Instrument_Drum   Instrument = 2
)

var Instrument_name = map[int32]string{
	0: "Voice",
	1: "Guitar",
	2: "Drum",
}
var Instrument_value = map[string]int32{
	"Voice":  0,
	"Guitar": 1,
	"Drum":   2,
}

func (x Instrument) Enum() *Instrument {
	p := new(Instrument)
	*p = x
	return p
}
func (x Instrument) String() string {
	return proto.EnumName(Instrument_name, int32(x))
}
func (x *Instrument) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Instrument_value, data, "Instrument")
	if err != nil {
		return err
	}
	*x = Instrument(value)
	return nil
}

type Genre int32

const (
	Genre_Pop          Genre = 0
	Genre_Rock         Genre = 1
	Genre_Jazz         Genre = 2
	Genre_NintendoCore Genre = 3
	Genre_Indie        Genre = 4
	Genre_Punk         Genre = 5
	Genre_Dance        Genre = 6
)

var Genre_name = map[int32]string{
	0: "Pop",
	1: "Rock",
	2: "Jazz",
	3: "NintendoCore",
	4: "Indie",
	5: "Punk",
	6: "Dance",
}
var Genre_value = map[string]int32{
	"Pop":          0,
	"Rock":         1,
	"Jazz":         2,
	"NintendoCore": 3,
	"Indie":        4,
	"Punk":         5,
	"Dance":        6,
}

func (x Genre) Enum() *Genre {
	p := new(Genre)
	*p = x
	return p
}
func (x Genre) String() string {
	return proto.EnumName(Genre_name, int32(x))
}
func (x *Genre) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Genre_value, data, "Genre")
	if err != nil {
		return err
	}
	*x = Genre(value)
	return nil
}

type Artist struct {
	Name             *string     `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Role             *Instrument `protobuf:"varint,2,opt,enum=proto2.Instrument" json:"Role,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Artist) Reset()         { *m = Artist{} }
func (m *Artist) String() string { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()    {}

func (m *Artist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Artist) GetRole() Instrument {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return Instrument_Voice
}

type Song struct {
	Name             *string   `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Track            *uint64   `protobuf:"varint,2,opt" json:"Track,omitempty"`
	Duration         *float64  `protobuf:"fixed64,3,opt" json:"Duration,omitempty"`
	Composer         []*Artist `protobuf:"bytes,4,rep" json:"Composer,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}

func (m *Song) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Song) GetTrack() uint64 {
	if m != nil && m.Track != nil {
		return *m.Track
	}
	return 0
}

func (m *Song) GetDuration() float64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Song) GetComposer() []*Artist {
	if m != nil {
		return m.Composer
	}
	return nil
}

type Album struct {
	Name             *string  `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Song             []*Song  `protobuf:"bytes,2,rep" json:"Song,omitempty"`
	Genre            *Genre   `protobuf:"varint,3,opt,enum=proto2.Genre" json:"Genre,omitempty"`
	Year             *string  `protobuf:"bytes,4,opt" json:"Year,omitempty"`
	Producer         []string `protobuf:"bytes,5,rep" json:"Producer,omitempty"`
	Mediocre         *bool    `protobuf:"varint,6,opt" json:"Mediocre,omitempty"`
	Rated            *bool    `protobuf:"varint,7,opt" json:"Rated,omitempty"`
	Epilogue         *string  `protobuf:"bytes,8,opt" json:"Epilogue,omitempty"`
	Likes            []bool   `protobuf:"varint,9,rep" json:"Likes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Album) Reset()         { *m = Album{} }
func (m *Album) String() string { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()    {}

func (m *Album) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Album) GetSong() []*Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *Album) GetGenre() Genre {
	if m != nil && m.Genre != nil {
		return *m.Genre
	}
	return Genre_Pop
}

func (m *Album) GetYear() string {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return ""
}

func (m *Album) GetProducer() []string {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *Album) GetMediocre() bool {
	if m != nil && m.Mediocre != nil {
		return *m.Mediocre
	}
	return false
}

func (m *Album) GetRated() bool {
	if m != nil && m.Rated != nil {
		return *m.Rated
	}
	return false
}

func (m *Album) GetEpilogue() string {
	if m != nil && m.Epilogue != nil {
		return *m.Epilogue
	}
	return ""
}

func (m *Album) GetLikes() []bool {
	if m != nil {
		return m.Likes
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto2.Instrument", Instrument_name, Instrument_value)
	proto.RegisterEnum("proto2.Genre", Genre_name, Genre_value)
}

// Client API for Proto2 service

type Proto2Client interface {
	Produce(ctx context.Context, in *Album, opts ...grpc.CallOption) (*Album, error)
}

type proto2Client struct {
	cc *grpc.ClientConn
}

func NewProto2Client(cc *grpc.ClientConn) Proto2Client {
	return &proto2Client{cc}
}

func (c *proto2Client) Produce(ctx context.Context, in *Album, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := grpc.Invoke(ctx, "/proto2.Proto2/Produce", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Proto2 service

type Proto2Server interface {
	Produce(context.Context, *Album) (*Album, error)
}

func RegisterProto2Server(s *grpc.Server, srv Proto2Server) {
	s.RegisterService(&_Proto2_serviceDesc, srv)
}

func _Proto2_Produce_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Album)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(Proto2Server).Produce(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Proto2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto2.Proto2",
	HandlerType: (*Proto2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Proto2_Produce_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
