// Code generated by protoc-gen-go.
// source: order.proto
// DO NOT EDIT!

package qpov

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Order struct {
	Package string   `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Dir     string   `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
	File    string   `protobuf:"bytes,3,opt,name=file" json:"file,omitempty"`
	Args    []string `protobuf:"bytes,4,rep,name=args" json:"args,omitempty"`
	OrderId string   `protobuf:"bytes,5,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Order)(nil), "qpov.Order")
}

var fileDescriptor1 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2c, 0xc8, 0x2f, 0x53, 0x2a, 0xe3,
	0x62, 0xf5, 0x07, 0x09, 0x0a, 0x49, 0x70, 0xb1, 0x17, 0x24, 0x26, 0x67, 0x27, 0xa6, 0xa7, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x02, 0x5c, 0xcc, 0x29, 0x99, 0x45, 0x12,
	0x4c, 0x60, 0x51, 0x10, 0x53, 0x48, 0x88, 0x8b, 0x25, 0x2d, 0x33, 0x27, 0x55, 0x82, 0x19, 0x2c,
	0x04, 0x66, 0x83, 0xc4, 0x12, 0x8b, 0xd2, 0x8b, 0x25, 0x58, 0x14, 0x98, 0x41, 0x62, 0x20, 0xb6,
	0x90, 0x24, 0x17, 0x07, 0xd8, 0xc6, 0xf8, 0xcc, 0x14, 0x09, 0x56, 0x88, 0xa1, 0x60, 0xbe, 0x67,
	0x4a, 0x12, 0x1b, 0xd8, 0x11, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xbb, 0x76, 0xb7,
	0x93, 0x00, 0x00, 0x00,
}
