// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	import.proto

It has these top-level messages:
	Empty
	Configure
	Get
	Close
	Value
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Type is an enum representing the type of the value. This isn't the
// full set of Sentinel types since some types cannot be sent via
// Protobufs such as rules or functions.
type Value_Type int32

const (
	Value_INVALID   Value_Type = 0
	Value_UNDEFINED Value_Type = 1
	Value_NULL      Value_Type = 2
	Value_BOOL      Value_Type = 3
	Value_INT       Value_Type = 4
	Value_FLOAT     Value_Type = 5
	Value_STRING    Value_Type = 6
	Value_LIST      Value_Type = 7
	Value_MAP       Value_Type = 8
)

var Value_Type_name = map[int32]string{
	0: "INVALID",
	1: "UNDEFINED",
	2: "NULL",
	3: "BOOL",
	4: "INT",
	5: "FLOAT",
	6: "STRING",
	7: "LIST",
	8: "MAP",
}
var Value_Type_value = map[string]int32{
	"INVALID":   0,
	"UNDEFINED": 1,
	"NULL":      2,
	"BOOL":      3,
	"INT":       4,
	"FLOAT":     5,
	"STRING":    6,
	"LIST":      7,
	"MAP":       8,
}

func (x Value_Type) String() string {
	return proto1.EnumName(Value_Type_name, int32(x))
}
func (Value_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// Empty is just an empty message.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Configure are the structures for Import.Configure
type Configure struct {
}

func (m *Configure) Reset()                    { *m = Configure{} }
func (m *Configure) String() string            { return proto1.CompactTextString(m) }
func (*Configure) ProtoMessage()               {}
func (*Configure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Configure_Request struct {
	Config *Value `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *Configure_Request) Reset()                    { *m = Configure_Request{} }
func (m *Configure_Request) String() string            { return proto1.CompactTextString(m) }
func (*Configure_Request) ProtoMessage()               {}
func (*Configure_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Configure_Request) GetConfig() *Value {
	if m != nil {
		return m.Config
	}
	return nil
}

type Configure_Response struct {
	InstanceId uint64 `protobuf:"varint,1,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
}

func (m *Configure_Response) Reset()                    { *m = Configure_Response{} }
func (m *Configure_Response) String() string            { return proto1.CompactTextString(m) }
func (*Configure_Response) ProtoMessage()               {}
func (*Configure_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Configure_Response) GetInstanceId() uint64 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

// Get are the structures for an Import.Get.
type Get struct {
}

func (m *Get) Reset()                    { *m = Get{} }
func (m *Get) String() string            { return proto1.CompactTextString(m) }
func (*Get) ProtoMessage()               {}
func (*Get) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Request is a single request for a Get.
type Get_Request struct {
	InstanceId   uint64   `protobuf:"varint,1,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	ExecId       uint64   `protobuf:"varint,2,opt,name=exec_id,json=execId" json:"exec_id,omitempty"`
	ExecDeadline uint64   `protobuf:"varint,3,opt,name=exec_deadline,json=execDeadline" json:"exec_deadline,omitempty"`
	Keys         []string `protobuf:"bytes,4,rep,name=keys" json:"keys,omitempty"`
	KeyId        uint64   `protobuf:"varint,5,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Call         bool     `protobuf:"varint,6,opt,name=call" json:"call,omitempty"`
	Args         []*Value `protobuf:"bytes,7,rep,name=args" json:"args,omitempty"`
}

func (m *Get_Request) Reset()                    { *m = Get_Request{} }
func (m *Get_Request) String() string            { return proto1.CompactTextString(m) }
func (*Get_Request) ProtoMessage()               {}
func (*Get_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Get_Request) GetInstanceId() uint64 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *Get_Request) GetExecId() uint64 {
	if m != nil {
		return m.ExecId
	}
	return 0
}

func (m *Get_Request) GetExecDeadline() uint64 {
	if m != nil {
		return m.ExecDeadline
	}
	return 0
}

func (m *Get_Request) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Get_Request) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Get_Request) GetCall() bool {
	if m != nil {
		return m.Call
	}
	return false
}

func (m *Get_Request) GetArgs() []*Value {
	if m != nil {
		return m.Args
	}
	return nil
}

