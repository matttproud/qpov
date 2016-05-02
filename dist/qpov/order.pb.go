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
	BatchId string   `protobuf:"bytes,6,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Order)(nil), "qpov.Order")
}

var fileDescriptor1 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2c, 0xc8, 0x2f, 0x53, 0x9a, 0xc8,
	0xc8, 0xc5, 0xea, 0x0f, 0x12, 0x15, 0x92, 0xe0, 0x62, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x4c, 0x4f,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x04, 0xb8, 0x98, 0x53, 0x32, 0x8b,
	0x24, 0x98, 0xc0, 0xa2, 0x20, 0xa6, 0x90, 0x10, 0x17, 0x4b, 0x5a, 0x66, 0x4e, 0xaa, 0x04, 0x33,
	0x58, 0x08, 0xcc, 0x06, 0x89, 0x25, 0x16, 0xa5, 0x17, 0x4b, 0xb0, 0x28, 0x30, 0x83, 0xc4, 0x40,
	0x6c, 0x21, 0x49, 0x2e, 0x0e, 0xb0, 0x95, 0xf1, 0x99, 0x29, 0x12, 0xac, 0x10, 0x43, 0xc1, 0x7c,
	0xcf, 0x14, 0x90, 0x54, 0x52, 0x62, 0x49, 0x72, 0x06, 0x48, 0x8a, 0x0d, 0x22, 0x05, 0xe6, 0x7b,
	0xa6, 0x24, 0xb1, 0x81, 0x1d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xb7, 0xcf, 0x25,
	0xaf, 0x00, 0x00, 0x00,
}
