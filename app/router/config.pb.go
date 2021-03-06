package router

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_net1 "v2ray.com/core/common/net"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Type of domain value.
type Domain_Type int32

const (
	// The value is used as is.
	Domain_Plain Domain_Type = 0
	// The value is used as a regular expression.
	Domain_Regex Domain_Type = 1
	// The value is a domain.
	Domain_Domain Domain_Type = 2
)

var Domain_Type_name = map[int32]string{
	0: "Plain",
	1: "Regex",
	2: "Domain",
}
var Domain_Type_value = map[string]int32{
	"Plain":  0,
	"Regex":  1,
	"Domain": 2,
}

func (x Domain_Type) String() string {
	return proto.EnumName(Domain_Type_name, int32(x))
}
func (Domain_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Config_DomainStrategy int32

const (
	// Use domain as is.
	Config_AsIs Config_DomainStrategy = 0
	// Always resolve IP for domains.
	Config_UseIp Config_DomainStrategy = 1
	// Resolve to IP if the domain doesn't match any rules.
	Config_IpIfNonMatch Config_DomainStrategy = 2
)

var Config_DomainStrategy_name = map[int32]string{
	0: "AsIs",
	1: "UseIp",
	2: "IpIfNonMatch",
}
var Config_DomainStrategy_value = map[string]int32{
	"AsIs":         0,
	"UseIp":        1,
	"IpIfNonMatch": 2,
}

func (x Config_DomainStrategy) String() string {
	return proto.EnumName(Config_DomainStrategy_name, int32(x))
}
func (Config_DomainStrategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

// Domain for routing decision.
type Domain struct {
	// Domain matching type.
	Type Domain_Type `protobuf:"varint,1,opt,name=type,enum=v2ray.core.app.router.Domain_Type" json:"type,omitempty"`
	// Domain value.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Domain) GetType() Domain_Type {
	if m != nil {
		return m.Type
	}
	return Domain_Plain
}

func (m *Domain) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IP for routing decision, in CIDR form.
type CIDR struct {
	// IP address, should be either 4 or 16 bytes.
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Number of leading ones in the network mask.
	Prefix uint32 `protobuf:"varint,2,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *CIDR) Reset()                    { *m = CIDR{} }
func (m *CIDR) String() string            { return proto.CompactTextString(m) }
func (*CIDR) ProtoMessage()               {}
func (*CIDR) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CIDR) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *CIDR) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

type GeoIP struct {
	CountryCode string  `protobuf:"bytes,1,opt,name=country_code,json=countryCode" json:"country_code,omitempty"`
	Cidr        []*CIDR `protobuf:"bytes,2,rep,name=cidr" json:"cidr,omitempty"`
}

func (m *GeoIP) Reset()                    { *m = GeoIP{} }
func (m *GeoIP) String() string            { return proto.CompactTextString(m) }
func (*GeoIP) ProtoMessage()               {}
func (*GeoIP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GeoIP) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *GeoIP) GetCidr() []*CIDR {
	if m != nil {
		return m.Cidr
	}
	return nil
}

type GeoIPList struct {
	Entry []*GeoIP `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *GeoIPList) Reset()                    { *m = GeoIPList{} }
func (m *GeoIPList) String() string            { return proto.CompactTextString(m) }
func (*GeoIPList) ProtoMessage()               {}
func (*GeoIPList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GeoIPList) GetEntry() []*GeoIP {
	if m != nil {
		return m.Entry
	}
	return nil
}

type RoutingRule struct {
	Tag         string                              `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	Domain      []*Domain                           `protobuf:"bytes,2,rep,name=domain" json:"domain,omitempty"`
	Cidr        []*CIDR                             `protobuf:"bytes,3,rep,name=cidr" json:"cidr,omitempty"`
	PortRange   *v2ray_core_common_net.PortRange    `protobuf:"bytes,4,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	NetworkList *v2ray_core_common_net1.NetworkList `protobuf:"bytes,5,opt,name=network_list,json=networkList" json:"network_list,omitempty"`
	SourceCidr  []*CIDR                             `protobuf:"bytes,6,rep,name=source_cidr,json=sourceCidr" json:"source_cidr,omitempty"`
	UserEmail   []string                            `protobuf:"bytes,7,rep,name=user_email,json=userEmail" json:"user_email,omitempty"`
	InboundTag  []string                            `protobuf:"bytes,8,rep,name=inbound_tag,json=inboundTag" json:"inbound_tag,omitempty"`
}

func (m *RoutingRule) Reset()                    { *m = RoutingRule{} }
func (m *RoutingRule) String() string            { return proto.CompactTextString(m) }
func (*RoutingRule) ProtoMessage()               {}
func (*RoutingRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoutingRule) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *RoutingRule) GetDomain() []*Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *RoutingRule) GetCidr() []*CIDR {
	if m != nil {
		return m.Cidr
	}
	return nil
}

func (m *RoutingRule) GetPortRange() *v2ray_core_common_net.PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *RoutingRule) GetNetworkList() *v2ray_core_common_net1.NetworkList {
	if m != nil {
		return m.NetworkList
	}
	return nil
}

