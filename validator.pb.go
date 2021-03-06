// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator.proto

/*
Package validator is a generated protocol buffer package.

It is generated from these files:
	validator.proto

It has these top-level messages:
	FieldValidator
*/
package validator

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,5,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of integer strictly equal this value.
	LengthEq         *int64 `protobuf:"varint,16,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FieldValidator) Reset()                    { *m = FieldValidator{} }
func (m *FieldValidator) String() string            { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()               {}
func (*FieldValidator) Descriptor() ([]byte, []int) { return fileDescriptorValidator, []int{0} }

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

func (m *FieldValidator) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidator) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil && m.StringNotEmpty != nil {
		return *m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() int64 {
	if m != nil && m.RepeatedCountMin != nil {
		return *m.RepeatedCountMin
	}
	return 0
}

func (m *FieldValidator) GetRepeatedCountMax() int64 {
	if m != nil && m.RepeatedCountMax != nil {
		return *m.RepeatedCountMax
	}
	return 0
}

func (m *FieldValidator) GetLengthGt() int64 {
	if m != nil && m.LengthGt != nil {
		return *m.LengthGt
	}
	return 0
}

func (m *FieldValidator) GetLengthLt() int64 {
	if m != nil && m.LengthLt != nil {
		return *m.LengthLt
	}
	return 0
}

func (m *FieldValidator) GetLengthEq() int64 {
	if m != nil && m.LengthEq != nil {
		return *m.LengthEq
	}
	return 0
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptorValidator) }

var fileDescriptorValidator = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x8f, 0x93, 0x40,
	0x14, 0xc7, 0x83, 0xbb, 0xec, 0xc2, 0xb0, 0xdb, 0x6d, 0x26, 0x9a, 0x4c, 0x35, 0x8d, 0x44, 0x2f,
	0x1c, 0x0c, 0x4d, 0x3c, 0x7a, 0xd4, 0x60, 0x2f, 0xf8, 0x23, 0x1c, 0x3c, 0x78, 0x21, 0x58, 0x5e,
	0xa7, 0x93, 0x0c, 0x33, 0x74, 0xe6, 0xd5, 0xe0, 0x3f, 0xe0, 0x3f, 0xad, 0x07, 0xc3, 0x20, 0x85,
	0x4d, 0x7a, 0x9c, 0xcf, 0xe7, 0xcb, 0x1b, 0x78, 0x7c, 0xc9, 0xc3, 0xcf, 0x4a, 0x8a, 0xba, 0x42,
	0x6d, 0xd2, 0xd6, 0x68, 0xd4, 0x34, 0x3c, 0x83, 0xe7, 0x31, 0xd7, 0x9a, 0x4b, 0xd8, 0x38, 0xf1,
	0xe3, 0xb4, 0xdf, 0xd4, 0x60, 0x77, 0x46, 0xb4, 0xe7, 0xf0, 0xab, 0xdf, 0xd7, 0x64, 0xf1, 0x51,
	0x80, 0xac, 0xbf, 0x8d, 0x0f, 0xd1, 0xa7, 0xc4, 0x37, 0xc0, 0xa1, 0x63, 0x5e, 0xec, 0x25, 0x61,
	0x31, 0x1c, 0xe8, 0x33, 0x72, 0x23, 0x14, 0x96, 0x1c, 0xd9, 0x93, 0xd8, 0x4b, 0xae, 0x0a, 0x5f,
	0x28, 0xdc, 0xe2, 0x88, 0x25, 0xb2, 0xab, 0x33, 0xce, 0x91, 0xae, 0x09, 0x69, 0x2c, 0x2f, 0xa1,
	0x13, 0x16, 0x2d, 0xbb, 0x8e, 0xbd, 0x24, 0x28, 0xc2, 0xc6, 0xf2, 0xcc, 0x01, 0xfa, 0x92, 0x44,
	0x87, 0x53, 0x53, 0xa9, 0x12, 0x8c, 0xd1, 0x86, 0xf9, 0xee, 0x22, 0xe2, 0x50, 0xd6, 0x13, 0xba,
	0x22, 0xc1, 0x5e, 0xea, 0xca, 0xdd, 0x77, 0x13, 0x7b, 0x89, 0x57, 0xdc, 0xba, 0xf3, 0x16, 0x27,
	0x25, 0x91, 0xdd, 0xce, 0x54, 0x8e, 0xf4, 0x35, 0xb9, 0x1f, 0x14, 0xb4, 0x56, 0x48, 0xad, 0x58,
	0xe0, 0xfc, 0x9d, 0x83, 0xd9, 0xc0, 0xe8, 0x0b, 0x12, 0x8e, 0xa3, 0x81, 0x85, 0x2e, 0x10, 0xfc,
	0x9f, 0x0d, 0x93, 0x94, 0x08, 0x8c, 0xcc, 0x64, 0x8e, 0x40, 0x13, 0xb2, 0xb4, 0x68, 0x84, 0xe2,
	0xa5, 0xd2, 0x58, 0x42, 0xd3, 0xe2, 0x2f, 0x16, 0xb9, 0x4f, 0x5b, 0x0c, 0xfc, 0xb3, 0xc6, 0xac,
	0xa7, 0xf4, 0x0d, 0xa1, 0x06, 0x5a, 0xa8, 0x10, 0xea, 0x72, 0xa7, 0x4f, 0x0a, 0xcb, 0x46, 0x28,
	0x76, 0xe7, 0x36, 0xb4, 0x1c, 0xcd, 0x87, 0x5e, 0x7c, 0x12, 0xea, 0x52, 0xba, 0xea, 0xd8, 0xfd,
	0xa5, 0x74, 0xd5, 0xf5, 0xaf, 0x28, 0x41, 0x71, 0x3c, 0xf4, 0xbb, 0x59, 0xb8, 0x50, 0x30, 0x80,
	0x2d, 0xce, 0xa4, 0x44, 0xf6, 0x30, 0x97, 0xf9, 0x5c, 0xc2, 0x91, 0x2d, 0xe7, 0x32, 0x3b, 0xbe,
	0xfb, 0x4a, 0xfc, 0x7d, 0xdf, 0x03, 0xba, 0x4e, 0x87, 0xd2, 0xa4, 0x63, 0x69, 0x52, 0xd7, 0x8f,
	0x2f, 0x2d, 0x0a, 0xad, 0x2c, 0xfb, 0xfb, 0xa7, 0xff, 0xd1, 0xd1, 0xdb, 0x55, 0x3a, 0xf5, 0xee,
	0x71, 0x81, 0x8a, 0x61, 0xd0, 0xfb, 0xe8, 0xfb, 0xd4, 0xc4, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x56, 0x30, 0x88, 0xa6, 0x02, 0x00, 0x00,
}
