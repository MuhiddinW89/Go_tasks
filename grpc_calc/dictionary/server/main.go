package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "app/dict"

	"google.golang.org/grpc"
)

  var maps = map[string]string{}

type server struct {
	pb.UnimplementedTranslateServiceServer
}


func (s *server) Dictionary(ctx context.Context, req *pb.DictionaryRequest) (*pb.DictionaryResponse, error) {

	log.Println("Dictionary >>>>>>>> ", req)
	if req.Word == "Apple" {
		return &pb.DictionaryResponse{Definition: "Olma"}, nil
	}

	return &pb.DictionaryResponse{}, nil
}


func (s *server) CreateDictionary(ctx context.Context, req *pb.CreateDictionaryRequest)  (*pb.CreateDictionaryResponse, error) {

	log.Println("createDictionary >>>>>>>> ", req)
	maps[req.Word] = req.Definition

	return &pb.CreateDictionaryResponse{}, nil
}


func (s *server) GetDictionary(ctx context.Context, req *pb.GetDictionaryRequest)  (*pb.GetDictionaryResponse, error) {

	log.Println( req)
	
	for k, v := range maps{
		if req.Key == k {
			return &pb.GetDictionaryResponse{Value: v}, nil
		}
}
			return &pb.GetDictionaryResponse{}, nil
}


func (s *server) SumArr(ctx context.Context, req *pb.SumArrRequest)  (*pb.SumArrResponse, error) {
	
	var sum int32 

	for _, v := range req.Data{
		sum += v
	}
	
	return &pb.SumArrResponse{Sum: sum}, nil	
}


func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error){
	a := req.FirsNum
	b := req.SecondNum 
	
	return &pb.AddResponse{Result: a+b }, nil	
}

func (s *server) Div(ctx context.Context, req *pb.DivRequest) (*pb.DivResponse, error){
	a := req.FirsNum
	b := req.SecondNum 
	
	return &pb.DivResponse{Result: a-b }, nil	
}

func (s *server) Mult(ctx context.Context, req *pb.MultRequest) (*pb.MultResponse, error){
	a := req.FirsNum
	b := req.SecondNum 
	
	return &pb.MultResponse{Result: a*b }, nil	
}

func (s *server) Sub(ctx context.Context, req *pb.SubRequest) (*pb.SubResponse, error){
	a := req.FirsNum
	b := req.SecondNum 

	if b == 0 {
		return nil, fmt.Errorf("request should non null")
	}
	
	return &pb.SubResponse{Result: a/b }, nil	
}

func (s *server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error){
	
	a := req.FirsNum 

	if a < 0 {
		return nil, fmt.Errorf("request should positive")
	}

	var gues float64
	gues = 1;

	for{
		gues += 1

		if gues*gues == a {
			return &pb.SqrtResponse{Result: gues }, nil
			break
		}

	}
	
	return &pb.SqrtResponse{}, nil	
}

func (s *server) Pow(ctx context.Context, req *pb.PowRequest) (*pb.PowResponse, error){
	a := req.FirsNum
	b := req.SecondNum 
	var i int64
	var result int64
	result = 1
	for i = 1; i <= b; i++ {
		result *= a
	}

	return &pb.PowResponse{Result: result }, nil	
}

func (s *server) Min(ctx context.Context, req *pb.MinRequest) (*pb.MinResponse, error){
	a := req.FirsNum
	b := req.SecondNum 
	
	if a < b {
		return &pb.MinResponse{Result: a }, nil	
	}else if a > b {
		return &pb.MinResponse{Result: b }, nil	
	}

	return &pb.MinResponse{Result: a }, nil	
}

func main() {

	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTranslateServiceServer(s, &server{})

	fmt.Println("Listen RPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
