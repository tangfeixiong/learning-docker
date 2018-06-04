// Code generated by protoc-gen-gogo.
// source: pb/demo.proto
// DO NOT EDIT!

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import raftpb "github.com/coreos/etcd/raft/raftpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DemoReqResp struct {
	DemoMessages []*raftpb.Message `protobuf:"bytes,1,rep,name=demo_messages,json=demoMessages" json:"demo_messages,omitempty"`
	StateCode    int32             `protobuf:"varint,2,opt,name=state_code,json=stateCode,proto3" json:"state_code,omitempty"`
	StateMessage string            `protobuf:"bytes,3,opt,name=state_message,json=stateMessage,proto3" json:"state_message,omitempty"`
}

func (m *DemoReqResp) Reset()                    { *m = DemoReqResp{} }
func (m *DemoReqResp) String() string            { return proto.CompactTextString(m) }
func (*DemoReqResp) ProtoMessage()               {}
func (*DemoReqResp) Descriptor() ([]byte, []int) { return fileDescriptorDemo, []int{0} }

func (m *DemoReqResp) GetDemoMessages() []*raftpb.Message {
	if m != nil {
		return m.DemoMessages
	}
	return nil
}

func (m *DemoReqResp) GetStateCode() int32 {
	if m != nil {
		return m.StateCode
	}
	return 0
}

func (m *DemoReqResp) GetStateMessage() string {
	if m != nil {
		return m.StateMessage
	}
	return ""
}

type BootstrapReqResp struct {
}

func (m *BootstrapReqResp) Reset()                    { *m = BootstrapReqResp{} }
func (m *BootstrapReqResp) String() string            { return proto.CompactTextString(m) }
func (*BootstrapReqResp) ProtoMessage()               {}
func (*BootstrapReqResp) Descriptor() ([]byte, []int) { return fileDescriptorDemo, []int{1} }

func init() {
	proto.RegisterType((*DemoReqResp)(nil), "pb.DemoReqResp")
	proto.RegisterType((*BootstrapReqResp)(nil), "pb.BootstrapReqResp")
}

func init() { proto.RegisterFile("pb/demo.proto", fileDescriptorDemo) }

var fileDescriptorDemo = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0x41, 0xae, 0x82, 0x40,
	0x0c, 0x86, 0x33, 0x90, 0xf7, 0x12, 0x0a, 0x44, 0x9d, 0x15, 0x31, 0x31, 0x21, 0xb8, 0x61, 0x05,
	0x89, 0x7a, 0x02, 0x75, 0xeb, 0x66, 0x2e, 0x40, 0x66, 0xa4, 0xba, 0x22, 0x1d, 0x69, 0xef, 0xe0,
	0xb5, 0x0d, 0xcc, 0xac, 0x9a, 0x7e, 0xed, 0xff, 0xe7, 0x83, 0xd2, 0xbb, 0x7e, 0xc4, 0x89, 0x3a,
	0x3f, 0x93, 0x90, 0x4e, 0xbc, 0xdb, 0xef, 0x66, 0xfb, 0x12, 0xef, 0xfa, 0x65, 0x04, 0xdc, 0x7c,
	0x15, 0xe4, 0x77, 0x9c, 0xc8, 0xe0, 0xc7, 0x20, 0x7b, 0x7d, 0x81, 0x72, 0x09, 0x0d, 0x13, 0x32,
	0xdb, 0x37, 0x72, 0xa5, 0xea, 0xb4, 0xcd, 0x4f, 0x9b, 0x2e, 0x44, 0xbb, 0x47, 0xe0, 0xa6, 0x58,
	0xbe, 0xe2, 0xc2, 0xfa, 0x00, 0xc0, 0x62, 0x05, 0x87, 0x27, 0x8d, 0x58, 0x25, 0xb5, 0x6a, 0xff,
	0x4c, 0xb6, 0x92, 0x1b, 0x8d, 0xa8, 0x8f, 0x50, 0x86, 0x73, 0x6c, 0xad, 0xd2, 0x5a, 0xb5, 0x99,
	0x29, 0x56, 0x18, 0x4b, 0x1a, 0x0d, 0xdb, 0x2b, 0x91, 0xb0, 0xcc, 0xd6, 0x47, 0x1b, 0xf7, 0xbf,
	0x4a, 0x9e, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0x5c, 0xf1, 0xc9, 0xcc, 0x00, 0x00, 0x00,
}
