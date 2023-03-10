// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: pkg/pb/employee.proto

package pb

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

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*AddEmployeeResponse, error)
	RemoveEmployee(ctx context.Context, in *RemoveEmployeeRequest, opts ...grpc.CallOption) (*RemoveEmployeeResponse, error)
	LongOp(ctx context.Context, in *LongOpRequest, opts ...grpc.CallOption) (*LongOpResponse, error)
	SortBy(ctx context.Context, in *SortByRequest, opts ...grpc.CallOption) (*SortByResponse, error)
	AvgMedianSalary(ctx context.Context, in *AvgMedianSalaryRequest, opts ...grpc.CallOption) (*AvgMedianSalaryResponse, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*AddEmployeeResponse, error) {
	out := new(AddEmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/AddEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) RemoveEmployee(ctx context.Context, in *RemoveEmployeeRequest, opts ...grpc.CallOption) (*RemoveEmployeeResponse, error) {
	out := new(RemoveEmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/RemoveEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) LongOp(ctx context.Context, in *LongOpRequest, opts ...grpc.CallOption) (*LongOpResponse, error) {
	out := new(LongOpResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/LongOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) SortBy(ctx context.Context, in *SortByRequest, opts ...grpc.CallOption) (*SortByResponse, error) {
	out := new(SortByResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/SortBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) AvgMedianSalary(ctx context.Context, in *AvgMedianSalaryRequest, opts ...grpc.CallOption) (*AvgMedianSalaryResponse, error) {
	out := new(AvgMedianSalaryResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/AvgMedianSalary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	AddEmployee(context.Context, *AddEmployeeRequest) (*AddEmployeeResponse, error)
	RemoveEmployee(context.Context, *RemoveEmployeeRequest) (*RemoveEmployeeResponse, error)
	LongOp(context.Context, *LongOpRequest) (*LongOpResponse, error)
	SortBy(context.Context, *SortByRequest) (*SortByResponse, error)
	AvgMedianSalary(context.Context, *AvgMedianSalaryRequest) (*AvgMedianSalaryResponse, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) AddEmployee(context.Context, *AddEmployeeRequest) (*AddEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) RemoveEmployee(context.Context, *RemoveEmployeeRequest) (*RemoveEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) LongOp(context.Context, *LongOpRequest) (*LongOpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LongOp not implemented")
}
func (UnimplementedEmployeeServiceServer) SortBy(context.Context, *SortByRequest) (*SortByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortBy not implemented")
}
func (UnimplementedEmployeeServiceServer) AvgMedianSalary(context.Context, *AvgMedianSalaryRequest) (*AvgMedianSalaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvgMedianSalary not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/AddEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).AddEmployee(ctx, req.(*AddEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_RemoveEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).RemoveEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/RemoveEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).RemoveEmployee(ctx, req.(*RemoveEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_LongOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongOpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).LongOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/LongOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).LongOp(ctx, req.(*LongOpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_SortBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).SortBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/SortBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).SortBy(ctx, req.(*SortByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_AvgMedianSalary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvgMedianSalaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).AvgMedianSalary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/AvgMedianSalary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).AvgMedianSalary(ctx, req.(*AvgMedianSalaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEmployee",
			Handler:    _EmployeeService_AddEmployee_Handler,
		},
		{
			MethodName: "RemoveEmployee",
			Handler:    _EmployeeService_RemoveEmployee_Handler,
		},
		{
			MethodName: "LongOp",
			Handler:    _EmployeeService_LongOp_Handler,
		},
		{
			MethodName: "SortBy",
			Handler:    _EmployeeService_SortBy_Handler,
		},
		{
			MethodName: "AvgMedianSalary",
			Handler:    _EmployeeService_AvgMedianSalary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/employee.proto",
}
