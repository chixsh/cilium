// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v2alpha/rbac.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _type "github.com/cilium/cilium/pkg/envoy/envoy/type"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Should we do white-list or black-list style access control.
type RBAC_Action int32

const (
	// The policies grant access to principals. The rest is denied. This is white-list style
	// access control. This is the default type.
	RBAC_ALLOW RBAC_Action = 0
	// The policies deny access to principals. The rest is allowed. This is black-list style
	// access control.
	RBAC_DENY RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}
var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}
func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{0, 0}
}

// Role Based Access Control (RBAC) provides service-level and method-level access control for a service.
// The RBAC engine authorizes a request by evaluating the request context (expressed in the form of
// :ref: `AttributeContext <envoy_api_msg_service.auth.v2alpha.AttributeContext>`) against the RBAC policies.
//
// RBAC policies are additive. The policies are examined in order. A request is allowed once a matching policy
// is found (suppose the `action` is ALLOW).
//
// Here is an example of RBAC configuration. It has two policies:
// * Service account "cluster.local/ns/default/sa/admin" has full access (empty permission entry means full access)
//   to the service.
// * Any user (empty principal entry means any user) can read ("GET") the service at paths with prefix "/products" or
//   suffix "/reviews" when request header "version" set to either "v1" or "v2".
//
//   action: ALLOW
//   policies:
//     "service-admin":
//       permissions:
//       -
//       principals:
//         authenticated:
//           name: "cluster.local/ns/default/sa/admin"
//     "product-viewer":
//       permissions:
//       - paths: [prefix: "/products", suffix: "/reviews"]
//         methods: ["GET"]
//         conditions:
//         - header:
//             key: "version"
//             values: [simple: "v1", simple: "v2"]
//       principals:
//       -
//
type RBAC struct {
	Action RBAC_Action `protobuf:"varint,1,opt,name=action,enum=envoy.config.rbac.v2alpha.RBAC_Action" json:"action,omitempty"`
	// Maps from policy name to policy.
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{0}
}
func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (dst *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(dst, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy specifies a role and the principals that are assigned/denied the role.
type Policy struct {
	// Required. The set of permissions that define a role.
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
	// Required. List of principals that are assigned/denied the role based on “action”.
	Principals           []*Principal `protobuf:"bytes,2,rep,name=principals" json:"principals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{1}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

// Specifies how to match an entry in a map.
type MapEntryMatch struct {
	// The key to select an entry from the map.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// A list of matched values.
	Values               []*_type.StringMatch `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MapEntryMatch) Reset()         { *m = MapEntryMatch{} }
func (m *MapEntryMatch) String() string { return proto.CompactTextString(m) }
func (*MapEntryMatch) ProtoMessage()    {}
func (*MapEntryMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{2}
}
func (m *MapEntryMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapEntryMatch.Unmarshal(m, b)
}
func (m *MapEntryMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapEntryMatch.Marshal(b, m, deterministic)
}
func (dst *MapEntryMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapEntryMatch.Merge(dst, src)
}
func (m *MapEntryMatch) XXX_Size() int {
	return xxx_messageInfo_MapEntryMatch.Size(m)
}
func (m *MapEntryMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MapEntryMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MapEntryMatch proto.InternalMessageInfo

func (m *MapEntryMatch) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MapEntryMatch) GetValues() []*_type.StringMatch {
	if m != nil {
		return m.Values
	}
	return nil
}

// Specifies how to match IP addresses.
type IpMatch struct {
	// IP addresses in CIDR notation.
	Cidrs                []*core.CidrRange `protobuf:"bytes,1,rep,name=cidrs" json:"cidrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IpMatch) Reset()         { *m = IpMatch{} }
func (m *IpMatch) String() string { return proto.CompactTextString(m) }
func (*IpMatch) ProtoMessage()    {}
func (*IpMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{3}
}
func (m *IpMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpMatch.Unmarshal(m, b)
}
func (m *IpMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpMatch.Marshal(b, m, deterministic)
}
func (dst *IpMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpMatch.Merge(dst, src)
}
func (m *IpMatch) XXX_Size() int {
	return xxx_messageInfo_IpMatch.Size(m)
}
func (m *IpMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_IpMatch.DiscardUnknown(m)
}

var xxx_messageInfo_IpMatch proto.InternalMessageInfo

func (m *IpMatch) GetCidrs() []*core.CidrRange {
	if m != nil {
		return m.Cidrs
	}
	return nil
}

// Specifies how to match ports.
type PortMatch struct {
	// Port numbers.
	Ports                []uint32 `protobuf:"varint,1,rep,packed,name=ports" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortMatch) Reset()         { *m = PortMatch{} }
func (m *PortMatch) String() string { return proto.CompactTextString(m) }
func (*PortMatch) ProtoMessage()    {}
func (*PortMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{4}
}
func (m *PortMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortMatch.Unmarshal(m, b)
}
func (m *PortMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortMatch.Marshal(b, m, deterministic)
}
func (dst *PortMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortMatch.Merge(dst, src)
}
func (m *PortMatch) XXX_Size() int {
	return xxx_messageInfo_PortMatch.Size(m)
}
func (m *PortMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_PortMatch.DiscardUnknown(m)
}

var xxx_messageInfo_PortMatch proto.InternalMessageInfo

func (m *PortMatch) GetPorts() []uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

// Permission defines a permission to access the service.
type Permission struct {
	// Optional. A list of HTTP paths or gRPC methods.
	// gRPC methods must be presented as fully-qualified name in the form of
	// packageName.serviceName/methodName.
	// If this field is unset, it applies to any path.
	Paths []*_type.StringMatch `protobuf:"bytes,1,rep,name=paths" json:"paths,omitempty"`
	// Required. A list of HTTP methods (e.g., "GET", "POST").
	// If this field is unset, it applies to any method.
	Methods []string `protobuf:"bytes,2,rep,name=methods" json:"methods,omitempty"`
	// Optional. Custom conditions.
	Conditions           []*Permission_Condition `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{5}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (dst *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(dst, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetPaths() []*_type.StringMatch {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *Permission) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Permission) GetConditions() []*Permission_Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

// Definition of a custom condition.
type Permission_Condition struct {
	// Types that are valid to be assigned to ConditionSpec:
	//	*Permission_Condition_Header
	//	*Permission_Condition_DestinationIps
	//	*Permission_Condition_DestinationPorts
	ConditionSpec        isPermission_Condition_ConditionSpec `protobuf_oneof:"condition_spec"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *Permission_Condition) Reset()         { *m = Permission_Condition{} }
func (m *Permission_Condition) String() string { return proto.CompactTextString(m) }
func (*Permission_Condition) ProtoMessage()    {}
func (*Permission_Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{5, 0}
}
func (m *Permission_Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Condition.Unmarshal(m, b)
}
func (m *Permission_Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Condition.Marshal(b, m, deterministic)
}
func (dst *Permission_Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Condition.Merge(dst, src)
}
func (m *Permission_Condition) XXX_Size() int {
	return xxx_messageInfo_Permission_Condition.Size(m)
}
func (m *Permission_Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Condition proto.InternalMessageInfo

type isPermission_Condition_ConditionSpec interface {
	isPermission_Condition_ConditionSpec()
}

type Permission_Condition_Header struct {
	Header *MapEntryMatch `protobuf:"bytes,1,opt,name=header,oneof"`
}
type Permission_Condition_DestinationIps struct {
	DestinationIps *IpMatch `protobuf:"bytes,2,opt,name=destination_ips,json=destinationIps,oneof"`
}
type Permission_Condition_DestinationPorts struct {
	DestinationPorts *PortMatch `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,oneof"`
}

func (*Permission_Condition_Header) isPermission_Condition_ConditionSpec()           {}
func (*Permission_Condition_DestinationIps) isPermission_Condition_ConditionSpec()   {}
func (*Permission_Condition_DestinationPorts) isPermission_Condition_ConditionSpec() {}

func (m *Permission_Condition) GetConditionSpec() isPermission_Condition_ConditionSpec {
	if m != nil {
		return m.ConditionSpec
	}
	return nil
}

func (m *Permission_Condition) GetHeader() *MapEntryMatch {
	if x, ok := m.GetConditionSpec().(*Permission_Condition_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission_Condition) GetDestinationIps() *IpMatch {
	if x, ok := m.GetConditionSpec().(*Permission_Condition_DestinationIps); ok {
		return x.DestinationIps
	}
	return nil
}

func (m *Permission_Condition) GetDestinationPorts() *PortMatch {
	if x, ok := m.GetConditionSpec().(*Permission_Condition_DestinationPorts); ok {
		return x.DestinationPorts
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Permission_Condition) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Permission_Condition_OneofMarshaler, _Permission_Condition_OneofUnmarshaler, _Permission_Condition_OneofSizer, []interface{}{
		(*Permission_Condition_Header)(nil),
		(*Permission_Condition_DestinationIps)(nil),
		(*Permission_Condition_DestinationPorts)(nil),
	}
}

func _Permission_Condition_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Permission_Condition)
	// condition_spec
	switch x := m.ConditionSpec.(type) {
	case *Permission_Condition_Header:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case *Permission_Condition_DestinationIps:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DestinationIps); err != nil {
			return err
		}
	case *Permission_Condition_DestinationPorts:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DestinationPorts); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Permission_Condition.ConditionSpec has unexpected type %T", x)
	}
	return nil
}

