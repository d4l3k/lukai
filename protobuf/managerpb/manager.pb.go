// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/managerpb/manager.proto

/*
	Package managerpb is a generated protocol buffer package.

	It is generated from these files:
		protobuf/managerpb/manager.proto

	It has these top-level messages:
		Model
		UploadModelRequest
		UploadModelResponse
*/
package managerpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import aggregatorpb "github.com/luk-ai/lukai/protobuf/aggregatorpb"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Model struct {
	Domain      string                   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	ModelType   string                   `protobuf:"bytes,3,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Name        string                   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string                   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	HyperParams aggregatorpb.HyperParams `protobuf:"bytes,8,opt,name=hyper_params,json=hyperParams" json:"hyper_params"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{0} }

func (m *Model) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Model) GetModelType() string {
	if m != nil {
		return m.ModelType
	}
	return ""
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Model) GetHyperParams() aggregatorpb.HyperParams {
	if m != nil {
		return m.HyperParams
	}
	return aggregatorpb.HyperParams{}
}

// UploadModelRequest is a request to upload a model and start training it.
type UploadModelRequest struct {
	Meta  Model  `protobuf:"bytes,1,opt,name=meta" json:"meta"`
	Model []byte `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
}

func (m *UploadModelRequest) Reset()                    { *m = UploadModelRequest{} }
func (*UploadModelRequest) ProtoMessage()               {}
func (*UploadModelRequest) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{1} }

func (m *UploadModelRequest) GetMeta() Model {
	if m != nil {
		return m.Meta
	}
	return Model{}
}

func (m *UploadModelRequest) GetModel() []byte {
	if m != nil {
		return m.Model
	}
	return nil
}

type UploadModelResponse struct {
	ModelId  uint64 `protobuf:"varint,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	ModelUrl string `protobuf:"bytes,2,opt,name=model_url,json=modelUrl,proto3" json:"model_url,omitempty"`
}

func (m *UploadModelResponse) Reset()                    { *m = UploadModelResponse{} }
func (*UploadModelResponse) ProtoMessage()               {}
func (*UploadModelResponse) Descriptor() ([]byte, []int) { return fileDescriptorManager, []int{2} }

func (m *UploadModelResponse) GetModelId() uint64 {
	if m != nil {
		return m.ModelId
	}
	return 0
}

