// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/policies.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	msp "github.com/hyperledger/fabric-protos-go/msp"
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

type Policy_PolicyType int32

const (
	Policy_UNKNOWN       Policy_PolicyType = 0
	Policy_SIGNATURE     Policy_PolicyType = 1
	Policy_MSP           Policy_PolicyType = 2
	Policy_IMPLICIT_META Policy_PolicyType = 3
)

var Policy_PolicyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SIGNATURE",
	2: "MSP",
	3: "IMPLICIT_META",
}

var Policy_PolicyType_value = map[string]int32{
	"UNKNOWN":       0,
	"SIGNATURE":     1,
	"MSP":           2,
	"IMPLICIT_META": 3,
}

func (x Policy_PolicyType) String() string {
	return proto.EnumName(Policy_PolicyType_name, int32(x))
}

func (Policy_PolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{0, 0}
}

type ImplicitMetaPolicy_Rule int32

const (
	ImplicitMetaPolicy_ANY      ImplicitMetaPolicy_Rule = 0
	ImplicitMetaPolicy_ALL      ImplicitMetaPolicy_Rule = 1
	ImplicitMetaPolicy_MAJORITY ImplicitMetaPolicy_Rule = 2
)

var ImplicitMetaPolicy_Rule_name = map[int32]string{
	0: "ANY",
	1: "ALL",
	2: "MAJORITY",
}

var ImplicitMetaPolicy_Rule_value = map[string]int32{
	"ANY":      0,
	"ALL":      1,
	"MAJORITY": 2,
}

func (x ImplicitMetaPolicy_Rule) String() string {
	return proto.EnumName(ImplicitMetaPolicy_Rule_name, int32(x))
}

func (ImplicitMetaPolicy_Rule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{3, 0}
}

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{0}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Policy) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
type SignaturePolicyEnvelope struct {
	Version              int32               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Rule                 *SignaturePolicy    `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
	Identities           []*msp.MSPPrincipal `protobuf:"bytes,3,rep,name=identities,proto3" json:"identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SignaturePolicyEnvelope) Reset()         { *m = SignaturePolicyEnvelope{} }
func (m *SignaturePolicyEnvelope) String() string { return proto.CompactTextString(m) }
func (*SignaturePolicyEnvelope) ProtoMessage()    {}
func (*SignaturePolicyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{1}
}

func (m *SignaturePolicyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignaturePolicyEnvelope.Unmarshal(m, b)
}
func (m *SignaturePolicyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignaturePolicyEnvelope.Marshal(b, m, deterministic)
}
func (m *SignaturePolicyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignaturePolicyEnvelope.Merge(m, src)
}
func (m *SignaturePolicyEnvelope) XXX_Size() int {
	return xxx_messageInfo_SignaturePolicyEnvelope.Size(m)
}
func (m *SignaturePolicyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SignaturePolicyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SignaturePolicyEnvelope proto.InternalMessageInfo

func (m *SignaturePolicyEnvelope) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SignaturePolicyEnvelope) GetRule() *SignaturePolicy {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *SignaturePolicyEnvelope) GetIdentities() []*msp.MSPPrincipal {
	if m != nil {
		return m.Identities
	}
	return nil
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
type SignaturePolicy struct {
	// Types that are valid to be assigned to Type:
	//	*SignaturePolicy_SignedBy
	//	*SignaturePolicy_NOutOf_
	Type                 isSignaturePolicy_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SignaturePolicy) Reset()         { *m = SignaturePolicy{} }
func (m *SignaturePolicy) String() string { return proto.CompactTextString(m) }
func (*SignaturePolicy) ProtoMessage()    {}
func (*SignaturePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{2}
}

func (m *SignaturePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignaturePolicy.Unmarshal(m, b)
}
func (m *SignaturePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignaturePolicy.Marshal(b, m, deterministic)
}
func (m *SignaturePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignaturePolicy.Merge(m, src)
}
func (m *SignaturePolicy) XXX_Size() int {
	return xxx_messageInfo_SignaturePolicy.Size(m)
}
func (m *SignaturePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_SignaturePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_SignaturePolicy proto.InternalMessageInfo

type isSignaturePolicy_Type interface {
	isSignaturePolicy_Type()
}

type SignaturePolicy_SignedBy struct {
	SignedBy int32 `protobuf:"varint,1,opt,name=signed_by,json=signedBy,proto3,oneof"`
}

type SignaturePolicy_NOutOf_ struct {
	NOutOf *SignaturePolicy_NOutOf `protobuf:"bytes,2,opt,name=n_out_of,json=nOutOf,proto3,oneof"`
}

func (*SignaturePolicy_SignedBy) isSignaturePolicy_Type() {}

func (*SignaturePolicy_NOutOf_) isSignaturePolicy_Type() {}

func (m *SignaturePolicy) GetType() isSignaturePolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SignaturePolicy) GetSignedBy() int32 {
	if x, ok := m.GetType().(*SignaturePolicy_SignedBy); ok {
		return x.SignedBy
	}
	return 0
}

func (m *SignaturePolicy) GetNOutOf() *SignaturePolicy_NOutOf {
	if x, ok := m.GetType().(*SignaturePolicy_NOutOf_); ok {
		return x.NOutOf
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignaturePolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignaturePolicy_SignedBy)(nil),
		(*SignaturePolicy_NOutOf_)(nil),
	}
}

type SignaturePolicy_NOutOf struct {
	N                    int32              `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	Rules                []*SignaturePolicy `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SignaturePolicy_NOutOf) Reset()         { *m = SignaturePolicy_NOutOf{} }
func (m *SignaturePolicy_NOutOf) String() string { return proto.CompactTextString(m) }
func (*SignaturePolicy_NOutOf) ProtoMessage()    {}
func (*SignaturePolicy_NOutOf) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{2, 0}
}