func _Permission_Condition_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Permission_Condition)
	switch tag {
	case 1: // condition_spec.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MapEntryMatch)
		err := b.DecodeMessage(msg)
		m.ConditionSpec = &Permission_Condition_Header{msg}
		return true, err
	case 2: // condition_spec.destination_ips
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IpMatch)
		err := b.DecodeMessage(msg)
		m.ConditionSpec = &Permission_Condition_DestinationIps{msg}
		return true, err
	case 3: // condition_spec.destination_ports
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PortMatch)
		err := b.DecodeMessage(msg)
		m.ConditionSpec = &Permission_Condition_DestinationPorts{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Permission_Condition_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Permission_Condition)
	// condition_spec
	switch x := m.ConditionSpec.(type) {
	case *Permission_Condition_Header:
		s := proto.Size(x.Header)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Permission_Condition_DestinationIps:
		s := proto.Size(x.DestinationIps)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Permission_Condition_DestinationPorts:
		s := proto.Size(x.DestinationPorts)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Principal defines an identity or a group of identities.
type Principal struct {
	// Optional. Authenticated attributes that identify the principal.
	Authenticated *Principal_Authenticated `protobuf:"bytes,1,opt,name=authenticated" json:"authenticated,omitempty"`
	// Optional. Custom attributes that identify the principal.
	Attributes           []*Principal_Attribute `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{6}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (dst *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(dst, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if m != nil {
		return m.Authenticated
	}
	return nil
}

func (m *Principal) GetAttributes() []*Principal_Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Authentication attributes for principal. These could be filled out inside RBAC filter.
// Or if an authentication filter is used, they can be provided by the authentication filter.
type Principal_Authenticated struct {
	// Optional. The name of the principal. This matches to the "source.principal" field in
	// ":ref: `AttributeContext <envoy_api_msg_service.auth.v2alpha.AttributeContext>`.
	// If unset, it applies to any user.
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{6, 0}
}
func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (dst *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(dst, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Definition of a custom attribute to identify the principal.
type Principal_Attribute struct {
	// Types that are valid to be assigned to AttributeSpec:
	//	*Principal_Attribute_Service
	//	*Principal_Attribute_SourceIps
	//	*Principal_Attribute_Header
	AttributeSpec        isPrincipal_Attribute_AttributeSpec `protobuf_oneof:"attribute_spec"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Principal_Attribute) Reset()         { *m = Principal_Attribute{} }
func (m *Principal_Attribute) String() string { return proto.CompactTextString(m) }
func (*Principal_Attribute) ProtoMessage()    {}
func (*Principal_Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_09b42bf12f38e636, []int{6, 1}
}
func (m *Principal_Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Attribute.Unmarshal(m, b)
}
func (m *Principal_Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Attribute.Marshal(b, m, deterministic)
}
func (dst *Principal_Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Attribute.Merge(dst, src)
}
func (m *Principal_Attribute) XXX_Size() int {
	return xxx_messageInfo_Principal_Attribute.Size(m)
}
func (m *Principal_Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Attribute proto.InternalMessageInfo

type isPrincipal_Attribute_AttributeSpec interface {
	isPrincipal_Attribute_AttributeSpec()
}

type Principal_Attribute_Service struct {
	Service string `protobuf:"bytes,1,opt,name=service,oneof"`
}
type Principal_Attribute_SourceIps struct {
	SourceIps *IpMatch `protobuf:"bytes,2,opt,name=source_ips,json=sourceIps,oneof"`
}
type Principal_Attribute_Header struct {
	Header *MapEntryMatch `protobuf:"bytes,3,opt,name=header,oneof"`
}

func (*Principal_Attribute_Service) isPrincipal_Attribute_AttributeSpec()   {}
func (*Principal_Attribute_SourceIps) isPrincipal_Attribute_AttributeSpec() {}
func (*Principal_Attribute_Header) isPrincipal_Attribute_AttributeSpec()    {}

func (m *Principal_Attribute) GetAttributeSpec() isPrincipal_Attribute_AttributeSpec {
	if m != nil {
		return m.AttributeSpec
	}
	return nil
}

func (m *Principal_Attribute) GetService() string {
	if x, ok := m.GetAttributeSpec().(*Principal_Attribute_Service); ok {
		return x.Service
	}
	return ""
}

func (m *Principal_Attribute) GetSourceIps() *IpMatch {
	if x, ok := m.GetAttributeSpec().(*Principal_Attribute_SourceIps); ok {
		return x.SourceIps
	}
	return nil
}

func (m *Principal_Attribute) GetHeader() *MapEntryMatch {
	if x, ok := m.GetAttributeSpec().(*Principal_Attribute_Header); ok {
		return x.Header
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Principal_Attribute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Principal_Attribute_OneofMarshaler, _Principal_Attribute_OneofUnmarshaler, _Principal_Attribute_OneofSizer, []interface{}{
		(*Principal_Attribute_Service)(nil),
		(*Principal_Attribute_SourceIps)(nil),
		(*Principal_Attribute_Header)(nil),
	}
}

func _Principal_Attribute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Principal_Attribute)
	// attribute_spec
	switch x := m.AttributeSpec.(type) {
	case *Principal_Attribute_Service:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Service)
	case *Principal_Attribute_SourceIps:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SourceIps); err != nil {
			return err
		}
	case *Principal_Attribute_Header:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Principal_Attribute.AttributeSpec has unexpected type %T", x)
	}
	return nil
}