// Response is a single response for a Get.
type Get_Response struct {
	InstanceId uint64   `protobuf:"varint,1,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	KeyId      uint64   `protobuf:"varint,2,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Keys       []string `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	Value      *Value   `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Get_Response) Reset()                    { *m = Get_Response{} }
func (m *Get_Response) String() string            { return proto1.CompactTextString(m) }
func (*Get_Response) ProtoMessage()               {}
func (*Get_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *Get_Response) GetInstanceId() uint64 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *Get_Response) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Get_Response) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Get_Response) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// MultiRequest allows multiple requests in a single Get.
type Get_MultiRequest struct {
	Requests []*Get_Request `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *Get_MultiRequest) Reset()                    { *m = Get_MultiRequest{} }
func (m *Get_MultiRequest) String() string            { return proto1.CompactTextString(m) }
func (*Get_MultiRequest) ProtoMessage()               {}
func (*Get_MultiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 2} }

func (m *Get_MultiRequest) GetRequests() []*Get_Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

// MultiResponse allows multiple responses in a single Get.
type Get_MultiResponse struct {
	Responses []*Get_Response `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
}

func (m *Get_MultiResponse) Reset()                    { *m = Get_MultiResponse{} }
func (m *Get_MultiResponse) String() string            { return proto1.CompactTextString(m) }
func (*Get_MultiResponse) ProtoMessage()               {}
func (*Get_MultiResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 3} }

func (m *Get_MultiResponse) GetResponses() []*Get_Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

// Close contains the structures for Close RPC calls.
type Close struct {
}

