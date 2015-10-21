// Code generated by protoc-gen-go.
// source: qpov.proto
// DO NOT EDIT!

/*
Package qpov is a generated protocol buffer package.

It is generated from these files:
	qpov.proto

It has these top-level messages:
	GetRequest
	GetReply
	RenewRequest
	RenewReply
	DoneRequest
	DoneReply
	AddRequest
	AddReply
*/
package qpov

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}

type GetReply struct {
	LeaseId          *string `protobuf:"bytes,1,req,name=lease_id" json:"lease_id,omitempty"`
	OrderDefinition  *string `protobuf:"bytes,2,req,name=order_definition" json:"order_definition,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}

func (m *GetReply) GetLeaseId() string {
	if m != nil && m.LeaseId != nil {
		return *m.LeaseId
	}
	return ""
}

func (m *GetReply) GetOrderDefinition() string {
	if m != nil && m.OrderDefinition != nil {
		return *m.OrderDefinition
	}
	return ""
}

type RenewRequest struct {
	LeaseId          *string `protobuf:"bytes,1,req,name=lease_id" json:"lease_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RenewRequest) Reset()         { *m = RenewRequest{} }
func (m *RenewRequest) String() string { return proto.CompactTextString(m) }
func (*RenewRequest) ProtoMessage()    {}

func (m *RenewRequest) GetLeaseId() string {
	if m != nil && m.LeaseId != nil {
		return *m.LeaseId
	}
	return ""
}

type RenewReply struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RenewReply) Reset()         { *m = RenewReply{} }
func (m *RenewReply) String() string { return proto.CompactTextString(m) }
func (*RenewReply) ProtoMessage()    {}

type DoneRequest struct {
	LeaseId          *string `protobuf:"bytes,1,req,name=lease_id" json:"lease_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DoneRequest) Reset()         { *m = DoneRequest{} }
func (m *DoneRequest) String() string { return proto.CompactTextString(m) }
func (*DoneRequest) ProtoMessage()    {}

func (m *DoneRequest) GetLeaseId() string {
	if m != nil && m.LeaseId != nil {
		return *m.LeaseId
	}
	return ""
}

type DoneReply struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DoneReply) Reset()         { *m = DoneReply{} }
func (m *DoneReply) String() string { return proto.CompactTextString(m) }
func (*DoneReply) ProtoMessage()    {}

type AddRequest struct {
	OrderDefinition  *string `protobuf:"bytes,1,req,name=order_definition" json:"order_definition,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}

func (m *AddRequest) GetOrderDefinition() string {
	if m != nil && m.OrderDefinition != nil {
		return *m.OrderDefinition
	}
	return ""
}

type AddReply struct {
	OrderId          *string `protobuf:"bytes,1,req,name=order_id" json:"order_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}

func (m *AddReply) GetOrderId() string {
	if m != nil && m.OrderId != nil {
		return *m.OrderId
	}
	return ""
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Scheduler service

type SchedulerClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewReply, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewReply, error) {
	out := new(RenewReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error) {
	out := new(DoneReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Done", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Renew(context.Context, *RenewRequest) (*RenewReply, error)
	Done(context.Context, *DoneRequest) (*DoneReply, error)
	Add(context.Context, *AddRequest) (*AddReply, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Renew(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Done(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qpov.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Scheduler_Get_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _Scheduler_Renew_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Scheduler_Done_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Scheduler_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