func (m *UploadModelResponse) GetModelUrl() string {
	if m != nil {
		return m.ModelUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Model)(nil), "managerpb.Model")
	proto.RegisterType((*UploadModelRequest)(nil), "managerpb.UploadModelRequest")
	proto.RegisterType((*UploadModelResponse)(nil), "managerpb.UploadModelResponse")
}
func (this *Model) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Model)
	if !ok {
		that2, ok := that.(Model)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Domain != that1.Domain {
		return false
	}
	if this.ModelType != that1.ModelType {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.HyperParams.Equal(&that1.HyperParams) {
		return false
	}
	return true
}
func (this *UploadModelRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UploadModelRequest)
	if !ok {
		that2, ok := that.(UploadModelRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Meta.Equal(&that1.Meta) {
		return false
	}
	if !bytes.Equal(this.Model, that1.Model) {
		return false
	}
	return true
}
func (this *UploadModelResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UploadModelResponse)
	if !ok {
		that2, ok := that.(UploadModelResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ModelId != that1.ModelId {
		return false
	}
	if this.ModelUrl != that1.ModelUrl {
		return false
	}
	return true
}
func (this *Model) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&managerpb.Model{")
	s = append(s, "Domain: "+fmt.Sprintf("%#v", this.Domain)+",\n")
	s = append(s, "ModelType: "+fmt.Sprintf("%#v", this.ModelType)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "HyperParams: "+strings.Replace(this.HyperParams.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UploadModelRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&managerpb.UploadModelRequest{")
	s = append(s, "Meta: "+strings.Replace(this.Meta.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Model: "+fmt.Sprintf("%#v", this.Model)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UploadModelResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&managerpb.UploadModelResponse{")
	s = append(s, "ModelId: "+fmt.Sprintf("%#v", this.ModelId)+",\n")
	s = append(s, "ModelUrl: "+fmt.Sprintf("%#v", this.ModelUrl)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringManager(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Manager service

type ManagerClient interface {
	UploadModel(ctx context.Context, in *UploadModelRequest, opts ...grpc.CallOption) (*UploadModelResponse, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) UploadModel(ctx context.Context, in *UploadModelRequest, opts ...grpc.CallOption) (*UploadModelResponse, error) {
	out := new(UploadModelResponse)
	err := grpc.Invoke(ctx, "/managerpb.Manager/UploadModel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerServer interface {
	UploadModel(context.Context, *UploadModelRequest) (*UploadModelResponse, error)
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_UploadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UploadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/managerpb.Manager/UploadModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UploadModel(ctx, req.(*UploadModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "managerpb.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadModel",
			Handler:    _Manager_UploadModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/managerpb/manager.proto",
}

func (m *Model) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Model) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Domain) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.ModelType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.ModelType)))
		i += copy(dAtA[i:], m.ModelType)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	dAtA[i] = 0x42
	i++
	i = encodeVarintManager(dAtA, i, uint64(m.HyperParams.Size()))
	n1, err := m.HyperParams.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *UploadModelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadModelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintManager(dAtA, i, uint64(m.Meta.Size()))
	n2, err := m.Meta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Model) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.Model)))
		i += copy(dAtA[i:], m.Model)
	}
	return i, nil
}

func (m *UploadModelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadModelResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ModelId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintManager(dAtA, i, uint64(m.ModelId))
	}
	if len(m.ModelUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintManager(dAtA, i, uint64(len(m.ModelUrl)))
		i += copy(dAtA[i:], m.ModelUrl)
	}
	return i, nil
}

func encodeVarintManager(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Model) Size() (n int) {
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	l = len(m.ModelType)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	l = m.HyperParams.Size()
	n += 1 + l + sovManager(uint64(l))
	return n
}

func (m *UploadModelRequest) Size() (n int) {
	var l int
	_ = l
	l = m.Meta.Size()
	n += 1 + l + sovManager(uint64(l))
	l = len(m.Model)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	return n
}

func (m *UploadModelResponse) Size() (n int) {
	var l int
	_ = l
	if m.ModelId != 0 {
		n += 1 + sovManager(uint64(m.ModelId))
	}
	l = len(m.ModelUrl)
	if l > 0 {
		n += 1 + l + sovManager(uint64(l))
	}
	return n
}

func sovManager(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozManager(x uint64) (n int) {
	return sovManager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Model) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Model{`,
		`Domain:` + fmt.Sprintf("%v", this.Domain) + `,`,
		`ModelType:` + fmt.Sprintf("%v", this.ModelType) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`HyperParams:` + strings.Replace(strings.Replace(this.HyperParams.String(), "HyperParams", "aggregatorpb.HyperParams", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UploadModelRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UploadModelRequest{`,
		`Meta:` + strings.Replace(strings.Replace(this.Meta.String(), "Model", "Model", 1), `&`, ``, 1) + `,`,
		`Model:` + fmt.Sprintf("%v", this.Model) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UploadModelResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UploadModelResponse{`,
		`ModelId:` + fmt.Sprintf("%v", this.ModelId) + `,`,
		`ModelUrl:` + fmt.Sprintf("%v", this.ModelUrl) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringManager(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Model) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Model: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Model: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HyperParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HyperParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UploadModelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadModelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadModelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Model = append(m.Model[:0], dAtA[iNdEx:postIndex]...)
			if m.Model == nil {
				m.Model = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UploadModelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadModelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadModelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelId", wireType)
			}
			m.ModelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModelId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipManager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManager
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthManager
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowManager
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipManager(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthManager = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManager   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protobuf/managerpb/manager.proto", fileDescriptorManager) }

var fileDescriptorManager = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x31, 0xae, 0xd3, 0x40,
	0x10, 0xf5, 0x82, 0xf3, 0xff, 0xcf, 0xf8, 0x17, 0x68, 0x41, 0xc8, 0x09, 0xca, 0x62, 0xb9, 0x8a,
	0x10, 0x71, 0xa4, 0xd0, 0x53, 0xa4, 0x82, 0x22, 0x08, 0x59, 0x04, 0x89, 0x2a, 0x5a, 0xc7, 0x8b,
	0x63, 0x61, 0x7b, 0x97, 0xb5, 0x5d, 0xa4, 0xe3, 0x08, 0x1c, 0x83, 0x23, 0x70, 0x84, 0x94, 0x29,
	0xa9, 0x10, 0x31, 0x0d, 0x65, 0x8e, 0x80, 0x3c, 0x1b, 0x12, 0x23, 0x44, 0xe5, 0x79, 0xf3, 0xde,
	0x3c, 0xbf, 0x99, 0x05, 0x4f, 0x69, 0x59, 0xc9, 0xa8, 0x7e, 0x3f, 0xcd, 0x79, 0xc1, 0x13, 0xa1,
	0x55, 0xf4, 0xa7, 0x0a, 0x90, 0xa2, 0xfd, 0x33, 0x31, 0x9c, 0x24, 0x69, 0xb5, 0xa9, 0xa3, 0x60,
	0x2d, 0xf3, 0x69, 0x22, 0x13, 0x39, 0x3d, 0x0f, 0xb7, 0x08, 0x01, 0x56, 0x66, 0x72, 0xf8, 0xbc,
	0x23, 0xcf, 0xea, 0x0f, 0x13, 0x9e, 0xb6, 0x1f, 0x9e, 0x5e, 0xc6, 0x78, 0x92, 0x68, 0x91, 0xf0,
	0x4a, 0xb6, 0xbf, 0xbd, 0x00, 0x33, 0xef, 0x7f, 0x25, 0xd0, 0x5b, 0xc8, 0x58, 0x64, 0xf4, 0x21,
	0x5c, 0xc5, 0x32, 0xe7, 0x69, 0xe1, 0xde, 0xf1, 0xc8, 0xb8, 0x1f, 0x9e, 0x10, 0x1d, 0x01, 0xe4,
	0xad, 0x60, 0x55, 0x6d, 0x95, 0x70, 0xef, 0x22, 0xd7, 0xc7, 0xce, 0x9b, 0xad, 0x12, 0x94, 0x82,
	0x5d, 0xf0, 0x5c, 0xb8, 0x36, 0x12, 0x58, 0x53, 0x0f, 0x9c, 0x58, 0x94, 0x6b, 0x9d, 0xaa, 0x2a,
	0x95, 0x85, 0xdb, 0x43, 0xaa, 0xdb, 0xa2, 0x73, 0xb8, 0xdd, 0x6c, 0x95, 0xd0, 0x2b, 0xc5, 0x35,
	0xcf, 0x4b, 0xf7, 0xc6, 0x23, 0x63, 0x67, 0x36, 0x08, 0xba, 0x61, 0x83, 0x17, 0xad, 0xe2, 0x35,
	0x0a, 0xe6, 0xf6, 0xee, 0xfb, 0x63, 0x2b, 0x74, 0x36, 0x97, 0x96, 0xff, 0x16, 0xe8, 0x52, 0x65,
	0x92, 0xc7, 0x98, 0x3f, 0x14, 0x1f, 0x6b, 0x51, 0x56, 0xf4, 0x09, 0xd8, 0xb9, 0xa8, 0xb8, 0x4b,
	0xd0, 0xf1, 0x5e, 0x70, 0xbe, 0x6c, 0x80, 0xb2, 0x93, 0x11, 0x6a, 0xe8, 0x03, 0xe8, 0xe1, 0x22,
	0xb8, 0xf1, 0x6d, 0x68, 0x80, 0xbf, 0x80, 0xfb, 0x7f, 0xf9, 0x96, 0x4a, 0x16, 0xa5, 0xa0, 0x03,
	0xb8, 0x31, 0x77, 0x48, 0x63, 0x34, 0xb7, 0xc3, 0x6b, 0xc4, 0x2f, 0x63, 0xfa, 0x08, 0xcc, 0x41,
	0x56, 0xb5, 0xce, 0x4e, 0xd7, 0x33, 0xda, 0xa5, 0xce, 0x66, 0xef, 0xe0, 0x7a, 0x61, 0x32, 0xd0,
	0x57, 0xe0, 0x74, 0x9c, 0xe9, 0xa8, 0x13, 0xee, 0xdf, 0x4d, 0x86, 0xec, 0x7f, 0xb4, 0x09, 0xe4,
	0x5b, 0xf3, 0xa7, 0xfb, 0x03, 0xb3, 0xbe, 0x1d, 0x98, 0x75, 0x3c, 0x30, 0xf2, 0xa9, 0x61, 0xe4,
	0x4b, 0xc3, 0xc8, 0xae, 0x61, 0x64, 0xdf, 0x30, 0xf2, 0xa3, 0x61, 0xe4, 0x57, 0xc3, 0xac, 0x63,
	0xc3, 0xc8, 0xe7, 0x9f, 0xcc, 0x8a, 0xae, 0xf0, 0xc5, 0x9f, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x8b, 0xa0, 0xd9, 0x8f, 0x02, 0x00, 0x00,
}