func (m *Close) Reset()                    { *m = Close{} }
func (m *Close) String() string            { return proto1.CompactTextString(m) }
func (*Close) ProtoMessage()               {}
func (*Close) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Close_Request struct {
	InstanceId uint64 `protobuf:"varint,1,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
}

func (m *Close_Request) Reset()                    { *m = Close_Request{} }
func (m *Close_Request) String() string            { return proto1.CompactTextString(m) }
func (*Close_Request) ProtoMessage()               {}
func (*Close_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Close_Request) GetInstanceId() uint64 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

// Value represents a Sentinel value.
type Value struct {
	// type is the type of this value
	Type Value_Type `protobuf:"varint,1,opt,name=type,enum=proto.Value_Type" json:"type,omitempty"`
	// value is the value only if the type is not UNDEFINED or NULL.
	// If the value is UNDEFINED or NULL, then the value is known.
	//
	// Types that are valid to be assigned to Value:
	//	*Value_ValueBool
	//	*Value_ValueInt
	//	*Value_ValueFloat
	//	*Value_ValueString
	//	*Value_ValueList
	//	*Value_ValueMap
	Value isValue_Value `protobuf_oneof:"value"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto1.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isValue_Value interface {
	isValue_Value()
}

type Value_ValueBool struct {
	ValueBool bool `protobuf:"varint,2,opt,name=value_bool,json=valueBool,oneof"`
}
type Value_ValueInt struct {
	ValueInt int64 `protobuf:"varint,3,opt,name=value_int,json=valueInt,oneof"`
}
type Value_ValueFloat struct {
	ValueFloat float64 `protobuf:"fixed64,4,opt,name=value_float,json=valueFloat,oneof"`
}
type Value_ValueString struct {
	ValueString string `protobuf:"bytes,5,opt,name=value_string,json=valueString,oneof"`
}
type Value_ValueList struct {
	ValueList *Value_List `protobuf:"bytes,6,opt,name=value_list,json=valueList,oneof"`
}
type Value_ValueMap struct {
	ValueMap *Value_Map `protobuf:"bytes,7,opt,name=value_map,json=valueMap,oneof"`
}

func (*Value_ValueBool) isValue_Value()   {}
func (*Value_ValueInt) isValue_Value()    {}
func (*Value_ValueFloat) isValue_Value()  {}
func (*Value_ValueString) isValue_Value() {}
func (*Value_ValueList) isValue_Value()   {}
func (*Value_ValueMap) isValue_Value()    {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetType() Value_Type {
	if m != nil {
		return m.Type
	}
	return Value_INVALID
}

func (m *Value) GetValueBool() bool {
	if x, ok := m.GetValue().(*Value_ValueBool); ok {
		return x.ValueBool
	}
	return false
}

func (m *Value) GetValueInt() int64 {
	if x, ok := m.GetValue().(*Value_ValueInt); ok {
		return x.ValueInt
	}
	return 0
}

func (m *Value) GetValueFloat() float64 {
	if x, ok := m.GetValue().(*Value_ValueFloat); ok {
		return x.ValueFloat
	}
	return 0
}

func (m *Value) GetValueString() string {
	if x, ok := m.GetValue().(*Value_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *Value) GetValueList() *Value_List {
	if x, ok := m.GetValue().(*Value_ValueList); ok {
		return x.ValueList
	}
	return nil
}

func (m *Value) GetValueMap() *Value_Map {
	if x, ok := m.GetValue().(*Value_ValueMap); ok {
		return x.ValueMap
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_ValueBool)(nil),
		(*Value_ValueInt)(nil),
		(*Value_ValueFloat)(nil),
		(*Value_ValueString)(nil),
		(*Value_ValueList)(nil),
		(*Value_ValueMap)(nil),
	}
}

func _Value_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_ValueBool:
		t := uint64(0)
		if x.ValueBool {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto1.WireVarint)
		b.EncodeVarint(t)
	case *Value_ValueInt:
		b.EncodeVarint(3<<3 | proto1.WireVarint)
		b.EncodeVarint(uint64(x.ValueInt))
	case *Value_ValueFloat:
		b.EncodeVarint(4<<3 | proto1.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.ValueFloat))
	case *Value_ValueString:
		b.EncodeVarint(5<<3 | proto1.WireBytes)
		b.EncodeStringBytes(x.ValueString)
	case *Value_ValueList:
		b.EncodeVarint(6<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.ValueList); err != nil {
			return err
		}
	case *Value_ValueMap:
		b.EncodeVarint(7<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.ValueMap); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 2: // value.value_bool
		if wire != proto1.WireVarint {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_ValueBool{x != 0}
		return true, err
	case 3: // value.value_int
		if wire != proto1.WireVarint {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_ValueInt{int64(x)}
		return true, err
	case 4: // value.value_float
		if wire != proto1.WireFixed64 {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Value_ValueFloat{math.Float64frombits(x)}
		return true, err
	case 5: // value.value_string
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Value_ValueString{x}
		return true, err
	case 6: // value.value_list
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(Value_List)
		err := b.DecodeMessage(msg)
		m.Value = &Value_ValueList{msg}
		return true, err
	case 7: // value.value_map
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(Value_Map)
		err := b.DecodeMessage(msg)
		m.Value = &Value_ValueMap{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_ValueBool:
		n += proto1.SizeVarint(2<<3 | proto1.WireVarint)
		n += 1
	case *Value_ValueInt:
		n += proto1.SizeVarint(3<<3 | proto1.WireVarint)
		n += proto1.SizeVarint(uint64(x.ValueInt))
	case *Value_ValueFloat:
		n += proto1.SizeVarint(4<<3 | proto1.WireFixed64)
		n += 8
	case *Value_ValueString:
		n += proto1.SizeVarint(5<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(len(x.ValueString)))
		n += len(x.ValueString)
	case *Value_ValueList:
		s := proto1.Size(x.ValueList)
		n += proto1.SizeVarint(6<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *Value_ValueMap:
		s := proto1.Size(x.ValueMap)
		n += proto1.SizeVarint(7<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Value_KV struct {
	Key   *Value `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Value_KV) Reset()                    { *m = Value_KV{} }
func (m *Value_KV) String() string            { return proto1.CompactTextString(m) }
func (*Value_KV) ProtoMessage()               {}
func (*Value_KV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Value_KV) GetKey() *Value {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Value_KV) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Value_Map struct {
	Elems []*Value_KV `protobuf:"bytes,1,rep,name=elems" json:"elems,omitempty"`
}

func (m *Value_Map) Reset()                    { *m = Value_Map{} }
func (m *Value_Map) String() string            { return proto1.CompactTextString(m) }
func (*Value_Map) ProtoMessage()               {}
func (*Value_Map) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *Value_Map) GetElems() []*Value_KV {
	if m != nil {
		return m.Elems
	}
	return nil
}

type Value_List struct {
	Elems []*Value `protobuf:"bytes,1,rep,name=elems" json:"elems,omitempty"`
}

func (m *Value_List) Reset()                    { *m = Value_List{} }
func (m *Value_List) String() string            { return proto1.CompactTextString(m) }
func (*Value_List) ProtoMessage()               {}
func (*Value_List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 2} }

func (m *Value_List) GetElems() []*Value {
	if m != nil {
		return m.Elems
	}
	return nil
}

func init() {
	proto1.RegisterType((*Empty)(nil), "proto.Empty")
	proto1.RegisterType((*Configure)(nil), "proto.Configure")
	proto1.RegisterType((*Configure_Request)(nil), "proto.Configure.Request")
	proto1.RegisterType((*Configure_Response)(nil), "proto.Configure.Response")
	proto1.RegisterType((*Get)(nil), "proto.Get")
	proto1.RegisterType((*Get_Request)(nil), "proto.Get.Request")
	proto1.RegisterType((*Get_Response)(nil), "proto.Get.Response")
	proto1.RegisterType((*Get_MultiRequest)(nil), "proto.Get.MultiRequest")
	proto1.RegisterType((*Get_MultiResponse)(nil), "proto.Get.MultiResponse")
	proto1.RegisterType((*Close)(nil), "proto.Close")
	proto1.RegisterType((*Close_Request)(nil), "proto.Close.Request")
	proto1.RegisterType((*Value)(nil), "proto.Value")
	proto1.RegisterType((*Value_KV)(nil), "proto.Value.KV")
	proto1.RegisterType((*Value_Map)(nil), "proto.Value.Map")
	proto1.RegisterType((*Value_List)(nil), "proto.Value.List")
	proto1.RegisterEnum("proto.Value_Type", Value_Type_name, Value_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Import service

type ImportClient interface {
	Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error)
	Get(ctx context.Context, in *Get_MultiRequest, opts ...grpc.CallOption) (*Get_MultiResponse, error)
	Close(ctx context.Context, in *Close_Request, opts ...grpc.CallOption) (*Empty, error)
}

type importClient struct {
	cc *grpc.ClientConn
}

func NewImportClient(cc *grpc.ClientConn) ImportClient {
	return &importClient{cc}
}

func (c *importClient) Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error) {
	out := new(Configure_Response)
	err := grpc.Invoke(ctx, "/proto.Import/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importClient) Get(ctx context.Context, in *Get_MultiRequest, opts ...grpc.CallOption) (*Get_MultiResponse, error) {
	out := new(Get_MultiResponse)
	err := grpc.Invoke(ctx, "/proto.Import/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importClient) Close(ctx context.Context, in *Close_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.Import/Close", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Import service

type ImportServer interface {
	Configure(context.Context, *Configure_Request) (*Configure_Response, error)
	Get(context.Context, *Get_MultiRequest) (*Get_MultiResponse, error)
	Close(context.Context, *Close_Request) (*Empty, error)
}

func RegisterImportServer(s *grpc.Server, srv ImportServer) {
	s.RegisterService(&_Import_serviceDesc, srv)
}

func _Import_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Import/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).Configure(ctx, req.(*Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Import_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Get_MultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Import/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).Get(ctx, req.(*Get_MultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Import_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Close_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Import/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).Close(ctx, req.(*Close_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Import_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Import",
	HandlerType: (*ImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _Import_Configure_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Import_Get_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Import_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "import.proto",
}

func init() { proto1.RegisterFile("import.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xda, 0x58,
	0x10, 0xc6, 0xf8, 0x0f, 0x06, 0xb2, 0xeb, 0x9d, 0xdd, 0x55, 0xbc, 0x96, 0xb6, 0xa1, 0x4e, 0x23,
	0xa1, 0xa4, 0x22, 0x2a, 0xb9, 0xe9, 0x55, 0xd5, 0x10, 0x92, 0x60, 0x05, 0x48, 0x75, 0x42, 0xb8,
	0x8d, 0x1c, 0x38, 0x89, 0x2c, 0x8c, 0xed, 0xe2, 0x43, 0x55, 0xfa, 0x56, 0x55, 0x9f, 0xa2, 0x97,
	0x7d, 0xa3, 0xca, 0x63, 0x9b, 0x00, 0x49, 0xd5, 0xf6, 0xca, 0x67, 0xe6, 0x9b, 0x9f, 0x6f, 0xe6,
	0x7c, 0x3e, 0x50, 0xf5, 0xa6, 0x51, 0x38, 0x13, 0x8d, 0x68, 0x16, 0x8a, 0x10, 0x55, 0xfa, 0xd8,
	0x3a, 0xa8, 0xa7, 0xd3, 0x48, 0x2c, 0x6c, 0x0f, 0xca, 0x27, 0x61, 0x70, 0xe7, 0xdd, 0xcf, 0x67,
	0xdc, 0x3a, 0x04, 0x9d, 0xf1, 0xf7, 0x73, 0x1e, 0x0b, 0x7c, 0x01, 0xda, 0x88, 0xfc, 0xa6, 0x5c,
	0x93, 0xea, 0x95, 0x66, 0x35, 0xcd, 0x6f, 0x0c, 0x5d, 0x7f, 0xce, 0x59, 0x86, 0x59, 0x07, 0x50,
	0x62, 0x3c, 0x8e, 0xc2, 0x20, 0xe6, 0xb8, 0x03, 0x15, 0x2f, 0x88, 0x85, 0x1b, 0x8c, 0xf8, 0x8d,
	0x37, 0x36, 0xa5, 0x9a, 0x54, 0x57, 0x18, 0xe4, 0x2e, 0x67, 0x6c, 0x7f, 0x93, 0x41, 0x3e, 0xe7,
	0xc2, 0xfa, 0x2a, 0x3d, 0xb4, 0xf9, 0x59, 0x12, 0x6e, 0x83, 0xce, 0x3f, 0xf2, 0x51, 0x02, 0x16,
	0x09, 0xd4, 0x12, 0xd3, 0x19, 0xe3, 0x2e, 0x6c, 0x11, 0x30, 0xe6, 0xee, 0xd8, 0xf7, 0x02, 0x4e,
	0x3c, 0x15, 0x56, 0x4d, 0x9c, 0xed, 0xcc, 0x87, 0x08, 0xca, 0x84, 0x2f, 0x62, 0x53, 0xa9, 0xc9,
	0xf5, 0x32, 0xa3, 0x33, 0xfe, 0x0b, 0xda, 0x84, 0x2f, 0x92, 0x82, 0x2a, 0x65, 0xa8, 0x13, 0xbe,
	0x70, 0xc6, 0x49, 0xe8, 0xc8, 0xf5, 0x7d, 0x53, 0xab, 0x49, 0xf5, 0x12, 0xa3, 0x33, 0xd6, 0x40,
	0x71, 0x67, 0xf7, 0xb1, 0xa9, 0xd7, 0xe4, 0x47, 0x2b, 0x20, 0xc4, 0xfa, 0xf4, 0x1b, 0x0b, 0x58,
	0xe9, 0x5c, 0xdc, 0xe8, 0x4c, 0x24, 0xe5, 0x15, 0x92, 0x36, 0xa8, 0x1f, 0x92, 0x36, 0xa6, 0xf2,
	0xc4, 0xf6, 0x53, 0xc8, 0x7a, 0x03, 0xd5, 0xde, 0xdc, 0x17, 0x5e, 0xbe, 0xcb, 0x06, 0x94, 0x66,
	0xe9, 0x31, 0x36, 0x25, 0x62, 0x8c, 0x59, 0xda, 0x39, 0x17, 0x8d, 0x2c, 0x8a, 0x2d, 0x63, 0xac,
	0x16, 0x6c, 0x65, 0xf9, 0xd9, 0x00, 0xaf, 0xa0, 0x3c, 0xcb, 0xce, 0x79, 0x85, 0xbf, 0xd7, 0x2a,
	0xa4, 0x18, 0x7b, 0x88, 0xb2, 0x8f, 0x40, 0x3d, 0xf1, 0xc3, 0x98, 0x5b, 0xfb, 0xbf, 0x7e, 0xa7,
	0xf6, 0x17, 0x05, 0x54, 0x9a, 0x04, 0xf7, 0x40, 0x11, 0x8b, 0x88, 0x53, 0xcc, 0x1f, 0xcd, 0xbf,
	0x56, 0xa7, 0x6c, 0x0c, 0x16, 0x11, 0x67, 0x04, 0xe3, 0x0e, 0x00, 0x8d, 0x7c, 0x73, 0x1b, 0x86,
	0x3e, 0x2d, 0xaf, 0xd4, 0x29, 0xb0, 0x32, 0xf9, 0x5a, 0x61, 0xe8, 0xe3, 0xff, 0x90, 0x1a, 0x37,
	0x5e, 0x20, 0x48, 0x08, 0x72, 0xa7, 0xc0, 0x4a, 0xe4, 0x72, 0x02, 0x81, 0xcf, 0xa1, 0x92, 0xc2,
	0x77, 0x7e, 0xe8, 0x0a, 0xda, 0xa9, 0xd4, 0x29, 0xb0, 0xb4, 0xe8, 0x59, 0xe2, 0xc3, 0x5d, 0xa8,
	0xa6, 0x21, 0xb1, 0x98, 0x79, 0xc1, 0x3d, 0x69, 0xa3, 0xdc, 0x29, 0xb0, 0x34, 0xf1, 0x8a, 0x9c,
	0xd8, 0xcc, 0x79, 0xf8, 0x5e, 0x2c, 0x48, 0x29, 0x95, 0x0d, 0xd2, 0x5d, 0x2f, 0x16, 0x4b, 0x6a,
	0x89, 0x81, 0x87, 0x39, 0xb5, 0xa9, 0x1b, 0x99, 0x3a, 0xa5, 0x18, 0x6b, 0x29, 0x3d, 0x37, 0x5a,
	0x92, 0xed, 0xb9, 0x91, 0xd5, 0x81, 0xe2, 0xc5, 0x10, 0x9f, 0x81, 0x3c, 0xe1, 0x0b, 0x5a, 0xcc,
	0xe6, 0xf5, 0x27, 0xc0, 0x83, 0x40, 0x8a, 0x3f, 0x16, 0xc8, 0x4b, 0x90, 0x7b, 0x6e, 0x84, 0x7b,
	0xa0, 0x72, 0x9f, 0x4f, 0xf3, 0x2b, 0xfd, 0x73, 0xad, 0xfb, 0xc5, 0x90, 0xa5, 0xa8, 0xb5, 0x0f,
	0x0a, 0x11, 0xb6, 0xd7, 0xc3, 0x37, 0x2a, 0x13, 0x64, 0x7b, 0xa0, 0x24, 0xd7, 0x83, 0x15, 0xd0,
	0x9d, 0xfe, 0xf0, 0xb8, 0xeb, 0xb4, 0x8d, 0x02, 0x6e, 0x41, 0xf9, 0xba, 0xdf, 0x3e, 0x3d, 0x73,
	0xfa, 0xa7, 0x6d, 0x43, 0xc2, 0x12, 0x28, 0xfd, 0xeb, 0x6e, 0xd7, 0x28, 0x26, 0xa7, 0xd6, 0xe5,
	0x65, 0xd7, 0x90, 0x51, 0x07, 0xd9, 0xe9, 0x0f, 0x0c, 0x05, 0xcb, 0xa0, 0x9e, 0x75, 0x2f, 0x8f,
	0x07, 0x86, 0x8a, 0x00, 0xda, 0xd5, 0x80, 0x39, 0xfd, 0x73, 0x43, 0x4b, 0x22, 0xbb, 0xce, 0xd5,
	0xc0, 0xd0, 0x93, 0xc8, 0xde, 0xf1, 0x3b, 0xa3, 0xd4, 0xd2, 0xb3, 0x41, 0x9b, 0x9f, 0x25, 0xd0,
	0x1c, 0x7a, 0xca, 0xf0, 0xed, 0xca, 0xa3, 0x85, 0x66, 0x46, 0x70, 0xe9, 0xc9, 0xa5, 0x6e, 0xfd,
	0xf7, 0x04, 0x92, 0x49, 0xfd, 0x35, 0x3d, 0x45, 0xb8, 0xbd, 0x22, 0xef, 0xd5, 0x7f, 0xc9, 0x32,
	0x1f, 0x03, 0x59, 0xe6, 0x41, 0xa6, 0x78, 0xfc, 0x27, 0xaf, 0x9e, 0x58, 0xcb, 0x9e, 0xf9, 0xba,
	0xe8, 0x75, 0xbd, 0xd5, 0xc8, 0x38, 0xfa, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x24, 0x8f, 0x67, 0x0d,
	0x84, 0x05, 0x00, 0x00,
}
