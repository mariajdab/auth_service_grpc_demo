// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	SignupWithPhoneNumber(ctx context.Context, in *SignupWithPhoneNumberRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	VerifyPhoneNumber(ctx context.Context, in *VerifyPhoneNumberRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	LoginWithPhoneNumber(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*GenericResponse, error)
	ValidatePhoneNumberLogin(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*GenericResponse, error)
	GetProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GenericResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignupWithPhoneNumber(ctx context.Context, in *SignupWithPhoneNumberRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SignupWithPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyPhoneNumber(ctx context.Context, in *VerifyPhoneNumberRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/VerifyPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LoginWithPhoneNumber(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/LoginWithPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidatePhoneNumberLogin(ctx context.Context, in *OTP, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ValidatePhoneNumberLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	SignupWithPhoneNumber(context.Context, *SignupWithPhoneNumberRequest) (*GenericResponse, error)
	VerifyPhoneNumber(context.Context, *VerifyPhoneNumberRequest) (*GenericResponse, error)
	LoginWithPhoneNumber(context.Context, *Phone) (*GenericResponse, error)
	ValidatePhoneNumberLogin(context.Context, *OTP) (*GenericResponse, error)
	GetProfile(context.Context, *Empty) (*GenericResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) SignupWithPhoneNumber(context.Context, *SignupWithPhoneNumberRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupWithPhoneNumber not implemented")
}
func (UnimplementedAuthServiceServer) VerifyPhoneNumber(context.Context, *VerifyPhoneNumberRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPhoneNumber not implemented")
}
func (UnimplementedAuthServiceServer) LoginWithPhoneNumber(context.Context, *Phone) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithPhoneNumber not implemented")
}
func (UnimplementedAuthServiceServer) ValidatePhoneNumberLogin(context.Context, *OTP) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePhoneNumberLogin not implemented")
}
func (UnimplementedAuthServiceServer) GetProfile(context.Context, *Empty) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_SignupWithPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupWithPhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignupWithPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SignupWithPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignupWithPhoneNumber(ctx, req.(*SignupWithPhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/VerifyPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyPhoneNumber(ctx, req.(*VerifyPhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LoginWithPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoginWithPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/LoginWithPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoginWithPhoneNumber(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidatePhoneNumberLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidatePhoneNumberLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ValidatePhoneNumberLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidatePhoneNumberLogin(ctx, req.(*OTP))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetProfile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignupWithPhoneNumber",
			Handler:    _AuthService_SignupWithPhoneNumber_Handler,
		},
		{
			MethodName: "VerifyPhoneNumber",
			Handler:    _AuthService_VerifyPhoneNumber_Handler,
		},
		{
			MethodName: "LoginWithPhoneNumber",
			Handler:    _AuthService_LoginWithPhoneNumber_Handler,
		},
		{
			MethodName: "ValidatePhoneNumberLogin",
			Handler:    _AuthService_ValidatePhoneNumberLogin_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _AuthService_GetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

// OTPServiceClient is the client API for OTPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OTPServiceClient interface {
	CreateTwillioOTP(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*GenericResponse, error)
}

type oTPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOTPServiceClient(cc grpc.ClientConnInterface) OTPServiceClient {
	return &oTPServiceClient{cc}
}

func (c *oTPServiceClient) CreateTwillioOTP(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/auth.OTPService/CreateTwillioOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OTPServiceServer is the server API for OTPService service.
// All implementations must embed UnimplementedOTPServiceServer
// for forward compatibility
type OTPServiceServer interface {
	CreateTwillioOTP(context.Context, *Phone) (*GenericResponse, error)
	mustEmbedUnimplementedOTPServiceServer()
}

// UnimplementedOTPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOTPServiceServer struct {
}

func (UnimplementedOTPServiceServer) CreateTwillioOTP(context.Context, *Phone) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTwillioOTP not implemented")
}
func (UnimplementedOTPServiceServer) mustEmbedUnimplementedOTPServiceServer() {}

// UnsafeOTPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OTPServiceServer will
// result in compilation errors.
type UnsafeOTPServiceServer interface {
	mustEmbedUnimplementedOTPServiceServer()
}

func RegisterOTPServiceServer(s grpc.ServiceRegistrar, srv OTPServiceServer) {
	s.RegisterService(&OTPService_ServiceDesc, srv)
}

func _OTPService_CreateTwillioOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServiceServer).CreateTwillioOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.OTPService/CreateTwillioOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServiceServer).CreateTwillioOTP(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

// OTPService_ServiceDesc is the grpc.ServiceDesc for OTPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OTPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.OTPService",
	HandlerType: (*OTPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTwillioOTP",
			Handler:    _OTPService_CreateTwillioOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
