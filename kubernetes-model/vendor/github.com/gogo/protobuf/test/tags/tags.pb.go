/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by protoc-gen-gogo.
// source: tags.proto
// DO NOT EDIT!

/*
Package tags is a generated protocol buffer package.

It is generated from these files:
	tags.proto

It has these top-level messages:
	Outside
	Inside
*/
package tags

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Outside struct {
	*Inside          `protobuf:"bytes,1,opt,name=Inside,embedded=Inside" json:""`
	Field2           *string `protobuf:"bytes,2,opt,name=Field2" json:"MyField2" xml:",comment"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Outside) Reset()                    { *m = Outside{} }
func (m *Outside) String() string            { return proto.CompactTextString(m) }
func (*Outside) ProtoMessage()               {}
func (*Outside) Descriptor() ([]byte, []int) { return fileDescriptorTags, []int{0} }

func (m *Outside) GetField2() string {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return ""
}

type Inside struct {
	Field1           *string `protobuf:"bytes,1,opt,name=Field1" json:"MyField1" xml:",chardata"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Inside) Reset()                    { *m = Inside{} }
func (m *Inside) String() string            { return proto.CompactTextString(m) }
func (*Inside) ProtoMessage()               {}
func (*Inside) Descriptor() ([]byte, []int) { return fileDescriptorTags, []int{1} }

func (m *Inside) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

func init() {
	proto.RegisterType((*Outside)(nil), "tags.Outside")
	proto.RegisterType((*Inside)(nil), "tags.Inside")
}
func NewPopulatedOutside(r randyTags, easy bool) *Outside {
	this := &Outside{}
	if r.Intn(10) != 0 {
		this.Inside = NewPopulatedInside(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := string(randStringTags(r))
		this.Field2 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 3)
	}
	return this
}

func NewPopulatedInside(r randyTags, easy bool) *Inside {
	this := &Inside{}
	if r.Intn(10) != 0 {
		v2 := string(randStringTags(r))
		this.Field1 = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 2)
	}
	return this
}

type randyTags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTags(r randyTags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTags(r randyTags) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTags(r)
	}
	return string(tmps)
}
func randUnrecognizedTags(r randyTags, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTags(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTags(dAtA []byte, r randyTags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTags(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTags(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTags(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("tags.proto", fileDescriptorTags) }

var fileDescriptorTags = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4c, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x92, 0x49, 0xa5,
	0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x34, 0x29, 0x15, 0x72, 0xb1, 0xfb, 0x97, 0x96, 0x14,
	0x67, 0xa6, 0xa4, 0x0a, 0xe9, 0x71, 0xb1, 0x79, 0xe6, 0x81, 0x58, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x3c, 0x7a, 0x60, 0xc3, 0x21, 0x62, 0x4e, 0x1c, 0x17, 0xee, 0xc9, 0x33, 0xbe, 0xba,
	0x27, 0xcf, 0x10, 0x04, 0x55, 0x25, 0x64, 0xc6, 0xc5, 0xe6, 0x96, 0x99, 0x9a, 0x93, 0x62, 0x24,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0xf7, 0xea, 0x9e, 0x3c, 0x87, 0x6f, 0x25, 0x44, 0xec,
	0xd3, 0x3d, 0x79, 0xbe, 0x8a, 0xdc, 0x1c, 0x2b, 0x25, 0x9d, 0xe4, 0xfc, 0xdc, 0xdc, 0xd4, 0xbc,
	0x12, 0xa5, 0x20, 0xa8, 0x6a, 0x25, 0x47, 0x98, 0x3d, 0x42, 0xe6, 0x50, 0x13, 0x0c, 0xc1, 0x36,
	0x72, 0x3a, 0xc9, 0x23, 0x99, 0x60, 0xf8, 0xe9, 0x9e, 0x3c, 0x3f, 0xd4, 0x84, 0x8c, 0xc4, 0xa2,
	0x94, 0xc4, 0x92, 0x44, 0x98, 0x11, 0x86, 0x4e, 0x2c, 0x3f, 0x1e, 0xca, 0x31, 0x02, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x57, 0x12, 0x09, 0x10, 0xfd, 0x00, 0x00, 0x00,
}
