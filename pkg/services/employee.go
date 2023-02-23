package services

import (
	"context"
	"go-employees/pkg/pb"
	"sort"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type employeeService struct {
	pb.UnimplementedEmployeeServiceServer

	// Not thread safe
	employees map[int]*pb.Employee
}

func InitEmployeeService() *employeeService {
	return &employeeService{employees: make(map[int]*pb.Employee, 50)}
}

func (s *employeeService) AddEmployee(ctx context.Context, req *pb.AddEmployeeRequest) (*pb.AddEmployeeResponse, error) {
	var empl *pb.AddEmployeeBody

	if empl = req.GetEmployee(); empl == nil {
		return nil, status.Error(codes.InvalidArgument, "Payload is missing")
	}
	idx := len(s.employees) + 1
	newEmployee := &pb.Employee{
		Id:       int64(idx),
		Name:     empl.Name,
		Birthday: empl.Birthday,
		Salary:   empl.Salary,
	}
	s.employees[idx] = newEmployee
	return &pb.AddEmployeeResponse{Status: int64(codes.OK), Id: int64(idx)}, nil
}

func (s *employeeService) RemoveEmployee(ctx context.Context, req *pb.RemoveEmployeeRequest) (*pb.RemoveEmployeeResponse, error) {
	var key int64
	if key = req.GetId(); key == 0 {
		return nil, status.Error(codes.InvalidArgument, "Missing Id parameter")
	}
	delete(s.employees, int(key))
	return &pb.RemoveEmployeeResponse{
		Status: int64(codes.OK),
	}, nil
}

func (s *employeeService) LongOp(ctx context.Context, req *pb.LongOpRequest) (*pb.LongOpResponse, error) {
	time.Sleep(time.Duration(req.GetDuration()) * time.Second)
	return &pb.LongOpResponse{Ok: int64(codes.OK)}, nil
}

func (s *employeeService) SortBy(ctx context.Context, req *pb.SortByRequest) (*pb.SortByResponse, error) {
	// Had to use this ugly structure, instead of default case, to avoid unneccassary mem allocation
	if req.SortField != pb.SortField_BIRTHDAY &&
		req.SortField != pb.SortField_SALARY &&
		req.SortField != pb.SortField_NAME {
		return nil, status.Error(codes.InvalidArgument, "Wrong field for sorting")
	}

	emplKeys := make([]int, 0, len(s.employees))
	for k := range s.employees {
		emplKeys = append(emplKeys, k)
	}
	switch req.GetSortField() {
	case pb.SortField_NAME:
		sort.Slice(emplKeys, func(i, j int) bool {
			return s.employees[emplKeys[i]].Name < s.employees[emplKeys[j]].Name
		})
	case pb.SortField_BIRTHDAY:
		sort.Slice(emplKeys, func(i, j int) bool {
			ti, _ := ptypes.Timestamp(s.employees[emplKeys[i]].Birthday)
			tj, _ := ptypes.Timestamp(s.employees[emplKeys[j]].Birthday)
			return ti.Before(tj)
		})
	case pb.SortField_SALARY:
		sort.Slice(emplKeys, func(i, j int) bool {
			return s.employees[emplKeys[i]].Salary < s.employees[emplKeys[j]].Salary
		})
	}
	sortedEmpl := make([]*pb.Employee, 0, len(emplKeys))
	for _, el := range emplKeys {
		sortedEmpl = append(sortedEmpl, s.employees[el])
	}
	return &pb.SortByResponse{Employees: sortedEmpl}, nil
}

func (s *employeeService) AvgMedianSalary(ctx context.Context, req *pb.AvgMedianSalaryRequest) (*pb.AvgMedianSalaryResponse, error) {
	if len(s.employees) == 0 {
		return nil, status.Error(codes.NotFound, "No records")
	}

	salaries := make(sort.Float64Slice, 0, len(s.employees))
	sum := float64(0)
	for _, v := range s.employees {
		sum += v.Salary
		salaries = append(salaries, v.Salary)
	}
	avg := sum / float64(len(s.employees))

	sort.Float64s(salaries)
	med := salaries[len(salaries)/2]
	return &pb.AvgMedianSalaryResponse{AvgSalary: avg, MedianSalary: med}, nil
}