func _Principal_Attribute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Principal_Attribute)
	switch tag {
	case 1: // attribute_spec.service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AttributeSpec = &Principal_Attribute_Service{x}
		return true, err
	case 2: // attribute_spec.source_ips
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IpMatch)
		err := b.DecodeMessage(msg)
		m.AttributeSpec = &Principal_Attribute_SourceIps{msg}
		return true, err
	case 3: // attribute_spec.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MapEntryMatch)
		err := b.DecodeMessage(msg)
		m.AttributeSpec = &Principal_Attribute_Header{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Principal_Attribute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Principal_Attribute)
	// attribute_spec
	switch x := m.AttributeSpec.(type) {
	case *Principal_Attribute_Service:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Service)))
		n += len(x.Service)
	case *Principal_Attribute_SourceIps:
		s := proto.Size(x.SourceIps)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Principal_Attribute_Header:
		s := proto.Size(x.Header)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v2alpha.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v2alpha.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v2alpha.Policy")
	proto.RegisterType((*MapEntryMatch)(nil), "envoy.config.rbac.v2alpha.MapEntryMatch")
	proto.RegisterType((*IpMatch)(nil), "envoy.config.rbac.v2alpha.IpMatch")
	proto.RegisterType((*PortMatch)(nil), "envoy.config.rbac.v2alpha.PortMatch")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v2alpha.Permission")
	proto.RegisterType((*Permission_Condition)(nil), "envoy.config.rbac.v2alpha.Permission.Condition")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v2alpha.Principal")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v2alpha.Principal.Authenticated")
	proto.RegisterType((*Principal_Attribute)(nil), "envoy.config.rbac.v2alpha.Principal.Attribute")
	proto.RegisterEnum("envoy.config.rbac.v2alpha.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
}

