package services

import (
	"context"
	"go-employees/pkg/pb"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EmployeeTestSuite struct {
	suite.Suite
	service    *employeeService
	grpcServer *grpc.Server
	grpcClient pb.EmployeeServiceClient
}

func (suite *EmployeeTestSuite) SetupSuite() {
	server := grpc.NewServer()
	suite.grpcServer = server

	service := InitEmployeeService()
	suite.service = service
	pb.RegisterEmployeeServiceServer(server, service)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		suite.FailNow("Test failed to Listen: %v\n", err)
	}
	go func() {
		err = server.Serve(listener)
		if err != nil {
			suite.FailNow("Test failed to Serve: %v\n", err)
		}
	}()

	// Creating a client
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		suite.FailNow("Test failed to establish a grpc connection: %v\n", err)
	}
	client := pb.NewEmployeeServiceClient(conn)
	if client == nil {
		suite.FailNow("SUITE grpcClient failed - is nil")
	}
	suite.grpcClient = client
}
func (suite *EmployeeTestSuite) TestAddEmployee() {
	if suite.grpcClient == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	if suite.grpcServer == nil {
		suite.FailNow("grpcClient failed - is nil")
	}

	// Testing valid case
	res, _ := suite.grpcClient.AddEmployee(context.Background(), &pb.AddEmployeeRequest{
		Employee: &pb.AddEmployeeBody{
			Name:     "TestEmployee",
			Birthday: timestamppb.Now(),
			Salary:   float64(1000),
		},
	})
	_, ok := suite.service.employees[int(res.GetId())]
	suite.Assert().Equal(true, ok)
	suite.T().Cleanup(func() {
		delete(suite.service.employees, int(res.GetId()))
	})

	// Testing invalid case
	_, err := suite.grpcClient.AddEmployee(context.Background(), &pb.AddEmployeeRequest{})
	suite.Assert().Equal(codes.InvalidArgument, status.Code(err))
}

func (suite *EmployeeTestSuite) TestRemoveEmployee() {
	if suite.grpcClient == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	if suite.grpcServer == nil {
		suite.FailNow("grpcClient failed - is nil")
	}

	testIdx := 123

	suite.service.employees[testIdx] = &pb.Employee{
		Id:       int64(testIdx),
		Name:     "Test Bob",
		Birthday: timestamppb.Now(),
		Salary:   42.0,
	}
	_, err := suite.grpcClient.RemoveEmployee(context.Background(), &pb.RemoveEmployeeRequest{
		Id: int64(testIdx),
	})
	suite.Assert().Equal(codes.OK, status.Code(err))

	_, ok := suite.service.employees[testIdx]
	suite.Assert().Equal(false, ok)

	_, err = suite.grpcClient.RemoveEmployee(context.Background(), &pb.RemoveEmployeeRequest{})
	suite.Assert().Equal(codes.InvalidArgument, status.Code(err))
}

func (suite *EmployeeTestSuite) TestLongOp() {
	if suite.grpcClient == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	if suite.grpcServer == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	_, err := suite.grpcClient.LongOp(context.Background(), &pb.LongOpRequest{
		Duration: 3,
	})
	suite.Assert().NoError(err)
	suite.Assert().Equal(codes.OK, status.Code(err))
	_, err = suite.grpcClient.LongOp(context.Background(), nil)
	suite.Assert().Equal(codes.Internal, status.Code(err))
}

func (suite *EmployeeTestSuite) TestSortBy() {
	if suite.grpcClient == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	if suite.grpcServer == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	now := ptypes.TimestampNow()
	employees := map[int]*pb.Employee{
		3: {Id: 3, Name: "John Smith", Birthday: now, Salary: 300},
		1: {Id: 1, Name: "Alice Johnson", Birthday: timestamppb.New(time.Now().Add(time.Second * 2)), Salary: 200},
		2: {Id: 2, Name: "Bob Brown", Birthday: timestamppb.New(time.Now().Add(time.Second)), Salary: 50},
	}
	suite.T().Cleanup(func() {
		delete(suite.service.employees, 1)
		delete(suite.service.employees, 2)
		delete(suite.service.employees, 3)
	})
	suite.service.employees = employees

	// sort by name
	resp, _ := suite.grpcClient.SortBy(context.Background(), &pb.SortByRequest{SortField: pb.SortField_NAME})
	suite.Assert().Equal(int64(1), resp.Employees[0].Id)
	suite.Assert().Equal(int64(2), resp.Employees[1].Id)
	suite.Assert().Equal(int64(3), resp.Employees[2].Id)

	// sort by birthday
	resp, _ = suite.grpcClient.SortBy(context.Background(), &pb.SortByRequest{SortField: pb.SortField_BIRTHDAY})
	suite.Assert().Equal(int64(3), resp.Employees[0].Id)
	suite.Assert().Equal(int64(2), resp.Employees[1].Id)
	suite.Assert().Equal(int64(1), resp.Employees[2].Id)

	// sort by salary
	resp, _ = suite.grpcClient.SortBy(context.Background(), &pb.SortByRequest{SortField: pb.SortField_SALARY})
	suite.Assert().Equal(int64(2), resp.Employees[0].Id)
	suite.Assert().Equal(int64(1), resp.Employees[1].Id)
	suite.Assert().Equal(int64(3), resp.Employees[2].Id)

	// invalid sort
	_, err := suite.grpcClient.SortBy(context.Background(), &pb.SortByRequest{SortField: 123})
	suite.Assert().Error(err)
	suite.Assert().Equal(codes.InvalidArgument, status.Code(err))
}

func (suite *EmployeeTestSuite) TestAvgMedianSalary() {
	if suite.grpcClient == nil {
		suite.FailNow("grpcClient failed - is nil")
	}
	if suite.grpcServer == nil {
		suite.FailNow("grpcClient failed - is nil")
	}

	// Assertion failed call
	_, err := suite.grpcClient.AvgMedianSalary(context.Background(), &pb.AvgMedianSalaryRequest{})
	suite.Assert().Equal(codes.NotFound, status.Code(err))

	now := ptypes.TimestampNow()
	employees := map[int]*pb.Employee{
		3: {Id: 3, Name: "John Smith", Birthday: now, Salary: 300},
		1: {Id: 1, Name: "Alice Johnson", Birthday: timestamppb.New(time.Now().Add(time.Second * 2)), Salary: 200},
		2: {Id: 2, Name: "Bob Brown", Birthday: timestamppb.New(time.Now().Add(time.Second)), Salary: 50},
		4: {Id: 4, Name: "Peter Owl", Birthday: timestamppb.New(time.Now().Add(time.Second)), Salary: 50},
		5: {Id: 5, Name: "Tommy Tos", Birthday: timestamppb.New(time.Now().Add(time.Second)), Salary: 150},
	}
	suite.service.employees = employees

	// Assertion successfull call
	resp, _ := suite.grpcClient.AvgMedianSalary(context.Background(), &pb.AvgMedianSalaryRequest{})
	suite.Assert().Equal(float64(150), resp.AvgSalary)
	suite.Assert().Equal(float64(150), resp.MedianSalary)

}

func TestEmployeeTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeTestSuite))
}