func (m *RoutingRule) GetSourceCidr() []*CIDR {
	if m != nil {
		return m.SourceCidr
	}
	return nil
}

func (m *RoutingRule) GetUserEmail() []string {
	if m != nil {
		return m.UserEmail
	}
	return nil
}

func (m *RoutingRule) GetInboundTag() []string {
	if m != nil {
		return m.InboundTag
	}
	return nil
}

type Config struct {
	DomainStrategy Config_DomainStrategy `protobuf:"varint,1,opt,name=domain_strategy,json=domainStrategy,enum=v2ray.core.app.router.Config_DomainStrategy" json:"domain_strategy,omitempty"`
	Rule           []*RoutingRule        `protobuf:"bytes,2,rep,name=rule" json:"rule,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Config) GetDomainStrategy() Config_DomainStrategy {
	if m != nil {
		return m.DomainStrategy
	}
	return Config_AsIs
}

func (m *Config) GetRule() []*RoutingRule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func init() {
	proto.RegisterType((*Domain)(nil), "v2ray.core.app.router.Domain")
	proto.RegisterType((*CIDR)(nil), "v2ray.core.app.router.CIDR")
	proto.RegisterType((*GeoIP)(nil), "v2ray.core.app.router.GeoIP")
	proto.RegisterType((*GeoIPList)(nil), "v2ray.core.app.router.GeoIPList")
	proto.RegisterType((*RoutingRule)(nil), "v2ray.core.app.router.RoutingRule")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.router.Config")
	proto.RegisterEnum("v2ray.core.app.router.Domain_Type", Domain_Type_name, Domain_Type_value)
	proto.RegisterEnum("v2ray.core.app.router.Config_DomainStrategy", Config_DomainStrategy_name, Config_DomainStrategy_value)
}

func init() { proto.RegisterFile("v2ray.com/core/app/router/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0x3f, 0x3b, 0x89, 0xbf, 0x7a, 0x1c, 0x82, 0xb5, 0xa2, 0xc8, 0x14, 0x2a, 0x82, 0x85,
	0x20, 0x07, 0x64, 0x4b, 0x41, 0xc0, 0x05, 0x54, 0x95, 0xb4, 0x42, 0x91, 0xa0, 0x8a, 0x96, 0x96,
	0x03, 0x1c, 0xac, 0xad, 0xb3, 0x35, 0x2b, 0x92, 0xdd, 0xd5, 0x7a, 0x5d, 0x9a, 0x1b, 0x2f, 0xc0,
	0x8b, 0xf0, 0x34, 0x3c, 0x12, 0xda, 0x5d, 0x17, 0x5a, 0xd4, 0x00, 0xb7, 0x9d, 0xf1, 0x6f, 0x66,
	0xfe, 0x33, 0x9e, 0x81, 0x07, 0xa7, 0x63, 0x45, 0x56, 0x59, 0x29, 0x96, 0x79, 0x29, 0x14, 0xcd,
	0x89, 0x94, 0xb9, 0x12, 0x8d, 0xa6, 0x2a, 0x2f, 0x05, 0x3f, 0x61, 0x55, 0x26, 0x95, 0xd0, 0x02,
	0x6d, 0x9e, 0x73, 0x8a, 0x66, 0x44, 0xca, 0xcc, 0x31, 0x5b, 0xf7, 0x7f, 0x0b, 0x2f, 0xc5, 0x72,
	0x29, 0x78, 0xce, 0xa9, 0xce, 0xa5, 0x50, 0xda, 0x05, 0x6f, 0x3d, 0x5c, 0x4f, 0x71, 0xaa, 0x3f,
	0x0b, 0xf5, 0xc9, 0x81, 0xe9, 0x17, 0x0f, 0x82, 0x3d, 0xb1, 0x24, 0x8c, 0xa3, 0xa7, 0xd0, 0xd5,
	0x2b, 0x49, 0x13, 0x6f, 0xe8, 0x8d, 0x06, 0xe3, 0x34, 0xbb, 0xb2, 0x7e, 0xe6, 0xe0, 0xec, 0x70,
	0x25, 0x29, 0xb6, 0x3c, 0xba, 0x01, 0xbd, 0x53, 0xb2, 0x68, 0x68, 0xe2, 0x0f, 0xbd, 0x51, 0x88,
	0x9d, 0x91, 0x8e, 0xa0, 0x6b, 0x18, 0x14, 0x42, 0x6f, 0xb6, 0x20, 0x8c, 0xc7, 0xff, 0x99, 0x27,
	0xa6, 0x15, 0x3d, 0x8b, 0x3d, 0x04, 0xe7, 0x55, 0x63, 0x3f, 0xcd, 0xa0, 0x3b, 0x99, 0xee, 0x61,
	0x34, 0x00, 0x9f, 0x49, 0x5b, 0xbd, 0x8f, 0x7d, 0x26, 0xd1, 0x4d, 0x08, 0xa4, 0xa2, 0x27, 0xec,
	0xcc, 0x26, 0xbe, 0x86, 0x5b, 0x2b, 0xfd, 0x00, 0xbd, 0x57, 0x54, 0x4c, 0x67, 0xe8, 0x1e, 0xf4,
	0x4b, 0xd1, 0x70, 0xad, 0x56, 0x45, 0x29, 0xe6, 0x4e, 0x78, 0x88, 0xa3, 0xd6, 0x37, 0x11, 0x73,
	0x8a, 0x72, 0xe8, 0x96, 0x6c, 0xae, 0x12, 0x7f, 0xd8, 0x19, 0x45, 0xe3, 0xdb, 0x6b, 0x7a, 0x32,
	0xe5, 0xb1, 0x05, 0xd3, 0x1d, 0x08, 0x6d, 0xf2, 0xd7, 0xac, 0xd6, 0x68, 0x0c, 0x3d, 0x6a, 0x52,
	0x25, 0x9e, 0x0d, 0xbf, 0xb3, 0x26, 0xdc, 0x06, 0x60, 0x87, 0xa6, 0x5f, 0x3b, 0x10, 0x61, 0xd1,
	0x68, 0xc6, 0x2b, 0xdc, 0x2c, 0x28, 0x8a, 0xa1, 0xa3, 0x49, 0xd5, 0x6a, 0x33, 0x4f, 0xf4, 0x04,
	0x82, 0xb9, 0xed, 0xbd, 0x55, 0xb5, 0xfd, 0xc7, 0x49, 0xe3, 0x16, 0xfe, 0xd9, 0x4a, 0xe7, 0x1f,
	0x5b, 0x41, 0x3b, 0x00, 0x66, 0x23, 0x0a, 0x45, 0x78, 0x45, 0x93, 0xee, 0xd0, 0x1b, 0x45, 0xe3,
	0xe1, 0xc5, 0x30, 0xb7, 0x14, 0x19, 0xa7, 0x3a, 0x9b, 0x09, 0xa5, 0xb1, 0xe1, 0x70, 0x28, 0xcf,
	0x9f, 0x68, 0x1f, 0xfa, 0xed, 0xb2, 0x14, 0x0b, 0x56, 0xeb, 0xa4, 0x67, 0x53, 0xa4, 0x6b, 0x52,
	0x1c, 0x38, 0xd4, 0x0c, 0x0e, 0x47, 0xfc, 0x97, 0x81, 0x9e, 0x43, 0x54, 0x8b, 0x46, 0x95, 0xb4,
	0xb0, 0xfa, 0x83, 0xbf, 0xeb, 0x07, 0xc7, 0x4f, 0x4c, 0x17, 0xdb, 0x00, 0x4d, 0x4d, 0x55, 0x41,
	0x97, 0x84, 0x2d, 0x92, 0xff, 0x87, 0x9d, 0x51, 0x88, 0x43, 0xe3, 0xd9, 0x37, 0x0e, 0x74, 0x17,
	0x22, 0xc6, 0x8f, 0x45, 0xc3, 0xe7, 0x85, 0x19, 0xf3, 0x86, 0xfd, 0x0e, 0xad, 0xeb, 0x90, 0x54,
	0xe9, 0x77, 0x0f, 0x82, 0x89, 0xbd, 0x2b, 0x74, 0x04, 0xd7, 0xdd, 0x2c, 0x8b, 0x5a, 0x2b, 0xa2,
	0x69, 0xb5, 0x6a, 0x77, 0xfd, 0xd1, 0x3a, 0x31, 0xee, 0x1e, 0xdd, 0x8f, 0x78, 0xdb, 0xc6, 0xe0,
	0xc1, 0xfc, 0x92, 0x6d, 0xee, 0x46, 0x35, 0x0b, 0xda, 0xfe, 0xcd, 0x75, 0x77, 0x73, 0x61, 0x27,
	0xb0, 0xe5, 0xd3, 0x67, 0x30, 0xb8, 0x9c, 0x19, 0x6d, 0x40, 0x77, 0xb7, 0x9e, 0xd6, 0xee, 0x54,
	0x8e, 0x6a, 0x3a, 0x95, 0xb1, 0x87, 0x62, 0xe8, 0x4f, 0xe5, 0xf4, 0xe4, 0x40, 0xf0, 0x37, 0x44,
	0x97, 0x1f, 0x63, 0xff, 0xe5, 0x0b, 0xb8, 0x55, 0x8a, 0xe5, 0xd5, 0x75, 0x66, 0xde, 0xfb, 0xc0,
	0xbd, 0xbe, 0xf9, 0x9b, 0xef, 0xc6, 0x98, 0xac, 0xb2, 0x89, 0x21, 0x76, 0xa5, 0xb4, 0x12, 0xa8,
	0x3a, 0x0e, 0xec, 0xe5, 0x3f, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x12, 0x0e, 0x89, 0x89,
	0x04, 0x00, 0x00,
}
