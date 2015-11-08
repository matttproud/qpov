// Code generated by protoc-gen-go.
// source: scheduler.proto
// DO NOT EDIT!

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
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}

type GetReply struct {
	LeaseId         string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	OrderDefinition string `protobuf:"bytes,2,opt,name=order_definition" json:"order_definition,omitempty"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}

type RenewRequest struct {
	LeaseId   string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	ExtendSec int32  `protobuf:"varint,2,opt,name=extend_sec" json:"extend_sec,omitempty"`
}

func (m *RenewRequest) Reset()         { *m = RenewRequest{} }
func (m *RenewRequest) String() string { return proto.CompactTextString(m) }
func (*RenewRequest) ProtoMessage()    {}

type RenewReply struct {
	NewTimeoutSec int64 `protobuf:"varint,1,opt,name=new_timeout_sec" json:"new_timeout_sec,omitempty"`
}

func (m *RenewReply) Reset()         { *m = RenewReply{} }
func (m *RenewReply) String() string { return proto.CompactTextString(m) }
func (*RenewReply) ProtoMessage()    {}

type DoneRequest struct {
	LeaseId      string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	Image        []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Stdout       []byte `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr       []byte `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
	JsonMetadata string `protobuf:"bytes,5,opt,name=json_metadata" json:"json_metadata,omitempty"`
}

func (m *DoneRequest) Reset()         { *m = DoneRequest{} }
func (m *DoneRequest) String() string { return proto.CompactTextString(m) }
func (*DoneRequest) ProtoMessage()    {}

type DoneReply struct {
}

func (m *DoneReply) Reset()         { *m = DoneReply{} }
func (m *DoneReply) String() string { return proto.CompactTextString(m) }
func (*DoneReply) ProtoMessage()    {}

type AddRequest struct {
	OrderDefinition string `protobuf:"bytes,1,opt,name=order_definition" json:"order_definition,omitempty"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}

type AddReply struct {
	OrderId string `protobuf:"bytes,1,opt,name=order_id" json:"order_id,omitempty"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}

type Lease struct {
	OrderId   string `protobuf:"bytes,1,opt,name=order_id" json:"order_id,omitempty"`
	LeaseId   string `protobuf:"bytes,2,opt,name=lease_id" json:"lease_id,omitempty"`
	Done      bool   `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	UserId    int64  `protobuf:"varint,4,opt,name=user_id" json:"user_id,omitempty"`
	CreatedMs int64  `protobuf:"varint,5,opt,name=created_ms" json:"created_ms,omitempty"`
	UpdatedMs int64  `protobuf:"varint,6,opt,name=updated_ms" json:"updated_ms,omitempty"`
	ExpiresMs int64  `protobuf:"varint,7,opt,name=expires_ms" json:"expires_ms,omitempty"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}

type LeasesRequest struct {
	Done bool `protobuf:"varint,1,opt,name=done" json:"done,omitempty"`
}

func (m *LeasesRequest) Reset()         { *m = LeasesRequest{} }
func (m *LeasesRequest) String() string { return proto.CompactTextString(m) }
func (*LeasesRequest) ProtoMessage()    {}

type LeasesReply struct {
	Leases []*Lease `protobuf:"bytes,1,rep,name=leases" json:"leases,omitempty"`
}

func (m *LeasesReply) Reset()         { *m = LeasesReply{} }
func (m *LeasesReply) String() string { return proto.CompactTextString(m) }
func (*LeasesReply) ProtoMessage()    {}

func (m *LeasesReply) GetLeases() []*Lease {
	if m != nil {
		return m.Leases
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Scheduler service

type SchedulerClient interface {
	// Render client API.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewReply, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error)
	// Order handling API. Restricted.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	// Stats API. Restricted.
	Leases(ctx context.Context, in *LeasesRequest, opts ...grpc.CallOption) (*LeasesReply, error)
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

func (c *schedulerClient) Leases(ctx context.Context, in *LeasesRequest, opts ...grpc.CallOption) (*LeasesReply, error) {
	out := new(LeasesReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Leases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	// Render client API.
	Get(context.Context, *GetRequest) (*GetReply, error)
	Renew(context.Context, *RenewRequest) (*RenewReply, error)
	Done(context.Context, *DoneRequest) (*DoneReply, error)
	// Order handling API. Restricted.
	Add(context.Context, *AddRequest) (*AddReply, error)
	// Stats API. Restricted.
	Leases(context.Context, *LeasesRequest) (*LeasesReply, error)
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

func _Scheduler_Leases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LeasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Leases(ctx, in)
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
		{
			MethodName: "Leases",
			Handler:    _Scheduler_Leases_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