func (m *SignaturePolicy_NOutOf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignaturePolicy_NOutOf.Unmarshal(m, b)
}
func (m *SignaturePolicy_NOutOf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignaturePolicy_NOutOf.Marshal(b, m, deterministic)
}
func (m *SignaturePolicy_NOutOf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignaturePolicy_NOutOf.Merge(m, src)
}
func (m *SignaturePolicy_NOutOf) XXX_Size() int {
	return xxx_messageInfo_SignaturePolicy_NOutOf.Size(m)
}
func (m *SignaturePolicy_NOutOf) XXX_DiscardUnknown() {
	xxx_messageInfo_SignaturePolicy_NOutOf.DiscardUnknown(m)
}

var xxx_messageInfo_SignaturePolicy_NOutOf proto.InternalMessageInfo

func (m *SignaturePolicy_NOutOf) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *SignaturePolicy_NOutOf) GetRules() []*SignaturePolicy {
	if m != nil {
		return m.Rules
	}
	return nil
}

// ImplicitMetaPolicy is a policy type which depends on the hierarchical nature of the configuration
// It is implicit because the rule is generate implicitly based on the number of sub policies
// It is meta because it depends only on the result of other policies
// When evaluated, this policy iterates over all immediate child sub-groups, retrieves the policy
// of name sub_policy, evaluates the collection and applies the rule.
// For example, with 4 sub-groups, and a policy name of "foo", ImplicitMetaPolicy retrieves
// each sub-group, retrieves policy "foo" for each subgroup, evaluates it, and, in the case of ANY
// 1 satisfied is sufficient, ALL would require 4 signatures, and MAJORITY would require 3 signatures.
type ImplicitMetaPolicy struct {
	SubPolicy            string                  `protobuf:"bytes,1,opt,name=sub_policy,json=subPolicy,proto3" json:"sub_policy,omitempty"`
	Rule                 ImplicitMetaPolicy_Rule `protobuf:"varint,2,opt,name=rule,proto3,enum=common.ImplicitMetaPolicy_Rule" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ImplicitMetaPolicy) Reset()         { *m = ImplicitMetaPolicy{} }
func (m *ImplicitMetaPolicy) String() string { return proto.CompactTextString(m) }
func (*ImplicitMetaPolicy) ProtoMessage()    {}
func (*ImplicitMetaPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d02cf0d453425a3, []int{3}
}

func (m *ImplicitMetaPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImplicitMetaPolicy.Unmarshal(m, b)
}
func (m *ImplicitMetaPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImplicitMetaPolicy.Marshal(b, m, deterministic)
}
func (m *ImplicitMetaPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplicitMetaPolicy.Merge(m, src)
}
func (m *ImplicitMetaPolicy) XXX_Size() int {
	return xxx_messageInfo_ImplicitMetaPolicy.Size(m)
}
func (m *ImplicitMetaPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplicitMetaPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ImplicitMetaPolicy proto.InternalMessageInfo

func (m *ImplicitMetaPolicy) GetSubPolicy() string {
	if m != nil {
		return m.SubPolicy
	}
	return ""
}

func (m *ImplicitMetaPolicy) GetRule() ImplicitMetaPolicy_Rule {
	if m != nil {
		return m.Rule
	}
	return ImplicitMetaPolicy_ANY
}

func init() {
	proto.RegisterEnum("common.Policy_PolicyType", Policy_PolicyType_name, Policy_PolicyType_value)
	proto.RegisterEnum("common.ImplicitMetaPolicy_Rule", ImplicitMetaPolicy_Rule_name, ImplicitMetaPolicy_Rule_value)
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "common.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "common.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "common.SignaturePolicy.NOutOf")
	proto.RegisterType((*ImplicitMetaPolicy)(nil), "common.ImplicitMetaPolicy")
}

func init() { proto.RegisterFile("common/policies.proto", fileDescriptor_0d02cf0d453425a3) }

var fileDescriptor_0d02cf0d453425a3 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0x8f, 0xd2, 0x40,
	0x10, 0xc7, 0x29, 0x3f, 0x0a, 0x0c, 0x9c, 0xd6, 0xcd, 0x19, 0xc8, 0x25, 0x2a, 0x69, 0x8c, 0x21,
	0x31, 0x94, 0x84, 0xf3, 0xc9, 0x37, 0x50, 0xe2, 0x55, 0x69, 0x21, 0x0b, 0xa7, 0x39, 0x5f, 0x1a,
	0x0a, 0x4b, 0x6f, 0x93, 0xb2, 0xbb, 0xe9, 0x6e, 0x89, 0xfc, 0x17, 0x3e, 0xf9, 0xcf, 0xf8, 0xcf,
	0x99, 0x76, 0xa9, 0x21, 0x77, 0xb9, 0xb7, 0x99, 0xe9, 0x67, 0xa6, 0xdf, 0xef, 0xec, 0xc0, 0xcb,
	0x0d, 0xdf, 0xef, 0x39, 0x1b, 0x0a, 0x1e, 0xd3, 0x0d, 0x25, 0xd2, 0x11, 0x09, 0x57, 0x1c, 0x99,
	0xba, 0x7c, 0xd5, 0xd9, 0x4b, 0x31, 0xdc, 0x4b, 0x11, 0x88, 0x84, 0xb2, 0x0d, 0x15, 0xeb, 0x58,
	0x03, 0xf6, 0x2f, 0x30, 0x17, 0x59, 0xcb, 0x11, 0x21, 0xa8, 0xaa, 0xa3, 0x20, 0x5d, 0xa3, 0x67,
	0xf4, 0x6b, 0x38, 0x8f, 0xd1, 0x25, 0xd4, 0x0e, 0xeb, 0x38, 0x25, 0xdd, 0x72, 0xcf, 0xe8, 0xb7,
	0xb1, 0x4e, 0xec, 0xcf, 0x00, 0xba, 0x67, 0x95, 0x31, 0x2d, 0xa8, 0xdf, 0xfa, 0xdf, 0xfc, 0xf9,
	0x0f, 0xdf, 0x2a, 0xa1, 0x0b, 0x68, 0x2e, 0xdd, 0x2f, 0xfe, 0x78, 0x75, 0x8b, 0xa7, 0x96, 0x81,
	0xea, 0x50, 0xf1, 0x96, 0x0b, 0xab, 0x8c, 0x5e, 0xc0, 0x85, 0xeb, 0x2d, 0x66, 0xee, 0x27, 0x77,
	0x15, 0x78, 0xd3, 0xd5, 0xd8, 0xaa, 0xd8, 0x7f, 0x0c, 0xe8, 0x2c, 0x69, 0xc4, 0xd6, 0x2a, 0x4d,
	0x88, 0x9e, 0x37, 0x65, 0x07, 0x12, 0x73, 0x41, 0x50, 0x17, 0xea, 0x07, 0x92, 0x48, 0xca, 0xd9,
	0x49, 0x4e, 0x91, 0xa2, 0xf7, 0x50, 0x4d, 0xd2, 0x58, 0x0b, 0x6a, 0x8d, 0x3a, 0x8e, 0xf6, 0xe7,
	0x3c, 0x18, 0x84, 0x73, 0x08, 0x7d, 0x00, 0xa0, 0x5b, 0xc2, 0x14, 0x55, 0x94, 0xc8, 0x6e, 0xa5,
	0x57, 0xe9, 0xb7, 0x46, 0x97, 0x45, 0x8b, 0xb7, 0x5c, 0x2c, 0x8a, 0x65, 0xe0, 0x33, 0xce, 0xfe,
	0x6b, 0xc0, 0xf3, 0x07, 0xf3, 0xd0, 0x2b, 0x68, 0x4a, 0x1a, 0x31, 0xb2, 0x0d, 0xc2, 0xa3, 0x96,
	0x74, 0x53, 0xc2, 0x0d, 0x5d, 0x9a, 0x1c, 0xd1, 0x47, 0x68, 0xb0, 0x80, 0xa7, 0x2a, 0xe0, 0xbb,
	0x93, 0xb2, 0xd7, 0x4f, 0x28, 0x73, 0xfc, 0x79, 0xaa, 0xe6, 0xbb, 0x9b, 0x12, 0x36, 0x59, 0x1e,
	0x5d, 0x4d, 0xc1, 0xd4, 0x35, 0xd4, 0x06, 0xa3, 0xf0, 0x6b, 0x30, 0x34, 0x80, 0x5a, 0x66, 0x42,
	0x76, 0xcb, 0xb9, 0xee, 0x27, 0xad, 0x6a, 0x6a, 0x62, 0x42, 0x35, 0x7b, 0x0e, 0xfb, 0xb7, 0x01,
	0xc8, 0xdd, 0x8b, 0xec, 0x0a, 0x94, 0x47, 0xd4, 0xfa, 0xbf, 0x01, 0x90, 0x69, 0x18, 0xe4, 0xe7,
	0xa1, 0x1d, 0x34, 0x71, 0x53, 0xa6, 0xe1, 0xe9, 0xf3, 0xf5, 0xd9, 0x5a, 0x9f, 0x8d, 0xde, 0x14,
	0xff, 0x7a, 0x3c, 0xc8, 0xc1, 0x69, 0x4c, 0xf4, 0x7a, 0xed, 0x77, 0x50, 0xcd, 0xb2, 0xec, 0x95,
	0xc7, 0xfe, 0x9d, 0x55, 0xca, 0x83, 0xd9, 0xcc, 0x32, 0x50, 0x1b, 0x1a, 0xde, 0xf8, 0xeb, 0x1c,
	0xbb, 0xab, 0x3b, 0xab, 0x3c, 0xf9, 0x0e, 0x6f, 0x79, 0x12, 0x39, 0xf7, 0x47, 0x41, 0x92, 0x98,
	0x6c, 0x23, 0x92, 0x38, 0xbb, 0x75, 0x98, 0xd0, 0x8d, 0xbe, 0x41, 0x79, 0xfa, 0xdb, 0x4f, 0x27,
	0xa2, 0xea, 0x3e, 0x0d, 0xb3, 0x74, 0x78, 0x06, 0x0f, 0x35, 0x3c, 0xd0, 0xf0, 0x20, 0xe2, 0x43,
	0xcd, 0x87, 0x66, 0x5e, 0xb9, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x38, 0xcf, 0x03, 0x1e, 0xfc,
	0x02, 0x00, 0x00,
}