func init() {
	proto.RegisterFile("envoy/config/rbac/v2alpha/rbac.proto", fileDescriptor_rbac_09b42bf12f38e636)
}

var fileDescriptor_rbac_09b42bf12f38e636 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x9b, 0x66, 0x6d, 0x97, 0x53, 0x75, 0x14, 0x0b, 0x69, 0xa5, 0x62, 0x62, 0x0b, 0x03,
	0xf5, 0x66, 0x89, 0x14, 0x2e, 0x40, 0x48, 0x20, 0xb5, 0x65, 0x52, 0x27, 0xed, 0x9f, 0xbc, 0x0b,
	0xfe, 0x5c, 0x30, 0x79, 0x89, 0x59, 0x2d, 0xda, 0xc4, 0x72, 0xdc, 0x4a, 0x7d, 0x00, 0xee, 0x11,
	0x2f, 0xc0, 0x1b, 0xc0, 0x3d, 0x57, 0x88, 0xb7, 0xe1, 0x2d, 0x50, 0x6c, 0x27, 0x4b, 0x25, 0x56,
	0x06, 0x77, 0x76, 0xfc, 0x7d, 0xbf, 0x63, 0x9f, 0xf3, 0xb5, 0xb0, 0x4b, 0xe3, 0x79, 0xb2, 0xf0,
	0xc3, 0x24, 0x7e, 0xcf, 0x2e, 0x7d, 0x71, 0x41, 0x42, 0x7f, 0x1e, 0x90, 0x09, 0x1f, 0x13, 0xb5,
	0xf1, 0xb8, 0x48, 0x64, 0x82, 0xee, 0x2a, 0x95, 0xa7, 0x55, 0x9e, 0x3a, 0x30, 0xaa, 0xee, 0xe6,
	0x9c, 0x4c, 0x58, 0x44, 0x24, 0xf5, 0xf3, 0x85, 0xf6, 0x74, 0xef, 0x6b, 0x32, 0xe1, 0xcc, 0x9f,
	0x07, 0x7e, 0x98, 0x08, 0xea, 0x93, 0x28, 0x12, 0x34, 0x4d, 0x8d, 0x60, 0x4b, 0x0b, 0xe4, 0x82,
	0x53, 0x3f, 0x95, 0x82, 0xc5, 0x97, 0xe7, 0x53, 0x22, 0xc3, 0xb1, 0x3e, 0x76, 0x3f, 0x55, 0x61,
	0x0d, 0x0f, 0xfa, 0x43, 0xf4, 0x02, 0xea, 0x24, 0x94, 0x2c, 0x89, 0x3b, 0xd6, 0xb6, 0xd5, 0xdb,
	0x08, 0x1e, 0x79, 0xd7, 0xde, 0xc6, 0xcb, 0x0c, 0x5e, 0x5f, 0xa9, 0xb1, 0x71, 0xa1, 0x03, 0x58,
	0xe7, 0xc9, 0x84, 0x85, 0x8c, 0xa6, 0x9d, 0xea, 0xb6, 0xdd, 0x6b, 0x06, 0x7b, 0x7f, 0x23, 0x9c,
	0x1a, 0xfd, 0x7e, 0x2c, 0xc5, 0x02, 0x17, 0xf6, 0xee, 0x3b, 0x68, 0x2d, 0x1d, 0xa1, 0x36, 0xd8,
	0x1f, 0xe8, 0x42, 0x5d, 0xcc, 0xc1, 0xd9, 0x12, 0x3d, 0x81, 0xda, 0x9c, 0x4c, 0x66, 0xb4, 0x53,
	0xdd, 0xb6, 0x7a, 0xcd, 0x60, 0x67, 0x45, 0x29, 0x85, 0x5a, 0x60, 0xad, 0x7f, 0x56, 0x7d, 0x6a,
	0xb9, 0x5b, 0x50, 0xd7, 0x97, 0x47, 0x0e, 0xd4, 0xfa, 0x87, 0x87, 0x27, 0xaf, 0xda, 0x15, 0xb4,
	0x0e, 0x6b, 0x2f, 0xf7, 0x8f, 0xdf, 0xb4, 0x2d, 0xf7, 0x9b, 0x05, 0x75, 0x6d, 0x42, 0x67, 0xd0,
	0xe4, 0x54, 0x4c, 0x59, 0x9a, 0xb2, 0x24, 0x4e, 0x3b, 0x96, 0x7a, 0xd7, 0xc3, 0x55, 0xc5, 0x0a,
	0xf5, 0x00, 0xbe, 0xff, 0xfa, 0x61, 0xd7, 0x3e, 0x5b, 0xd5, 0x75, 0x0b, 0x97, 0x29, 0xe8, 0x14,
	0x80, 0x0b, 0x16, 0x87, 0x8c, 0x93, 0x49, 0xde, 0xab, 0xdd, 0x55, 0xcc, 0x5c, 0xbc, 0x84, 0x2c,
	0x31, 0x5c, 0x0c, 0xad, 0x23, 0xc2, 0x55, 0xaf, 0x8e, 0xb2, 0xd9, 0xfe, 0xa1, 0x61, 0x3e, 0xd4,
	0x55, 0x03, 0xf2, 0x82, 0x9b, 0xa6, 0x60, 0x96, 0x0b, 0xef, 0x4c, 0xe5, 0x42, 0x59, 0xb1, 0x91,
	0xb9, 0xcf, 0xa1, 0x71, 0xc0, 0x35, 0x2d, 0x80, 0x5a, 0xc8, 0x22, 0x91, 0xbf, 0xff, 0x9e, 0xb1,
	0x12, 0xce, 0xbc, 0x79, 0xe0, 0x65, 0x99, 0xf3, 0x86, 0x2c, 0x12, 0x98, 0xc4, 0x97, 0x14, 0x6b,
	0xa9, 0xbb, 0x03, 0xce, 0x69, 0x22, 0xa4, 0x06, 0xdc, 0x81, 0x1a, 0x4f, 0x84, 0xd4, 0x80, 0x16,
	0xd6, 0x1b, 0xf7, 0xab, 0x0d, 0x70, 0xd5, 0x2f, 0xb4, 0x07, 0x35, 0x4e, 0xe4, 0x38, 0xaf, 0x72,
	0xed, 0x05, 0xb5, 0x0a, 0x75, 0xa0, 0x31, 0xa5, 0x72, 0x9c, 0x44, 0xfa, 0x45, 0x0e, 0xce, 0xb7,
	0xe8, 0x04, 0x20, 0x4c, 0xe2, 0x88, 0x49, 0x35, 0x33, 0x5b, 0xd1, 0xfc, 0x1b, 0xcd, 0xcc, 0x1b,
	0xe6, 0x3e, 0x5c, 0x42, 0x74, 0x3f, 0x56, 0xc1, 0x29, 0x4e, 0xd0, 0x00, 0xea, 0x63, 0x4a, 0x22,
	0x2a, 0x54, 0x7b, 0x9b, 0x41, 0x6f, 0x05, 0x7a, 0x69, 0x2a, 0xa3, 0x0a, 0x36, 0x4e, 0x74, 0x04,
	0xb7, 0x22, 0x9a, 0x4a, 0x16, 0x93, 0x0c, 0x79, 0xce, 0x78, 0x6a, 0x82, 0xec, 0xae, 0x80, 0x99,
	0x71, 0x8c, 0x2a, 0x78, 0xa3, 0x64, 0x3e, 0xe0, 0x29, 0x3a, 0x83, 0xdb, 0x65, 0x9c, 0xee, 0xb5,
	0xad, 0x80, 0x2b, 0x83, 0x95, 0x0f, 0x68, 0x54, 0xc1, 0xed, 0x12, 0x20, 0xfb, 0x9e, 0x0e, 0xda,
	0xb0, 0x51, 0xf4, 0xe0, 0x3c, 0xe5, 0x34, 0x74, 0xbf, 0xd8, 0xe0, 0x14, 0x61, 0x44, 0xaf, 0xa1,
	0x45, 0x66, 0x72, 0x4c, 0x63, 0xc9, 0x42, 0x22, 0x69, 0x64, 0xda, 0x11, 0xdc, 0x24, 0xc9, 0x5e,
	0xbf, 0xec, 0xc4, 0xcb, 0x20, 0x74, 0x0c, 0x40, 0xa4, 0x14, 0xec, 0x62, 0x26, 0x8b, 0xbc, 0x7a,
	0x37, 0xc3, 0xe6, 0x36, 0x5c, 0x22, 0x74, 0x1f, 0x40, 0x6b, 0xa9, 0x1e, 0x42, 0xb0, 0x16, 0x93,
	0x29, 0x35, 0xbf, 0x0f, 0xb5, 0xee, 0xfe, 0xb4, 0xc0, 0x29, 0xec, 0xa8, 0x0b, 0x8d, 0x94, 0x8a,
	0x39, 0x0b, 0x8d, 0x68, 0x54, 0xc1, 0xf9, 0x07, 0x34, 0x04, 0x48, 0x93, 0x99, 0x08, 0xe9, 0x3f,
	0xcf, 0xcd, 0xd1, 0xbe, 0x6c, 0x64, 0x57, 0x29, 0xb2, 0xff, 0x37, 0x45, 0xd9, 0x84, 0x8a, 0x57,
	0xaa, 0x09, 0x0d, 0x9c, 0xb7, 0x0d, 0x63, 0xba, 0xa8, 0xab, 0xff, 0xf7, 0xc7, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xad, 0x33, 0x87, 0x36, 0x7b, 0x06, 0x00, 0x00,
}