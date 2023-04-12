// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: dict/dict.proto

package dict

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TranslateServiceClient is the client API for TranslateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslateServiceClient interface {
	Dictionary(ctx context.Context, in *DictionaryRequest, opts ...grpc.CallOption) (*DictionaryResponse, error)
	CreateDictionary(ctx context.Context, in *CreateDictionaryRequest, opts ...grpc.CallOption) (*CreateDictionaryResponse, error)
	GetDictionary(ctx context.Context, in *GetDictionaryRequest, opts ...grpc.CallOption) (*GetDictionaryResponse, error)
	SumArr(ctx context.Context, in *SumArrRequest, opts ...grpc.CallOption) (*SumArrResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Div(ctx context.Context, in *DivRequest, opts ...grpc.CallOption) (*DivResponse, error)
	Mult(ctx context.Context, in *MultRequest, opts ...grpc.CallOption) (*MultResponse, error)
	Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
	Pow(ctx context.Context, in *PowRequest, opts ...grpc.CallOption) (*PowResponse, error)
	Min(ctx context.Context, in *MinRequest, opts ...grpc.CallOption) (*MinResponse, error)
}

type translateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslateServiceClient(cc grpc.ClientConnInterface) TranslateServiceClient {
	return &translateServiceClient{cc}
}

func (c *translateServiceClient) Dictionary(ctx context.Context, in *DictionaryRequest, opts ...grpc.CallOption) (*DictionaryResponse, error) {
	out := new(DictionaryResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Dictionary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) CreateDictionary(ctx context.Context, in *CreateDictionaryRequest, opts ...grpc.CallOption) (*CreateDictionaryResponse, error) {
	out := new(CreateDictionaryResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/CreateDictionary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) GetDictionary(ctx context.Context, in *GetDictionaryRequest, opts ...grpc.CallOption) (*GetDictionaryResponse, error) {
	out := new(GetDictionaryResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/GetDictionary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) SumArr(ctx context.Context, in *SumArrRequest, opts ...grpc.CallOption) (*SumArrResponse, error) {
	out := new(SumArrResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/SumArr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Div(ctx context.Context, in *DivRequest, opts ...grpc.CallOption) (*DivResponse, error) {
	out := new(DivResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Mult(ctx context.Context, in *MultRequest, opts ...grpc.CallOption) (*MultResponse, error) {
	out := new(MultResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Mult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error) {
	out := new(SubResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Pow(ctx context.Context, in *PowRequest, opts ...grpc.CallOption) (*PowResponse, error) {
	out := new(PowResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Pow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translateServiceClient) Min(ctx context.Context, in *MinRequest, opts ...grpc.CallOption) (*MinResponse, error) {
	out := new(MinResponse)
	err := c.cc.Invoke(ctx, "/dict.TranslateService/Min", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslateServiceServer is the server API for TranslateService service.
// All implementations must embed UnimplementedTranslateServiceServer
// for forward compatibility
type TranslateServiceServer interface {
	Dictionary(context.Context, *DictionaryRequest) (*DictionaryResponse, error)
	CreateDictionary(context.Context, *CreateDictionaryRequest) (*CreateDictionaryResponse, error)
	GetDictionary(context.Context, *GetDictionaryRequest) (*GetDictionaryResponse, error)
	SumArr(context.Context, *SumArrRequest) (*SumArrResponse, error)
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Div(context.Context, *DivRequest) (*DivResponse, error)
	Mult(context.Context, *MultRequest) (*MultResponse, error)
	Sub(context.Context, *SubRequest) (*SubResponse, error)
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
	Pow(context.Context, *PowRequest) (*PowResponse, error)
	Min(context.Context, *MinRequest) (*MinResponse, error)
	mustEmbedUnimplementedTranslateServiceServer()
}

// UnimplementedTranslateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTranslateServiceServer struct {
}

func (UnimplementedTranslateServiceServer) Dictionary(context.Context, *DictionaryRequest) (*DictionaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dictionary not implemented")
}
func (UnimplementedTranslateServiceServer) CreateDictionary(context.Context, *CreateDictionaryRequest) (*CreateDictionaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDictionary not implemented")
}
func (UnimplementedTranslateServiceServer) GetDictionary(context.Context, *GetDictionaryRequest) (*GetDictionaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictionary not implemented")
}
func (UnimplementedTranslateServiceServer) SumArr(context.Context, *SumArrRequest) (*SumArrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumArr not implemented")
}
func (UnimplementedTranslateServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTranslateServiceServer) Div(context.Context, *DivRequest) (*DivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (UnimplementedTranslateServiceServer) Mult(context.Context, *MultRequest) (*MultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mult not implemented")
}
func (UnimplementedTranslateServiceServer) Sub(context.Context, *SubRequest) (*SubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (UnimplementedTranslateServiceServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedTranslateServiceServer) Pow(context.Context, *PowRequest) (*PowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pow not implemented")
}
func (UnimplementedTranslateServiceServer) Min(context.Context, *MinRequest) (*MinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Min not implemented")
}
func (UnimplementedTranslateServiceServer) mustEmbedUnimplementedTranslateServiceServer() {}

// UnsafeTranslateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslateServiceServer will
// result in compilation errors.
type UnsafeTranslateServiceServer interface {
	mustEmbedUnimplementedTranslateServiceServer()
}

func RegisterTranslateServiceServer(s grpc.ServiceRegistrar, srv TranslateServiceServer) {
	s.RegisterService(&TranslateService_ServiceDesc, srv)
}

func _TranslateService_Dictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Dictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Dictionary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Dictionary(ctx, req.(*DictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_CreateDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).CreateDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/CreateDictionary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).CreateDictionary(ctx, req.(*CreateDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_GetDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).GetDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/GetDictionary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).GetDictionary(ctx, req.(*GetDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_SumArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumArrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).SumArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/SumArr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).SumArr(ctx, req.(*SumArrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Div(ctx, req.(*DivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Mult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Mult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Mult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Mult(ctx, req.(*MultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Sub(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Pow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Pow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Pow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Pow(ctx, req.(*PowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslateService_Min_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).Min(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dict.TranslateService/Min",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).Min(ctx, req.(*MinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TranslateService_ServiceDesc is the grpc.ServiceDesc for TranslateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranslateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dict.TranslateService",
	HandlerType: (*TranslateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dictionary",
			Handler:    _TranslateService_Dictionary_Handler,
		},
		{
			MethodName: "CreateDictionary",
			Handler:    _TranslateService_CreateDictionary_Handler,
		},
		{
			MethodName: "GetDictionary",
			Handler:    _TranslateService_GetDictionary_Handler,
		},
		{
			MethodName: "SumArr",
			Handler:    _TranslateService_SumArr_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _TranslateService_Add_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _TranslateService_Div_Handler,
		},
		{
			MethodName: "Mult",
			Handler:    _TranslateService_Mult_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _TranslateService_Sub_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _TranslateService_Sqrt_Handler,
		},
		{
			MethodName: "Pow",
			Handler:    _TranslateService_Pow_Handler,
		},
		{
			MethodName: "Min",
			Handler:    _TranslateService_Min_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dict/dict.proto",
}
